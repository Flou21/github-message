package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	log.Println("writing to github pull request")
	repositoryName := readEnvVar("REPOSITORY")
	username := readEnvVar("USERNAME")
	token := readEnvVar("TOKEN")
	msg := parseMessage()
	newState := readEnvVar("NEW_STATE")

	if newState != "opened" {
		log.Println("state is not created, going to stop now")
		log.Println(newState)
		return
	}

	pullRequestNumber, err := strconv.Atoi(readEnvVar("PULL_REQUEST_NUMBER"))
	if err != nil {
		log.Fatalln("could not parse pull request number")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pullRequest, _, err := client.PullRequests.Get(ctx, username, repositoryName, pullRequestNumber)

	if err != nil {
		log.Fatalln("could not get pull request ,some values: ", username, repositoryName, pullRequestNumber)
	}

	githubUser, _, err := client.Users.Get(ctx, "")

	if err != nil {
		log.Fatalln("could not get user ", "")
	}

	url := pullRequest.GetHTMLURL()
	log.Println("pull request url: ", url)
	log.Println("username: ", username)
	log.Println("repositoryName: ", repositoryName)
	log.Println("pull request number: ", strconv.Itoa(pullRequestNumber))

	_, githubResponse, err := client.Issues.CreateComment(ctx, username, repositoryName, pullRequestNumber, &github.IssueComment{
		Body: &msg,
		User: githubUser,
		URL:  pullRequest.URL,
	})

	if err != nil {
		log.Fatalln("could not write pull request comment", err, githubResponse)
	}

	log.Println("finshed writing comment")

}

func readEnvVar(key string) string {
	val := os.Getenv(key)

	if val == "" {
		log.Fatalln(key + " is not given")
	}

	return val
}

func parseMessage() string {
	index := 1
	buffer := []string{}
	buffer = append(buffer, "[CI] ")
	for {
		key := "MESSAGE_" + strconv.Itoa(index)

		val := os.Getenv(key)

		if val == "" {
			return strings.Join(buffer, "")
		}

		buffer = append(buffer, val)

		index++
	}
}
