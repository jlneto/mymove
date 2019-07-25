package paperwork

import (
	"errors"
	"time"

	"github.com/transcom/mymove/pkg/route"

	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/rateengine"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/unit"
)

type ppmComputerParams struct {
	Weight                          unit.Pound
	PickupOriginZip5                string
	DutyStationZip5                 string
	DestinationZip5                 string
	DistanceMilesFromPickupZip      int
	DistanceMilesFromDutyStationZip int
	Date                            time.Time
	DaysInSIT                       int
}

type mockPPMComputer struct {
	costComputation   rateengine.CostComputation
	err               error
	ppmComputerParams []ppmComputerParams
}

func (mppmc *mockPPMComputer) ComputePPMIncludingLHDiscount(weight unit.Pound, originPickupZip5 string, originDutyStationZip5 string, destinationZip5 string, distanceMilesFromPickupZip int, distanceMilesFromDutyStationZip int, date time.Time, daysInSIT int) (cost rateengine.CostComputation, err error) {
	mppmc.ppmComputerParams = append(mppmc.ppmComputerParams, ppmComputerParams{
		Weight:                          weight,
		PickupOriginZip5:                originPickupZip5,
		DutyStationZip5:                 originDutyStationZip5,
		DestinationZip5:                 destinationZip5,
		DistanceMilesFromPickupZip:      distanceMilesFromPickupZip,
		DistanceMilesFromDutyStationZip: distanceMilesFromDutyStationZip,
		Date:                            date,
		DaysInSIT:                       daysInSIT,
	})
	return mppmc.costComputation, mppmc.err
}

func (mppmc *mockPPMComputer) CalledWith() []ppmComputerParams {
	return mppmc.ppmComputerParams
}

func (suite *PaperworkSuite) TestComputeObligationsParams() {
	ppmComputer := NewSSWPPMComputer(&mockPPMComputer{})
	pickupPostalCode := "85369"
	destinationPostalCode := "31905"
	ppm := models.PersonallyProcuredMove{
		PickupPostalCode:      &pickupPostalCode,
		DestinationPostalCode: &destinationPostalCode,
	}
	noPPM := models.ShipmentSummaryFormData{PersonallyProcuredMoves: models.PersonallyProcuredMoves{}}
	missingZip := models.ShipmentSummaryFormData{PersonallyProcuredMoves: models.PersonallyProcuredMoves{{}}}
	missingActualMoveDate := models.ShipmentSummaryFormData{PersonallyProcuredMoves: models.PersonallyProcuredMoves{ppm}}
	officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
	session := &auth.Session{
		ApplicationName: auth.OfficeApp,
		UserID:          *officeUser.UserID,
		OfficeUserID:    officeUser.ID,
	}

	_, err1 := ppmComputer.ComputeObligations(noPPM, route.NewTestingPlanner(10), suite.DB(), session)
	_, err2 := ppmComputer.ComputeObligations(missingZip, route.NewTestingPlanner(10), suite.DB(), session)
	_, err3 := ppmComputer.ComputeObligations(missingActualMoveDate, route.NewTestingPlanner(10), suite.DB(), session)

	suite.NotNil(err1)
	suite.NotNil(err2)
	suite.NotNil(err3)
}

func (suite *PaperworkSuite) TestTestComputeObligations() {
	miles := 100
	totalWeightEntitlement := unit.Pound(1000)
	ppmRemainingEntitlement := unit.Pound(2000)
	planner := route.NewTestingPlanner(miles)
	origMoveDate := time.Date(2018, 12, 11, 0, 0, 0, 0, time.UTC)
	actualDate := time.Date(2018, 12, 15, 0, 0, 0, 0, time.UTC)
	pickupPostalCode := "85369"
	dutyStationPostalCode := "50309"
	destinationPostalCode := "31905"
	cents := unit.Cents(1000)
	ppm := testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
		PersonallyProcuredMove: models.PersonallyProcuredMove{
			OriginalMoveDate:      &origMoveDate,
			ActualMoveDate:        &actualDate,
			PickupPostalCode:      &pickupPostalCode,
			DestinationPostalCode: &destinationPostalCode,
			TotalSITCost:          &cents,
		},
	})

	params := models.ShipmentSummaryFormData{
		PersonallyProcuredMoves: models.PersonallyProcuredMoves{ppm},
		WeightAllotment:         models.SSWMaxWeightEntitlement{TotalWeight: totalWeightEntitlement},
		PPMRemainingEntitlement: ppmRemainingEntitlement,
	}
	suite.Run("TestComputeObligations", func() {
		mockComputer := mockPPMComputer{
			costComputation: rateengine.CostComputation{GCC: 100, SITMax: 20000},
		}
		ppmComputer := NewSSWPPMComputer(&mockComputer)
		expectMaxObligationParams := ppmComputerParams{
			Weight:                          totalWeightEntitlement,
			PickupOriginZip5:                pickupPostalCode,
			DutyStationZip5:                 dutyStationPostalCode,
			DestinationZip5:                 destinationPostalCode,
			DistanceMilesFromPickupZip:      miles,
			DistanceMilesFromDutyStationZip: miles,
			Date:                            actualDate,
			DaysInSIT:                       0,
		}
		expectActualObligationParams := ppmComputerParams{
			Weight:                          ppmRemainingEntitlement,
			PickupOriginZip5:                pickupPostalCode,
			DutyStationZip5:                 dutyStationPostalCode,
			DestinationZip5:                 destinationPostalCode,
			DistanceMilesFromPickupZip:      miles,
			DistanceMilesFromDutyStationZip: miles,
			Date:                            actualDate,
			DaysInSIT:                       0,
		}
		officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
		session := &auth.Session{
			ApplicationName: auth.OfficeApp,
			UserID:          *officeUser.UserID,
			OfficeUserID:    officeUser.ID,
		}
		cost, err := ppmComputer.ComputeObligations(params, planner, suite.DB(), session)

		suite.NoError(err)
		calledWith := mockComputer.CalledWith()
		suite.Equal(*ppm.TotalSITCost, cost.ActualObligation.SIT)
		suite.Equal(expectMaxObligationParams, calledWith[0])
		suite.Equal(expectActualObligationParams, calledWith[1])
	})

	suite.Run("TestComputeObligations when actual PPM SIT exceeds MaxSIT", func() {
		mockComputer := mockPPMComputer{
			costComputation: rateengine.CostComputation{SITMax: unit.Cents(500)},
		}
		ppmComputer := NewSSWPPMComputer(&mockComputer)
		officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
		session := &auth.Session{
			ApplicationName: auth.OfficeApp,
			UserID:          *officeUser.UserID,
			OfficeUserID:    officeUser.ID,
		}
		obligations, err := ppmComputer.ComputeObligations(params, planner, suite.DB(), session)

		suite.NoError(err)
		suite.Equal(unit.Cents(500), obligations.ActualObligation.SIT)
	})

	suite.Run("TestComputeObligations when there is no actual PPM SIT", func() {
		ppm := testdatagen.MakePPM(suite.DB(), testdatagen.Assertions{
			PersonallyProcuredMove: models.PersonallyProcuredMove{
				OriginalMoveDate:      &origMoveDate,
				ActualMoveDate:        &actualDate,
				PickupPostalCode:      &pickupPostalCode,
				DestinationPostalCode: &destinationPostalCode,
			},
		})
		shipmentSummaryFormParams := models.ShipmentSummaryFormData{
			PersonallyProcuredMoves: models.PersonallyProcuredMoves{ppm},
			WeightAllotment:         models.SSWMaxWeightEntitlement{TotalWeight: totalWeightEntitlement},
		}
		mockComputer := mockPPMComputer{
			costComputation: rateengine.CostComputation{SITMax: unit.Cents(500)},
		}
		ppmComputer := NewSSWPPMComputer(&mockComputer)
		officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
		session := &auth.Session{
			ApplicationName: auth.OfficeApp,
			UserID:          *officeUser.UserID,
			OfficeUserID:    officeUser.ID,
		}
		obligations, err := ppmComputer.ComputeObligations(shipmentSummaryFormParams, planner, suite.DB(), session)

		suite.NoError(err)
		suite.Equal(unit.Cents(0), obligations.ActualObligation.SIT)
	})

	suite.Run("TestCalcError", func() {
		mockComputer := mockPPMComputer{err: errors.New("ERROR")}
		ppmComputer := SSWPPMComputer{&mockComputer}
		officeUser := testdatagen.MakeDefaultOfficeUser(suite.DB())
		session := &auth.Session{
			ApplicationName: auth.OfficeApp,
			UserID:          *officeUser.UserID,
			OfficeUserID:    officeUser.ID,
		}

		_, err := ppmComputer.ComputeObligations(params, planner, suite.DB(), session)

		suite.NotNil(err)
	})
}
