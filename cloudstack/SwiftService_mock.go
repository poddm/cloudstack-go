//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/SwiftService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSwiftServiceIface is a mock of SwiftServiceIface interface.
type MockSwiftServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockSwiftServiceIfaceMockRecorder
}

// MockSwiftServiceIfaceMockRecorder is the mock recorder for MockSwiftServiceIface.
type MockSwiftServiceIfaceMockRecorder struct {
	mock *MockSwiftServiceIface
}

// NewMockSwiftServiceIface creates a new mock instance.
func NewMockSwiftServiceIface(ctrl *gomock.Controller) *MockSwiftServiceIface {
	mock := &MockSwiftServiceIface{ctrl: ctrl}
	mock.recorder = &MockSwiftServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSwiftServiceIface) EXPECT() *MockSwiftServiceIfaceMockRecorder {
	return m.recorder
}

// AddSwift mocks base method.
func (m *MockSwiftServiceIface) AddSwift(p *AddSwiftParams) (*AddSwiftResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSwift", p)
	ret0, _ := ret[0].(*AddSwiftResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSwift indicates an expected call of AddSwift.
func (mr *MockSwiftServiceIfaceMockRecorder) AddSwift(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSwift", reflect.TypeOf((*MockSwiftServiceIface)(nil).AddSwift), p)
}

// GetSwiftID mocks base method.
func (m *MockSwiftServiceIface) GetSwiftID(keyword string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{keyword}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSwiftID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSwiftID indicates an expected call of GetSwiftID.
func (mr *MockSwiftServiceIfaceMockRecorder) GetSwiftID(keyword interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{keyword}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSwiftID", reflect.TypeOf((*MockSwiftServiceIface)(nil).GetSwiftID), varargs...)
}

// ListSwifts mocks base method.
func (m *MockSwiftServiceIface) ListSwifts(p *ListSwiftsParams) (*ListSwiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSwifts", p)
	ret0, _ := ret[0].(*ListSwiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSwifts indicates an expected call of ListSwifts.
func (mr *MockSwiftServiceIfaceMockRecorder) ListSwifts(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSwifts", reflect.TypeOf((*MockSwiftServiceIface)(nil).ListSwifts), p)
}

// NewAddSwiftParams mocks base method.
func (m *MockSwiftServiceIface) NewAddSwiftParams(url string) *AddSwiftParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddSwiftParams", url)
	ret0, _ := ret[0].(*AddSwiftParams)
	return ret0
}

// NewAddSwiftParams indicates an expected call of NewAddSwiftParams.
func (mr *MockSwiftServiceIfaceMockRecorder) NewAddSwiftParams(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddSwiftParams", reflect.TypeOf((*MockSwiftServiceIface)(nil).NewAddSwiftParams), url)
}

// NewListSwiftsParams mocks base method.
func (m *MockSwiftServiceIface) NewListSwiftsParams() *ListSwiftsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListSwiftsParams")
	ret0, _ := ret[0].(*ListSwiftsParams)
	return ret0
}

// NewListSwiftsParams indicates an expected call of NewListSwiftsParams.
func (mr *MockSwiftServiceIfaceMockRecorder) NewListSwiftsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListSwiftsParams", reflect.TypeOf((*MockSwiftServiceIface)(nil).NewListSwiftsParams))
}