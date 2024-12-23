// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ports/s3.go

package mocks

import (
	"io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientMockRecorder
}

type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

func (m *MockS3Client) UploadFile(bucketName, key string, file io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", bucketName, key, file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) UploadFile(bucketName, key, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockS3Client)(nil).UploadFile), bucketName, key, file)
}

