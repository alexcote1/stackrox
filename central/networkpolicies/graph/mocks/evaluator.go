// Code generated by MockGen. DO NOT EDIT.
// Source: evaluator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	tree "github.com/stackrox/rox/pkg/networkgraph/tree"
	reflect "reflect"
)

// MockEvaluator is a mock of Evaluator interface
type MockEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockEvaluatorMockRecorder
}

// MockEvaluatorMockRecorder is the mock recorder for MockEvaluator
type MockEvaluatorMockRecorder struct {
	mock *MockEvaluator
}

// NewMockEvaluator creates a new mock instance
func NewMockEvaluator(ctrl *gomock.Controller) *MockEvaluator {
	mock := &MockEvaluator{ctrl: ctrl}
	mock.recorder = &MockEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEvaluator) EXPECT() *MockEvaluatorMockRecorder {
	return m.recorder
}

// GetGraph mocks base method
func (m *MockEvaluator) GetGraph(clusterID string, deployments []*storage.Deployment, networkTree tree.ReadOnlyNetworkTree, networkPolicies []*storage.NetworkPolicy, includePorts bool) *v1.NetworkGraph {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGraph", clusterID, deployments, networkTree, networkPolicies, includePorts)
	ret0, _ := ret[0].(*v1.NetworkGraph)
	return ret0
}

// GetGraph indicates an expected call of GetGraph
func (mr *MockEvaluatorMockRecorder) GetGraph(clusterID, deployments, networkTree, networkPolicies, includePorts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGraph", reflect.TypeOf((*MockEvaluator)(nil).GetGraph), clusterID, deployments, networkTree, networkPolicies, includePorts)
}

// GetAppliedPolicies mocks base method
func (m *MockEvaluator) GetAppliedPolicies(deployments []*storage.Deployment, networkTree tree.ReadOnlyNetworkTree, networkPolicies []*storage.NetworkPolicy) []*storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppliedPolicies", deployments, networkTree, networkPolicies)
	ret0, _ := ret[0].([]*storage.NetworkPolicy)
	return ret0
}

// GetAppliedPolicies indicates an expected call of GetAppliedPolicies
func (mr *MockEvaluatorMockRecorder) GetAppliedPolicies(deployments, networkTree, networkPolicies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppliedPolicies", reflect.TypeOf((*MockEvaluator)(nil).GetAppliedPolicies), deployments, networkTree, networkPolicies)
}

// IncrementEpoch mocks base method
func (m *MockEvaluator) IncrementEpoch(clusterID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncrementEpoch", clusterID)
}

// IncrementEpoch indicates an expected call of IncrementEpoch
func (mr *MockEvaluatorMockRecorder) IncrementEpoch(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementEpoch", reflect.TypeOf((*MockEvaluator)(nil).IncrementEpoch), clusterID)
}

// Epoch mocks base method
func (m *MockEvaluator) Epoch(clusterID string) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Epoch", clusterID)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Epoch indicates an expected call of Epoch
func (mr *MockEvaluatorMockRecorder) Epoch(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Epoch", reflect.TypeOf((*MockEvaluator)(nil).Epoch), clusterID)
}

// MocknamespaceProvider is a mock of namespaceProvider interface
type MocknamespaceProvider struct {
	ctrl     *gomock.Controller
	recorder *MocknamespaceProviderMockRecorder
}

// MocknamespaceProviderMockRecorder is the mock recorder for MocknamespaceProvider
type MocknamespaceProviderMockRecorder struct {
	mock *MocknamespaceProvider
}

// NewMocknamespaceProvider creates a new mock instance
func NewMocknamespaceProvider(ctrl *gomock.Controller) *MocknamespaceProvider {
	mock := &MocknamespaceProvider{ctrl: ctrl}
	mock.recorder = &MocknamespaceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocknamespaceProvider) EXPECT() *MocknamespaceProviderMockRecorder {
	return m.recorder
}

// GetNamespaces mocks base method
func (m *MocknamespaceProvider) GetNamespaces(ctx context.Context) ([]*storage.NamespaceMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaces", ctx)
	ret0, _ := ret[0].([]*storage.NamespaceMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaces indicates an expected call of GetNamespaces
func (mr *MocknamespaceProviderMockRecorder) GetNamespaces(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaces", reflect.TypeOf((*MocknamespaceProvider)(nil).GetNamespaces), ctx)
}
