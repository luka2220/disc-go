package github

import (
	"bufio"
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
	if t1 == false {
		t.Fatal("Unexcpeted values... Results not equal")
	}

	t.Log(*testUserResponse)
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

	res, err := http.Get("https://api.github.com/users/luka2220")

	if err != nil {
		t.Fatalf("An error occured requesting github api URL = %v", err)
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)

	var responseBytes []byte
	var jsonResponse models.GithubUser

	// NOTE:: This test will fail if the user changes any profile info from below
	excpectedResult := &models.GithubUser{
		Name:      "Luka Piplica",
		AvatarURL: "https://avatars.githubusercontent.com/u/42144047?v=4",
		Url:       "https://api.github.com/users/luka2220",
		Blog:      "https://luka2220.github.io/cv/",
		Email:     "",
		Bio:       "Computer Science student at Trent University 💻\r\nContact me at piplicaluka64@gmail.com ",
		Followers: 14,
	}

	for scanner.Scan() {
		responseBytes = scanner.Bytes()
	}

	err = json.Unmarshal(responseBytes, &jsonResponse)
	if err != nil {
		t.Fatalf("Error unpacking bytes from json response = %v", err)
	}

	tc1 := assert.Equal(t, excpectedResult, &jsonResponse)
	if tc1 == false {
		t.Fatal("Unexcpected values... Result are not equal")
	}

	t.Logf("Result = %v", jsonResponse)
}
