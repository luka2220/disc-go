package github

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/luka2220/discGo/models"
	"github.com/stretchr/testify/assert"
)

func createTestGithubUser(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	testResponse := &models.GithubUser{
		Name:      "Luka Piplica",
		Url:       "https://github.com/luka2220",
		AvatarURL: "https://avatars.githubusercontent.com/u/42144047?v=4",
		Followers: 14,
		Bio:       "Computer Science student at Trent University 💻 Contact me at piplicaluka64@gmail.com",
		Email:     "piplicaluka64@gmail.com",
	}

	testResponseBytes, err := json.Marshal(testResponse)
	if err != nil {
		log.Fatalf("Error marshalling test response data: %v", err)
	}

	_, err = w.Write(testResponseBytes)
	if err != nil {
		log.Fatalf("Error responding to github service client: %v", err)
	}
}

func startTestServer() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/", createTestGithubUser)

	if err := http.ListenAndServe(":42069", nil); err != nil {
		log.Fatalf("Error starting http test server: %v", err)
	}
}

func TestGithubUser(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go startTestServer()

	testInterface := &GithubUserService{
		URL: "http://127.0.0.1:42069/",
	}

	githubTestingService := NewGithubUserService("", *testInterface)
	testUserResponse, err := githubTestingService.FetchGithubUser()

	testResponseExcpected := &models.GithubUser{
		Name:      "Luka Piplica",
		Url:       "http://127.0.0.1:42069/",
		AvatarURL: "https://avatars.githubusercontent.com/u/42144047?v=4",
		Followers: 14,
		Bio:       "Computer Science student at Trent University 💻 Contact me at piplicaluka64@gmail.com",
		Email:     "piplicaluka64@gmail.com",
	}

	if err != nil {
		t.Fatalf("Error fetching test user: %v", err)
	}

	t1 := assert.Equal(t, testResponseExcpected, testUserResponse)
	if t1 == true {
		t.Log(*testUserResponse)
	}
}

/*
	TODO:
	- Test the github API
	- Use a sample username as base case
	- Fetch the API in the TestGithubAPI rather than in the Github Service
	- Useful for checking if the api is working/ the github service interface is correct
*/

func TestGithubAPI(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}
