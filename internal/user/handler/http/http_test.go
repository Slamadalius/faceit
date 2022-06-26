package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Slamadalius/faceit/internal/entity"
	"github.com/Slamadalius/faceit/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

var (
	user = entity.User{
		ID:        "test",
		FirstName: "test",
		LastName:  "test",
		Nickname:  "test",
		Password:  "test",
		Email:     "test@test.com",
		Country:   "test",
	}
)

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

type UserHandlerTestSuite struct {
	suite.Suite
	userService *mocks.MockUserService
	underTest   Handler
}

func (suite *UserHandlerTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.userService = mocks.NewMockUserService(mockCtrl)
	suite.underTest = Handler{
		Service: suite.userService,
	}
}

func (suite *UserHandlerTestSuite) TestFindUsers() {
	users := []entity.User{
		user,
	}

	request, _ := json.Marshal(map[string]string{})

	suite.userService.EXPECT().FindUsers(context.Background(), gomock.Any(), gomock.Any()).Return(users, nil)

	r, _ := http.NewRequest(http.MethodPost, "/user/find", bytes.NewBuffer(request))

	w := httptest.NewRecorder()
	suite.underTest.findUsers(w, r)

	response := w.Result()
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)
	result := []entity.User{}
	_ = json.NewDecoder(response.Body).Decode(&result)

	suite.Equal(len(result), 1)
}

func (suite *UserHandlerTestSuite) TestCreateUser() {
	request, _ := json.Marshal(map[string]string{})

	suite.userService.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(nil)

	r, _ := http.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer(request))

	w := httptest.NewRecorder()
	suite.underTest.createUser(w, r)

	response := w.Result()
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)
}

func (suite *UserHandlerTestSuite) TestUpdateUser() {
	request, _ := json.Marshal(map[string]string{})

	suite.userService.EXPECT().UpdateUser(context.Background(), gomock.Any(), gomock.Any()).Return(nil)

	r, _ := http.NewRequest(http.MethodPut, "/user/update", bytes.NewBuffer(request))

	w := httptest.NewRecorder()
	suite.underTest.updateUser(w, r)

	response := w.Result()
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)
}

func (suite *UserHandlerTestSuite) TestDeleteUser() {
	suite.userService.EXPECT().DeleteUser(context.Background(), gomock.Any()).Return(nil)

	r, _ := http.NewRequest(http.MethodDelete, "/user/delete", nil)

	w := httptest.NewRecorder()
	suite.underTest.deleteUser(w, r)

	response := w.Result()
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode)
}
