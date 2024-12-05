// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	crypto "github.com/cometbft/cometbft/crypto"
	mock "github.com/stretchr/testify/mock"
)

// PubKey is an autogenerated mock type for the PubKey type
type PubKey struct {
	mock.Mock
}

// Address provides a mock function with no fields
func (_m *PubKey) Address() crypto.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 crypto.Address
	if rf, ok := ret.Get(0).(func() crypto.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.Address)
		}
	}

	return r0
}

// Bytes provides a mock function with no fields
func (_m *PubKey) Bytes() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Bytes")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Type provides a mock function with no fields
func (_m *PubKey) Type() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VerifySignature provides a mock function with given fields: msg, sig
func (_m *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	ret := _m.Called(msg, sig)

	if len(ret) == 0 {
		panic("no return value specified for VerifySignature")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte, []byte) bool); ok {
		r0 = rf(msg, sig)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewPubKey creates a new instance of PubKey. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPubKey(t interface {
	mock.TestingT
	Cleanup(func())
}) *PubKey {
	mock := &PubKey{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
