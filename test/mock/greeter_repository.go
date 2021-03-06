// Code generated by MockGen. DO NOT EDIT.
// Source: server/repository/greeter.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	"github.com/imind-lab/greeter/domain/greeter/repository"
	"github.com/imind-lab/greeter/domain/greeter/repository/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGreeterRepository is a mock of GreeterRepository interface.
type MockGreeterRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterRepositoryMockRecorder
}

// MockGreeterRepositoryMockRecorder is the mock recorder for MockGreeterRepository.
type MockGreeterRepositoryMockRecorder struct {
	mock *MockGreeterRepository
}

// NewMockGreeterRepository creates a new mock instance.
func NewMockGreeterRepository(ctrl *gomock.Controller) *MockGreeterRepository {
	mock := &MockGreeterRepository{ctrl: ctrl}
	mock.recorder = &MockGreeterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGreeterRepository) EXPECT() *MockGreeterRepositoryMockRecorder {
	return m.recorder
}

// CreateGreeter mocks base method.
func (m_2 *MockGreeterRepository) CreateGreeter(ctx context.Context, m model.Greeter) (model.Greeter, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "CreateGreeter", ctx, m)
	ret0, _ := ret[0].(model.Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGreeter indicates an expected call of CreateGreeter.
func (mr *MockGreeterRepositoryMockRecorder) CreateGreeter(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGreeter", reflect.TypeOf((*MockGreeterRepository)(nil).CreateGreeter), ctx, m)
}

// DeleteGreeterById mocks base method.
func (m *MockGreeterRepository) DeleteGreeterById(ctx context.Context, id int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGreeterById", ctx, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGreeterById indicates an expected call of DeleteGreeterById.
func (mr *MockGreeterRepositoryMockRecorder) DeleteGreeterById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGreeterById", reflect.TypeOf((*MockGreeterRepository)(nil).DeleteGreeterById), ctx, id)
}

// FindGreeterById mocks base method.
func (m *MockGreeterRepository) FindGreeterById(ctx context.Context, id int32) (model.Greeter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindGreeterById", ctx, id)
	ret0, _ := ret[0].(model.Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindGreeterById indicates an expected call of FindGreeterById.
func (mr *MockGreeterRepositoryMockRecorder) FindGreeterById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindGreeterById", reflect.TypeOf((*MockGreeterRepository)(nil).FindGreeterById), ctx, id)
}

// GetGreeterById mocks base method.
func (m *MockGreeterRepository) GetGreeterById(ctx context.Context, id int32, opt ...repository.GreeterByIdOption) (model.Greeter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range opt {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGreeterById", varargs...)
	ret0, _ := ret[0].(model.Greeter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGreeterById indicates an expected call of GetGreeterById.
func (mr *MockGreeterRepositoryMockRecorder) GetGreeterById(ctx, id interface{}, opt ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, opt...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGreeterById", reflect.TypeOf((*MockGreeterRepository)(nil).GetGreeterById), varargs...)
}

// GetGreeterList mocks base method.
func (m *MockGreeterRepository) GetGreeterList(ctx context.Context, status, lastId, pageSize, page int32) ([]model.Greeter, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGreeterList", ctx, status, lastId, pageSize, page)
	ret0, _ := ret[0].([]model.Greeter)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGreeterList indicates an expected call of GetGreeterList.
func (mr *MockGreeterRepositoryMockRecorder) GetGreeterList(ctx, status, lastId, pageSize, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGreeterList", reflect.TypeOf((*MockGreeterRepository)(nil).GetGreeterList), ctx, status, lastId, pageSize, page)
}

// UpdateGreeterCount mocks base method.
func (m *MockGreeterRepository) UpdateGreeterCount(ctx context.Context, id, num int32, column string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGreeterCount", ctx, id, num, column)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGreeterCount indicates an expected call of UpdateGreeterCount.
func (mr *MockGreeterRepositoryMockRecorder) UpdateGreeterCount(ctx, id, num, column interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGreeterCount", reflect.TypeOf((*MockGreeterRepository)(nil).UpdateGreeterCount), ctx, id, num, column)
}

// UpdateGreeterStatus mocks base method.
func (m *MockGreeterRepository) UpdateGreeterStatus(ctx context.Context, id, status int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGreeterStatus", ctx, id, status)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGreeterStatus indicates an expected call of UpdateGreeterStatus.
func (mr *MockGreeterRepositoryMockRecorder) UpdateGreeterStatus(ctx, id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGreeterStatus", reflect.TypeOf((*MockGreeterRepository)(nil).UpdateGreeterStatus), ctx, id, status)
}
