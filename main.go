package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	owner      = flag.String("owner", "", "Owner of repository")
	repository = flag.String("repository", "", "Name of repository")
	tag        = flag.String("tag", "", "Release tag")
	token      = flag.String("token", "", "Auth token, overrides env var GITHUB_AUTH_TOKEN")
)

func main() {
	// Use custom usage function
	flag.Usage = func() {
		fmt.Println("go-get-release fetches all assets from a Github release tag.")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	// Parse options and environment
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	// if token == "" {
	// 	log.Fatal("No token given")
	// }
	if *owner == "" {
		log.Fatal("No owner given")
	}
	if *repository == "" {
		log.Fatal("No tag given")
	}
	if *tag == "" {
		log.Fatal("No token given")
	}

	// Only authenticate when necessary
	client := github.NewClient(nil)
	ctx := context.Background()
	// Authenticate
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	}

	// Get release
	release, _, err := client.Repositories.GetReleaseByTag(ctx, *owner, *repository, *tag)
	if err != nil {
		log.Fatal("Error getting release by tag.", tag, err)
	}

	// Get release assets
	var opt github.ListOptions
	assets, _, err := client.Repositories.ListReleaseAssets(ctx, *owner, *repository, *release.ID, &opt)

	for _, asset := range assets {
		// Create the local file
		outFile, err := os.Create(*asset.Name)
		if err != nil {
			log.Println("Error creating file: ", asset.Name, err)
			continue
		}
		defer outFile.Close()

		// Download the asset
		content, redirectURL, err := client.Repositories.DownloadReleaseAsset(ctx, *owner, *repository, *asset.ID)
		if err != nil {
			log.Println("Error while downloading asset: ", asset.Name, err)
			continue
		}

		// Check for redirect
		if redirectURL != "" {
			// If we received a redirect we should overwrite our content with the body
			// of the redirect URL.
			response, err := http.Get(redirectURL)
			if err != nil {
				log.Println("Failed to download content from redirect URL: ", redirectURL)
				continue
			}
			content = response.Body
		}
		defer content.Close()

		// Write contents
		_, err = io.Copy(outFile, content)
		if err != nil {
			log.Println("Error while writing to file: ", asset.Name, err)
		}
	}
}
