// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/julianjca/julian-golang-training-beginner/internal/rest (interfaces: StarWarsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	golangtraining "github.com/julianjca/julian-golang-training-beginner"
)

// MockStarWarsClient is a mock of StarWarsClient interface.
type MockStarWarsClient struct {
	ctrl     *gomock.Controller
	recorder *MockStarWarsClientMockRecorder
}

// MockStarWarsClientMockRecorder is the mock recorder for MockStarWarsClient.
type MockStarWarsClientMockRecorder struct {
	mock *MockStarWarsClient
}

// NewMockStarWarsClient creates a new mock instance.
func NewMockStarWarsClient(ctrl *gomock.Controller) *MockStarWarsClient {
	mock := &MockStarWarsClient{ctrl: ctrl}
	mock.recorder = &MockStarWarsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStarWarsClient) EXPECT() *MockStarWarsClientMockRecorder {
	return m.recorder
}

// GetCharacters mocks base method.
func (m *MockStarWarsClient) GetCharacters() (*golangtraining.StarWarsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharacters")
	ret0, _ := ret[0].(*golangtraining.StarWarsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharacters indicates an expected call of GetCharacters.
func (mr *MockStarWarsClientMockRecorder) GetCharacters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharacters", reflect.TypeOf((*MockStarWarsClient)(nil).GetCharacters))
}
