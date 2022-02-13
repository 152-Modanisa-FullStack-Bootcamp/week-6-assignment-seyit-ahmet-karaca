// Code generated by MockGen. DO NOT EDIT.
// Source: .\controller\walletController.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIController is a mock of IController interface.
type MockIController struct {
	ctrl     *gomock.Controller
	recorder *MockIControllerMockRecorder
}

// MockIControllerMockRecorder is the mock recorder for MockIController.
type MockIControllerMockRecorder struct {
	mock *MockIController
}

// NewMockIController creates a new mock instance.
func NewMockIController(ctrl *gomock.Controller) *MockIController {
	mock := &MockIController{ctrl: ctrl}
	mock.recorder = &MockIControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIController) EXPECT() *MockIControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIController) Create(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", w, r)
}

// Create indicates an expected call of Create.
func (mr *MockIControllerMockRecorder) Create(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIController)(nil).Create), w, r)
}

// Get mocks base method.
func (m *MockIController) Get(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", w, r)
}

// Get indicates an expected call of Get.
func (mr *MockIControllerMockRecorder) Get(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIController)(nil).Get), w, r)
}

// GetAll mocks base method.
func (m *MockIController) GetAll(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAll", w, r)
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIControllerMockRecorder) GetAll(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIController)(nil).GetAll), w, r)
}

// Handle mocks base method.
func (m *MockIController) Handle(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", w, r)
}

// Handle indicates an expected call of Handle.
func (mr *MockIControllerMockRecorder) Handle(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIController)(nil).Handle), w, r)
}

// UpdateByUsername mocks base method.
func (m *MockIController) UpdateByUsername(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateByUsername", w, r)
}

// UpdateByUsername indicates an expected call of UpdateByUsername.
func (mr *MockIControllerMockRecorder) UpdateByUsername(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByUsername", reflect.TypeOf((*MockIController)(nil).UpdateByUsername), w, r)
}