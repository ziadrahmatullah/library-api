// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	mock "github.com/stretchr/testify/mock"

	valueobject "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: book
func (_m *BookRepository) CreateBook(book *entity.Book) (*entity.Book, error) {
	ret := _m.Called(book)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(*entity.Book) *entity.Book); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: clause, conditions
func (_m *BookRepository) FindAll(clause valueobject.Clause, conditions []valueobject.Condition) []*entity.Book {
	ret := _m.Called(clause, conditions)

	var r0 []*entity.Book
	if rf, ok := ret.Get(0).(func(valueobject.Clause, []valueobject.Condition) []*entity.Book); ok {
		r0 = rf(clause, conditions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Book)
		}
	}

	return r0
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
