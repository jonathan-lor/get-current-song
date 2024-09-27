package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"get-current-song/internal/spotify"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading from .env")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	token, err := spotify.GetSpotifyToken(clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
}
