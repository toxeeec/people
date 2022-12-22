package image_test

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/service/image"
)

type ImageSuite struct {
	suite.Suite
	is image.Service
	ir repository.Image
	ur repository.User
	pr repository.Post
}

func (s *ImageSuite) TestDateString() {
	t1 := time.UnixMilli(0)
	assert.Equal(s.T(), "1970/1/1", image.DateStr(t1))
	t2, _ := time.Parse("2006-01-02", "2022-12-31")
	assert.Equal(s.T(), "2022/12/31", image.DateStr(t2))
}

func (s *ImageSuite) TestCreate() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	const formData = `
--Boundary
Content-Disposition: form-data; name="image"; filename="file1.jpg"
Content-Type: image/jpeg

image-data
--Boundary--
`
	r := multipart.NewReader(strings.NewReader(formData), "Boundary")

	ir, err := s.is.Create(u.ID, r)
	assert.NoError(s.T(), err)
	i, _ := s.ir.Get(ir.ID)
	assert.Equal(s.T(), u.ID, i.UserID)

	bytes, err := os.ReadFile("images/" + fmt.Sprintf("%s/%s", image.DateStr(i.CreatedAt), i.Name))
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "image-data", string(bytes))
}

func (s *ImageSuite) TestAddToPost() {
	var au1 people.AuthUser
	var au2 people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au1)
	gofakeit.Struct(&au2)
	gofakeit.Struct(&np)
	u1, _ := s.ur.Create(au1)
	u2, _ := s.ur.Create(au2)
	p, _ := s.pr.Create(np, u1.ID, nil)
	i1, _ := s.ir.Create(gofakeit.LetterN(40), u1.ID)
	i2, _ := s.ir.Create(gofakeit.LetterN(40), u2.ID)

	notFoundError := people.NotFoundError
	authError := people.AuthError

	tests := map[string]struct {
		imageIDs []uint
		valid    bool
		kind     *people.ErrorKind
	}{
		"image not found": {[]uint{i2.ID + 5}, false, &notFoundError},
		"image not owned": {[]uint{i2.ID}, false, &authError},
		"valid":           {[]uint{i1.ID}, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			paths, err := s.is.AddToPost(tc.imageIDs, p.ID, u1.ID)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Len(s.T(), paths, 1)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *ImageSuite) TestListPostImages() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p, _ := s.pr.Create(np, u.ID, nil)
	var ids [3]uint
	for i := range ids {
		ir, _ := s.ir.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}
	s.ir.CreatePostImages(ids[:], p.ID)

	paths, err := s.is.ListPostImages(p.ID)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), paths, len(ids))
}

func (s *ImageSuite) TestListPostsImages() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var postIDs [3]uint
	var ids [len(postIDs)]uint
	for i := range ids {
		var np people.NewPost
		gofakeit.Struct(&np)
		p, _ := s.pr.Create(np, u.ID, nil)
		postIDs[i] = p.ID
		ir, _ := s.ir.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}
	for i := range postIDs {
		s.ir.CreatePostImages(ids[:], postIDs[i])
	}

	pathMap, err := s.is.ListPostsImages(postIDs[:])
	assert.NoError(s.T(), err)
	for k, paths := range pathMap {
		assert.Contains(s.T(), postIDs, k)
		assert.Len(s.T(), paths, len(ids))
	}
}

func (s *ImageSuite) SetupTest() {
	os.RemoveAll("images/")
	um := map[uint]people.User{}
	im := map[uint]people.Image{}
	pm := map[uint]people.Post{}
	s.ir = inmem.NewImageRepository(im)
	s.ur = inmem.NewUserRepository(um)
	s.pr = inmem.NewPostRepository(pm)
	s.is = image.NewService(s.ir)
}

func TestImageSuite(t *testing.T) {
	suite.Run(t, new(ImageSuite))
}
