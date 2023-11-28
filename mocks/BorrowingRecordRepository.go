// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity"
	mock "github.com/stretchr/testify/mock"

	valueobject "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/valueobject"
)

// BorrowingRecordRepository is an autogenerated mock type for the BorrowingRecordRepository type
type BorrowingRecordRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, t
func (_m *BorrowingRecordRepository) Create(ctx context.Context, t *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	ret := _m.Called(ctx, t)

	var r0 *entity.BorrowingRecords
	if rf, ok := ret.Get(0).(func(context.Context, *entity.BorrowingRecords) *entity.BorrowingRecords); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowingRecords)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.BorrowingRecords) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, t
func (_m *BorrowingRecordRepository) Delete(ctx context.Context, t *entity.BorrowingRecords) error {
	ret := _m.Called(ctx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.BorrowingRecords) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, query
func (_m *BorrowingRecordRepository) Find(ctx context.Context, query *valueobject.Query) []*entity.BorrowingRecords {
	ret := _m.Called(ctx, query)

	var r0 []*entity.BorrowingRecords
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.Query) []*entity.BorrowingRecords); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.BorrowingRecords)
		}
	}

	return r0
}

// First provides a mock function with given fields: ctx, query
func (_m *BorrowingRecordRepository) First(ctx context.Context, query *valueobject.Query) *entity.BorrowingRecords {
	ret := _m.Called(ctx, query)

	var r0 *entity.BorrowingRecords
	if rf, ok := ret.Get(0).(func(context.Context, *valueobject.Query) *entity.BorrowingRecords); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowingRecords)
		}
	}

	return r0
}

// Run provides a mock function with given fields: ctx, runner
func (_m *BorrowingRecordRepository) Run(ctx context.Context, runner func(context.Context) error) error {
	ret := _m.Called(ctx, runner)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, runner)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, t
func (_m *BorrowingRecordRepository) Update(ctx context.Context, t *entity.BorrowingRecords) (*entity.BorrowingRecords, error) {
	ret := _m.Called(ctx, t)

	var r0 *entity.BorrowingRecords
	if rf, ok := ret.Get(0).(func(context.Context, *entity.BorrowingRecords) *entity.BorrowingRecords); ok {
		r0 = rf(ctx, t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.BorrowingRecords)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entity.BorrowingRecords) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBorrowingRecordRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBorrowingRecordRepository creates a new instance of BorrowingRecordRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBorrowingRecordRepository(t mockConstructorTestingTNewBorrowingRecordRepository) *BorrowingRecordRepository {
	mock := &BorrowingRecordRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}