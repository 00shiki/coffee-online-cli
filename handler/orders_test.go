package handler

import (
	"coffee-online-cli/entity"
	MOCK_ORDERS "coffee-online-cli/repository/orders/mocks"
	MOCK_PRODUCTS "coffee-online-cli/repository/products/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestHandler_CoffeeOrdersSuccess(t *testing.T) {
	productsMock := []entity.Product{
		{
			ID:    1,
			Name:  "Cappuccino",
			Price: 15000,
			Stock: 15,
		},
	}
	productsRepo := new(MOCK_PRODUCTS.Repo)
	productsRepo.On("FetchProducts").Return(productsMock, nil)
	products, errFetchProducts := productsRepo.FetchProducts()
	assert.Equal(t, products, productsMock)
	assert.Equal(t, len(products), len(productsMock))
	assert.Nil(t, errFetchProducts)

	order := &entity.Order{
		ID: 1,
	}
	ordersRepo := new(MOCK_ORDERS.Repo)
	ordersRepo.On("OrderPayment", mock.Anything).Return(nil)
	errOrderPayment := ordersRepo.OrderPayment(order)
	assert.Nil(t, errOrderPayment)

	ordersRepo.On("CreateOrder", mock.Anything).Return(nil)
	errCreateOrder := ordersRepo.CreateOrder(order)
	assert.Nil(t, errCreateOrder)
}

func TestHandler_CoffeeOrdersFetchProductsFail(t *testing.T) {
	errFetchProductsMock := errors.New("fetch products error")
	productsRepo := new(MOCK_PRODUCTS.Repo)
	productsRepo.On("FetchProducts").Return(nil, errFetchProductsMock)
	products, errFetchProducts := productsRepo.FetchProducts()
	assert.Nil(t, products)
	assert.Equal(t, errFetchProducts, errFetchProductsMock)
}

func TestHandler_ShipOrdersSuccess(t *testing.T) {
	ordersMock := []entity.Order{
		{
			ID: 1,
		},
	}
	ordersRepo := new(MOCK_ORDERS.Repo)
	ordersRepo.On("FetchPendingOrders").Return(ordersMock, nil)
	orders, errFetchPendingOrders := ordersRepo.FetchPendingOrders()
	assert.Nil(t, errFetchPendingOrders)
	assert.Equal(t, len(orders), len(ordersMock))
	assert.Equal(t, orders[0].ID, 1)

	ordersRepo.On(
		"UpdateOrderShippingStatus",
		mock.AnythingOfType("int"),
		mock.Anything,
	).Return(nil)
	errUpdateOrder := ordersRepo.UpdateOrderShippingStatus(orders[0].ID, entity.Shipped)
	assert.Nil(t, errUpdateOrder)
}

func TestHandler_ShipOrdersFail(t *testing.T) {
	errFetchPendingOrdersMock := errors.New("fetch pending orders error")
	ordersRepo := new(MOCK_ORDERS.Repo)
	ordersRepo.On("FetchPendingOrders").Return(nil, errFetchPendingOrdersMock)
	orders, errFetchPendingOrders := ordersRepo.FetchPendingOrders()
	assert.Nil(t, orders)
	assert.NotNil(t, errFetchPendingOrders)
	assert.Equal(t, errFetchPendingOrders, errFetchPendingOrdersMock)
}

func TestHandler_UserOrdersSuccess(t *testing.T) {
	ordersMock := []entity.Order{
		{
			ID: 1,
		},
	}
	ordersRepo := new(MOCK_ORDERS.Repo)
	ordersRepo.On("FetchUserOrders", mock.AnythingOfType("int")).Return(ordersMock, nil)
	orders, errFetchUserOrders := ordersRepo.FetchUserOrders(1)
	assert.Nil(t, errFetchUserOrders)
	assert.Equal(t, len(orders), len(ordersMock))
	assert.Equal(t, orders[0].ID, 1)

	ordersRepo.On(
		"UpdateOrderShippingStatus",
		mock.AnythingOfType("int"),
		mock.Anything,
	).Return(nil)
	errUpdateOrder := ordersRepo.UpdateOrderShippingStatus(orders[0].ID, entity.Delivered)
	assert.Nil(t, errUpdateOrder)
}

func TestHandler_UserOrderFail(t *testing.T) {
	errFetchUserOrdersMock := errors.New("fetch user orders error")
	ordersRepo := new(MOCK_ORDERS.Repo)
	ordersRepo.On("FetchUserOrders", mock.AnythingOfType("int")).Return(nil, errFetchUserOrdersMock)
	orders, errFetchUserOrders := ordersRepo.FetchUserOrders(1)
	assert.Nil(t, orders)
	assert.NotNil(t, errFetchUserOrders)
	assert.Equal(t, errFetchUserOrders, errFetchUserOrdersMock)
}
