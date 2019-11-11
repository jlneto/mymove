// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// MoveTaskOrderActualWeightUpdater is an autogenerated mock type for the MoveTaskOrderActualWeightUpdater type
type MoveTaskOrderActualWeightUpdater struct {
	mock.Mock
}

// UpdateMoveTaskOrderActualWeight provides a mock function with given fields: moveTaskOrderID, actualWeight
func (_m *MoveTaskOrderActualWeightUpdater) UpdateMoveTaskOrderActualWeight(moveTaskOrderID uuid.UUID, actualWeight int64) (*models.MoveTaskOrder, error) {
	ret := _m.Called(moveTaskOrderID, actualWeight)

	var r0 *models.MoveTaskOrder
	if rf, ok := ret.Get(0).(func(uuid.UUID, int64) *models.MoveTaskOrder); ok {
		r0 = rf(moveTaskOrderID, actualWeight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MoveTaskOrder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, int64) error); ok {
		r1 = rf(moveTaskOrderID, actualWeight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}