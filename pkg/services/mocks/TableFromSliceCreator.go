// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TableFromSliceCreator is an autogenerated mock type for the TableFromSliceCreator type
type TableFromSliceCreator struct {
	mock.Mock
}

// CreateTableFromSlice provides a mock function with given fields: slice
func (_m *TableFromSliceCreator) CreateTableFromSlice(slice interface{}) error {
	ret := _m.Called(slice)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(slice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}