// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quay/claircore/updater (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package=mock_updater -destination=./mocks.go github.com/quay/claircore/updater Store
//

// Package mock_updater is a generated GoMock package.
package mock_updater

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	driver "github.com/quay/claircore/updater/driver/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetLatestUpdateOperations mocks base method.
func (m *MockStore) GetLatestUpdateOperations(arg0 context.Context) ([]driver.UpdateOperation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestUpdateOperations", arg0)
	ret0, _ := ret[0].([]driver.UpdateOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestUpdateOperations indicates an expected call of GetLatestUpdateOperations.
func (mr *MockStoreMockRecorder) GetLatestUpdateOperations(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestUpdateOperations", reflect.TypeOf((*MockStore)(nil).GetLatestUpdateOperations), arg0)
}

// UpdateEnrichments mocks base method.
func (m *MockStore) UpdateEnrichments(arg0 context.Context, arg1 uuid.UUID, arg2 string, arg3 driver.Fingerprint, arg4 []driver.EnrichmentRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEnrichments", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEnrichments indicates an expected call of UpdateEnrichments.
func (mr *MockStoreMockRecorder) UpdateEnrichments(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnrichments", reflect.TypeOf((*MockStore)(nil).UpdateEnrichments), arg0, arg1, arg2, arg3, arg4)
}

// UpdateVulnerabilities mocks base method.
func (m *MockStore) UpdateVulnerabilities(arg0 context.Context, arg1 uuid.UUID, arg2 string, arg3 driver.Fingerprint, arg4 *driver.ParsedVulnerabilities) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVulnerabilities", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVulnerabilities indicates an expected call of UpdateVulnerabilities.
func (mr *MockStoreMockRecorder) UpdateVulnerabilities(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVulnerabilities", reflect.TypeOf((*MockStore)(nil).UpdateVulnerabilities), arg0, arg1, arg2, arg3, arg4)
}
