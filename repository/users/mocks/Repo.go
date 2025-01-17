// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	entity "coffee-online-cli/entity"

	mock "github.com/stretchr/testify/mock"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// CheckEmailExists provides a mock function with given fields: email
func (_m *Repo) CheckEmailExists(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for CheckEmailExists")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: user
func (_m *Repo) CreateUser(user entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditUser provides a mock function with given fields: user
func (_m *Repo) EditUser(user entity.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for EditUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *Repo) GetUserByEmail(email string) (*entity.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: id
func (_m *Repo) GetUserByID(id int) (*entity.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*entity.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *entity.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoyalCustomer provides a mock function with given fields:
func (_m *Repo) LoyalCustomer() ([]entity.UserLoyal, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for LoyalCustomer")
	}

	var r0 []entity.UserLoyal
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.UserLoyal, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.UserLoyal); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.UserLoyal)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
