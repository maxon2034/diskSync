package gdrive

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func New(ctx context.Context, clientID, clientSecret, tokenPath string) (*Client, error) {
	client := Client{}
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{drive.DriveReadonlyScope},
		RedirectURL:  "http://localhost:8080",
	}

	_, err := os.Open(tokenPath)
	if err != nil {
		token, err := getToken(ctx, &config, tokenPath)
		if err != nil {
			return nil, fmt.Errorf("Error in getting token: %w", err)
		}
		httpClient := config.Client(ctx, token)
		client.Service, err = drive.NewService(ctx, option.WithHTTPClient(httpClient))
		if err != nil {
			return nil, fmt.Errorf("Error in init Google Drive Service")
		}
		return &client, nil
	} else {
		token, err := getTokenFromFile(tokenPath)
		if err != nil {
			return nil, fmt.Errorf("Error in getting token: %w", err)
		}
		httpClient := config.Client(ctx, token)
		client.Service, err = drive.NewService(ctx, option.WithHTTPClient(httpClient))
		if err != nil {
			return nil, fmt.Errorf("Error in init Google Drive Service")
		}
		return &client, nil
	}
}

func getToken(ctx context.Context, config *oauth2.Config, path string) (*oauth2.Token, error) {
	var code string
	authURL := config.AuthCodeURL("state-token")

	fmt.Println(authURL)
	fmt.Println("Register via the URL and send the code back")
	fmt.Scan(&code)

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Error in exchanging token: %w", err)
	}
	err = saveTokenToFile(path, token)
	if err != nil {
		return nil, fmt.Errorf("Error in writing in file: %w", err)
	}
	return token, nil
}

func getTokenFromFile(path string) (*oauth2.Token, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	if err := json.NewDecoder(f).Decode(tok); err != nil {
		return nil, err
	}
	return tok, nil
}
func saveTokenToFile(path string, token *oauth2.Token) error {
	fmt.Printf("Сохранение токена в файл: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(token)
}
