// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"
)

// ServiceItemListFetcher is an autogenerated mock type for the ServiceItemListFetcher type
type ServiceItemListFetcher struct {
	mock.Mock
}

// FetchServiceItemList provides a mock function with given fields: params
func (_m *ServiceItemListFetcher) FetchServiceItemList(params interface{}) (models.ServiceItems, error) {
	ret := _m.Called(params)

	var r0 models.ServiceItems
	if rf, ok := ret.Get(0).(func(interface{}) models.ServiceItems); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.ServiceItems)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
