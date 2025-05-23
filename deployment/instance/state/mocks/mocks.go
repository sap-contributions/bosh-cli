// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudfoundry/bosh-cli/v7/deployment/instance/state (interfaces: BuilderFactory,Builder,State)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	agentclient "github.com/cloudfoundry/bosh-agent/v2/agentclient"
	applyspec "github.com/cloudfoundry/bosh-agent/v2/agentclient/applyspec"
	blobstore "github.com/cloudfoundry/bosh-cli/v7/blobstore"
	state "github.com/cloudfoundry/bosh-cli/v7/deployment/instance/state"
	manifest "github.com/cloudfoundry/bosh-cli/v7/deployment/manifest"
	ui "github.com/cloudfoundry/bosh-cli/v7/ui"
	gomock "github.com/golang/mock/gomock"
)

// MockBuilderFactory is a mock of BuilderFactory interface.
type MockBuilderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderFactoryMockRecorder
}

// MockBuilderFactoryMockRecorder is the mock recorder for MockBuilderFactory.
type MockBuilderFactoryMockRecorder struct {
	mock *MockBuilderFactory
}

// NewMockBuilderFactory creates a new mock instance.
func NewMockBuilderFactory(ctrl *gomock.Controller) *MockBuilderFactory {
	mock := &MockBuilderFactory{ctrl: ctrl}
	mock.recorder = &MockBuilderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilderFactory) EXPECT() *MockBuilderFactoryMockRecorder {
	return m.recorder
}

// NewBuilder mocks base method.
func (m *MockBuilderFactory) NewBuilder(arg0 blobstore.Blobstore, arg1 agentclient.AgentClient) state.Builder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBuilder", arg0, arg1)
	ret0, _ := ret[0].(state.Builder)
	return ret0
}

// NewBuilder indicates an expected call of NewBuilder.
func (mr *MockBuilderFactoryMockRecorder) NewBuilder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBuilder", reflect.TypeOf((*MockBuilderFactory)(nil).NewBuilder), arg0, arg1)
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockBuilder) Build(arg0 string, arg1 int, arg2 manifest.Manifest, arg3 ui.Stage, arg4 agentclient.AgentState) (state.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(state.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockBuilderMockRecorder) Build(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockBuilder)(nil).Build), arg0, arg1, arg2, arg3, arg4)
}

// BuildInitialState mocks base method.
func (m *MockBuilder) BuildInitialState(arg0 string, arg1 int, arg2 manifest.Manifest) (state.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildInitialState", arg0, arg1, arg2)
	ret0, _ := ret[0].(state.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildInitialState indicates an expected call of BuildInitialState.
func (mr *MockBuilderMockRecorder) BuildInitialState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildInitialState", reflect.TypeOf((*MockBuilder)(nil).BuildInitialState), arg0, arg1, arg2)
}

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// CompiledPackages mocks base method.
func (m *MockState) CompiledPackages() []state.PackageRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompiledPackages")
	ret0, _ := ret[0].([]state.PackageRef)
	return ret0
}

// CompiledPackages indicates an expected call of CompiledPackages.
func (mr *MockStateMockRecorder) CompiledPackages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompiledPackages", reflect.TypeOf((*MockState)(nil).CompiledPackages))
}

// NetworkInterfaces mocks base method.
func (m *MockState) NetworkInterfaces() []state.NetworkRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkInterfaces")
	ret0, _ := ret[0].([]state.NetworkRef)
	return ret0
}

// NetworkInterfaces indicates an expected call of NetworkInterfaces.
func (mr *MockStateMockRecorder) NetworkInterfaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkInterfaces", reflect.TypeOf((*MockState)(nil).NetworkInterfaces))
}

// RenderedJobListArchive mocks base method.
func (m *MockState) RenderedJobListArchive() state.BlobRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderedJobListArchive")
	ret0, _ := ret[0].(state.BlobRef)
	return ret0
}

// RenderedJobListArchive indicates an expected call of RenderedJobListArchive.
func (mr *MockStateMockRecorder) RenderedJobListArchive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderedJobListArchive", reflect.TypeOf((*MockState)(nil).RenderedJobListArchive))
}

// RenderedJobs mocks base method.
func (m *MockState) RenderedJobs() []state.JobRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderedJobs")
	ret0, _ := ret[0].([]state.JobRef)
	return ret0
}

// RenderedJobs indicates an expected call of RenderedJobs.
func (mr *MockStateMockRecorder) RenderedJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderedJobs", reflect.TypeOf((*MockState)(nil).RenderedJobs))
}

// ToApplySpec mocks base method.
func (m *MockState) ToApplySpec() applyspec.ApplySpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToApplySpec")
	ret0, _ := ret[0].(applyspec.ApplySpec)
	return ret0
}

// ToApplySpec indicates an expected call of ToApplySpec.
func (mr *MockStateMockRecorder) ToApplySpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToApplySpec", reflect.TypeOf((*MockState)(nil).ToApplySpec))
}
