package interactors

import (
	"fmt"
	"sync"
	"encoding/json"
	"golang-netflix-hexagonal-arch/internal/repositories"
	"golang-netflix-hexagonal-arch/internal/entities"
)

type getSoftwareByAuthorUseCase struct {
	softwareRepositories []repositories.SoftwareRepository
}

func NewGetSoftwareByAuthorUseCase(softwareRepositories []repositories.SoftwareRepository) *getSoftwareByAuthorUseCase {
	return &getSoftwareByAuthorUseCase{
		softwareRepositories: softwareRepositories,
	}
}

func (useCase *getSoftwareByAuthorUseCase) Execute(author entities.Author) ([]entities.Software, error) {
	var waitGroup sync.WaitGroup

	softwares := []entities.Software{}
	for _, element := range useCase.softwareRepositories {
		repository := element

		waitGroup.Add(1)
		go func() error {
			
			results, err := repository.GetAllSoftwareByAuthor(author)
			if err != nil {
				return err
			}
			for _, software := range results {
				softwares = append(softwares, software)
			}
			
			defer waitGroup.Done()

			return nil
		} ()
	}

	waitGroup.Wait()

	fmt.Println("Softwares from all repos:\n" + PrettyPrint(softwares))
	return softwares, nil
}

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}