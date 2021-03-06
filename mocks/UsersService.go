// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	dtos "github.com/moisesmorillo/golang-api-example/db/dtos"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UsersService is an autogenerated mock type for the UsersService type
type UsersService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UsersService) Create(ctx context.Context, user *dtos.Users) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dtos.Users) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx
func (_m *UsersService) Get(ctx context.Context) (*[]dtos.Users, error) {
	ret := _m.Called(ctx)

	var r0 *[]dtos.Users
	if rf, ok := ret.Get(0).(func(context.Context) *[]dtos.Users); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]dtos.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
