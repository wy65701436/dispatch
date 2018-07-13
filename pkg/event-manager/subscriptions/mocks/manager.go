// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import entities "github.com/vmware/dispatch/pkg/event-manager/subscriptions/entities"
import mock "github.com/stretchr/testify/mock"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Manager) Create(_a0 context.Context, _a1 *entities.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *Manager) Delete(_a0 context.Context, _a1 *entities.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields: _a0, _a1
func (_m *Manager) Run(_a0 context.Context, _a1 []*entities.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*entities.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Manager) Update(_a0 context.Context, _a1 *entities.Subscription) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Subscription) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}