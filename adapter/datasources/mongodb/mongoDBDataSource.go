package mongodb

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-netflix-hexagonal-arch/bootstrap/config"
	"golang-netflix-hexagonal-arch/internal/entities"
	"golang-netflix-hexagonal-arch/adapter/datasources/mongodb/models"
)

type MongoDBDataSource struct {  
	projectsDatabase *mongo.Database
	ctx context.Context
}

func New(dbConfig *config.MongoDBConfig) *MongoDBDataSource {
	return &MongoDBDataSource{
		projectsDatabase: dbConfig.NewMongoDBClient().Database("projects"),
		ctx: dbConfig.Context(),
	}
}

func (datasource MongoDBDataSource) GetAllSoftwareByAuthor(author entities.Author) ([]entities.Software, error) {
	return nil, nil;
}

func (datasource MongoDBDataSource) RegisterSoftwareToAuthor(author entities.Author, softwares []entities.Software) error {
	collection := datasource.projectsDatabase.Collection("authors-softwares-golang")

	softwareModels := []models.SoftwareModel{}
	authorModel := models.AuthorModel{}
	for _, software := range softwares {
		authorModel = models.AuthorModel{
			UserName: software.Author.UserName,
			Name: software.Author.Name,
			AvatarUrl: software.Author.AvatarUrl,
		}

		softwareModels = append(softwareModels, models.SoftwareModel{
			Name: software.Name,
			Description: software.Description,
			Score: software.Score,
			Host: int(software.Host),
		})
	}

	_, err := collection.InsertOne(datasource.ctx, models.SoftwaresByAuthorModel{
		Author: &authorModel,
		Softwares: &softwareModels,
	});
	
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

	return nil
}