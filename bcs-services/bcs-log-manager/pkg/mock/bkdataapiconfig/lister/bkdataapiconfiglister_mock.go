// Tencent is pleased to support the open source community by making Blueking Container Service available.
// Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except
// in compliance with the License. You may obtain a copy of the License at
// http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under
// the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/listers/bkbcs.tencent.com/v1 (interfaces: BKDataApiConfigLister)

// Package lister is a generated GoMock package.
package lister

import (
	v1 "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/apis/bkbcs.tencent.com/v1"
	v10 "github.com/Tencent/bk-bcs/bcs-services/bcs-log-manager/pkg/generated/listers/bkbcs.tencent.com/v1"
	gomock "github.com/golang/mock/gomock"
	labels "k8s.io/apimachinery/pkg/labels"
	reflect "reflect"
)

// MockBKDataApiConfigLister is a mock of BKDataApiConfigLister interface
type MockBKDataApiConfigLister struct {
	ctrl     *gomock.Controller
	recorder *MockBKDataApiConfigListerMockRecorder
}

// MockBKDataApiConfigListerMockRecorder is the mock recorder for MockBKDataApiConfigLister
type MockBKDataApiConfigListerMockRecorder struct {
	mock *MockBKDataApiConfigLister
}

// NewMockBKDataApiConfigLister creates a new mock instance
func NewMockBKDataApiConfigLister(ctrl *gomock.Controller) *MockBKDataApiConfigLister {
	mock := &MockBKDataApiConfigLister{ctrl: ctrl}
	mock.recorder = &MockBKDataApiConfigListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBKDataApiConfigLister) EXPECT() *MockBKDataApiConfigListerMockRecorder {
	return m.recorder
}

// BKDataApiConfigs mocks base method
func (m *MockBKDataApiConfigLister) BKDataApiConfigs(arg0 string) v10.BKDataApiConfigNamespaceLister {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BKDataApiConfigs", arg0)
	ret0, _ := ret[0].(v10.BKDataApiConfigNamespaceLister)
	return ret0
}

// BKDataApiConfigs indicates an expected call of BKDataApiConfigs
func (mr *MockBKDataApiConfigListerMockRecorder) BKDataApiConfigs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BKDataApiConfigs", reflect.TypeOf((*MockBKDataApiConfigLister)(nil).BKDataApiConfigs), arg0)
}

// List mocks base method
func (m *MockBKDataApiConfigLister) List(arg0 labels.Selector) ([]*v1.BKDataApiConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]*v1.BKDataApiConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockBKDataApiConfigListerMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBKDataApiConfigLister)(nil).List), arg0)
}
