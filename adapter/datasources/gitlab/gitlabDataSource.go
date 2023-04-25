package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"golang-netflix-hexagonal-arch/internal/entities"
	"golang-netflix-hexagonal-arch/adapter/datasources/gitlab/services/data/response"
)

const gitlabURL = "https://gitlab.com";
const gitlabUserProjectsURIPath = gitlabURL + "/api/v4/users/%s/projects";

type GitlabDataSource struct {  

}

func New() *GitlabDataSource {
	return &GitlabDataSource{}
}

func (datasource GitlabDataSource) GetAllSoftwareByAuthor(author entities.Author) ([]entities.Software, error) {
	requestURL := fmt.Sprintf(gitlabUserProjectsURIPath, author.UserName)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}

	var results []response.GitlabUserProjectsResponse
	if err := json.Unmarshal(resBody, &results); err != nil {
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	softwares := []entities.Software{}
	for _, gitlabProject := range results {
		softwares = append(softwares, entities.Software {
			Name: gitlabProject.Name,
			Description: gitlabProject.Description,
			Score: gitlabProject.StarCount,
			Host: entities.GITLAB_HOST,
			Author: &entities.Author{
				UserName: gitlabProject.Namespace.Path,
				Name: gitlabProject.Namespace.Name,
				AvatarUrl: gitlabURL + gitlabProject.Namespace.AvatarURL,
			},
		})
	}

	fmt.Println("Gitlab Projects:\n" + PrettyPrint(softwares))

	return softwares, nil;
}

func (datasource GitlabDataSource) RegisterSoftwareToAuthor(author entities.Author, softwares []entities.Software) error {
	return nil;
}

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}