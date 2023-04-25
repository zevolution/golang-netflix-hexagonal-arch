package main

import (
	"fmt"
	"time"
	"golang-netflix-hexagonal-arch/bootstrap/config"
	"golang-netflix-hexagonal-arch/internal/repositories"
	"golang-netflix-hexagonal-arch/internal/interactors"
	"golang-netflix-hexagonal-arch/adapter/datasources/gitlab"
	"golang-netflix-hexagonal-arch/adapter/datasources/github"
	"golang-netflix-hexagonal-arch/adapter/datasources/mongodb"
	"golang-netflix-hexagonal-arch/adapter/transportlayers/event"
)

func main() {
	repositories := []repositories.SoftwareRepository{}
	repositories = append(repositories, gitlab.New(), github.New())
	getSoftwareByAuthorUseCase := interactors.NewGetSoftwareByAuthorUseCase(repositories)

	mongoDBDataSource := mongodb.New(config.NewMongoDBConfig())
	registerSoftwareByAuthorUseCase := interactors.NewRegisterSoftwareByAuthorUseCase(mongoDBDataSource)

	event := event.New(
		getSoftwareByAuthorUseCase,
		registerSoftwareByAuthorUseCase,
	)

	started := time.Now()
	event.OnStartupApplication()
	elapsed := time.Since(started)

    fmt.Printf("Get author repositories from Gitlab, Github and Bitbucket to save into MongoDB took %s\n", elapsed)
}