// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: workflow_rebuilder.go
//
// Generated by this command:
//
//	mockgen -copyright_file ../../LICENSE -package history -source workflow_rebuilder.go -destination workflow_rebuilder_mock.go
//

// Package history is a generated GoMock package.
package history

import (
	context "context"
	reflect "reflect"

	definition "go.temporal.io/server/common/definition"
	gomock "go.uber.org/mock/gomock"
)

// MockworkflowRebuilder is a mock of workflowRebuilder interface.
type MockworkflowRebuilder struct {
	ctrl     *gomock.Controller
	recorder *MockworkflowRebuilderMockRecorder
	isgomock struct{}
}

// MockworkflowRebuilderMockRecorder is the mock recorder for MockworkflowRebuilder.
type MockworkflowRebuilderMockRecorder struct {
	mock *MockworkflowRebuilder
}

// NewMockworkflowRebuilder creates a new mock instance.
func NewMockworkflowRebuilder(ctrl *gomock.Controller) *MockworkflowRebuilder {
	mock := &MockworkflowRebuilder{ctrl: ctrl}
	mock.recorder = &MockworkflowRebuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockworkflowRebuilder) EXPECT() *MockworkflowRebuilderMockRecorder {
	return m.recorder
}

// rebuild mocks base method.
func (m *MockworkflowRebuilder) rebuild(ctx context.Context, workflowKey definition.WorkflowKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "rebuild", ctx, workflowKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// rebuild indicates an expected call of rebuild.
func (mr *MockworkflowRebuilderMockRecorder) rebuild(ctx, workflowKey any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "rebuild", reflect.TypeOf((*MockworkflowRebuilder)(nil).rebuild), ctx, workflowKey)
}
