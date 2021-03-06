// Code generated by mockery v2.9.2. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	mock "github.com/stretchr/testify/mock"

	user "github.com/blackhorseya/godutch/internal/pkg/entity/user"
)

// IBiz is an autogenerated mock type for the IBiz type
type IBiz struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IBiz) GetByID(ctx contextx.Contextx, id int64) (*user.Profile, error) {
	ret := _m.Called(ctx, id)

	var r0 *user.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64) *user.Profile); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByToken provides a mock function with given fields: ctx, token
func (_m *IBiz) GetByToken(ctx contextx.Contextx, token string) (*user.Profile, error) {
	ret := _m.Called(ctx, token)

	var r0 *user.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string) *user.Profile); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, email, password
func (_m *IBiz) Login(ctx contextx.Contextx, email string, password string) (*user.Profile, error) {
	ret := _m.Called(ctx, email, password)

	var r0 *user.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string) *user.Profile); ok {
		r0 = rf(ctx, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx
func (_m *IBiz) Logout(ctx contextx.Contextx) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Signup provides a mock function with given fields: ctx, email, password, name
func (_m *IBiz) Signup(ctx contextx.Contextx, email string, password string, name string) (*user.Profile, error) {
	ret := _m.Called(ctx, email, password, name)

	var r0 *user.Profile
	if rf, ok := ret.Get(0).(func(contextx.Contextx, string, string, string) *user.Profile); ok {
		r0 = rf(ctx, email, password, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Profile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, string, string, string) error); ok {
		r1 = rf(ctx, email, password, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
