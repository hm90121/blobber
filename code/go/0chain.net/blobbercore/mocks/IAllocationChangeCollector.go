// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	allocation "github.com/0chain/blobber/code/go/0chain.net/blobbercore/allocation"

	mock "github.com/stretchr/testify/mock"
)

// IAllocationChangeCollector is an autogenerated mock type for the IAllocationChangeCollector type
type IAllocationChangeCollector struct {
	mock.Mock
}

// AddChange provides a mock function with given fields: allocationChange, changeProcessor
func (_m *IAllocationChangeCollector) AddChange(allocationChange *allocation.AllocationChange, changeProcessor allocation.AllocationChangeProcessor) {
	_m.Called(allocationChange, changeProcessor)
}

// ApplyChanges provides a mock function with given fields: ctx, allocationRoot
func (_m *IAllocationChangeCollector) ApplyChanges(ctx context.Context, allocationRoot string) error {
	ret := _m.Called(ctx, allocationRoot)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, allocationRoot)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommitToFileStore provides a mock function with given fields: ctx
func (_m *IAllocationChangeCollector) CommitToFileStore(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ComputeProperties provides a mock function with given fields:
func (_m *IAllocationChangeCollector) ComputeProperties() {
	_m.Called()
}

// DeleteChanges provides a mock function with given fields: ctx
func (_m *IAllocationChangeCollector) DeleteChanges(ctx context.Context) {
	_m.Called(ctx)
}

// GetAllocationID provides a mock function with given fields:
func (_m *IAllocationChangeCollector) GetAllocationID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetConnectionID provides a mock function with given fields:
func (_m *IAllocationChangeCollector) GetConnectionID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Save provides a mock function with given fields: ctx
func (_m *IAllocationChangeCollector) Save(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TableName provides a mock function with given fields:
func (_m *IAllocationChangeCollector) TableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
