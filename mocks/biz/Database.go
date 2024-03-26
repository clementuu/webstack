// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "webstack/models"

	mock "github.com/stretchr/testify/mock"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// AddTodo provides a mock function with given fields: td
func (_m *Database) AddTodo(td models.Todo) error {
	ret := _m.Called(td)

	if len(ret) == 0 {
		panic("no return value specified for AddTodo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Todo) error); ok {
		r0 = rf(td)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTodo provides a mock function with given fields: td
func (_m *Database) DeleteTodo(td models.Todo) error {
	ret := _m.Called(td)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTodo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Todo) error); ok {
		r0 = rf(td)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTodos provides a mock function with given fields:
func (_m *Database) GetTodos() ([]models.Todo, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTodos")
	}

	var r0 []models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyTodo provides a mock function with given fields: td
func (_m *Database) ModifyTodo(td models.Todo) error {
	ret := _m.Called(td)

	if len(ret) == 0 {
		panic("no return value specified for ModifyTodo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Todo) error); ok {
		r0 = rf(td)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}