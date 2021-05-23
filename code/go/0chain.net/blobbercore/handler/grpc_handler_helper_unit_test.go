package handler

import (
	"context"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/allocation"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/reference"
	"github.com/stretchr/testify/mock"
)

// StorageHandlerI is an autogenerated mock type for the StorageHandlerI type
type storageHandlerI struct {
	mock.Mock
}

func (_m *storageHandlerI) readPreRedeem(ctx context.Context, alloc *allocation.Allocation, numBlocks, pendNumBlocks int64, payerID string) (err error) {
	ret := _m.Called(ctx, alloc, numBlocks, pendNumBlocks, payerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *allocation.Allocation, int64, int64, string) error); ok {
		r0 = rf(ctx, alloc, numBlocks, pendNumBlocks, payerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// verifyAllocation provides a mock function with given fields: ctx, tx, readonly
func (_m *storageHandlerI) verifyAllocation(ctx context.Context, tx string, readonly bool) (*allocation.Allocation, error) {
	ret := _m.Called(ctx, tx, readonly)

	var r0 *allocation.Allocation
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *allocation.Allocation); ok {
		r0 = rf(ctx, tx, readonly)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*allocation.Allocation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, tx, readonly)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// verifyAuthTicket provides a mock function with given fields: ctx, authTokenString, allocationObj, refRequested, clientID
func (_m *storageHandlerI) verifyAuthTicket(ctx context.Context, authTokenString string, allocationObj *allocation.Allocation, refRequested *reference.Ref, clientID string) (bool, error) {
	ret := _m.Called(ctx, authTokenString, allocationObj, refRequested, clientID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, *allocation.Allocation, *reference.Ref, string) bool); ok {
		r0 = rf(ctx, authTokenString, allocationObj, refRequested, clientID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *allocation.Allocation, *reference.Ref, string) error); ok {
		r1 = rf(ctx, authTokenString, allocationObj, refRequested, clientID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
