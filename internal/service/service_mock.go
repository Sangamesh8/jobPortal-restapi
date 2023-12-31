// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock.go -package=service
//
// Package service is a generated GoMock package.
package service

import (
	context "context"
	models "job-portal-api/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockJobPortalService is a mock of JobPortalService interface.
type MockJobPortalService struct {
	ctrl     *gomock.Controller
	recorder *MockJobPortalServiceMockRecorder
}

// MockJobPortalServiceMockRecorder is the mock recorder for MockJobPortalService.
type MockJobPortalServiceMockRecorder struct {
	mock *MockJobPortalService
}

// NewMockJobPortalService creates a new mock instance.
func NewMockJobPortalService(ctrl *gomock.Controller) *MockJobPortalService {
	mock := &MockJobPortalService{ctrl: ctrl}
	mock.recorder = &MockJobPortalServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobPortalService) EXPECT() *MockJobPortalServiceMockRecorder {
	return m.recorder
}

// AddCompanyDetails mocks base method.
func (m *MockJobPortalService) AddCompanyDetails(ctx context.Context, companyData models.Company) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCompanyDetails", ctx, companyData)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCompanyDetails indicates an expected call of AddCompanyDetails.
func (mr *MockJobPortalServiceMockRecorder) AddCompanyDetails(ctx, companyData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCompanyDetails", reflect.TypeOf((*MockJobPortalService)(nil).AddCompanyDetails), ctx, companyData)
}

// AddJobDetails mocks base method.
func (m *MockJobPortalService) AddJobDetails(ctx context.Context, jobData models.CreateJobs, CompanyID uint64) (models.ResponseForJobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddJobDetails", ctx, jobData, CompanyID)
	ret0, _ := ret[0].(models.ResponseForJobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddJobDetails indicates an expected call of AddJobDetails.
func (mr *MockJobPortalServiceMockRecorder) AddJobDetails(ctx, jobData, CompanyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddJobDetails", reflect.TypeOf((*MockJobPortalService)(nil).AddJobDetails), ctx, jobData, CompanyID)
}

// ProcessJobApplication mocks base method.
func (m *MockJobPortalService) ProcessJobApplication(ctx context.Context, jobData []models.JobApplicantResponse) ([]models.JobApplicantResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessJobApplication", ctx, jobData)
	ret0, _ := ret[0].([]models.JobApplicantResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessJobApplication indicates an expected call of ProcessJobApplication.
func (mr *MockJobPortalServiceMockRecorder) ProcessJobApplication(ctx, jobData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessJobApplication", reflect.TypeOf((*MockJobPortalService)(nil).ProcessJobApplication), ctx, jobData)
}

// UserSignIn mocks base method.
func (m *MockJobPortalService) UserSignIn(ctx context.Context, userData models.NewUser) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignIn", ctx, userData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignIn indicates an expected call of UserSignIn.
func (mr *MockJobPortalServiceMockRecorder) UserSignIn(ctx, userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignIn", reflect.TypeOf((*MockJobPortalService)(nil).UserSignIn), ctx, userData)
}

// UserSignup mocks base method.
func (m *MockJobPortalService) UserSignup(ctx context.Context, userData models.NewUser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignup", ctx, userData)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignup indicates an expected call of UserSignup.
func (mr *MockJobPortalServiceMockRecorder) UserSignup(ctx, userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignup", reflect.TypeOf((*MockJobPortalService)(nil).UserSignup), ctx, userData)
}

// ViewAllCompanies mocks base method.
func (m *MockJobPortalService) ViewAllCompanies(ctx context.Context) ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewAllCompanies", ctx)
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewAllCompanies indicates an expected call of ViewAllCompanies.
func (mr *MockJobPortalServiceMockRecorder) ViewAllCompanies(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewAllCompanies", reflect.TypeOf((*MockJobPortalService)(nil).ViewAllCompanies), ctx)
}

// ViewAllJobs mocks base method.
func (m *MockJobPortalService) ViewAllJobs(ctx context.Context) ([]models.Jobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewAllJobs", ctx)
	ret0, _ := ret[0].([]models.Jobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewAllJobs indicates an expected call of ViewAllJobs.
func (mr *MockJobPortalServiceMockRecorder) ViewAllJobs(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewAllJobs", reflect.TypeOf((*MockJobPortalService)(nil).ViewAllJobs), ctx)
}

// ViewCompanyDetails mocks base method.
func (m *MockJobPortalService) ViewCompanyDetails(ctx context.Context, cid uint64) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompanyDetails", ctx, cid)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewCompanyDetails indicates an expected call of ViewCompanyDetails.
func (mr *MockJobPortalServiceMockRecorder) ViewCompanyDetails(ctx, cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompanyDetails", reflect.TypeOf((*MockJobPortalService)(nil).ViewCompanyDetails), ctx, cid)
}

// ViewJobByCompanyID mocks base method.
func (m *MockJobPortalService) ViewJobByCompanyID(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobByCompanyID", ctx, cid)
	ret0, _ := ret[0].([]models.Jobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJobByCompanyID indicates an expected call of ViewJobByCompanyID.
func (mr *MockJobPortalServiceMockRecorder) ViewJobByCompanyID(ctx, cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobByCompanyID", reflect.TypeOf((*MockJobPortalService)(nil).ViewJobByCompanyID), ctx, cid)
}

// ViewJobById mocks base method.
func (m *MockJobPortalService) ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobById", ctx, jid)
	ret0, _ := ret[0].(models.Jobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJobById indicates an expected call of ViewJobById.
func (mr *MockJobPortalServiceMockRecorder) ViewJobById(ctx, jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobById", reflect.TypeOf((*MockJobPortalService)(nil).ViewJobById), ctx, jid)
}
