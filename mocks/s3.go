// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/s3.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	s3 "github.com/aws/aws-sdk-go/service/s3"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockS3 is a mock of S3 interface
type MockS3 struct {
	ctrl     *gomock.Controller
	recorder *MockS3MockRecorder
}

// MockS3MockRecorder is the mock recorder for MockS3
type MockS3MockRecorder struct {
	mock *MockS3
}

// NewMockS3 creates a new mock instance
func NewMockS3(ctrl *gomock.Controller) *MockS3 {
	mock := &MockS3{ctrl: ctrl}
	mock.recorder = &MockS3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3) EXPECT() *MockS3MockRecorder {
	return m.recorder
}

// DeleteObject mocks base method
func (m *MockS3) DeleteObject(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockS3MockRecorder) DeleteObject(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockS3)(nil).DeleteObject), key)
}

// PutObjectRequest mocks base method
func (m *MockS3) PutObjectRequest(key, acl string) (string, http.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObjectRequest", key, acl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(http.Header)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PutObjectRequest indicates an expected call of PutObjectRequest
func (mr *MockS3MockRecorder) PutObjectRequest(key, acl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectRequest", reflect.TypeOf((*MockS3)(nil).PutObjectRequest), key, acl)
}

// PutObject mocks base method
func (m *MockS3) PutObject(key string, body *[]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", key, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject
func (mr *MockS3MockRecorder) PutObject(key, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3)(nil).PutObject), key, body)
}

// PutObjectInput mocks base method
func (m *MockS3) PutObjectInput(params *s3.PutObjectInput, body *[]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObjectInput", params, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObjectInput indicates an expected call of PutObjectInput
func (mr *MockS3MockRecorder) PutObjectInput(params, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectInput", reflect.TypeOf((*MockS3)(nil).PutObjectInput), params, body)
}