// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	entity "coffee-online-cli/entity"

	mock "github.com/stretchr/testify/mock"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// FetchProducts provides a mock function with given fields:
func (_m *Reader) FetchProducts() ([]entity.Product, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FetchProducts")
	}

	var r0 []entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByID provides a mock function with given fields: id
func (_m *Reader) GetProductByID(id int) (*entity.Product, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetProductByID")
	}

	var r0 *entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*entity.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *entity.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PopularProduct provides a mock function with given fields:
func (_m *Reader) PopularProduct() ([]entity.ProductPopular, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PopularProduct")
	}

	var r0 []entity.ProductPopular
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.ProductPopular, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.ProductPopular); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.ProductPopular)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}