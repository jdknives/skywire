// Code generated by mockery v1.0.0. DO NOT EDIT.

package arclient

import cipher "github.com/skycoin/dmsg/cipher"
import context "context"
import mock "github.com/stretchr/testify/mock"
import pfilter "github.com/AudriusButkevicius/pfilter"

// MockAPIClient is an autogenerated mock type for the APIClient type
type MockAPIClient struct {
	mock.Mock
}

// BindSTCPR provides a mock function with given fields: ctx, port
func (_m *MockAPIClient) BindSTCPR(ctx context.Context, port string) error {
	ret := _m.Called(ctx, port)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BindSUDPH provides a mock function with given fields: filter
func (_m *MockAPIClient) BindSUDPH(filter *pfilter.PacketFilter) (<-chan RemoteVisor, error) {
	ret := _m.Called(filter)

	var r0 <-chan RemoteVisor
	if rf, ok := ret.Get(0).(func(*pfilter.PacketFilter) <-chan RemoteVisor); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan RemoteVisor)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pfilter.PacketFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *MockAPIClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Health provides a mock function with given fields: ctx
func (_m *MockAPIClient) Health(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Resolve provides a mock function with given fields: ctx, tType, pk
func (_m *MockAPIClient) Resolve(ctx context.Context, tType string, pk cipher.PubKey) (VisorData, error) {
	ret := _m.Called(ctx, tType, pk)

	var r0 VisorData
	if rf, ok := ret.Get(0).(func(context.Context, string, cipher.PubKey) VisorData); ok {
		r0 = rf(ctx, tType, pk)
	} else {
		r0 = ret.Get(0).(VisorData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, cipher.PubKey) error); ok {
		r1 = rf(ctx, tType, pk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
