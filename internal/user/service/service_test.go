package service

import (
	"context"
	"testing"
	"time"

	"github.com/Slamadalius/faceit/internal/entity"
	"github.com/Slamadalius/faceit/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

var (
	contextTimeout = time.Second * 5
	user           = entity.User{
		ID:        "test",
		FirstName: "test",
		LastName:  "test",
		Nickname:  "test",
		Password:  "test",
		Email:     "test@test.com",
		Country:   "test",
	}
)

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

type UserServiceTestSuite struct {
	suite.Suite
	userRepository *mocks.MockUserRepository
	underTest      entity.UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()

	suite.userRepository = mocks.NewMockUserRepository(mockCtrl)
	suite.underTest = NewUserService(suite.userRepository)
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	suite.userRepository.EXPECT().Insert(context.Background(), user).Return(nil)

	err := suite.underTest.CreateUser(context.Background(), user)

	suite.NoError(err, "Should be no error")
}

func (suite *UserServiceTestSuite) TestUpdateUser() {
	suite.userRepository.EXPECT().Update(context.Background(), gomock.Any(), user).Return(nil)

	err := suite.underTest.UpdateUser(context.Background(), gomock.Any().String(), user)

	suite.NoError(err, "Should be no error")
}

func (suite *UserServiceTestSuite) TestDeleteUser() {
	suite.userRepository.EXPECT().Delete(context.Background(), gomock.Any()).Return(nil)

	err := suite.underTest.DeleteUser(context.Background(), gomock.Any().String())

	suite.NoError(err, "Should be no error")
}

func (suite *UserServiceTestSuite) TestFindUsers() {
	suite.userRepository.EXPECT().FindAll(context.Background(), gomock.Any(), gomock.Any()).Return([]entity.User{user}, nil)

	users, err := suite.underTest.FindUsers(context.Background(), map[string]string{}, 1)

	suite.NoError(err, "Should be no error")
	suite.Equal(user, users[0])
}
