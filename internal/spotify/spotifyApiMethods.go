package spotify

import (
	"io"
	"log"
	"net/http"
)

func GetCurrentlyListening(bearerToken string) (string, error) {
	endpoint := "https://api.spotify.com/v1/me/player/currently-playing"
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	scope := "user-read-currently-playing"
	req.Header.Set("Authorization", "Bearer " + bearerToken)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
