package handler

import (
	"coffee-online-cli/entity"
	MOCK_PRODUCTS "coffee-online-cli/repository/products/mocks"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_CreateProductSuccess(t *testing.T) {
	product := entity.Product{
		Name:  "Cappucino",
		Stock: 25,
		Price: 19000,
	}
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("CreateProduct", mock.Anything).Return(nil)

	err := productRepo.CreateProduct(product)
	assert.Nil(t, err)
}

func TestHandler_CreateProductFail(t *testing.T) {
	product := entity.Product{
		Name:  "Cappucino",
		Stock: 25,
		Price: 19000,
	}
	errMock := errors.New("Error creaating product")
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("CreateProduct", mock.Anything).Return(errMock)

	err := productRepo.CreateProduct(product)
	assert.NotNil(t, err)
	assert.Equal(t, err, errMock)
}


func TestHandler_PopularProductSuccess(t *testing.T) {
	popularMock := []entity.ProductPopular{
		{Name: "Product 1", TotalOrder: 100},
		{Name: "Product 2", TotalOrder: 50},
	}
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("PopularProduct").Return(popularMock,nil)
	popular, err := productRepo.PopularProduct()
	assert.Nil(t, err)
	if !reflect.DeepEqual(popular, popularMock) {
		t.Errorf("PopularProduct() = %v, want %v", popular, popularMock)
	}
}

func TestHandler_PopularProductFail(t *testing.T) {
	errMock := errors.New("error popular product")
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("PopularProduct").Return(nil, errMock)
	popular, err := productRepo.PopularProduct()
	assert.Nil(t, popular)
	assert.Equal(t, errMock, err)
}

func TestHandler_ProductRestockSuccess(t *testing.T) {
	restockMock := []entity.Product{
		{
			ID: 4,
			Name: "Product 1", 
			Stock: 100,
			Price: 3000,
		},
	}	
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts").Return(restockMock, nil)
	products, err := productRepo.FetchProducts()
	assert.Nil(t, err)
	assert.Equal(t, products[0], restockMock[0])

	productRepo.On("ProductStockUpdate", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
	newStock := 11
	err = productRepo.ProductStockUpdate(products[0].ID, products[0].Stock+newStock)
	assert.Nil(t, err)
}

func TestHandler_ProductRestockFail(t *testing.T) {
	errMock := errors.New("error products")
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts").Return(nil, errMock)
	products, err := productRepo.FetchProducts()
	assert.Nil(t, products)
	assert.Equal(t, errMock, err)
}

func TestHandler_ProductStockSuccess (t *testing.T) {
	restockMock := []entity.Product{
		{
			ID: 4,
			Name: "Product 1", 
			Stock: 100,
		},
	}	
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts").Return(restockMock, nil)
	products, err := productRepo.FetchProducts()
	assert.Nil(t, err)
	assert.Equal(t, products[0], restockMock[0])
}
