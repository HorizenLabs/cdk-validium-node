// Code generated by mockery v2.22.1. DO NOT EDIT.

package synchronizer

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/0xPolygon/cdk-data-availability/types"
)

// dataCommitteeClientMock is an autogenerated mock type for the ClientInterface type
type dataCommitteeClientMock struct {
	mock.Mock
}

// GetOffChainData provides a mock function with given fields: ctx, hash
func (_m *dataCommitteeClientMock) GetOffChainData(ctx context.Context, hash common.Hash) ([]byte, error) {
	ret := _m.Called(ctx, hash)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) ([]byte, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) []byte); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignSequence provides a mock function with given fields: signedSequence
func (_m *dataCommitteeClientMock) SignSequence(signedSequence types.SignedSequence) ([]byte, error) {
	ret := _m.Called(signedSequence)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(types.SignedSequence) ([]byte, error)); ok {
		return rf(signedSequence)
	}
	if rf, ok := ret.Get(0).(func(types.SignedSequence) []byte); ok {
		r0 = rf(signedSequence)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(types.SignedSequence) error); ok {
		r1 = rf(signedSequence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewDataCommitteeClientMock interface {
	mock.TestingT
	Cleanup(func())
}

// newDataCommitteeClientMock creates a new instance of dataCommitteeClientMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newDataCommitteeClientMock(t mockConstructorTestingTnewDataCommitteeClientMock) *dataCommitteeClientMock {
	mock := &dataCommitteeClientMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
