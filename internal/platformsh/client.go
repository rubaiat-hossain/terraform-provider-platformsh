package platformsh

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	restyClient *resty.Client
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type Project struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProjectsResponse struct {
	Projects []Project `json:"projects"`
}

func NewClient(apiToken string) (*Client, error) {
	client := resty.New()
	tokenResp, err := client.R().
		SetBasicAuth("platform-api-user", "").
		SetFormData(map[string]string{
			"grant_type": "api_token",
			"api_token":  apiToken,
		}).
		SetResult(&TokenResponse{}).
		Post("https://auth.api.platform.sh/oauth2/token")

	if err != nil {
		return nil, err
	}

	tokenResponse := tokenResp.Result().(*TokenResponse)
	client.SetAuthToken(tokenResponse.AccessToken)

	return &Client{restyClient: client}, nil
}

func (c *Client) GetSession() *resty.Client {
	return c.restyClient
}

func (c *Client) GetProjects() ([]Project, error) {
	var projectsResponse ProjectsResponse
	_, err := c.restyClient.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.restyClient.Token)).
		SetResult(&projectsResponse).
		Get("https://api.platform.sh/projects")

	if err != nil {
		return nil, err
	}

	return projectsResponse.Projects, nil
}
