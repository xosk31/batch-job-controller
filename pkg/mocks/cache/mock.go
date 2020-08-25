// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bakito/batch-job-controller/pkg/lifecycle (interfaces: Cache)

// Package mock_lifecycle is a generated GoMock package.
package mock_lifecycle

import (
	config "github.com/bakito/batch-job-controller/pkg/config"
	lifecycle "github.com/bakito/batch-job-controller/pkg/lifecycle"
	metrics "github.com/bakito/batch-job-controller/pkg/metrics"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockCache is a mock of Cache interface
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// AddPod mocks base method
func (m *MockCache) AddPod(arg0 lifecycle.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPod", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPod indicates an expected call of AddPod
func (mr *MockCacheMockRecorder) AddPod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPod", reflect.TypeOf((*MockCache)(nil).AddPod), arg0)
}

// AllAdded mocks base method
func (m *MockCache) AllAdded(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllAdded", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AllAdded indicates an expected call of AllAdded
func (mr *MockCacheMockRecorder) AllAdded(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllAdded", reflect.TypeOf((*MockCache)(nil).AllAdded), arg0)
}

// Config mocks base method
func (m *MockCache) Config() config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(config.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockCacheMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCache)(nil).Config))
}

// Has mocks base method
func (m *MockCache) Has(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockCacheMockRecorder) Has(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockCache)(nil).Has), arg0, arg1)
}

// NewExecution mocks base method
func (m *MockCache) NewExecution() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewExecution")
	ret0, _ := ret[0].(string)
	return ret0
}

// NewExecution indicates an expected call of NewExecution
func (mr *MockCacheMockRecorder) NewExecution() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewExecution", reflect.TypeOf((*MockCache)(nil).NewExecution))
}

// PodTerminated mocks base method
func (m *MockCache) PodTerminated(arg0, arg1 string, arg2 v1.PodPhase) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodTerminated", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PodTerminated indicates an expected call of PodTerminated
func (mr *MockCacheMockRecorder) PodTerminated(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodTerminated", reflect.TypeOf((*MockCache)(nil).PodTerminated), arg0, arg1, arg2)
}

// ReportReceived mocks base method
func (m *MockCache) ReportReceived(arg0, arg1 string, arg2 error, arg3 metrics.Results) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportReceived", arg0, arg1, arg2, arg3)
}

// ReportReceived indicates an expected call of ReportReceived
func (mr *MockCacheMockRecorder) ReportReceived(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportReceived", reflect.TypeOf((*MockCache)(nil).ReportReceived), arg0, arg1, arg2, arg3)
}
