// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nguyenvanxuanvu/register_course_check/pkg/modulefx/repository (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/nguyenvanxuanvu/register_course_check/pkg/dto"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetListCourseOfTeachingPlan mocks base method.
func (m *MockRepository) GetListCourseOfTeachingPlan(arg0, arg1, arg2 string, arg3 int) ([]string, []dto.FreeCreditInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListCourseOfTeachingPlan", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].([]dto.FreeCreditInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetListCourseOfTeachingPlan indicates an expected call of GetListCourseOfTeachingPlan.
func (mr *MockRepositoryMockRecorder) GetListCourseOfTeachingPlan(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListCourseOfTeachingPlan", reflect.TypeOf((*MockRepository)(nil).GetListCourseOfTeachingPlan), arg0, arg1, arg2, arg3)
}

// GetMinMaxCredit mocks base method.
func (m *MockRepository) GetMinMaxCredit(arg0, arg1 string, arg2 int) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinMaxCredit", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMinMaxCredit indicates an expected call of GetMinMaxCredit.
func (mr *MockRepositoryMockRecorder) GetMinMaxCredit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinMaxCredit", reflect.TypeOf((*MockRepository)(nil).GetMinMaxCredit), arg0, arg1, arg2)
}

// UpdateCourseCondition mocks base method.
func (m *MockRepository) UpdateCourseCondition(arg0 []dto.CourseConditionConfig) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourseCondition", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourseCondition indicates an expected call of UpdateCourseCondition.
func (mr *MockRepositoryMockRecorder) UpdateCourseCondition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourseCondition", reflect.TypeOf((*MockRepository)(nil).UpdateCourseCondition), arg0)
}
