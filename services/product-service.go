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
	UpdateProduct(string, *models.ProductUpdateRequest) (*models.Product, error)
	DeleteProduct(string) error
}

func NewProductService(productRepo repositories.IProductRepository) *ProductService {
	productService := ProductService{}
	productService.productRepo = productRepo
	return &productService
}

func (s *ProductService) CreateProduct(dataRequest *models.ProductRequest) (*models.Product, error) {

	fileName, err := helpers.GenerateFileNameImage()
	if err != nil {
		return nil, err
	}

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

	resultProducts := []models.ProductResponse{}
	for _, value := range *dataProducts {

		listVariants := []models.VariantResponse{}

		for _, variant := range value.Variants {
			dataVariant := models.VariantResponse{
				UUID:        variant.UUID,
				VariantName: variant.VariantName,
				Quantity:    variant.Quantity,
				ProductID:   variant.ProductID,
			}
			listVariants = append(listVariants, dataVariant)
		}

		dataProduct := models.ProductResponse{
			UUID:     value.UUID,
			Name:     value.Name,
			ImageUrl: value.ImageUrl,
			AdminId:  value.AdminID,
			Variants: listVariants,
		}

		resultProducts = append(resultProducts, dataProduct)
	}

	result := models.Pagination{
		Data:       resultProducts,
		Offset:     offset,
		Limit:      limit,
		Total:      totalProducts,
		PrevOffset: helpers.GetPreviousOffset(offset, limit),
		NextOffset: helpers.GetNextOffset(offset, limit, totalProducts),
	}

	return &result, nil
}

func (s *ProductService) UpdateProduct(uuid string, dataRequest *models.ProductUpdateRequest) (*models.Product, error) {

	dataProduct, err := s.productRepo.GetProduct(uuid)
	if err != nil {
		return nil, err
	}

	dataProduct.Name = dataRequest.Name

	if dataRequest.Image.Filename != "" {

		errDeleteCloudinary := helpers.DeleteFile(dataProduct.ImageUrl)
		if errDeleteCloudinary != nil {
			return nil, err
		}

		fileName, err := helpers.GenerateFileNameImage()
		if err != nil {
			return nil, err
		}

		uploadResult, err := helpers.UploadFile(&dataRequest.Image, fileName)
		if err != nil {
			return nil, err
		}
		dataProduct.ImageUrl = uploadResult
	}

	resultProduct, err := s.productRepo.UpdateProduct(dataProduct)
	if err != nil {
		return nil, err
	}

	return resultProduct, nil
}

func (s *ProductService) DeleteProduct(uuid string) error {

	dataProduct, err := s.productRepo.GetProduct(uuid)
	if err != nil {
		return err
	}

	err = helpers.DeleteFile(dataProduct.ImageUrl)
	if err != nil {
		return err
	}

	err = s.productRepo.DeleteProduct(uuid)
	if err != nil {
		return err
	}

	return nil
}
