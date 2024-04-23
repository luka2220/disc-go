package github

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/luka2220/discGo/models"
)

type GithubUserService struct {
	URL string
}

func NewGithubUserService(username string, params ...GithubUserService) *GithubUserService {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	if len(params) > 0 {
		url = params[0].URL
	}

	return &GithubUserService{
		url,
	}
}

func (g *GithubUserService) FetchGithubUser() (*models.GithubUser, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	resp, err := http.Get(g.URL)
	if err != nil {
		log.Fatalf("Error fetching github user data %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	var githubDataBytes []byte
	var githubData *models.GithubUser

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		githubDataBytes = scanner.Bytes()
	}

	if err := json.Unmarshal(githubDataBytes, &githubData); err != nil {
		log.Fatalf("Error occured un-marshalling github data: %s", err)
		return nil, err
	}

	githubData.Url = fmt.Sprintf(g.URL)

	return githubData, nil
}
