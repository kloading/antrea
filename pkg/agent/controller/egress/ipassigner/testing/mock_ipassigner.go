// Copyright 2021 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/agent/controller/egress/ipassigner (interfaces: IPAssigner)

// Package testing is a generated GoMock package.
package testing

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPAssigner is a mock of IPAssigner interface
type MockIPAssigner struct {
	ctrl     *gomock.Controller
	recorder *MockIPAssignerMockRecorder
}

// MockIPAssignerMockRecorder is the mock recorder for MockIPAssigner
type MockIPAssignerMockRecorder struct {
	mock *MockIPAssigner
}

// NewMockIPAssigner creates a new mock instance
func NewMockIPAssigner(ctrl *gomock.Controller) *MockIPAssigner {
	mock := &MockIPAssigner{ctrl: ctrl}
	mock.recorder = &MockIPAssignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPAssigner) EXPECT() *MockIPAssignerMockRecorder {
	return m.recorder
}

// AssignEgressIP mocks base method
func (m *MockIPAssigner) AssignEgressIP(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignEgressIP", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignEgressIP indicates an expected call of AssignEgressIP
func (mr *MockIPAssignerMockRecorder) AssignEgressIP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignEgressIP", reflect.TypeOf((*MockIPAssigner)(nil).AssignEgressIP), arg0, arg1)
}

// AssignedIPs mocks base method
func (m *MockIPAssigner) AssignedIPs() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedIPs")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// AssignedIPs indicates an expected call of AssignedIPs
func (mr *MockIPAssignerMockRecorder) AssignedIPs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedIPs", reflect.TypeOf((*MockIPAssigner)(nil).AssignedIPs))
}

// UnassignEgressIP mocks base method
func (m *MockIPAssigner) UnassignEgressIP(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnassignEgressIP", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnassignEgressIP indicates an expected call of UnassignEgressIP
func (mr *MockIPAssignerMockRecorder) UnassignEgressIP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnassignEgressIP", reflect.TypeOf((*MockIPAssigner)(nil).UnassignEgressIP), arg0)
}