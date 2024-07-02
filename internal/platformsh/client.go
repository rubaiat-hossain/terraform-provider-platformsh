package platformsh

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty *resty.Client
}

func NewClient(apiToken string) (*Client, error) {
	client := resty.New()

	resp, err := client.R().
		SetBasicAuth("platform-api-user", "").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type": "api_token",
			"api_token":  apiToken,
		}).
		Post("https://auth.api.platform.sh/oauth2/token")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get token: %s", resp.Status())
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return nil, fmt.Errorf("access token not found in response")
	}

	return &Client{
		resty: client.SetAuthToken("Bearer " + token),
	}, nil
}

