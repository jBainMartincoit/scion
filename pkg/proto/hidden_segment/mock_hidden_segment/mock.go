// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/pkg/proto/hidden_segment (interfaces: AuthoritativeHiddenSegmentLookupServiceServer,HiddenSegmentRegistrationServiceServer,HiddenSegmentLookupServiceServer)

// Package mock_hidden_segment is a generated GoMock package.
package mock_hidden_segment

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hidden_segment "github.com/scionproto/scion/pkg/proto/hidden_segment"
)

// MockAuthoritativeHiddenSegmentLookupServiceServer is a mock of AuthoritativeHiddenSegmentLookupServiceServer interface.
type MockAuthoritativeHiddenSegmentLookupServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthoritativeHiddenSegmentLookupServiceServerMockRecorder
}

// MockAuthoritativeHiddenSegmentLookupServiceServerMockRecorder is the mock recorder for MockAuthoritativeHiddenSegmentLookupServiceServer.
type MockAuthoritativeHiddenSegmentLookupServiceServerMockRecorder struct {
	mock *MockAuthoritativeHiddenSegmentLookupServiceServer
}

// NewMockAuthoritativeHiddenSegmentLookupServiceServer creates a new mock instance.
func NewMockAuthoritativeHiddenSegmentLookupServiceServer(ctrl *gomock.Controller) *MockAuthoritativeHiddenSegmentLookupServiceServer {
	mock := &MockAuthoritativeHiddenSegmentLookupServiceServer{ctrl: ctrl}
	mock.recorder = &MockAuthoritativeHiddenSegmentLookupServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthoritativeHiddenSegmentLookupServiceServer) EXPECT() *MockAuthoritativeHiddenSegmentLookupServiceServerMockRecorder {
	return m.recorder
}

// AuthoritativeHiddenSegments mocks base method.
func (m *MockAuthoritativeHiddenSegmentLookupServiceServer) AuthoritativeHiddenSegments(arg0 context.Context, arg1 *hidden_segment.AuthoritativeHiddenSegmentsRequest) (*hidden_segment.AuthoritativeHiddenSegmentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthoritativeHiddenSegments", arg0, arg1)
	ret0, _ := ret[0].(*hidden_segment.AuthoritativeHiddenSegmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthoritativeHiddenSegments indicates an expected call of AuthoritativeHiddenSegments.
func (mr *MockAuthoritativeHiddenSegmentLookupServiceServerMockRecorder) AuthoritativeHiddenSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthoritativeHiddenSegments", reflect.TypeOf((*MockAuthoritativeHiddenSegmentLookupServiceServer)(nil).AuthoritativeHiddenSegments), arg0, arg1)
}

// MockHiddenSegmentRegistrationServiceServer is a mock of HiddenSegmentRegistrationServiceServer interface.
type MockHiddenSegmentRegistrationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockHiddenSegmentRegistrationServiceServerMockRecorder
}

// MockHiddenSegmentRegistrationServiceServerMockRecorder is the mock recorder for MockHiddenSegmentRegistrationServiceServer.
type MockHiddenSegmentRegistrationServiceServerMockRecorder struct {
	mock *MockHiddenSegmentRegistrationServiceServer
}

// NewMockHiddenSegmentRegistrationServiceServer creates a new mock instance.
func NewMockHiddenSegmentRegistrationServiceServer(ctrl *gomock.Controller) *MockHiddenSegmentRegistrationServiceServer {
	mock := &MockHiddenSegmentRegistrationServiceServer{ctrl: ctrl}
	mock.recorder = &MockHiddenSegmentRegistrationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHiddenSegmentRegistrationServiceServer) EXPECT() *MockHiddenSegmentRegistrationServiceServerMockRecorder {
	return m.recorder
}

// HiddenSegmentRegistration mocks base method.
func (m *MockHiddenSegmentRegistrationServiceServer) HiddenSegmentRegistration(arg0 context.Context, arg1 *hidden_segment.HiddenSegmentRegistrationRequest) (*hidden_segment.HiddenSegmentRegistrationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HiddenSegmentRegistration", arg0, arg1)
	ret0, _ := ret[0].(*hidden_segment.HiddenSegmentRegistrationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HiddenSegmentRegistration indicates an expected call of HiddenSegmentRegistration.
func (mr *MockHiddenSegmentRegistrationServiceServerMockRecorder) HiddenSegmentRegistration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HiddenSegmentRegistration", reflect.TypeOf((*MockHiddenSegmentRegistrationServiceServer)(nil).HiddenSegmentRegistration), arg0, arg1)
}

// MockHiddenSegmentLookupServiceServer is a mock of HiddenSegmentLookupServiceServer interface.
type MockHiddenSegmentLookupServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockHiddenSegmentLookupServiceServerMockRecorder
}

// MockHiddenSegmentLookupServiceServerMockRecorder is the mock recorder for MockHiddenSegmentLookupServiceServer.
type MockHiddenSegmentLookupServiceServerMockRecorder struct {
	mock *MockHiddenSegmentLookupServiceServer
}

// NewMockHiddenSegmentLookupServiceServer creates a new mock instance.
func NewMockHiddenSegmentLookupServiceServer(ctrl *gomock.Controller) *MockHiddenSegmentLookupServiceServer {
	mock := &MockHiddenSegmentLookupServiceServer{ctrl: ctrl}
	mock.recorder = &MockHiddenSegmentLookupServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHiddenSegmentLookupServiceServer) EXPECT() *MockHiddenSegmentLookupServiceServerMockRecorder {
	return m.recorder
}

// HiddenSegments mocks base method.
func (m *MockHiddenSegmentLookupServiceServer) HiddenSegments(arg0 context.Context, arg1 *hidden_segment.HiddenSegmentsRequest) (*hidden_segment.HiddenSegmentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HiddenSegments", arg0, arg1)
	ret0, _ := ret[0].(*hidden_segment.HiddenSegmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HiddenSegments indicates an expected call of HiddenSegments.
func (mr *MockHiddenSegmentLookupServiceServerMockRecorder) HiddenSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HiddenSegments", reflect.TypeOf((*MockHiddenSegmentLookupServiceServer)(nil).HiddenSegments), arg0, arg1)
}