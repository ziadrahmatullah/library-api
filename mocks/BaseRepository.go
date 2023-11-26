// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	valueobject "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

// BaseRepository is an autogenerated mock type for the BaseRepository type
type BaseRepository[T interface{}] struct {
	mock.Mock
}

// Create provides a mock function with given fields: t
func (_m *BaseRepository[T]) Create(t *T) (*T, error) {
	ret := _m.Called(t)

	var r0 *T
	if rf, ok := ret.Get(0).(func(*T) *T); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*T) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: query
func (_m *BaseRepository[T]) Find(query valueobject.Query) []*T {
	ret := _m.Called(query)

	var r0 []*T
	if rf, ok := ret.Get(0).(func(valueobject.Query) []*T); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*T)
		}
	}

	return r0
}

// First provides a mock function with given fields: query
func (_m *BaseRepository[T]) First(query valueobject.Query) *T {
	ret := _m.Called(query)

	var r0 *T
	if rf, ok := ret.Get(0).(func(valueobject.Query) *T); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	return r0
}

type mockConstructorTestingTNewBaseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBaseRepository creates a new instance of BaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBaseRepository[T interface{}](t mockConstructorTestingTNewBaseRepository) *BaseRepository[T] {
	mock := &BaseRepository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}