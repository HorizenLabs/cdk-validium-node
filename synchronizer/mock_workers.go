// Code generated by mockery v2.22.1. DO NOT EDIT.

package synchronizer

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// workersMock is an autogenerated mock type for the workersInterface type
type workersMock struct {
	mock.Mock
}

// String provides a mock function with given fields:
func (_m *workersMock) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// asyncRequestRollupInfoByBlockRange provides a mock function with given fields: ctx, request
func (_m *workersMock) asyncRequestRollupInfoByBlockRange(ctx context.Context, request requestRollupInfoByBlockRange) (chan responseRollupInfoByBlockRange, error) {
	ret := _m.Called(ctx, request)

	var r0 chan responseRollupInfoByBlockRange
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, requestRollupInfoByBlockRange) (chan responseRollupInfoByBlockRange, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, requestRollupInfoByBlockRange) chan responseRollupInfoByBlockRange); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan responseRollupInfoByBlockRange)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, requestRollupInfoByBlockRange) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getResponseChannelForRollupInfo provides a mock function with given fields:
func (_m *workersMock) getResponseChannelForRollupInfo() chan responseRollupInfoByBlockRange {
	ret := _m.Called()

	var r0 chan responseRollupInfoByBlockRange
	if rf, ok := ret.Get(0).(func() chan responseRollupInfoByBlockRange); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan responseRollupInfoByBlockRange)
		}
	}

	return r0
}

// howManyRunningWorkers provides a mock function with given fields:
func (_m *workersMock) howManyRunningWorkers() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// initialize provides a mock function with given fields:
func (_m *workersMock) initialize() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// requestLastBlockWithRetries provides a mock function with given fields: ctx, timeout, maxPermittedRetries
func (_m *workersMock) requestLastBlockWithRetries(ctx context.Context, timeout time.Duration, maxPermittedRetries int) responseL1LastBlock {
	ret := _m.Called(ctx, timeout, maxPermittedRetries)

	var r0 responseL1LastBlock
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, int) responseL1LastBlock); ok {
		r0 = rf(ctx, timeout, maxPermittedRetries)
	} else {
		r0 = ret.Get(0).(responseL1LastBlock)
	}

	return r0
}

// stop provides a mock function with given fields:
func (_m *workersMock) stop() {
	_m.Called()
}

// waitFinishAllWorkers provides a mock function with given fields:
func (_m *workersMock) waitFinishAllWorkers() {
	_m.Called()
}

type mockConstructorTestingTnewWorkersMock interface {
	mock.TestingT
	Cleanup(func())
}

// newWorkersMock creates a new instance of workersMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newWorkersMock(t mockConstructorTestingTnewWorkersMock) *workersMock {
	mock := &workersMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
