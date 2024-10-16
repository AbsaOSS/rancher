// Code generated by MockGen. DO NOT EDIT.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

// MockSubjectAccessReview is a mock of SubjectAccessReview interface.
type MockSubjectAccessReview struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectAccessReviewMockRecorder
}

// MockSubjectAccessReviewMockRecorder is the mock recorder for MockSubjectAccessReview.
type MockSubjectAccessReviewMockRecorder struct {
	mock *MockSubjectAccessReview
}

// NewMockSubjectAccessReview creates a new mock instance.
func NewMockSubjectAccessReview(ctrl *gomock.Controller) *MockSubjectAccessReview {
	mock := &MockSubjectAccessReview{ctrl: ctrl}
	mock.recorder = &MockSubjectAccessReviewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectAccessReview) EXPECT() *MockSubjectAccessReviewMockRecorder {
	return m.recorder
}

// UserCanImpersonateExtras mocks base method.
func (m *MockSubjectAccessReview) UserCanImpersonateExtras(req *http.Request, user string, impExtras map[string][]string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCanImpersonateExtras", req, user, impExtras)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCanImpersonateExtras indicates an expected call of UserCanImpersonateExtras.
func (mr *MockSubjectAccessReviewMockRecorder) UserCanImpersonateExtras(req, user, impExtras interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCanImpersonateExtras", reflect.TypeOf((*MockSubjectAccessReview)(nil).UserCanImpersonateExtras), req, user, impExtras)
}

// UserCanImpersonateGroup mocks base method.
func (m *MockSubjectAccessReview) UserCanImpersonateGroup(req *http.Request, user string, group string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCanImpersonateGroup", req, user, group)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCanImpersonateGroup indicates an expected call of UserCanImpersonateGroup.
func (mr *MockSubjectAccessReviewMockRecorder) UserCanImpersonateGroup(req, user, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCanImpersonateGroup", reflect.TypeOf((*MockSubjectAccessReview)(nil).UserCanImpersonateGroup), req, user, group)
}

// UserCanImpersonateUser mocks base method.
func (m *MockSubjectAccessReview) UserCanImpersonateUser(req *http.Request, user, impUser string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCanImpersonateUser", req, user, impUser)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserCanImpersonateUser indicates an expected call of UserCanImpersonateUser.
func (mr *MockSubjectAccessReviewMockRecorder) UserCanImpersonateUser(req, user, impUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCanImpersonateUser", reflect.TypeOf((*MockSubjectAccessReview)(nil).UserCanImpersonateUser), req, user, impUser)
}

// MockSubjectAccessReviewClientGetter is a mock of SubjectAccessReviewClientGetter interface.
type MockSubjectAccessReviewClientGetter struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectAccessReviewClientGetterMockRecorder
}

// MockSubjectAccessReviewClientGetterMockRecorder is the mock recorder for MockSubjectAccessReviewClientGetter.
type MockSubjectAccessReviewClientGetterMockRecorder struct {
	mock *MockSubjectAccessReviewClientGetter
}

// NewMockSubjectAccessReviewClientGetter creates a new mock instance.
func NewMockSubjectAccessReviewClientGetter(ctrl *gomock.Controller) *MockSubjectAccessReviewClientGetter {
	mock := &MockSubjectAccessReviewClientGetter{ctrl: ctrl}
	mock.recorder = &MockSubjectAccessReviewClientGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectAccessReviewClientGetter) EXPECT() *MockSubjectAccessReviewClientGetterMockRecorder {
	return m.recorder
}

// SubjectAccessReviewForCluster mocks base method.
func (m *MockSubjectAccessReviewClientGetter) SubjectAccessReviewForCluster(request *http.Request) (v1.SubjectAccessReviewInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubjectAccessReviewForCluster", request)
	ret0, _ := ret[0].(v1.SubjectAccessReviewInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubjectAccessReviewForCluster indicates an expected call of SubjectAccessReviewForCluster.
func (mr *MockSubjectAccessReviewClientGetterMockRecorder) SubjectAccessReviewForCluster(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubjectAccessReviewForCluster", reflect.TypeOf((*MockSubjectAccessReviewClientGetter)(nil).SubjectAccessReviewForCluster), request)
}
