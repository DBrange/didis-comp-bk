// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers (interfaces: ForProfile)
//
// Generated by this command:
//
//	mockgen -destination=tests/mocks/for_profile.mock.go -package=mocks github.com/DBrange/didis-comp-bk/domains/profile/ports/drivers ForProfile
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/DBrange/didis-comp-bk/cmd/api/models"
	dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockForProfile is a mock of ForProfile interface.
type MockForProfile struct {
	ctrl     *gomock.Controller
	recorder *MockForProfileMockRecorder
}

// MockForProfileMockRecorder is the mock recorder for MockForProfile.
type MockForProfileMockRecorder struct {
	mock *MockForProfile
}

// NewMockForProfile creates a new mock instance.
func NewMockForProfile(ctrl *gomock.Controller) *MockForProfile {
	mock := &MockForProfile{ctrl: ctrl}
	mock.recorder = &MockForProfileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockForProfile) EXPECT() *MockForProfileMockRecorder {
	return m.recorder
}

// CloseProfile mocks base method.
func (m *MockForProfile) CloseProfile(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseProfile indicates an expected call of CloseProfile.
func (mr *MockForProfileMockRecorder) CloseProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseProfile", reflect.TypeOf((*MockForProfile)(nil).CloseProfile), arg0, arg1)
}

// GetPersonalInfoByID mocks base method.
func (m *MockForProfile) GetPersonalInfoByID(arg0 context.Context, arg1 string) (*dto.GetPersonalInfoByIDDTORes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonalInfoByID", arg0, arg1)
	ret0, _ := ret[0].(*dto.GetPersonalInfoByIDDTORes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPersonalInfoByID indicates an expected call of GetPersonalInfoByID.
func (mr *MockForProfileMockRecorder) GetPersonalInfoByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonalInfoByID", reflect.TypeOf((*MockForProfile)(nil).GetPersonalInfoByID), arg0, arg1)
}

// GetProfileAvailabilityInfoByID mocks base method.
func (m *MockForProfile) GetProfileAvailabilityInfoByID(arg0 context.Context, arg1, arg2 string) (*dto.GetProfileDailyAvailabilityInfoByIDDTORes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileAvailabilityInfoByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.GetProfileDailyAvailabilityInfoByIDDTORes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileAvailabilityInfoByID indicates an expected call of GetProfileAvailabilityInfoByID.
func (mr *MockForProfileMockRecorder) GetProfileAvailabilityInfoByID(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileAvailabilityInfoByID", reflect.TypeOf((*MockForProfile)(nil).GetProfileAvailabilityInfoByID), arg0, arg1, arg2)
}

// ModifyPassword mocks base method.
func (m *MockForProfile) ModifyPassword(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyPassword", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyPassword indicates an expected call of ModifyPassword.
func (mr *MockForProfileMockRecorder) ModifyPassword(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyPassword", reflect.TypeOf((*MockForProfile)(nil).ModifyPassword), arg0, arg1, arg2, arg3)
}

// ModifyPersonalInfo mocks base method.
func (m *MockForProfile) ModifyPersonalInfo(arg0 context.Context, arg1 string, arg2 *dto.ModifyPersonalInfoDTOReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyPersonalInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyPersonalInfo indicates an expected call of ModifyPersonalInfo.
func (mr *MockForProfileMockRecorder) ModifyPersonalInfo(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyPersonalInfo", reflect.TypeOf((*MockForProfile)(nil).ModifyPersonalInfo), arg0, arg1, arg2)
}

// ModifyProfileAvailability mocks base method.
func (m *MockForProfile) ModifyProfileAvailability(arg0 context.Context, arg1 string, arg2 *dto.ModifyProfileDailyAvailabilityDTOReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyProfileAvailability", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyProfileAvailability indicates an expected call of ModifyProfileAvailability.
func (mr *MockForProfileMockRecorder) ModifyProfileAvailability(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyProfileAvailability", reflect.TypeOf((*MockForProfile)(nil).ModifyProfileAvailability), arg0, arg1, arg2)
}

// RegisterCompetitor mocks base method.
func (m *MockForProfile) RegisterCompetitor(arg0 context.Context, arg1 string, arg2 models.SPORT, arg3 models.COMPETITOR_TYPE) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCompetitor", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterCompetitor indicates an expected call of RegisterCompetitor.
func (mr *MockForProfileMockRecorder) RegisterCompetitor(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCompetitor", reflect.TypeOf((*MockForProfile)(nil).RegisterCompetitor), arg0, arg1, arg2, arg3)
}

// RegisterUser mocks base method.
func (m *MockForProfile) RegisterUser(arg0 context.Context, arg1 *dto.RegisterUserDTOReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockForProfileMockRecorder) RegisterUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockForProfile)(nil).RegisterUser), arg0, arg1)
}
