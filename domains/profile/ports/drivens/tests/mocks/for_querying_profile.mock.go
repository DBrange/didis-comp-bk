// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens (interfaces: ForQueryingProfile)
//
// Generated by this command:
//
//	mockgen -destination=tests/mocks/for_querying_profile.mock.go -package=mocks github.com/DBrange/didis-comp-bk/domains/profile/ports/drivens ForQueryingProfile
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/DBrange/didis-comp-bk/cmd/api/models"
	dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	models0 "github.com/DBrange/didis-comp-bk/domains/repository/models/role"
	mongo "go.mongodb.org/mongo-driver/mongo"
	gomock "go.uber.org/mock/gomock"
)

// MockForQueryingProfile is a mock of ForQueryingProfile interface.
type MockForQueryingProfile struct {
	ctrl     *gomock.Controller
	recorder *MockForQueryingProfileMockRecorder
}

// MockForQueryingProfileMockRecorder is the mock recorder for MockForQueryingProfile.
type MockForQueryingProfileMockRecorder struct {
	mock *MockForQueryingProfile
}

// NewMockForQueryingProfile creates a new mock instance.
func NewMockForQueryingProfile(ctrl *gomock.Controller) *MockForQueryingProfile {
	mock := &MockForQueryingProfile{ctrl: ctrl}
	mock.recorder = &MockForQueryingProfileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockForQueryingProfile) EXPECT() *MockForQueryingProfileMockRecorder {
	return m.recorder
}

// CloseProfile mocks base method.
func (m *MockForQueryingProfile) CloseProfile(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseProfile indicates an expected call of CloseProfile.
func (mr *MockForQueryingProfileMockRecorder) CloseProfile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseProfile", reflect.TypeOf((*MockForQueryingProfile)(nil).CloseProfile), arg0, arg1)
}

// CreateAvailability mocks base method.
func (m *MockForQueryingProfile) CreateAvailability(arg0 context.Context, arg1, arg2 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAvailability", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAvailability indicates an expected call of CreateAvailability.
func (mr *MockForQueryingProfileMockRecorder) CreateAvailability(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAvailability", reflect.TypeOf((*MockForQueryingProfile)(nil).CreateAvailability), arg0, arg1, arg2)
}

// CreateLocation mocks base method.
func (m *MockForQueryingProfile) CreateLocation(arg0 context.Context, arg1 *dto.CreateLocationDTOReq) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocation", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocation indicates an expected call of CreateLocation.
func (mr *MockForQueryingProfileMockRecorder) CreateLocation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocation", reflect.TypeOf((*MockForQueryingProfile)(nil).CreateLocation), arg0, arg1)
}

// CreateOrganizer mocks base method.
func (m *MockForQueryingProfile) CreateOrganizer(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganizer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrganizer indicates an expected call of CreateOrganizer.
func (mr *MockForQueryingProfileMockRecorder) CreateOrganizer(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganizer", reflect.TypeOf((*MockForQueryingProfile)(nil).CreateOrganizer), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockForQueryingProfile) CreateUser(arg0 context.Context, arg1 *dto.CreateUserDTOReq) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockForQueryingProfileMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockForQueryingProfile)(nil).CreateUser), arg0, arg1)
}

// GetDailyAvailabilityByID mocks base method.
func (m *MockForQueryingProfile) GetDailyAvailabilityByID(arg0 context.Context, arg1, arg2 string) (*models.GetDailyAvailabilityByIDDTORes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDailyAvailabilityByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.GetDailyAvailabilityByIDDTORes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDailyAvailabilityByID indicates an expected call of GetDailyAvailabilityByID.
func (mr *MockForQueryingProfileMockRecorder) GetDailyAvailabilityByID(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDailyAvailabilityByID", reflect.TypeOf((*MockForQueryingProfile)(nil).GetDailyAvailabilityByID), arg0, arg1, arg2)
}

// GetLocationByID mocks base method.
func (m *MockForQueryingProfile) GetLocationByID(arg0 context.Context, arg1 string) (*dto.GetLocationByIDDTORes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationByID", arg0, arg1)
	ret0, _ := ret[0].(*dto.GetLocationByIDDTORes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationByID indicates an expected call of GetLocationByID.
func (mr *MockForQueryingProfileMockRecorder) GetLocationByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationByID", reflect.TypeOf((*MockForQueryingProfile)(nil).GetLocationByID), arg0, arg1)
}

// GetRoleByNameAndType mocks base method.
func (m *MockForQueryingProfile) GetRoleByNameAndType(arg0 context.Context, arg1, arg2 string) (*models0.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleByNameAndType", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models0.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleByNameAndType indicates an expected call of GetRoleByNameAndType.
func (mr *MockForQueryingProfileMockRecorder) GetRoleByNameAndType(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleByNameAndType", reflect.TypeOf((*MockForQueryingProfile)(nil).GetRoleByNameAndType), arg0, arg1, arg2)
}

// GetUserByID mocks base method.
func (m *MockForQueryingProfile) GetUserByID(arg0 context.Context, arg1 string) (*dto.GetUserByIDDTORes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(*dto.GetUserByIDDTORes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockForQueryingProfileMockRecorder) GetUserByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockForQueryingProfile)(nil).GetUserByID), arg0, arg1)
}

// ModifyPassword mocks base method.
func (m *MockForQueryingProfile) ModifyPassword(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyPassword", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyPassword indicates an expected call of ModifyPassword.
func (mr *MockForQueryingProfileMockRecorder) ModifyPassword(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyPassword", reflect.TypeOf((*MockForQueryingProfile)(nil).ModifyPassword), arg0, arg1, arg2, arg3)
}

// RegisterCompetitor mocks base method.
func (m *MockForQueryingProfile) RegisterCompetitor(arg0 context.Context, arg1 string, arg2 models.SPORT, arg3 models.COMPETITOR_TYPE) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterCompetitor", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterCompetitor indicates an expected call of RegisterCompetitor.
func (mr *MockForQueryingProfileMockRecorder) RegisterCompetitor(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCompetitor", reflect.TypeOf((*MockForQueryingProfile)(nil).RegisterCompetitor), arg0, arg1, arg2, arg3)
}

// UpdateAvailability mocks base method.
func (m *MockForQueryingProfile) UpdateAvailability(arg0 context.Context, arg1 string, arg2 *dto.UpdateDailyAvailabilityDTOReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvailability", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvailability indicates an expected call of UpdateAvailability.
func (mr *MockForQueryingProfileMockRecorder) UpdateAvailability(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvailability", reflect.TypeOf((*MockForQueryingProfile)(nil).UpdateAvailability), arg0, arg1, arg2)
}

// UpdateLocation mocks base method.
func (m *MockForQueryingProfile) UpdateLocation(arg0 context.Context, arg1 string, arg2 *dto.UpdateLocationDTOReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLocation", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLocation indicates an expected call of UpdateLocation.
func (mr *MockForQueryingProfileMockRecorder) UpdateLocation(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLocation", reflect.TypeOf((*MockForQueryingProfile)(nil).UpdateLocation), arg0, arg1, arg2)
}

// UpdateUser mocks base method.
func (m *MockForQueryingProfile) UpdateUser(arg0 context.Context, arg1 string, arg2 *dto.UpdateUserDTOReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockForQueryingProfileMockRecorder) UpdateUser(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockForQueryingProfile)(nil).UpdateUser), arg0, arg1, arg2)
}

// WithTransaction mocks base method.
func (m *MockForQueryingProfile) WithTransaction(arg0 context.Context, arg1 func(mongo.SessionContext) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithTransaction indicates an expected call of WithTransaction.
func (mr *MockForQueryingProfileMockRecorder) WithTransaction(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTransaction", reflect.TypeOf((*MockForQueryingProfile)(nil).WithTransaction), arg0, arg1)
}
