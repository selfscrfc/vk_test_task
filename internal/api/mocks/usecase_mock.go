// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"
	api_models "vk_test_task/internal/api/models"

	gomock "github.com/golang/mock/gomock"
)

// MockUseCaseInterface is a mock of UseCaseInterface interface.
type MockUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseInterfaceMockRecorder
}

// MockUseCaseInterfaceMockRecorder is the mock recorder for MockUseCaseInterface.
type MockUseCaseInterfaceMockRecorder struct {
	mock *MockUseCaseInterface
}

// NewMockUseCaseInterface creates a new mock instance.
func NewMockUseCaseInterface(ctrl *gomock.Controller) *MockUseCaseInterface {
	mock := &MockUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCaseInterface) EXPECT() *MockUseCaseInterfaceMockRecorder {
	return m.recorder
}

// CreateActor mocks base method.
func (m *MockUseCaseInterface) CreateActor(params api_models.CreateActorParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActor", params)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActor indicates an expected call of CreateActor.
func (mr *MockUseCaseInterfaceMockRecorder) CreateActor(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActor", reflect.TypeOf((*MockUseCaseInterface)(nil).CreateActor), params)
}

// CreateFilm mocks base method.
func (m *MockUseCaseInterface) CreateFilm(params api_models.CreateFilmParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFilm", params)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFilm indicates an expected call of CreateFilm.
func (mr *MockUseCaseInterfaceMockRecorder) CreateFilm(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilm", reflect.TypeOf((*MockUseCaseInterface)(nil).CreateFilm), params)
}

// DeleteActor mocks base method.
func (m *MockUseCaseInterface) DeleteActor(params api_models.DeleteActorParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActor", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActor indicates an expected call of DeleteActor.
func (mr *MockUseCaseInterfaceMockRecorder) DeleteActor(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActor", reflect.TypeOf((*MockUseCaseInterface)(nil).DeleteActor), params)
}

// DeleteFilm mocks base method.
func (m *MockUseCaseInterface) DeleteFilm(params api_models.DeleteFilmParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFilm", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFilm indicates an expected call of DeleteFilm.
func (mr *MockUseCaseInterfaceMockRecorder) DeleteFilm(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilm", reflect.TypeOf((*MockUseCaseInterface)(nil).DeleteFilm), params)
}

// GetActors mocks base method.
func (m *MockUseCaseInterface) GetActors() (api_models.GetActorsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActors")
	ret0, _ := ret[0].(api_models.GetActorsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActors indicates an expected call of GetActors.
func (mr *MockUseCaseInterfaceMockRecorder) GetActors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActors", reflect.TypeOf((*MockUseCaseInterface)(nil).GetActors))
}

// GetFilms mocks base method.
func (m *MockUseCaseInterface) GetFilms(params api_models.GetFilmsParams) (api_models.GetFilmsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilms", params)
	ret0, _ := ret[0].(api_models.GetFilmsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilms indicates an expected call of GetFilms.
func (mr *MockUseCaseInterfaceMockRecorder) GetFilms(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilms", reflect.TypeOf((*MockUseCaseInterface)(nil).GetFilms), params)
}

// SearchFilm mocks base method.
func (m *MockUseCaseInterface) SearchFilm(params api_models.SearchFilmParams) (api_models.SearchFilmResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFilm", params)
	ret0, _ := ret[0].(api_models.SearchFilmResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFilm indicates an expected call of SearchFilm.
func (mr *MockUseCaseInterfaceMockRecorder) SearchFilm(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFilm", reflect.TypeOf((*MockUseCaseInterface)(nil).SearchFilm), params)
}

// SignIn mocks base method.
func (m *MockUseCaseInterface) SignIn(params api_models.AuthParams) (api_models.SignInUseCaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", params)
	ret0, _ := ret[0].(api_models.SignInUseCaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUseCaseInterfaceMockRecorder) SignIn(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUseCaseInterface)(nil).SignIn), params)
}

// SignUp mocks base method.
func (m *MockUseCaseInterface) SignUp(params api_models.AuthParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUseCaseInterfaceMockRecorder) SignUp(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUseCaseInterface)(nil).SignUp), params)
}

// UpdateActor mocks base method.
func (m *MockUseCaseInterface) UpdateActor(params api_models.UpdateActorParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActor", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActor indicates an expected call of UpdateActor.
func (mr *MockUseCaseInterfaceMockRecorder) UpdateActor(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActor", reflect.TypeOf((*MockUseCaseInterface)(nil).UpdateActor), params)
}

// UpdateFilm mocks base method.
func (m *MockUseCaseInterface) UpdateFilm(params api_models.UpdateFilmParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilm", params)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFilm indicates an expected call of UpdateFilm.
func (mr *MockUseCaseInterfaceMockRecorder) UpdateFilm(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilm", reflect.TypeOf((*MockUseCaseInterface)(nil).UpdateFilm), params)
}
