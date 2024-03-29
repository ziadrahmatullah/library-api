// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	mock "github.com/stretchr/testify/mock"

	valueobject "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, t
func (_m *BookRepository) Create(ctx context.Context, t *entity.Book) (*entity.Book, error) {
	ret := _m.Called(ctx, t)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) *entity.Book); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Book) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, t
func (_m *BookRepository) Delete(ctx context.Context, t *entity.Book) error {
	ret := _m.Called(ctx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, query
func (_m *BookRepository) Find(ctx context.Context, query *valueobject.Query) ([]*entity.Book, error) {
	ret := _m.Called(ctx, query)

	var r0 []*entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.Query) []*entity.Book); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *valueobject.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// First provides a mock function with given fields: ctx, query
func (_m *BookRepository) First(ctx context.Context, query *valueobject.Query) (*entity.Book, error) {
	ret := _m.Called(ctx, query)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.Query) *entity.Book); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *valueobject.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, t
func (_m *BookRepository) Update(ctx context.Context, t *entity.Book) (*entity.Book, error) {
	ret := _m.Called(ctx, t)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Book) *entity.Book); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.Book) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookRepository creates a new instance of BookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookRepository(t mockConstructorTestingTNewBookRepository) *BookRepository {
	mock := &BookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
