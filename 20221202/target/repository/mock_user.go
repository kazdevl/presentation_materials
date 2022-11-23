// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/kazdevl/presentation_materials/20221202/target/model"
)

// MockIFUserModelRepository is a mock of IFUserModelRepository interface.
type MockIFUserModelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFUserModelRepositoryMockRecorder
}

// MockIFUserModelRepositoryMockRecorder is the mock recorder for MockIFUserModelRepository.
type MockIFUserModelRepositoryMockRecorder struct {
	mock *MockIFUserModelRepository
}

// NewMockIFUserModelRepository creates a new mock instance.
func NewMockIFUserModelRepository(ctrl *gomock.Controller) *MockIFUserModelRepository {
	mock := &MockIFUserModelRepository{ctrl: ctrl}
	mock.recorder = &MockIFUserModelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFUserModelRepository) EXPECT() *MockIFUserModelRepositoryMockRecorder {
	return m.recorder
}

// FindByName mocks base method.
func (m *MockIFUserModelRepository) FindByName(name string) (model.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].(model.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockIFUserModelRepositoryMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockIFUserModelRepository)(nil).FindByName), name)
}

// Get mocks base method.
func (m *MockIFUserModelRepository) Get(uk model.UserPK) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", uk)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIFUserModelRepositoryMockRecorder) Get(uk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIFUserModelRepository)(nil).Get), uk)
}
