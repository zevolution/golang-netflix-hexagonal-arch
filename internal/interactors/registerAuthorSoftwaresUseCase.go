package interactors

import (
	"golang-netflix-hexagonal-arch/internal/repositories"
	"golang-netflix-hexagonal-arch/internal/entities"
)

type registerSoftwareByAuthorUseCase struct {
	softwareRepository repositories.SoftwareRepository
}

func NewRegisterSoftwareByAuthorUseCase(softwareRepository repositories.SoftwareRepository) *registerSoftwareByAuthorUseCase {
	return &registerSoftwareByAuthorUseCase{
		softwareRepository: softwareRepository,
	}
}

func (useCase *registerSoftwareByAuthorUseCase) Execute(author entities.Author, softwares []entities.Software) (error) {
	err := useCase.softwareRepository.RegisterSoftwareToAuthor(author, softwares);

	return err;
}