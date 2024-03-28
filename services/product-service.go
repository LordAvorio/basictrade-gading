package services

import (
	"basictrade-gading/models"
	"basictrade-gading/repositories"
	"basictrade-gading/utils"
	"basictrade-gading/utils/helpers"
)

type ProductService struct {
	productRepo repositories.IProductRepository
}

type IProductService interface {
	CreateProduct(*models.ProductRequest) (*models.Product, error)
	GetProduct(string) (*models.Product, error)
	GetProducts(int, int, string) (*models.Pagination, error)
}

func NewProductService(productRepo repositories.IProductRepository) *ProductService {
	productService := ProductService{}
	productService.productRepo = productRepo
	return &productService
}

func (s *ProductService) CreateProduct(dataRequest *models.ProductRequest) (*models.Product, error) {

	fileName := helpers.RemoveExtension(dataRequest.Image.Filename)

	uploadResult, err := helpers.UploadFile(&dataRequest.Image, fileName)
	if err != nil {
		return nil, err
	}

	generateKsuid, err := utils.GenerateKSUID()
	if err != nil {
		return nil, err
	}

	data := models.Product{
		UUID:     generateKsuid,
		Name:     dataRequest.Name,
		AdminID:  dataRequest.AdminId,
		ImageUrl: uploadResult,
	}

	resultProduct, err := s.productRepo.CreateProduct(&data)
	if err != nil {
		return nil, err
	}

	return resultProduct, nil

}

func (s *ProductService) GetProduct(uuid string) (*models.Product, error) {

	resultProduct, err := s.productRepo.GetProduct(uuid)
	if err != nil {
		return nil, err
	}

	return resultProduct, nil
}

func (s *ProductService) GetProducts(limit, offset int, nameFilter string) (*models.Pagination, error) {

	dataProducts, err := s.productRepo.GetProducts(offset, limit, nameFilter)
	if err != nil {
		return nil, err
	}

	totalProducts, err := s.productRepo.GetTotalProduct(nameFilter)
	if err != nil {
		return nil, err
	}

	totalPage := helpers.GetTotalPages(totalProducts, limit)

	resultProducts := []models.ProductResponse{}
	for _, value := range *dataProducts {
		dataProduct := models.ProductResponse {
			UUID: value.UUID,
			Name: value.Name,
			ImageUrl: value.ImageUrl,
			AdminId: value.AdminID,
		}

		resultProducts = append(resultProducts, dataProduct)
	}


	result := models.Pagination{
		Data:         resultProducts,
		TotalPage:    totalPage,
		NextPage:     helpers.GetNextPage(offset, limit, totalProducts),
		PreviousPage: helpers.GetPrevPage(offset, limit),
		CurrentPage:  helpers.GetCurrentPage(offset, limit),
	}

	return &result, nil

}
