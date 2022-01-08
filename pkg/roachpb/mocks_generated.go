// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cockroachdb/cockroach/pkg/roachpb (interfaces: InternalClient,Internal_RangeFeedClient)

// Package roachpb is a generated GoMock package.
package roachpb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockInternalClient is a mock of InternalClient interface.
type MockInternalClient struct {
	ctrl     *gomock.Controller
	recorder *MockInternalClientMockRecorder
}

// MockInternalClientMockRecorder is the mock recorder for MockInternalClient.
type MockInternalClientMockRecorder struct {
	mock *MockInternalClient
}

// NewMockInternalClient creates a new mock instance.
func NewMockInternalClient(ctrl *gomock.Controller) *MockInternalClient {
	mock := &MockInternalClient{ctrl: ctrl}
	mock.recorder = &MockInternalClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInternalClient) EXPECT() *MockInternalClientMockRecorder {
	return m.recorder
}

// Batch mocks base method.
func (m *MockInternalClient) Batch(arg0 context.Context, arg1 *BatchRequest, arg2 ...grpc.CallOption) (*BatchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Batch", varargs...)
	ret0, _ := ret[0].(*BatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batch indicates an expected call of Batch.
func (mr *MockInternalClientMockRecorder) Batch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockInternalClient)(nil).Batch), varargs...)
}

// GetSpanConfigs mocks base method.
func (m *MockInternalClient) GetSpanConfigs(arg0 context.Context, arg1 *GetSpanConfigsRequest, arg2 ...grpc.CallOption) (*GetSpanConfigsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSpanConfigs", varargs...)
	ret0, _ := ret[0].(*GetSpanConfigsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpanConfigs indicates an expected call of GetSpanConfigs.
func (mr *MockInternalClientMockRecorder) GetSpanConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpanConfigs", reflect.TypeOf((*MockInternalClient)(nil).GetSpanConfigs), varargs...)
}

// GossipSubscription mocks base method.
func (m *MockInternalClient) GossipSubscription(arg0 context.Context, arg1 *GossipSubscriptionRequest, arg2 ...grpc.CallOption) (Internal_GossipSubscriptionClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GossipSubscription", varargs...)
	ret0, _ := ret[0].(Internal_GossipSubscriptionClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GossipSubscription indicates an expected call of GossipSubscription.
func (mr *MockInternalClientMockRecorder) GossipSubscription(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GossipSubscription", reflect.TypeOf((*MockInternalClient)(nil).GossipSubscription), varargs...)
}

// Join mocks base method.
func (m *MockInternalClient) Join(arg0 context.Context, arg1 *JoinNodeRequest, arg2 ...grpc.CallOption) (*JoinNodeResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Join", varargs...)
	ret0, _ := ret[0].(*JoinNodeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Join indicates an expected call of Join.
func (mr *MockInternalClientMockRecorder) Join(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockInternalClient)(nil).Join), varargs...)
}

// RangeFeed mocks base method.
func (m *MockInternalClient) RangeFeed(arg0 context.Context, arg1 *RangeFeedRequest, arg2 ...grpc.CallOption) (Internal_RangeFeedClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RangeFeed", varargs...)
	ret0, _ := ret[0].(Internal_RangeFeedClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RangeFeed indicates an expected call of RangeFeed.
func (mr *MockInternalClientMockRecorder) RangeFeed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeFeed", reflect.TypeOf((*MockInternalClient)(nil).RangeFeed), varargs...)
}

// RangeLookup mocks base method.
func (m *MockInternalClient) RangeLookup(arg0 context.Context, arg1 *RangeLookupRequest, arg2 ...grpc.CallOption) (*RangeLookupResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RangeLookup", varargs...)
	ret0, _ := ret[0].(*RangeLookupResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RangeLookup indicates an expected call of RangeLookup.
func (mr *MockInternalClientMockRecorder) RangeLookup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeLookup", reflect.TypeOf((*MockInternalClient)(nil).RangeLookup), varargs...)
}

// ResetQuorum mocks base method.
func (m *MockInternalClient) ResetQuorum(arg0 context.Context, arg1 *ResetQuorumRequest, arg2 ...grpc.CallOption) (*ResetQuorumResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetQuorum", varargs...)
	ret0, _ := ret[0].(*ResetQuorumResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetQuorum indicates an expected call of ResetQuorum.
func (mr *MockInternalClientMockRecorder) ResetQuorum(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetQuorum", reflect.TypeOf((*MockInternalClient)(nil).ResetQuorum), varargs...)
}

// TenantSettings mocks base method.
func (m *MockInternalClient) TenantSettings(arg0 context.Context, arg1 *TenantSettingsRequest, arg2 ...grpc.CallOption) (Internal_TenantSettingsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TenantSettings", varargs...)
	ret0, _ := ret[0].(Internal_TenantSettingsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TenantSettings indicates an expected call of TenantSettings.
func (mr *MockInternalClientMockRecorder) TenantSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantSettings", reflect.TypeOf((*MockInternalClient)(nil).TenantSettings), varargs...)
}

// TokenBucket mocks base method.
func (m *MockInternalClient) TokenBucket(arg0 context.Context, arg1 *TokenBucketRequest, arg2 ...grpc.CallOption) (*TokenBucketResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TokenBucket", varargs...)
	ret0, _ := ret[0].(*TokenBucketResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenBucket indicates an expected call of TokenBucket.
func (mr *MockInternalClientMockRecorder) TokenBucket(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenBucket", reflect.TypeOf((*MockInternalClient)(nil).TokenBucket), varargs...)
}

// UpdateSpanConfigs mocks base method.
func (m *MockInternalClient) UpdateSpanConfigs(arg0 context.Context, arg1 *UpdateSpanConfigsRequest, arg2 ...grpc.CallOption) (*UpdateSpanConfigsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSpanConfigs", varargs...)
	ret0, _ := ret[0].(*UpdateSpanConfigsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSpanConfigs indicates an expected call of UpdateSpanConfigs.
func (mr *MockInternalClientMockRecorder) UpdateSpanConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSpanConfigs", reflect.TypeOf((*MockInternalClient)(nil).UpdateSpanConfigs), varargs...)
}

// MockInternal_RangeFeedClient is a mock of Internal_RangeFeedClient interface.
type MockInternal_RangeFeedClient struct {
	ctrl     *gomock.Controller
	recorder *MockInternal_RangeFeedClientMockRecorder
}

// MockInternal_RangeFeedClientMockRecorder is the mock recorder for MockInternal_RangeFeedClient.
type MockInternal_RangeFeedClientMockRecorder struct {
	mock *MockInternal_RangeFeedClient
}

// NewMockInternal_RangeFeedClient creates a new mock instance.
func NewMockInternal_RangeFeedClient(ctrl *gomock.Controller) *MockInternal_RangeFeedClient {
	mock := &MockInternal_RangeFeedClient{ctrl: ctrl}
	mock.recorder = &MockInternal_RangeFeedClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInternal_RangeFeedClient) EXPECT() *MockInternal_RangeFeedClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockInternal_RangeFeedClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockInternal_RangeFeedClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockInternal_RangeFeedClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockInternal_RangeFeedClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockInternal_RangeFeedClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockInternal_RangeFeedClient)(nil).Context))
}

// Header mocks base method.
func (m *MockInternal_RangeFeedClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockInternal_RangeFeedClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockInternal_RangeFeedClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockInternal_RangeFeedClient) Recv() (*RangeFeedEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*RangeFeedEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockInternal_RangeFeedClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockInternal_RangeFeedClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockInternal_RangeFeedClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockInternal_RangeFeedClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockInternal_RangeFeedClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockInternal_RangeFeedClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockInternal_RangeFeedClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockInternal_RangeFeedClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockInternal_RangeFeedClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockInternal_RangeFeedClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockInternal_RangeFeedClient)(nil).Trailer))
}
