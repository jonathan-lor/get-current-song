package spotify

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetSpotifyToken(clientID, clientSecret string) (string, error) {
	spotifyTokenURL := "https://accounts.spotify.com/api/token"

	// authHeader := base64.StdEncoding.EncodeToString([]byte("grant_type=client_credentials&client_id=" + clientID + "&client_secret=" + clientSecret))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", spotifyTokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

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
