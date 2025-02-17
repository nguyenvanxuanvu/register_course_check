// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/service (interfaces: RegisterCourseCheckService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

// MockRegisterCourseCheckService is a mock of RegisterCourseCheckService interface.
type MockRegisterCourseCheckService struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterCourseCheckServiceMockRecorder
}

// MockRegisterCourseCheckServiceMockRecorder is the mock recorder for MockRegisterCourseCheckService.
type MockRegisterCourseCheckServiceMockRecorder struct {
	mock *MockRegisterCourseCheckService
}

// NewMockRegisterCourseCheckService creates a new mock instance.
func NewMockRegisterCourseCheckService(ctrl *gomock.Controller) *MockRegisterCourseCheckService {
	mock := &MockRegisterCourseCheckService{ctrl: ctrl}
	mock.recorder = &MockRegisterCourseCheckServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegisterCourseCheckService) EXPECT() *MockRegisterCourseCheckServiceMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockRegisterCourseCheckService) Check(arg0 context.Context, arg1 *dto.CheckRequestDTO) (*dto.CheckResponseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", arg0, arg1)
	ret0, _ := ret[0].(*dto.CheckResponseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockRegisterCourseCheckServiceMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockRegisterCourseCheckService)(nil).Check), arg0, arg1)
}

// Suggestion mocks base method.
func (m *MockRegisterCourseCheckService) Suggestion(arg0 context.Context, arg1 *dto.SuggestionRequestDTO) (*dto.SuggestionResponseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suggestion", arg0, arg1)
	ret0, _ := ret[0].(*dto.SuggestionResponseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Suggestion indicates an expected call of Suggestion.
func (mr *MockRegisterCourseCheckServiceMockRecorder) Suggestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suggestion", reflect.TypeOf((*MockRegisterCourseCheckService)(nil).Suggestion), arg0, arg1)
}

// UpdateCourseCondition mocks base method.
func (m *MockRegisterCourseCheckService) UpdateCourseCondition(arg0 context.Context, arg1 []dto.CourseConditionConfig) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourseCondition", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourseCondition indicates an expected call of UpdateCourseCondition.
func (mr *MockRegisterCourseCheckServiceMockRecorder) UpdateCourseCondition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourseCondition", reflect.TypeOf((*MockRegisterCourseCheckService)(nil).UpdateCourseCondition), arg0, arg1)
}
