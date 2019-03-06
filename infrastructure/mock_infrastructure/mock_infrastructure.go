// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/ghcp/infrastructure/interfaces (interfaces: GitHubClient,GitHubClientInit)

// Package mock_infrastructure is a generated GoMock package.
package mock_infrastructure

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v24/github"
	interfaces "github.com/int128/ghcp/infrastructure/interfaces"
	reflect "reflect"
)

// MockGitHubClient is a mock of GitHubClient interface
type MockGitHubClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitHubClientMockRecorder
}

// MockGitHubClientMockRecorder is the mock recorder for MockGitHubClient
type MockGitHubClientMockRecorder struct {
	mock *MockGitHubClient
}

// NewMockGitHubClient creates a new mock instance
func NewMockGitHubClient(ctrl *gomock.Controller) *MockGitHubClient {
	mock := &MockGitHubClient{ctrl: ctrl}
	mock.recorder = &MockGitHubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitHubClient) EXPECT() *MockGitHubClientMockRecorder {
	return m.recorder
}

// CreateBlob mocks base method
func (m *MockGitHubClient) CreateBlob(arg0 context.Context, arg1, arg2 string, arg3 *github.Blob) (*github.Blob, *github.Response, error) {
	ret := m.ctrl.Call(m, "CreateBlob", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Blob)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateBlob indicates an expected call of CreateBlob
func (mr *MockGitHubClientMockRecorder) CreateBlob(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlob", reflect.TypeOf((*MockGitHubClient)(nil).CreateBlob), arg0, arg1, arg2, arg3)
}

// CreateCommit mocks base method
func (m *MockGitHubClient) CreateCommit(arg0 context.Context, arg1, arg2 string, arg3 *github.Commit) (*github.Commit, *github.Response, error) {
	ret := m.ctrl.Call(m, "CreateCommit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Commit)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCommit indicates an expected call of CreateCommit
func (mr *MockGitHubClientMockRecorder) CreateCommit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommit", reflect.TypeOf((*MockGitHubClient)(nil).CreateCommit), arg0, arg1, arg2, arg3)
}

// CreateRef mocks base method
func (m *MockGitHubClient) CreateRef(arg0 context.Context, arg1, arg2 string, arg3 *github.Reference) (*github.Reference, *github.Response, error) {
	ret := m.ctrl.Call(m, "CreateRef", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Reference)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRef indicates an expected call of CreateRef
func (mr *MockGitHubClientMockRecorder) CreateRef(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRef", reflect.TypeOf((*MockGitHubClient)(nil).CreateRef), arg0, arg1, arg2, arg3)
}

// CreateTree mocks base method
func (m *MockGitHubClient) CreateTree(arg0 context.Context, arg1, arg2, arg3 string, arg4 []github.TreeEntry) (*github.Tree, *github.Response, error) {
	ret := m.ctrl.Call(m, "CreateTree", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*github.Tree)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTree indicates an expected call of CreateTree
func (mr *MockGitHubClientMockRecorder) CreateTree(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockGitHubClient)(nil).CreateTree), arg0, arg1, arg2, arg3, arg4)
}

// Query mocks base method
func (m *MockGitHubClient) Query(arg0 context.Context, arg1 interface{}, arg2 map[string]interface{}) error {
	ret := m.ctrl.Call(m, "Query", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query
func (mr *MockGitHubClientMockRecorder) Query(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockGitHubClient)(nil).Query), arg0, arg1, arg2)
}

// UpdateRef mocks base method
func (m *MockGitHubClient) UpdateRef(arg0 context.Context, arg1, arg2 string, arg3 *github.Reference, arg4 bool) (*github.Reference, *github.Response, error) {
	ret := m.ctrl.Call(m, "UpdateRef", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*github.Reference)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateRef indicates an expected call of UpdateRef
func (mr *MockGitHubClientMockRecorder) UpdateRef(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRef", reflect.TypeOf((*MockGitHubClient)(nil).UpdateRef), arg0, arg1, arg2, arg3, arg4)
}

// MockGitHubClientInit is a mock of GitHubClientInit interface
type MockGitHubClientInit struct {
	ctrl     *gomock.Controller
	recorder *MockGitHubClientInitMockRecorder
}

// MockGitHubClientInitMockRecorder is the mock recorder for MockGitHubClientInit
type MockGitHubClientInitMockRecorder struct {
	mock *MockGitHubClientInit
}

// NewMockGitHubClientInit creates a new mock instance
func NewMockGitHubClientInit(ctrl *gomock.Controller) *MockGitHubClientInit {
	mock := &MockGitHubClientInit{ctrl: ctrl}
	mock.recorder = &MockGitHubClientInitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitHubClientInit) EXPECT() *MockGitHubClientInitMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockGitHubClientInit) Init(arg0 interfaces.GitHubClientInitOptions) {
	m.ctrl.Call(m, "Init", arg0)
}

// Init indicates an expected call of Init
func (mr *MockGitHubClientInitMockRecorder) Init(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockGitHubClientInit)(nil).Init), arg0)
}
