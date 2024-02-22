// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: producer_manager.go

// Package api is a generated GoMock package.
package api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	messaging "github.com/uber/cadence/common/messaging"
)

// MockProducerManager is a mock of ProducerManager interface.
type MockProducerManager struct {
	ctrl     *gomock.Controller
	recorder *MockProducerManagerMockRecorder
}

// MockProducerManagerMockRecorder is the mock recorder for MockProducerManager.
type MockProducerManagerMockRecorder struct {
	mock *MockProducerManager
}

// NewMockProducerManager creates a new mock instance.
func NewMockProducerManager(ctrl *gomock.Controller) *MockProducerManager {
	mock := &MockProducerManager{ctrl: ctrl}
	mock.recorder = &MockProducerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducerManager) EXPECT() *MockProducerManagerMockRecorder {
	return m.recorder
}

// GetProducerByDomain mocks base method.
func (m *MockProducerManager) GetProducerByDomain(domain string) (messaging.Producer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducerByDomain", domain)
	ret0, _ := ret[0].(messaging.Producer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducerByDomain indicates an expected call of GetProducerByDomain.
func (mr *MockProducerManagerMockRecorder) GetProducerByDomain(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducerByDomain", reflect.TypeOf((*MockProducerManager)(nil).GetProducerByDomain), domain)
}