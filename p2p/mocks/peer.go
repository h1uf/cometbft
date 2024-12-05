// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	log "github.com/cometbft/cometbft/libs/log"
	conn "github.com/cometbft/cometbft/p2p/transport/tcp/conn"

	mock "github.com/stretchr/testify/mock"

	net "net"

	netaddr "github.com/cometbft/cometbft/p2p/netaddr"

	nodeinfo "github.com/cometbft/cometbft/p2p/internal/nodeinfo"

	p2p "github.com/cometbft/cometbft/p2p"
)

// Peer is an autogenerated mock type for the Peer type
type Peer struct {
	mock.Mock
}

// Conn provides a mock function with no fields
func (_m *Peer) Conn() net.Conn {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Conn")
	}

	var r0 net.Conn
	if rf, ok := ret.Get(0).(func() net.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	return r0
}

// FlushStop provides a mock function with no fields
func (_m *Peer) FlushStop() {
	_m.Called()
}

// Get provides a mock function with given fields: key
func (_m *Peer) Get(key string) any {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 any
	if rf, ok := ret.Get(0).(func(string) any); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	return r0
}

// GetRemovalFailed provides a mock function with no fields
func (_m *Peer) GetRemovalFailed() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRemovalFailed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasChannel provides a mock function with given fields: chID
func (_m *Peer) HasChannel(chID byte) bool {
	ret := _m.Called(chID)

	if len(ret) == 0 {
		panic("no return value specified for HasChannel")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(byte) bool); ok {
		r0 = rf(chID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ID provides a mock function with no fields
func (_m *Peer) ID() p2p.ID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 p2p.ID
	if rf, ok := ret.Get(0).(func() p2p.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(p2p.ID)
	}

	return r0
}

// IsOutbound provides a mock function with no fields
func (_m *Peer) IsOutbound() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsOutbound")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsPersistent provides a mock function with no fields
func (_m *Peer) IsPersistent() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsPersistent")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRunning provides a mock function with no fields
func (_m *Peer) IsRunning() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRunning")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NodeInfo provides a mock function with no fields
func (_m *Peer) NodeInfo() nodeinfo.NodeInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NodeInfo")
	}

	var r0 nodeinfo.NodeInfo
	if rf, ok := ret.Get(0).(func() nodeinfo.NodeInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nodeinfo.NodeInfo)
		}
	}

	return r0
}

// OnReset provides a mock function with no fields
func (_m *Peer) OnReset() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OnReset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStart provides a mock function with no fields
func (_m *Peer) OnStart() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for OnStart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnStop provides a mock function with no fields
func (_m *Peer) OnStop() {
	_m.Called()
}

// Quit provides a mock function with no fields
func (_m *Peer) Quit() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Quit")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// RemoteAddr provides a mock function with no fields
func (_m *Peer) RemoteAddr() net.Addr {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteAddr")
	}

	var r0 net.Addr
	if rf, ok := ret.Get(0).(func() net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Addr)
		}
	}

	return r0
}

// RemoteIP provides a mock function with no fields
func (_m *Peer) RemoteIP() net.IP {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteIP")
	}

	var r0 net.IP
	if rf, ok := ret.Get(0).(func() net.IP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.IP)
		}
	}

	return r0
}

// Reset provides a mock function with no fields
func (_m *Peer) Reset() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Reset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: e
func (_m *Peer) Send(e p2p.Envelope) bool {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Envelope) bool); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Set provides a mock function with given fields: key, value
func (_m *Peer) Set(key string, value any) {
	_m.Called(key, value)
}

// SetLogger provides a mock function with given fields: l
func (_m *Peer) SetLogger(l log.Logger) {
	_m.Called(l)
}

// SetRemovalFailed provides a mock function with no fields
func (_m *Peer) SetRemovalFailed() {
	_m.Called()
}

// SocketAddr provides a mock function with no fields
func (_m *Peer) SocketAddr() *netaddr.NetAddr {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SocketAddr")
	}

	var r0 *netaddr.NetAddr
	if rf, ok := ret.Get(0).(func() *netaddr.NetAddr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*netaddr.NetAddr)
		}
	}

	return r0
}

// Start provides a mock function with no fields
func (_m *Peer) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with no fields
func (_m *Peer) Status() conn.ConnectionStatus {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 conn.ConnectionStatus
	if rf, ok := ret.Get(0).(func() conn.ConnectionStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(conn.ConnectionStatus)
	}

	return r0
}

// Stop provides a mock function with no fields
func (_m *Peer) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with no fields
func (_m *Peer) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TrySend provides a mock function with given fields: e
func (_m *Peer) TrySend(e p2p.Envelope) bool {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for TrySend")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Envelope) bool); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewPeer creates a new instance of Peer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPeer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Peer {
	mock := &Peer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
