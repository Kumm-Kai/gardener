// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/kubernetes (interfaces: Interface)

// Package kubernetes is a generated GoMock package.
package kubernetes

import (
	versioned "github.com/gardener/gardener/pkg/client/core/clientset/versioned"
	versioned0 "github.com/gardener/gardener/pkg/client/garden/clientset/versioned"
	kubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	gomock "github.com/golang/mock/gomock"
	clientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	meta "k8s.io/apimachinery/pkg/api/meta"
	kubernetes0 "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	clientset0 "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// APIExtension mocks base method
func (m *MockInterface) APIExtension() clientset.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIExtension")
	ret0, _ := ret[0].(clientset.Interface)
	return ret0
}

// APIExtension indicates an expected call of APIExtension
func (mr *MockInterfaceMockRecorder) APIExtension() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIExtension", reflect.TypeOf((*MockInterface)(nil).APIExtension))
}

// APIRegistration mocks base method
func (m *MockInterface) APIRegistration() clientset0.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIRegistration")
	ret0, _ := ret[0].(clientset0.Interface)
	return ret0
}

// APIRegistration indicates an expected call of APIRegistration
func (mr *MockInterfaceMockRecorder) APIRegistration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIRegistration", reflect.TypeOf((*MockInterface)(nil).APIRegistration))
}

// Applier mocks base method
func (m *MockInterface) Applier() kubernetes.ApplierInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Applier")
	ret0, _ := ret[0].(kubernetes.ApplierInterface)
	return ret0
}

// Applier indicates an expected call of Applier
func (mr *MockInterfaceMockRecorder) Applier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Applier", reflect.TypeOf((*MockInterface)(nil).Applier))
}

// CheckForwardPodPort mocks base method
func (m *MockInterface) CheckForwardPodPort(arg0, arg1 string, arg2, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckForwardPodPort", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckForwardPodPort indicates an expected call of CheckForwardPodPort
func (mr *MockInterfaceMockRecorder) CheckForwardPodPort(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckForwardPodPort", reflect.TypeOf((*MockInterface)(nil).CheckForwardPodPort), arg0, arg1, arg2, arg3)
}

// Client mocks base method
func (m *MockInterface) Client() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockInterfaceMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockInterface)(nil).Client))
}

// ForwardPodPort mocks base method
func (m *MockInterface) ForwardPodPort(arg0, arg1 string, arg2, arg3 int) (chan struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForwardPodPort", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForwardPodPort indicates an expected call of ForwardPodPort
func (mr *MockInterfaceMockRecorder) ForwardPodPort(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardPodPort", reflect.TypeOf((*MockInterface)(nil).ForwardPodPort), arg0, arg1, arg2, arg3)
}

// Garden mocks base method
func (m *MockInterface) Garden() versioned0.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Garden")
	ret0, _ := ret[0].(versioned0.Interface)
	return ret0
}

// Garden indicates an expected call of Garden
func (mr *MockInterfaceMockRecorder) Garden() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Garden", reflect.TypeOf((*MockInterface)(nil).Garden))
}

// GardenCore mocks base method
func (m *MockInterface) GardenCore() versioned.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GardenCore")
	ret0, _ := ret[0].(versioned.Interface)
	return ret0
}

// GardenCore indicates an expected call of GardenCore
func (mr *MockInterfaceMockRecorder) GardenCore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GardenCore", reflect.TypeOf((*MockInterface)(nil).GardenCore))
}

// Kubernetes mocks base method
func (m *MockInterface) Kubernetes() kubernetes0.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kubernetes")
	ret0, _ := ret[0].(kubernetes0.Interface)
	return ret0
}

// Kubernetes indicates an expected call of Kubernetes
func (mr *MockInterfaceMockRecorder) Kubernetes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kubernetes", reflect.TypeOf((*MockInterface)(nil).Kubernetes))
}

// RESTClient mocks base method
func (m *MockInterface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockInterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockInterface)(nil).RESTClient))
}

// RESTConfig mocks base method
func (m *MockInterface) RESTConfig() *rest.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTConfig")
	ret0, _ := ret[0].(*rest.Config)
	return ret0
}

// RESTConfig indicates an expected call of RESTConfig
func (mr *MockInterfaceMockRecorder) RESTConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTConfig", reflect.TypeOf((*MockInterface)(nil).RESTConfig))
}

// RESTMapper mocks base method
func (m *MockInterface) RESTMapper() meta.RESTMapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	return ret0
}

// RESTMapper indicates an expected call of RESTMapper
func (mr *MockInterfaceMockRecorder) RESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTMapper", reflect.TypeOf((*MockInterface)(nil).RESTMapper))
}

// Version mocks base method
func (m *MockInterface) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockInterfaceMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockInterface)(nil).Version))
}
