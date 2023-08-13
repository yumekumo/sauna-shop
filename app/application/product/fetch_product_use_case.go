package product

import (
	"context"
)

type FetchProductUseCase struct {
	FetchProductQueryService FetchProductQueryService
}

func NewFetchProductUseCase(
	fetchProductQueryService FetchProductQueryService,
) *FetchProductUseCase {
	return &FetchProductUseCase{
		FetchProductQueryService: fetchProductQueryService,
	}
}

type FetchProductUseCaseDto struct {
	ID          string
	Name        string
	Description string
	Price       int64
	Stock       int
	OwnerID     string
	OwnerName   string
}

func (uc *FetchProductUseCase) Run(ctx context.Context) ([]*FetchProductUseCaseDto, error) {
	qsDtos, err := uc.FetchProductQueryService.Run(ctx)
	var ucDtos []*FetchProductUseCaseDto

	for _, qsDto := range qsDtos {
		ucDtos = append(ucDtos, &FetchProductUseCaseDto{
			ID:          qsDto.ID,
			Name:        qsDto.Name,
			Description: qsDto.Description,
			Price:       qsDto.Price,
			Stock:       qsDto.Stock,
			OwnerID:     qsDto.OwnerID,
			OwnerName:   qsDto.OwnerName,
		})
	}
	return ucDtos, err
}
