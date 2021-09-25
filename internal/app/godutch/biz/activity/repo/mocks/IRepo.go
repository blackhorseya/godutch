// Code generated by mockery v2.9.2. DO NOT EDIT.

package mocks

import (
	contextx "github.com/blackhorseya/godutch/internal/pkg/base/contextx"
	event "github.com/blackhorseya/godutch/internal/pkg/entity/event"

	mock "github.com/stretchr/testify/mock"

	user "github.com/blackhorseya/godutch/internal/pkg/entity/user"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

// AddMembers provides a mock function with given fields: ctx, id, newUsers
func (_m *IRepo) AddMembers(ctx contextx.Contextx, id int64, newUsers []*user.Profile) (*event.Activity, error) {
	ret := _m.Called(ctx, id, newUsers)

	var r0 *event.Activity
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, []*user.Profile) *event.Activity); ok {
		r0 = rf(ctx, id, newUsers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64, []*user.Profile) error); ok {
		r1 = rf(ctx, id, newUsers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Count provides a mock function with given fields: ctx, userID
func (_m *IRepo) Count(ctx contextx.Contextx, userID int64) (int, error) {
	ret := _m.Called(ctx, userID)

	var r0 int
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64) int); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, created
func (_m *IRepo) Create(ctx contextx.Contextx, created *event.Activity) (*event.Activity, error) {
	ret := _m.Called(ctx, created)

	var r0 *event.Activity
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *event.Activity) *event.Activity); ok {
		r0 = rf(ctx, created)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *event.Activity) error); ok {
		r1 = rf(ctx, created)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id, userID
func (_m *IRepo) Delete(ctx contextx.Contextx, id int64, userID int64) error {
	ret := _m.Called(ctx, id, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, int64) error); ok {
		r0 = rf(ctx, id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id, userID
func (_m *IRepo) GetByID(ctx contextx.Contextx, id int64, userID int64) (*event.Activity, error) {
	ret := _m.Called(ctx, id, userID)

	var r0 *event.Activity
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, int64) *event.Activity); ok {
		r0 = rf(ctx, id, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64, int64) error); ok {
		r1 = rf(ctx, id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, userID, limit, offset
func (_m *IRepo) List(ctx contextx.Contextx, userID int64, limit int, offset int) ([]*event.Activity, error) {
	ret := _m.Called(ctx, userID, limit, offset)

	var r0 []*event.Activity
	if rf, ok := ret.Get(0).(func(contextx.Contextx, int64, int, int) []*event.Activity); ok {
		r0 = rf(ctx, userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*event.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, int64, int, int) error); ok {
		r1 = rf(ctx, userID, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, updated
func (_m *IRepo) Update(ctx contextx.Contextx, updated *event.Activity) (*event.Activity, error) {
	ret := _m.Called(ctx, updated)

	var r0 *event.Activity
	if rf, ok := ret.Get(0).(func(contextx.Contextx, *event.Activity) *event.Activity); ok {
		r0 = rf(ctx, updated)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(contextx.Contextx, *event.Activity) error); ok {
		r1 = rf(ctx, updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
