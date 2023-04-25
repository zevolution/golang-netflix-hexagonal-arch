package event

import (
	"fmt"
	"golang-netflix-hexagonal-arch/internal/entities"
)

type GetSoftwareByAuthorUseCase interface {
	Execute(author entities.Author) ([]entities.Software, error)
}

type RegisterAuthorSoftwaresUseCase interface {
	Execute(author entities.Author, softwares []entities.Software) (error)
}

type StartupApplicationEvent struct {
	getSoftwareByAuthorUseCase GetSoftwareByAuthorUseCase
	registerAuthorSoftwaresUseCase RegisterAuthorSoftwaresUseCase
}

func New(getSoftwareByAuthorUseCase GetSoftwareByAuthorUseCase, registerAuthorSoftwaresUseCase RegisterAuthorSoftwaresUseCase) *StartupApplicationEvent {
	return &StartupApplicationEvent{
		getSoftwareByAuthorUseCase: getSoftwareByAuthorUseCase,
		registerAuthorSoftwaresUseCase: registerAuthorSoftwaresUseCase,
	}
}

func (event *StartupApplicationEvent) OnStartupApplication() error {
	author := entities.Author {
		UserName: "zevolution",
	}

	softwares, err := event.getSoftwareByAuthorUseCase.Execute(author)
	if err != nil {
		return err
	}

	err = event.registerAuthorSoftwaresUseCase.Execute(author, softwares)
	if err != nil {
		fmt.Printf("Could not create record in mongodb: %s\n", err)
		return err
	}

	return nil;
}
