package user

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/post"
)

type UserSuite struct {
	suite.Suite
	db          *sqlx.DB
	us          service
	ps          people.PostService
	user1       people.AuthUser
	user2       people.AuthUser
	user3       people.AuthUser
	user4       people.AuthUser
	unknownUser people.AuthUser
	id1         uint
	id2         uint
	id3         uint
	id4         uint
	post1       people.Post
	postBody1   people.PostBody
}

func (suite *UserSuite) TestCreate() {
	rows, _ := suite.db.Queryx("SELECT user_id, handle FROM user_profile WHERE handle = $1", suite.user1.Handle)
	for rows.Next() {
		var actual people.AuthUser
		rows.StructScan(&actual)
		assert.Equal(suite.T(), suite.id1, *actual.ID)
	}
}

func (suite *UserSuite) TestExists() {
	assert.True(suite.T(), suite.us.Exists(suite.user1.Handle))
	assert.False(suite.T(), suite.us.Exists(suite.unknownUser.Handle))
}

func (suite *UserSuite) TestDelete() {
	suite.us.Delete(suite.user1.Handle)
	assert.False(suite.T(), suite.us.Exists(suite.user1.Handle))
}

func (suite *UserSuite) TestGetAuthUser() {
	actual, err := suite.us.GetAuthUser(suite.user1.Handle)
	assert.Equal(suite.T(), suite.user1.Handle, actual.Handle)
	assert.NoError(suite.T(), err)

	_, err = suite.us.GetAuthUser(suite.unknownUser.Handle)
	assert.Error(suite.T(), err)
}

func (suite *UserSuite) TestGet() {
	actual, err := suite.us.Get(suite.user1.Handle, nil)
	assert.Equal(suite.T(), suite.user1.Handle, actual.Handle)
	assert.NoError(suite.T(), err)

	_, err = suite.us.GetAuthUser(suite.unknownUser.Handle)
	assert.Error(suite.T(), err)
}

func (suite *UserSuite) TestLiked() {
	suite.ps.Like(suite.post1.ID, suite.id1)
	suite.ps.Like(suite.post1.ID, suite.id2)
	suite.ps.Like(suite.post1.ID, suite.id3)

	actual, err := suite.us.Liked(suite.post1.ID, nil, people.NewPagination[string](nil, nil, nil))
	assert.Equal(suite.T(), len(actual.Data), 3)
	assert.NoError(suite.T(), err)
}

func (suite *UserSuite) SetupSuite() {
	db, err := people.PostgresConnect()
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
	suite.us = service{db}
	suite.ps = post.NewService(db)
}

func (suite *UserSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *UserSuite) SetupTest() {
	suite.db.MustExec("TRUNCATE user_profile CASCADE")
	suite.db.MustExec("TRUNCATE post CASCADE")
	gofakeit.Struct(&suite.user1)
	gofakeit.Struct(&suite.user2)
	gofakeit.Struct(&suite.user3)
	gofakeit.Struct(&suite.user4)
	gofakeit.Struct(&suite.unknownUser)
	suite.id1, _ = suite.us.Create(suite.user1)
	suite.id2, _ = suite.us.Create(suite.user2)
	suite.id3, _ = suite.us.Create(suite.user3)
	suite.id4, _ = suite.us.Create(suite.user4)
	gofakeit.Struct(&suite.postBody1)
	suite.post1, _ = suite.ps.Create(suite.id1, suite.postBody1)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
