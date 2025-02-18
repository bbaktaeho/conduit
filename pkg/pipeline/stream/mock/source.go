// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/pipeline/stream (interfaces: Source)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	record "github.com/conduitio/conduit/pkg/record"
	gomock "github.com/golang/mock/gomock"
)

// Source is a mock of Source interface.
type Source struct {
	ctrl     *gomock.Controller
	recorder *SourceMockRecorder
}

// SourceMockRecorder is the mock recorder for Source.
type SourceMockRecorder struct {
	mock *Source
}

// NewSource creates a new mock instance.
func NewSource(ctrl *gomock.Controller) *Source {
	mock := &Source{ctrl: ctrl}
	mock.recorder = &SourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Source) EXPECT() *SourceMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *Source) Ack(arg0 context.Context, arg1 record.Position) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *SourceMockRecorder) Ack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*Source)(nil).Ack), arg0, arg1)
}

// Errors mocks base method.
func (m *Source) Errors() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Errors")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// Errors indicates an expected call of Errors.
func (mr *SourceMockRecorder) Errors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errors", reflect.TypeOf((*Source)(nil).Errors))
}

// ID mocks base method.
func (m *Source) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *SourceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*Source)(nil).ID))
}

// Open mocks base method.
func (m *Source) Open(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *SourceMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*Source)(nil).Open), arg0)
}

// Read mocks base method.
func (m *Source) Read(arg0 context.Context) (record.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(record.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *SourceMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*Source)(nil).Read), arg0)
}

// Stop mocks base method.
func (m *Source) Stop(arg0 context.Context) (record.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(record.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *SourceMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*Source)(nil).Stop), arg0)
}

// Teardown mocks base method.
func (m *Source) Teardown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teardown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Teardown indicates an expected call of Teardown.
func (mr *SourceMockRecorder) Teardown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*Source)(nil).Teardown), arg0)
}
