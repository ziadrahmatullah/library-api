// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/-/tree/ziad-rahmatullah/models"
	mock "github.com/stretchr/testify/mock"
)

// BorrowRepository is an autogenerated mock type for the BorrowRepository type
type BorrowRepository struct {
	mock.Mock
}

// FindBorrow provides a mock function with given fields: _a0, _a1
func (_m *BorrowRepository) FindBorrow(_a0 context.Context, _a1 models.BorrowBook) (uint, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, models.BorrowBook) uint); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.BorrowBook) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBorrows provides a mock function with given fields: _a0
func (_m *BorrowRepository) FindBorrows(_a0 context.Context) ([]models.BorrowBook, error) {
	ret := _m.Called(_a0)

	var r0 []models.BorrowBook
	if rf, ok := ret.Get(0).(func(context.Context) []models.BorrowBook); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.BorrowBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBorrow provides a mock function with given fields: _a0, _a1
func (_m *BorrowRepository) NewBorrow(_a0 context.Context, _a1 models.BorrowBook) (*models.BorrowBook, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.BorrowBook
	if rf, ok := ret.Get(0).(func(context.Context, models.BorrowBook) *models.BorrowBook); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BorrowBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.BorrowBook) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBorrowStatus provides a mock function with given fields: _a0, _a1
func (_m *BorrowRepository) UpdateBorrowStatus(_a0 context.Context, _a1 uint) (*models.BorrowBook, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.BorrowBook
	if rf, ok := ret.Get(0).(func(context.Context, uint) *models.BorrowBook); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.BorrowBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBorrowRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBorrowRepository creates a new instance of BorrowRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBorrowRepository(t mockConstructorTestingTNewBorrowRepository) *BorrowRepository {
	mock := &BorrowRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}