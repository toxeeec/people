package image

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
)

func DateStr(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day())
}

func path(t time.Time, name string) string {
	return baseDir + fmt.Sprintf("%s/%s", DateStr(t), name)
}

type Service interface {
	Create(userID uint, r *multipart.Reader) (people.ImageResponse, error)
	AddToPost(ids []uint, postID, userID uint) ([]string, error)
	UpdateUserImage(id *uint, userID uint) (*string, error)
	ListPostImages(postID uint) ([]string, error)
	ListPostsImages(postIDs []uint) (map[uint][]string, error)
	GetUserImage(userID uint) (*string, error)
	ListUsersImages(userIDs []uint) (map[uint]*string, error)
}

type imageService struct {
	ir repository.Image
}

func NewService(ir repository.Image) Service {
	err := os.Mkdir("images", os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(err.Error())
	}
	s := imageService{ir: ir}
	c := cron.New()
	c.AddFunc("@every 1h", s.wipeUnused)
	c.Start()
	return &s
}

const (
	baseDir = "images/"
)

func (s *imageService) Create(userID uint, r *multipart.Reader) (people.ImageResponse, error) {
	const MB = 1 << 20
	const maxSize = 5 * MB
	const maxSizeString = "5 MB"
	form, err := r.ReadForm(maxSize)
	if err != nil {
		return people.ImageResponse{}, err
	}
	fhs, ok := form.File["image"]
	if !ok || len(fhs) == 0 {
		return people.ImageResponse{}, errors.New("Image not found")
	}
	fh := fhs[0]
	if fh.Size > maxSize {
		return people.ImageResponse{}, service.NewError(people.ValidationError, "Image size cannot be larger than "+maxSizeString)
	}

	now := time.Now()
	dateStr := DateStr(now)
	err = os.MkdirAll(baseDir+dateStr, os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return people.ImageResponse{}, err
	}

	mf, err := fh.Open()
	if err != nil {
		return people.ImageResponse{}, err
	}
	defer mf.Close()

	ext := filepath.Ext(fh.Filename)
	name := uuid.New().String() + ext
	path := path(now, name)
	f, err := os.Create(path)
	if err != nil {
		return people.ImageResponse{}, err
	}
	defer f.Close()

	_, err = io.Copy(f, mf)
	if err != nil {
		return people.ImageResponse{}, err
	}

	ir, err := s.ir.Create(name, userID)
	if err != nil {
		go os.Remove(path)
		return people.ImageResponse{}, err
	}
	return ir, nil
}

func (s *imageService) AddToPost(ids []uint, postID, userID uint) ([]string, error) {
	imgs, err := s.ir.List(ids)
	if err != nil {
		return nil, err
	}
	paths := make([]string, len(ids))
	for i, id := range ids {
		found := false
		for _, img := range imgs {
			if img.ID != id {
				continue
			}
			if err := validate(img, userID); err != nil {
				return nil, err
			}
			paths[i] = "/" + path(img.CreatedAt, img.Name)
			found = true
		}
		if !found {
			return nil, service.NewError(people.NotFoundError, fmt.Sprintf("Image not found: %v", id))
		}
	}
	err = s.ir.CreatePostImages(ids, postID)
	if err != nil {
		return nil, err
	}
	err = s.ir.MarkUsed(ids)
	if err != nil {
		go s.ir.DeleteManyPostImages(ids)
		return nil, err
	}
	return paths, nil
}

func (s *imageService) UpdateUserImage(imgID *uint, userID uint) (*string, error) {
	if imgID == nil {
		s.ir.DeleteUserImage(userID)
		return nil, nil

	}
	id := *imgID
	img, err := s.ir.Get(id)
	if err != nil {
		return nil, service.NewError(people.NotFoundError, fmt.Sprintf("Image not found: %v", id))
	}
	if err := validate(img, userID); err != nil {
		return nil, err
	}
	w, h, err := getDimensions(img)
	if err != nil {
		return nil, err
	}
	if w != h {
		return nil, service.NewError(people.ResourceError, "Image must have a 1:1 aspect ratio")
	}
	// TODO: resize
	prev, _ := s.ir.GetUserImage(userID)
	if prev.ID == img.ID {
		path := "/" + path(img.CreatedAt, img.Name)
		return &path, nil
	}
	err = s.ir.CreateUserImage(id, userID)
	if err != nil {
		return nil, err
	}
	err = s.ir.MarkUsed([]uint{id})
	if err != nil {
		go s.ir.DeleteUserImage(id)
		return nil, err
	}
	s.ir.DeleteMany([]uint{prev.ID})
	path := "/" + path(img.CreatedAt, img.Name)
	return &path, nil
}

func (s *imageService) ListPostImages(postID uint) ([]string, error) {
	imgs, err := s.ir.ListPostImages(postID)
	if err != nil {
		return nil, err
	}
	paths := make([]string, len(imgs))
	for i, img := range imgs {
		paths[i] = "/" + path(img.CreatedAt, img.Name)
	}
	return paths, nil
}

func (s *imageService) ListPostsImages(postIDs []uint) (map[uint][]string, error) {
	idsMap, err := s.ir.ListPostsImageIDs(postIDs)
	if err != nil {
		return nil, err
	}
	idsSet := map[uint]struct{}{}
	var length uint
	for _, ids := range idsMap {
		for _, id := range ids {
			idsSet[id] = struct{}{}
			length++
		}
	}
	ids := make([]uint, 0, length)
	for k := range idsSet {
		ids = append(ids, k)
	}

	imgs, err := s.ir.List(ids)
	if err != nil {
		return nil, err
	}
	imgsMap := make(map[uint]people.Image, len(imgs))
	for _, img := range imgs {
		imgsMap[img.ID] = img
	}
	pathsMap := make(map[uint][]string, len(idsMap))
	for postID, ids := range idsMap {
		paths := make([]string, 0, 4)
		for _, id := range ids {
			paths = append(paths, "/"+path(imgsMap[id].CreatedAt, imgsMap[id].Name))
		}
		pathsMap[postID] = paths
	}
	return pathsMap, nil
}

func (s *imageService) GetUserImage(userID uint) (*string, error) {
	img, err := s.ir.GetUserImage(userID)
	if err != nil {
		return nil, nil
	}
	path := "/" + path(img.CreatedAt, img.Name)
	return &path, nil
}

func (s *imageService) ListUsersImages(userIDs []uint) (map[uint]*string, error) {
	idsMap, err := s.ir.ListUsersImageIDs(userIDs)
	if err != nil {
		return nil, err
	}
	ids := make([]uint, 0, len(idsMap))
	for _, id := range idsMap {
		if id != nil {
			ids = append(ids, *id)
		}
	}
	imgs, err := s.ir.List(ids)
	if err != nil {
		return nil, err
	}
	imgsMap := make(map[uint]people.Image, len(imgs))
	for _, img := range imgs {
		imgsMap[img.ID] = img
	}
	pathsMap := make(map[uint]*string, len(idsMap))
	for userID, imgID := range idsMap {
		if imgID == nil {
			pathsMap[userID] = nil
			continue
		}
		img := imgsMap[*imgID]
		path := "/" + path(img.CreatedAt, img.Name)
		pathsMap[userID] = &path
	}
	return pathsMap, nil
}

func (s *imageService) wipeUnused() {
	imgs, err := s.ir.ListUnusedBefore(time.Now().Add(-expiredAfter))
	if err != nil {
		return
	}
	for _, i := range imgs {
		os.Remove(path(i.CreatedAt, i.Name))
	}
	s.ir.DeleteMany(IDs(imgs))
}
