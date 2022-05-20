// Code generated by mockery v2.9.4. DO NOT EDIT.

package gerrit

import (
	mock "github.com/stretchr/testify/mock"

	platform "github.com/epam/edp-gerrit-operator/v2/pkg/service/platform"

	resty "gopkg.in/resty.v1"

	v1alpha1 "github.com/epam/edp-gerrit-operator/v2/pkg/apis/v2/v1alpha1"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterfaceMock struct {
	mock.Mock
}

// AddAccessRights provides a mock function with given fields: projectName, permissions
func (_m *ClientInterfaceMock) AddAccessRights(projectName string, permissions []AccessInfo) error {
	ret := _m.Called(projectName, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []AccessInfo) error); ok {
		r0 = rf(projectName, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUserToGroup provides a mock function with given fields: groupName, username
func (_m *ClientInterfaceMock) AddUserToGroup(groupName string, username string) error {
	ret := _m.Called(groupName, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(groupName, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUserToGroups provides a mock function with given fields: userName, groupNames
func (_m *ClientInterfaceMock) AddUserToGroups(userName string, groupNames []string) error {
	ret := _m.Called(userName, groupNames)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(userName, groupNames)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeAbandon provides a mock function with given fields: changeID
func (_m *ClientInterfaceMock) ChangeAbandon(changeID string) error {
	ret := _m.Called(changeID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(changeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeGet provides a mock function with given fields: changeID
func (_m *ClientInterfaceMock) ChangeGet(changeID string) (*Change, error) {
	ret := _m.Called(changeID)

	var r0 *Change
	if rf, ok := ret.Get(0).(func(string) *Change); ok {
		r0 = rf(changeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Change)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(changeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangePassword provides a mock function with given fields: username, password
func (_m *ClientInterfaceMock) ChangePassword(username string, password string) error {
	ret := _m.Called(username, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckCredentials provides a mock function with given fields:
func (_m *ClientInterfaceMock) CheckCredentials() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckGroup provides a mock function with given fields: groupName
func (_m *ClientInterfaceMock) CheckGroup(groupName string) (*int, error) {
	ret := _m.Called(groupName)

	var r0 *int
	if rf, ok := ret.Get(0).(func(string) *int); ok {
		r0 = rf(groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroup provides a mock function with given fields: name, description, visibleToAll
func (_m *ClientInterfaceMock) CreateGroup(name string, description string, visibleToAll bool) (*Group, error) {
	ret := _m.Called(name, description, visibleToAll)

	var r0 *Group
	if rf, ok := ret.Get(0).(func(string, string, bool) *Group); ok {
		r0 = rf(name, description, visibleToAll)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool) error); ok {
		r1 = rf(name, description, visibleToAll)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProject provides a mock function with given fields: prj
func (_m *ClientInterfaceMock) CreateProject(prj *Project) error {
	ret := _m.Called(prj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Project) error); ok {
		r0 = rf(prj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: username, password, fullname, publicKey
func (_m *ClientInterfaceMock) CreateUser(username string, password string, fullname string, publicKey string) error {
	ret := _m.Called(username, password, fullname, publicKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(username, password, fullname, publicKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAccessRights provides a mock function with given fields: projectName, permissions
func (_m *ClientInterfaceMock) DeleteAccessRights(projectName string, permissions []AccessInfo) error {
	ret := _m.Called(projectName, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []AccessInfo) error); ok {
		r0 = rf(projectName, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProject provides a mock function with given fields: name
func (_m *ClientInterfaceMock) DeleteProject(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserFromGroup provides a mock function with given fields: groupName, username
func (_m *ClientInterfaceMock) DeleteUserFromGroup(groupName string, username string) error {
	ret := _m.Called(groupName, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(groupName, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProject provides a mock function with given fields: name
func (_m *ClientInterfaceMock) GetProject(name string) (*Project, error) {
	ret := _m.Called(name)

	var r0 *Project
	if rf, ok := ret.Get(0).(func(string) *Project); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitAdminUser provides a mock function with given fields: instance, _a1, GerritScriptsPath, podName, gerritAdminPublicKey
func (_m *ClientInterfaceMock) InitAdminUser(instance v1alpha1.Gerrit, _a1 platform.PlatformService, GerritScriptsPath string, podName string, gerritAdminPublicKey string) (v1alpha1.Gerrit, error) {
	ret := _m.Called(instance, _a1, GerritScriptsPath, podName, gerritAdminPublicKey)

	var r0 v1alpha1.Gerrit
	if rf, ok := ret.Get(0).(func(v1alpha1.Gerrit, platform.PlatformService, string, string, string) v1alpha1.Gerrit); ok {
		r0 = rf(instance, _a1, GerritScriptsPath, podName, gerritAdminPublicKey)
	} else {
		r0 = ret.Get(0).(v1alpha1.Gerrit)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v1alpha1.Gerrit, platform.PlatformService, string, string, string) error); ok {
		r1 = rf(instance, _a1, GerritScriptsPath, podName, gerritAdminPublicKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitAllProjects provides a mock function with given fields: instance, _a1, GerritScriptsPath, podName, gerritAdminPublicKey
func (_m *ClientInterfaceMock) InitAllProjects(instance v1alpha1.Gerrit, _a1 platform.PlatformService, GerritScriptsPath string, podName string, gerritAdminPublicKey string) error {
	ret := _m.Called(instance, _a1, GerritScriptsPath, podName, gerritAdminPublicKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(v1alpha1.Gerrit, platform.PlatformService, string, string, string) error); ok {
		r0 = rf(instance, _a1, GerritScriptsPath, podName, gerritAdminPublicKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitNewRestClient provides a mock function with given fields: instance, url, user, password
func (_m *ClientInterfaceMock) InitNewRestClient(instance *v1alpha1.Gerrit, url string, user string, password string) error {
	ret := _m.Called(instance, url, user, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Gerrit, string, string, string) error); ok {
		r0 = rf(instance, url, user, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitNewSshClient provides a mock function with given fields: userName, privateKey, host, port
func (_m *ClientInterfaceMock) InitNewSshClient(userName string, privateKey []byte, host string, port int32) error {
	ret := _m.Called(userName, privateKey, host, port)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, string, int32) error); ok {
		r0 = rf(userName, privateKey, host, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListProjectBranches provides a mock function with given fields: projectName
func (_m *ClientInterfaceMock) ListProjectBranches(projectName string) ([]Branch, error) {
	ret := _m.Called(projectName)

	var r0 []Branch
	if rf, ok := ret.Get(0).(func(string) []Branch); ok {
		r0 = rf(projectName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Branch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(projectName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: _type
func (_m *ClientInterfaceMock) ListProjects(_type string) ([]Project, error) {
	ret := _m.Called(_type)

	var r0 []Project
	if rf, ok := ret.Get(0).(func(string) []Project); ok {
		r0 = rf(_type)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_type)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReloadPlugin provides a mock function with given fields: plugin
func (_m *ClientInterfaceMock) ReloadPlugin(plugin string) error {
	ret := _m.Called(plugin)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(plugin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Resty provides a mock function with given fields:
func (_m *ClientInterfaceMock) Resty() *resty.Client {
	ret := _m.Called()

	var r0 *resty.Client
	if rf, ok := ret.Get(0).(func() *resty.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resty.Client)
		}
	}

	return r0
}

// SetProjectParent provides a mock function with given fields: projectName, parentName
func (_m *ClientInterfaceMock) SetProjectParent(projectName string, parentName string) error {
	ret := _m.Called(projectName, parentName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(projectName, parentName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAccessRights provides a mock function with given fields: projectName, permissions
func (_m *ClientInterfaceMock) UpdateAccessRights(projectName string, permissions []AccessInfo) error {
	ret := _m.Called(projectName, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []AccessInfo) error); ok {
		r0 = rf(projectName, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateGroup provides a mock function with given fields: groupID, description, visibleToAll
func (_m *ClientInterfaceMock) UpdateGroup(groupID string, description string, visibleToAll bool) error {
	ret := _m.Called(groupID, description, visibleToAll)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool) error); ok {
		r0 = rf(groupID, description, visibleToAll)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProject provides a mock function with given fields: prj
func (_m *ClientInterfaceMock) UpdateProject(prj *Project) error {
	ret := _m.Called(prj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Project) error); ok {
		r0 = rf(prj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}