package usecase

import (
	domain "backend-test/domain/book"
	log "backend-test/helper/log"
)

type BookUsecase struct {
	infra domain.BookInfraService
}

func NewBookUsecase(infra domain.BookInfraService) *BookUsecase {
	return &BookUsecase{infra}
}

func (book *BookUsecase) Search(key string) (interface{}, error) {

	lists, err := book.infra.List(key)

	if err != nil {
		defer log.CreateLogResponse(&log.FormatLog{
			Event: "usecase|infra|search|book",
			Error: err.Error(),
		})

		return nil, err
	}

	return lists, nil
}
