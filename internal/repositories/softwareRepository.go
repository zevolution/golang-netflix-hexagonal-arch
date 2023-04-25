package repositories

import "golang-netflix-hexagonal-arch/internal/entities"

type SoftwareRepository interface {
	GetAllSoftwareByAuthor(author entities.Author) ([]entities.Software, error)
	RegisterSoftwareToAuthor(author entities.Author, softwares []entities.Software) error
}