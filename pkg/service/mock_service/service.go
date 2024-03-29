// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SimonRichardson/foodhygiene/pkg/service (interfaces: Service)

package mock_service

import (
	service "github.com/SimonRichardson/foodhygiene/pkg/service"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockService) EXPECT() *MockServiceMockRecorder {
	return _m.recorder
}

// Authorities mocks base method
func (_m *MockService) Authorities() ([]service.Authority, error) {
	ret := _m.ctrl.Call(_m, "Authorities")
	ret0, _ := ret[0].([]service.Authority)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorities indicates an expected call of Authorities
func (_mr *MockServiceMockRecorder) Authorities() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Authorities")
}

// EstablishmentsForAuthority mocks base method
func (_m *MockService) EstablishmentsForAuthority(_param0 string) ([]service.Establishment, error) {
	ret := _m.ctrl.Call(_m, "EstablishmentsForAuthority", _param0)
	ret0, _ := ret[0].([]service.Establishment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstablishmentsForAuthority indicates an expected call of EstablishmentsForAuthority
func (_mr *MockServiceMockRecorder) EstablishmentsForAuthority(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EstablishmentsForAuthority", arg0)
}
