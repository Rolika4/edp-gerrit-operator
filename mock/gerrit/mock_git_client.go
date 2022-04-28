// Code generated by mockery v2.9.4. DO NOT EDIT.

package mock

import (
	git "github.com/epam/edp-gerrit-operator/v2/pkg/client/git"

	mock "github.com/stretchr/testify/mock"
)

// GitClient is an autogenerated mock type for the GitClient type
type GitClient struct {
	mock.Mock
}

// CheckoutBranch provides a mock function with given fields: projectName, branch
func (_m *GitClient) CheckoutBranch(projectName string, branch string) error {
	ret := _m.Called(projectName, branch)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(projectName, branch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Clone provides a mock function with given fields: projectName
func (_m *GitClient) Clone(projectName string) (string, error) {
	ret := _m.Called(projectName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(projectName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: projectName, message, files, user
func (_m *GitClient) Commit(projectName string, message string, files []string, user *git.User) error {
	ret := _m.Called(projectName, message, files, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string, *git.User) error); ok {
		r0 = rf(projectName, message, files, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateChangeID provides a mock function with given fields:
func (_m *GitClient) GenerateChangeID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Merge provides a mock function with given fields: projectName, sourceBranch, targetBranch, options
func (_m *GitClient) Merge(projectName string, sourceBranch string, targetBranch string, options ...string) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, projectName, sourceBranch, targetBranch)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, ...string) error); ok {
		r0 = rf(projectName, sourceBranch, targetBranch, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Push provides a mock function with given fields: projectName, remote, refSpecs
func (_m *GitClient) Push(projectName string, remote string, refSpecs ...string) (string, error) {
	_va := make([]interface{}, len(refSpecs))
	for _i := range refSpecs {
		_va[_i] = refSpecs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, projectName, remote)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, ...string) string); ok {
		r0 = rf(projectName, remote, refSpecs...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, ...string) error); ok {
		r1 = rf(projectName, remote, refSpecs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetFileContents provides a mock function with given fields: projectName, filePath, contents
func (_m *GitClient) SetFileContents(projectName string, filePath string, contents string) error {
	ret := _m.Called(projectName, filePath, contents)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(projectName, filePath, contents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetProjectUser provides a mock function with given fields: projectName, user
func (_m *GitClient) SetProjectUser(projectName string, user *git.User) error {
	ret := _m.Called(projectName, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *git.User) error); ok {
		r0 = rf(projectName, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
