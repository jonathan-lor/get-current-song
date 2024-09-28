package spotify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetSpotifyToken(clientID, clientSecret string) (string, error) {
	spotifyTokenURL := "https://accounts.spotify.com/api/token"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", spotifyTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var authResponse AuthResponse
	err = json.NewDecoder(res.Body).Decode(&authResponse)
	if err != nil {
		return "", err
	}
	fmt.Println(authResponse.ExpiresIn)
	return authResponse.AccessToken, nil
}
