package github

import (
	"fmt"
	"sync"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"golang-netflix-hexagonal-arch/internal/entities"
	"golang-netflix-hexagonal-arch/adapter/datasources/github/services/data/response"
)

const githubAPIUrl = "https://api.github.com";
const githubUserURIPath = githubAPIUrl + "/users/%s";
const githubUserProjectsURIPath = githubUserURIPath + "/repos";

type GithubDataSource struct {  

}

func New() *GithubDataSource {
	return &GithubDataSource{}
}

func (datasource GithubDataSource) GetAllSoftwareByAuthor(author entities.Author) ([]entities.Software, error) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	githubUserDetail := response.GithubUserDetailResponse{}
	go func() error {
		githubUserDetail = getUserDetail(author.UserName)
		defer waitGroup.Done()

		return nil
	} ()

	requestURL := fmt.Sprintf(githubUserProjectsURIPath, author.UserName)

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

	var results []response.GithubUserProjectsResponse
	if err := json.Unmarshal(resBody, &results); err != nil {
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	waitGroup.Wait()

	softwares := []entities.Software{}
	for _, githubProject := range results {
		softwares = append(softwares, entities.Software {
			Name: githubProject.Name,
			Description: githubProject.Description,
			Score: githubProject.StargazersCount,
			Host: entities.GITHUB_HOST,
			Author: &entities.Author {
				UserName: githubUserDetail.Login,
				Name: githubUserDetail.Name,
				AvatarUrl: githubUserDetail.AvatarURL,
			},
		})
	}

	fmt.Println("Github Projects:\n" + PrettyPrint(softwares))

	return softwares, nil;
}

func (datasource GithubDataSource) RegisterSoftwareToAuthor(author entities.Author, softwares []entities.Software) error {
	return nil;
}

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

func getUserDetail(username string) (response.GithubUserDetailResponse) {
	requestURL := fmt.Sprintf(githubUserURIPath, username)

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

	var result response.GithubUserDetailResponse
	if err := json.Unmarshal(resBody, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	return result
}