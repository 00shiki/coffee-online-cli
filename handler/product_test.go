package handler

import (
	"coffee-online-cli/entity"
	MOCK_PRODUCTS "coffee-online-cli/repository/products/mocks"
	"errors"
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
	popularProductsMock := []entity.ProductPopular{
		{Name: "Espresso", Total: 50},
		{Name: "Latte", Total: 30},
	}
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("PopularProduct", mock.Anything).Return(popularProductsMock, nil)

	h := handler.NewHandler(productRepo)
	h.PopularProduct()

	productRepo.AssertExpectations(t)
}

func TestHandler_PopularProductFail(t *testing.T) {
	errMock := errors.New("Error fetching popular products")
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("PopularProduct", mock.Anything).Return(nil, errMock)

	h := handler.NewHandler(productRepo)
	h.PopularProduct()

	productRepo.AssertExpectations(t)
}

func TestHandler_ProductRestockSuccess(t *testing.T) {
	mockProducts := []entity.Product{
		{ID: 1, Name: "Espresso", Stock: 10, Price: 15000},
		{ID: 2, Name: "Latte", Stock: 5, Price: 20000},
	}
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts", mock.Anything).Return(mockProducts, nil)
	productRepo.On("ProductStockUpdate", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)

	h := handler.NewHandler(productRepo)

	// Mock user input
	input := "1\n10\n3\n"
	mockInput(input)

	h.ProductRestock()

	productRepo.AssertExpectations(t)
}

func TestHandler_ProductRestockFailFetchProducts(t *testing.T) {
	errMock := errors.New("Error fetching products")
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts", mock.Anything).Return(nil, errMock)

	h := handler.NewHandler(productRepo)

	// Mock user input
	input := "1\n10\n3\n"
	mockInput(input)

	h.ProductRestock()

	productRepo.AssertExpectations(t)
}

func TestHandler_ProductRestockFailProductStockUpdate(t *testing.T) {
	mockProducts := []entity.Product{
		{ID: 1, Name: "Espresso", Stock: 10, Price: 15000},
		{ID: 2, Name: "Latte", Stock: 5, Price: 20000},
	}
	errMock := errors.New("Error updating product stock")
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts", mock.Anything).Return(mockProducts, nil)
	productRepo.On("ProductStockUpdate", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errMock)

	h := handler.NewHandler(productRepo)

	// Mock user input
	input := "1\n10\n3\n"
	mockInput(input)

	h.ProductRestock()

	productRepo.AssertExpectations(t)
}

func TestHandler_ProductStockSuccess(t *testing.T) {
	mockProducts := []entity.Product{
		{ID: 1, Name: "Espresso", Stock: 10},
		{ID: 2, Name: "Latte", Stock: 5},
	}
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts", mock.Anything).Return(mockProducts, nil)

	h := handler.NewHandler(productRepo)
	h.ProductStock()

	productRepo.AssertExpectations(t)
}

func TestHandler_ProductStockFail(t *testing.T) {
	errMock := errors.New("Error fetching products")
	productRepo := new(MOCK_PRODUCTS.Repo)
	productRepo.On("FetchProducts", mock.Anything).Return(nil, errMock)

	h := handler.NewHandler(productRepo)
	h.ProductStock()

	productRepo.AssertExpectations(t)
}