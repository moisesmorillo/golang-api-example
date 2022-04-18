package services

import (
	"context"
	"errors"
	"github.com/moisesmorillo/golang-api-example/db/dtos"
	"github.com/moisesmorillo/golang-api-example/db/models"
	"github.com/moisesmorillo/golang-api-example/interfaces"
	"github.com/moisesmorillo/golang-api-example/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	ctx        = context.Background()
	er         = errors.New("error")
	user       = models.Users{Name: "a", ID: 1}
	users      = []models.Users{user}
	userCreate = &dtos.Users{Name: "a"}
)

type UsersServiceTestSuite struct {
	suite.Suite
	repo      *mocks.UserRepository
	underTest interfaces.UsersService
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UsersServiceTestSuite))
}

func (s *UsersServiceTestSuite) SetupTest() {
	s.repo = &mocks.UserRepository{}
	s.underTest = NewUsersService(s.repo)
}

func (s *UsersServiceTestSuite) TestGet_WhenRepoFails() {
	s.repo.Mock.On("Get", ctx).Return(nil, er)
	res, err := s.underTest.Get(ctx)
	s.Error(err)
	s.Nil(res)
}

func (s *UsersServiceTestSuite) TestGet_WhenSuccess() {
	s.repo.Mock.On("Get", ctx).Return(&users, nil)
	_, err := s.underTest.Get(ctx)
	s.NoError(err)
}

func (s *UsersServiceTestSuite) TestCreate_WhenRepoFails() {
	s.repo.Mock.On("Create", ctx, userCreate).Return(er)
	s.Error(s.underTest.Create(ctx, userCreate))
}

func (s *UsersServiceTestSuite) TestCreate_WhenSuccess() {
	s.repo.Mock.On("Create", ctx, userCreate).Return(nil)
	s.NoError(s.underTest.Create(ctx, userCreate))
}
