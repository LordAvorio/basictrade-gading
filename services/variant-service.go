package services

import (
	"basictrade-gading/models"
	"basictrade-gading/repositories"
	"basictrade-gading/utils"
	"basictrade-gading/utils/helpers"
)

type VariantService struct {
	variantRepo repositories.IVariantRepository
	productRepo repositories.IProductRepository
}

type IVariantService interface {
	CreateVariant(*models.VariantRequest) (*models.Variant, error)
	GetVariant(string) (*models.Variant, error)
	GetVariants(int, int, string) (*models.Pagination, error)
}

func NewVariantService(variantRepo repositories.IVariantRepository, productRepo repositories.IProductRepository) *VariantService {
	variantService := VariantService{}
	variantService.variantRepo = variantRepo
	variantService.productRepo = productRepo
	return &variantService
}

func (s *VariantService) CreateVariant(dataRequest *models.VariantRequest) (*models.Variant, error) {

	generateKsuid, err := utils.GenerateKSUID()
	if err != nil {
		return nil, err
	}

	dataProduct, err := s.productRepo.GetProduct(dataRequest.UUID)
	if err != nil {
		return nil, err
	}

	dataVariant := models.Variant{
		UUID:        generateKsuid,
		VariantName: dataRequest.VariantName,
		Quantity:    dataRequest.Quantity,
		ProductID:   dataProduct.ID,
	}

	result, err := s.variantRepo.CreateVariant(&dataVariant)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *VariantService) GetVariant(uuid string) (*models.Variant, error) {

	resultVariant, err := s.variantRepo.GetVariant(uuid)
	if err != nil {
		return nil, err
	}

	return resultVariant, nil
}

func (s *VariantService) GetVariants(limit, offset int, nameFilter string) (*models.Pagination, error) {

	dataVariants, err := s.variantRepo.GetVariants(offset, limit, nameFilter)
	if err != nil {
		return nil, err
	}

	totalVariants, err := s.variantRepo.GetTotalVariant(nameFilter)
	if err != nil {
		return nil, err
	}

	totalPage := helpers.GetTotalPages(totalVariants, limit)

	resultVariants := []models.VariantResponse{}
	for _, value := range *dataVariants {
		dataVariant := models.VariantResponse {
			UUID: value.UUID,
			VariantName: value.VariantName,
			Quantity: value.Quantity,
			ProductID: value.ProductID,
		}

		resultVariants = append(resultVariants, dataVariant)
	}


	result := models.Pagination{
		Data:         resultVariants,
		TotalPage:    totalPage,
		NextPage:     helpers.GetNextPage(offset, limit, totalVariants),
		PreviousPage: helpers.GetPrevPage(offset, limit),
		CurrentPage:  helpers.GetCurrentPage(offset, limit),
	}

	return &result, nil

}
