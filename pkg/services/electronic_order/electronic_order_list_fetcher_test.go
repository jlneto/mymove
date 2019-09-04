package electronicorder

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
	"github.com/transcom/mymove/pkg/services/pagination"
	"github.com/transcom/mymove/pkg/services/query"
)

type testElectronicOrderListQueryBuilder struct {
	fakeFetchMany func(model interface{}) error
}

func (t *testElectronicOrderListQueryBuilder) FetchMany(model interface{}, filters []services.QueryFilter, pagination services.Pagination) error {
	m := t.fakeFetchMany(model)
	return m
}

func (suite *ElectronicOrderServiceSuite) TestFetchElectronicOrderList() {
	suite.T().Run("if the transportation order is fetched, it should be returned", func(t *testing.T) {
		id, err := uuid.NewV4()
		suite.NoError(err)
		fakeFetchMany := func(model interface{}) error {
			value := reflect.ValueOf(model).Elem()
			value.Set(reflect.Append(value, reflect.ValueOf(models.ElectronicOrder{ID: id})))
			return nil
		}
		builder := &testElectronicOrderListQueryBuilder{
			fakeFetchMany: fakeFetchMany,
		}

		fetcher := NewElectronicOrderListFetcher(builder)
		filters := []services.QueryFilter{
			query.NewQueryFilter("id", "=", id.String()),
		}

		pagination := pagination.NewPagination(1, 25)

		electronicOrders, err := fetcher.FetchElectronicOrderList(filters, pagination)

		suite.NoError(err)
		suite.Equal(id, electronicOrders[0].ID)
	})

	suite.T().Run("if there is an error, we get it with no electronic orders", func(t *testing.T) {
		fakeFetchMany := func(model interface{}) error {
			return errors.New("Fetch error")
		}
		builder := &testElectronicOrderListQueryBuilder{
			fakeFetchMany: fakeFetchMany,
		}

		fetcher := NewElectronicOrderListFetcher(builder)

		pagination := pagination.NewPagination(1, 25)

		electronicOrders, err := fetcher.FetchElectronicOrderList([]services.QueryFilter{}, pagination)

		suite.Error(err)
		suite.Equal(err.Error(), "Fetch error")
		suite.Equal(models.ElectronicOrders(nil), electronicOrders)
	})
}
