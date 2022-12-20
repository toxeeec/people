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

type Service interface {
	DateStr(t time.Time) string
	Create(userID uint, r *multipart.Reader) (people.ImageResponse, error)
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

func (s *imageService) DateStr(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day())
}

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

	dateStr := s.DateStr(time.Now())
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
	fName := uuid.New().String() + ext
	path := baseDir + fmt.Sprintf("%s/%s", dateStr, fName)
	f, err := os.Create(path)
	if err != nil {
		return people.ImageResponse{}, err
	}
	defer f.Close()

	_, err = io.Copy(f, mf)
	if err != nil {
		return people.ImageResponse{}, err
	}

	ir, err := s.ir.Create(fName, userID)
	if err != nil {
		go os.Remove(path)
		return people.ImageResponse{}, err
	}
	return ir, nil
}

func (s *imageService) wipeUnused() {
	const expiredAfter = time.Hour
	is, err := s.ir.ListUnusedBefore(time.Now().Add(-expiredAfter))
	if err != nil {
		return
	}
	for _, i := range is {
		dateStr := s.DateStr(i.CreatedAt)
		os.Remove(baseDir + fmt.Sprintf("%s/%s", dateStr, i.Path))
	}
	s.ir.DeleteMany(Slice(is).IDs())
}
