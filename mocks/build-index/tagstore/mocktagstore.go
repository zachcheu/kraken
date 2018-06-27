// Code generated by MockGen. DO NOT EDIT.
// Source: code.uber.internal/infra/kraken/build-index/tagstore (interfaces: Store,FileStore)

// Package mocktagstore is a generated GoMock package.
package mocktagstore

import (
	core "code.uber.internal/infra/kraken/core"
	base "code.uber.internal/infra/kraken/lib/store/base"
	metadata "code.uber.internal/infra/kraken/lib/store/metadata"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
	time "time"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockStore) Get(arg0 string, arg1 bool) (core.Digest, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(core.Digest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockStoreMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), arg0, arg1)
}

// Has mocks base method
func (m *MockStore) Has(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has
func (mr *MockStoreMockRecorder) Has(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockStore)(nil).Has), arg0)
}

// Put mocks base method
func (m *MockStore) Put(arg0 string, arg1 core.Digest, arg2 time.Duration) error {
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockStoreMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockStore)(nil).Put), arg0, arg1, arg2)
}

// MockFileStore is a mock of FileStore interface
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return m.recorder
}

// CreateCacheFile mocks base method
func (m *MockFileStore) CreateCacheFile(arg0 string, arg1 io.Reader) error {
	ret := m.ctrl.Call(m, "CreateCacheFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCacheFile indicates an expected call of CreateCacheFile
func (mr *MockFileStoreMockRecorder) CreateCacheFile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCacheFile", reflect.TypeOf((*MockFileStore)(nil).CreateCacheFile), arg0, arg1)
}

// GetCacheFileReader mocks base method
func (m *MockFileStore) GetCacheFileReader(arg0 string) (base.FileReader, error) {
	ret := m.ctrl.Call(m, "GetCacheFileReader", arg0)
	ret0, _ := ret[0].(base.FileReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheFileReader indicates an expected call of GetCacheFileReader
func (mr *MockFileStoreMockRecorder) GetCacheFileReader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheFileReader", reflect.TypeOf((*MockFileStore)(nil).GetCacheFileReader), arg0)
}

// SetCacheFileMetadata mocks base method
func (m *MockFileStore) SetCacheFileMetadata(arg0 string, arg1 metadata.Metadata) (bool, error) {
	ret := m.ctrl.Call(m, "SetCacheFileMetadata", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetCacheFileMetadata indicates an expected call of SetCacheFileMetadata
func (mr *MockFileStoreMockRecorder) SetCacheFileMetadata(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCacheFileMetadata", reflect.TypeOf((*MockFileStore)(nil).SetCacheFileMetadata), arg0, arg1)
}