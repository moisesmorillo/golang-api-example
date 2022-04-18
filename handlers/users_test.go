package handlers

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/moisesmorillo/golang-api-example/db/dtos"
	"github.com/moisesmorillo/golang-api-example/interfaces"
	"github.com/moisesmorillo/golang-api-example/mocks"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

var requestCtx = context.Background()

type UserHandlerTestSuite struct {
	suite.Suite
	service   *mocks.UsersService
	underTest interfaces.UsersHandler
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.service = &mocks.UsersService{}
	s.underTest = NewUsersHandler(s.service)
}

func (s *UserHandlerTestSuite) TestGet_WhenServiceFails() {
	s.service.Mock.On("Get", requestCtx).Return(nil, errors.New("service error"))

	c := SetupControllerCase(http.MethodGet, "/users/", nil)
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.Error(s.underTest.Get(c.context))
}

func (s *UserHandlerTestSuite) TestGet_WhenOk() {
	s.service.Mock.On("Get", requestCtx).Return(&[]dtos.Users{
		{Name: "user"},
	}, nil)

	c := SetupControllerCase(http.MethodGet, "/users/", nil)
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(s.underTest.Get(c.context))
}

func (s *UserHandlerTestSuite) TestCreate_WhenBindMethodFails() {
	userJsonString := `this is not a json and should fail`

	c := SetupControllerCase(http.MethodPost, "/users/", strings.NewReader(userJsonString))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.Error(s.underTest.Create(c.context))
}

func (s *UserHandlerTestSuite) TestCreate_WhenValidateMethodFails() {
	userJsonString := `{}`

	c := SetupControllerCase(http.MethodPost, "/users/", strings.NewReader(userJsonString))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.Error(s.underTest.Create(c.context))
}

func (s *UserHandlerTestSuite) TestCreate_WhenServiceFails() {
	userJsonString := `{"name": "Random"}`

	s.service.Mock.On("Create", requestCtx, &dtos.Users{Name: "Random"}).Return(errors.New("some error"))
	c := SetupControllerCase(http.MethodPost, "/users/", strings.NewReader(userJsonString))
	c.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.Error(s.underTest.Create(c.context))
}
