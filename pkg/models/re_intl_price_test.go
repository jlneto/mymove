package models_test

import (
	"testing"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
)

func (suite *ModelSuite) TestReIntlPriceValidation() {
	suite.T().Run("test valid ReIntlPrice", func(t *testing.T) {

		validReIntlPrice := models.ReIntlPrice{
			ContractID:            uuid.Must(uuid.NewV4()),
			ServiceID:             uuid.Must(uuid.NewV4()),
			OriginRateAreaID:      uuid.Must(uuid.NewV4()),
			DestinationRateAreaID: uuid.Must(uuid.NewV4()),
			PriceCents:            1342,
		}

		expErrors := map[string][]string{}
		suite.verifyValidationErrors(&validReIntlPrice, expErrors)
	})

	suite.T().Run("test empty ReIntlPrice", func(t *testing.T) {
		invalidReIntlPrice := models.ReIntlPrice{}
		expErrors := map[string][]string{
			"contract_id":              {"ContractID can not be blank."},
			"service_id":               {"ServiceID can not be blank."},
			"destination_rate_area_id": {"DestinationRateAreaID can not be blank."},
			"origin_rate_area_id":      {"OriginRateAreaID can not be blank."},
			"price_cents":              {"PriceCents can not be blank."},
		}
		suite.verifyValidationErrors(&invalidReIntlPrice, expErrors)
	})
}