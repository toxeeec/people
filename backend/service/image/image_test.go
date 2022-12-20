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
}

func (s *ImageSuite) TestDateString() {
	t1 := time.UnixMilli(0)
	assert.Equal(s.T(), "1970/1/1", s.is.DateStr(t1))
	t2, _ := time.Parse("2006-01-02", "2022-12-31")
	assert.Equal(s.T(), "2022/12/31", s.is.DateStr(t2))
}

func (s *ImageSuite) TestCreate() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	const formData = `--Boundary
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

	bytes, err := os.ReadFile("images/" + fmt.Sprintf("%s/%s", s.is.DateStr(i.CreatedAt), i.Path))
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "image-data", string(bytes))
}

func (s *ImageSuite) SetupTest() {
	os.RemoveAll("images/")
	um := map[uint]people.User{}
	pm := map[uint]people.Image{}
	s.ir = inmem.NewImageRepository(pm)
	s.ur = inmem.NewUserRepository(um)
	s.is = image.NewService(s.ir)
}

func TestImageSuite(t *testing.T) {
	suite.Run(t, new(ImageSuite))
}
