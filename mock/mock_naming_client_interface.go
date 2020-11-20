/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: clients/naming_client/naming_client_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nacos-group/nacos-sdk-go/model"
	vo "github.com/nacos-group/nacos-sdk-go/vo"
)

// MockINamingClient is a mock of INamingClient interface
type MockINamingClient struct {
	ctrl     *gomock.Controller
	recorder *MockINamingClientMockRecorder
}

// MockINamingClientMockRecorder is the mock recorder for MockINamingClient
type MockINamingClientMockRecorder struct {
	mock *MockINamingClient
}

// NewMockINamingClient creates a new mock instance
func NewMockINamingClient(ctrl *gomock.Controller) *MockINamingClient {
	mock := &MockINamingClient{ctrl: ctrl}
	mock.recorder = &MockINamingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINamingClient) EXPECT() *MockINamingClientMockRecorder {
	return m.recorder
}

// RegisterInstance mocks base method
func (m *MockINamingClient) RegisterInstance(param vo.RegisterInstanceParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInstance", param)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInstance indicates an expected call of RegisterInstance
func (mr *MockINamingClientMockRecorder) RegisterInstance(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInstance", reflect.TypeOf((*MockINamingClient)(nil).RegisterInstance), param)
}

// DeregisterInstance mocks base method
func (m *MockINamingClient) DeregisterInstance(param vo.DeregisterInstanceParam) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterInstance", param)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterInstance indicates an expected call of DeregisterInstance
func (mr *MockINamingClientMockRecorder) DeregisterInstance(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterInstance", reflect.TypeOf((*MockINamingClient)(nil).DeregisterInstance), param)
}

// GetService mocks base method
func (m *MockINamingClient) GetService(param vo.GetServiceParam) (model.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", param)
	ret0, _ := ret[0].(model.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockINamingClientMockRecorder) GetService(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockINamingClient)(nil).GetService), param)
}

// SelectInstances mocks base method
func (m *MockINamingClient) SelectInstances(param vo.SelectInstancesParam) ([]model.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectInstances", param)
	ret0, _ := ret[0].([]model.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectInstances indicates an expected call of SelectInstances
func (mr *MockINamingClientMockRecorder) SelectInstances(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectInstances", reflect.TypeOf((*MockINamingClient)(nil).SelectInstances), param)
}

// SelectOneHealthyInstance mocks base method
func (m *MockINamingClient) SelectOneHealthyInstance(param vo.SelectOneHealthInstanceParam) (*model.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectOneHealthyInstance", param)
	ret0, _ := ret[0].(*model.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectOneHealthyInstance indicates an expected call of SelectOneHealthyInstance
func (mr *MockINamingClientMockRecorder) SelectOneHealthyInstance(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOneHealthyInstance", reflect.TypeOf((*MockINamingClient)(nil).SelectOneHealthyInstance), param)
}

// Subscribe mocks base method
func (m *MockINamingClient) Subscribe(param *vo.SubscribeParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockINamingClientMockRecorder) Subscribe(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockINamingClient)(nil).Subscribe), param)
}

// Unsubscribe mocks base method
func (m *MockINamingClient) Unsubscribe(param *vo.SubscribeParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockINamingClientMockRecorder) Unsubscribe(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockINamingClient)(nil).Unsubscribe), param)
}
