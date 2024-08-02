package handler

import (
	"coffee-online-cli/entity"
	MOCK_USERS "coffee-online-cli/repository/users/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestHandler_RegisterUsersSuccess(t *testing.T) {
	user := entity.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "123456",
		Location: "New York",
		Role:     entity.Role{ID: 1},
	}
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("CreateUser", mock.AnythingOfType("entity.User")).Return(nil)
	err := usersRepo.CreateUser(user)
	assert.Nil(t, err)
}

func TestHandler_RegisterUsersFail(t *testing.T) {
	user := entity.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "123456",
		Location: "New York",
		Role:     entity.Role{ID: 1},
	}
	errMock := errors.New("error creating user")
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("CreateUser", mock.AnythingOfType("entity.User")).Return(errMock)
	err := usersRepo.CreateUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, errMock, err)
}

func TestHandler_LoginUsersSuccess(t *testing.T) {
	email := "johndoe@example.com"
	userMock := &entity.User{
		ID:       1,
		Name:     "John Doe",
		Password: "1234",
	}
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(userMock, nil)
	user, errGetUserByEmail := usersRepo.GetUserByEmail(email)
	assert.Nil(t, errGetUserByEmail)
	if !reflect.DeepEqual(user, userMock) {
		t.Errorf("GetUserByEmail() = %v, want %v", user, userMock)
	}

	loggedUserMock := &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    email,
		Location: "New York",
		Role:     entity.Role{ID: 1, Name: "Customer"},
	}
	usersRepo.On("GetUserByID", mock.AnythingOfType("int")).Return(loggedUserMock, nil)
	loggedUser, errGetUserByID := usersRepo.GetUserByID(user.ID)
	assert.Nil(t, errGetUserByID)
	if !reflect.DeepEqual(loggedUser, loggedUserMock) {
		t.Errorf("GetUserByID() = %v, want %v", loggedUser, loggedUserMock)
	}
}

func TestHandler_LoginUsersFail(t *testing.T) {
	email := "johndoe@example.com"
	errGetUserByEmailMock := errors.New("error getting user by email")
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("GetUserByEmail", mock.AnythingOfType("string")).Return(nil, errGetUserByEmailMock)
	user, errGetUserByEmail := usersRepo.GetUserByEmail(email)
	assert.Nil(t, user)
	assert.Equal(t, errGetUserByEmail, errGetUserByEmailMock)
}

func TestHandler_ReportLoyalSuccess(t *testing.T) {
	loyalsMock := []entity.UserLoyal{
		{
			Name:          "John Doe",
			TotalOrder:    10,
			TotalSpending: 500000,
		},
	}
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("LoyalCustomer").Return(loyalsMock, nil)
	loyals, err := usersRepo.LoyalCustomer()
	assert.Nil(t, err)
	if !reflect.DeepEqual(loyals, loyalsMock) {
		t.Errorf("LoyalCustomer() = %v, want %v", loyals, loyalsMock)
	}
}

func TestHandler_ReportLoyalFail(t *testing.T) {
	errMock := errors.New("error loyal customer")
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("LoyalCustomer").Return(nil, errMock)
	loyals, err := usersRepo.LoyalCustomer()
	assert.Nil(t, loyals)
	assert.Equal(t, errMock, err)
}

func TestHandler_UserUpdateSuccess(t *testing.T) {
	user := entity.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "123456",
		Location: "New York",
	}
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("EditUser", mock.AnythingOfType("entity.User")).Return(nil)
	err := usersRepo.EditUser(user)
	assert.Nil(t, err)
}

func TestHandler_UserUpdateFail(t *testing.T) {
	user := entity.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "123456",
		Location: "New York",
	}
	errMock := errors.New("error updating user")
	usersRepo := new(MOCK_USERS.Repo)
	usersRepo.On("EditUser", mock.AnythingOfType("entity.User")).Return(errMock)
	err := usersRepo.EditUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, err, errMock)
}
