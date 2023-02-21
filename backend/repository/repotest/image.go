package repotest

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
)

type ImageSuite struct {
	suite.Suite
	repo repository.Image
	ur   repository.User
	pr   repository.Post
	fns  TestFns
}

func NewImageSuite(ir repository.Image, ur repository.User, pr repository.Post, fns TestFns) *ImageSuite {
	return &ImageSuite{repo: ir, ur: ur, pr: pr, fns: fns}
}

func (s *ImageSuite) TestCreate() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)

	name := gofakeit.LetterN(40)
	ir, err := s.repo.Create(name, u.ID)
	assert.NoError(s.T(), err)
	assert.NotZero(s.T(), ir.ID)
}

func (s *ImageSuite) TestGet() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	name := gofakeit.LetterN(40)
	ir, _ := s.repo.Create(name, u.ID)

	i, err := s.repo.Get(ir.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), ir.ID, i.ID)
	assert.Equal(s.T(), name, i.Name)
	assert.Equal(s.T(), u.ID, i.UserID)
}

func (s *ImageSuite) TestListUnusedBefore() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var irs [3]people.ImageResponse
	for i := range irs {
		irs[i], _ = s.repo.Create(gofakeit.LetterN(40), u.ID)
	}
	time.Sleep(1 * time.Second)

	imgs, err := s.repo.ListUnusedBefore(time.Now())
	assert.NoError(s.T(), err)
	for _, i := range imgs {
		assert.NotZero(s.T(), i.ID)
	}
}

func (s *ImageSuite) TestDeleteMany() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var ids [3]uint
	for i := range ids {
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}

	s.repo.DeleteMany(ids[:])
	for _, i := range ids {
		_, err := s.repo.Get(i)
		assert.Error(s.T(), err)
	}
}

func (s *ImageSuite) TestList() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	var ids [3]uint
	for i := range ids {
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}

	imgs, err := s.repo.List(ids[:])
	assert.NoError(s.T(), err)
	assert.Len(s.T(), imgs, len(ids))
}

func (s *ImageSuite) TestCreatePostImages() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p, _ := s.pr.Create(np, u.ID, nil)
	var ids [3]uint
	for i := range ids {
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}

	err := s.repo.CreatePostImages(ids[:], p.ID)
	assert.NoError(s.T(), err)
	imgs, _ := s.repo.ListPostImages(p.ID)
	assert.Len(s.T(), imgs, len(ids))
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
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}
	s.repo.CreatePostImages(ids[:], p.ID)

	imgs, err := s.repo.ListPostImages(p.ID)
	assert.NoError(s.T(), err)
	assert.Len(s.T(), imgs, len(ids))
}

func (s *ImageSuite) TestMarkUsed() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p, _ := s.pr.Create(np, u.ID, nil)
	var ids [3]uint
	for i := range ids {
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}
	s.repo.CreatePostImages(ids[:], p.ID)

	err := s.repo.MarkUsed(ids[:])
	assert.NoError(s.T(), err)
	imgs, _ := s.repo.ListPostImages(p.ID)
	for _, img := range imgs {
		assert.True(s.T(), img.InUse)
	}
}

func (s *ImageSuite) TestDeleteManyPostImages() {
	var au people.AuthUser
	var np people.NewPost
	gofakeit.Struct(&au)
	gofakeit.Struct(&np)
	u, _ := s.ur.Create(au)
	p, _ := s.pr.Create(np, u.ID, nil)
	var ids [3]uint
	for i := range ids {
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}
	s.repo.CreatePostImages(ids[:], p.ID)

	s.repo.DeleteManyPostImages(ids[:])
	imgs, _ := s.repo.ListPostImages(p.ID)
	assert.Empty(s.T(), imgs)
}

func (s *ImageSuite) TestListPostsImageIDs() {
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
		ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)
		ids[i] = ir.ID
	}
	for i := range ids {
		s.repo.CreatePostImages(ids[:], postIDs[i])
	}

	idsMap, err := s.repo.ListPostsImageIDs(postIDs[:])
	assert.NoError(s.T(), err)
	assert.Len(s.T(), idsMap, len(postIDs))
	for postID, imageIDs := range idsMap {
		assert.Equal(s.T(), ids[:], imageIDs)
		assert.Contains(s.T(), postIDs, postID)
	}
}

func (s *ImageSuite) SetupTest() {
	if s.fns.SetupTest != nil {
		s.fns.SetupTest()
	}
}

func (s *ImageSuite) TestCreateUserImage() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)

	err := s.repo.CreateUserImage(ir.ID, u.ID)
	assert.NoError(s.T(), err)
	img, _ := s.repo.GetUserImage(u.ID)
	assert.Equal(s.T(), ir.ID, img.ID)
}

func (s *ImageSuite) TestGetUserImage() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	ir, _ := s.repo.Create(gofakeit.LetterN(40), u.ID)

	s.repo.CreateUserImage(ir.ID, u.ID)
	img, err := s.repo.GetUserImage(u.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), ir.ID, img.ID)
}
