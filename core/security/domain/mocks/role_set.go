// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "flamingo.me/flamingo/core/security/domain"
import mock "github.com/stretchr/testify/mock"

// RoleSet is an autogenerated mock type for the RoleSet type
type RoleSet struct {
	mock.Mock
}

// Roles provides a mock function with given fields:
func (_m *RoleSet) Roles() []domain.Role {
	ret := _m.Called()

	var r0 []domain.Role
	if rf, ok := ret.Get(0).(func() []domain.Role); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Role)
		}
	}

	return r0
}
