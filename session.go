package iland

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type TokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	GrantType    string `json:"grant_type"`
}

type RefreshTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type APIError struct {
	Error         string `json:"error"`
	Message       string `json:"message"`
	DetailMessage string `json:"detail_message"`
}

func (c *Client) getToken() error {
	tokenRequest := TokenRequest{c.clientID, c.clientSecret, c.username, c.password, "password"}
	form := url.Values{}
	form.Add("client_id", tokenRequest.ClientID)
	form.Add("client_secret", tokenRequest.ClientSecret)
	form.Add("username", tokenRequest.Username)
	form.Add("password", tokenRequest.Password)
	form.Add("grant_type", tokenRequest.GrantType)
	resp, err := http.Post(accessURL, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Could not retrieve a token.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}
	c.Token = t
	c.setTokenExpiration()
	return nil
}

func (c *Client) refreshToken() error {
	tokenRequest := RefreshTokenRequest{c.clientID, c.clientSecret, c.Token.RefreshToken, "refresh_token"}
	form := url.Values{}
	form.Add("client_id", tokenRequest.ClientID)
	form.Add("client_secret", tokenRequest.ClientSecret)
	form.Add("refresh_token", tokenRequest.RefreshToken)
	form.Add("grant_type", tokenRequest.GrantType)
	resp, err := http.Post(refreshURL, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("Could not refresh current token.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var t Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		return err
	}
	c.Token = t
	c.setTokenExpiration()
	return nil
}

func (c *Client) setTokenExpiration() {
	c.tokenExpiration = time.Now().Add(time.Duration(c.Token.ExpiresIn-10) * time.Second)
}

func removeJSONHijackingPrefix(b []byte) []byte {
	return bytes.TrimPrefix(b, []byte(")]}'"))
}

func (c *Client) request(relPath, verb string, payload []byte) ([]byte, error) {
	err := c.RefreshTokenIfNecessary()
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{}
	path := apiBaseURL + relPath
	bytesJSON := bytes.NewBuffer(payload)
	req, err := http.NewRequest(verb, path, bytesJSON)
	req.Header.Add("Authorization", "Bearer "+c.Token.AccessToken)
	req.Header.Add("Accept", "application/vnd.ilandcloud.api.v0.8+json")
	req.Header.Add("Content-Type", "application/vnd.ilandcloud.api.v0.8+json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	statusCode := resp.StatusCode
	responseBody := removeJSONHijackingPrefix(body)
	if statusCode >= 300 {

		var e APIError
		err = json.Unmarshal(responseBody, &e)
		if err != nil {
			return []byte{}, fmt.Errorf("Error marshaling ApiError: %s", err.Error())
		}
		var errMsg string
		if e.DetailMessage != "" {
			errMsg = e.DetailMessage
		} else {
			errMsg = e.Message
		}
		return responseBody, errors.New(errMsg)
	}

	return responseBody, nil
}

func (c *Client) getBinary(relPath string) ([]byte, error) {
	err := c.RefreshTokenIfNecessary()
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{}
	path := apiBaseURL + relPath
	req, err := http.NewRequest("GET", path, nil)
	req.Header.Add("Authorization", "Bearer "+c.Token.AccessToken)
	req.Header.Add("Content-Type", "application/vnd.ilandcloud.api.v0.8+json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	statusCode := resp.StatusCode
	responseBody := removeJSONHijackingPrefix(body)
	if statusCode >= 300 {

		var e APIError
		err = json.Unmarshal(responseBody, &e)
		if err != nil {
			return []byte{}, fmt.Errorf("Error marshaling ApiError: %s", err.Error())
		}
		var errMsg string
		if e.DetailMessage != "" {
			errMsg = e.DetailMessage
		} else {
			errMsg = e.Message
		}
		return responseBody, errors.New(errMsg)
	}

	return responseBody, nil
}

func (c *Client) postForm(relPath, contentType string, payload []byte) ([]byte, error) {
	err := c.RefreshTokenIfNecessary()
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{}
	path := apiBaseURL + relPath
	bytesJSON := bytes.NewBuffer(payload)
	req, err := http.NewRequest("POST", path, bytesJSON)
	req.Header.Add("Authorization", "Bearer "+c.Token.AccessToken)
	req.Header.Add("Accept", "application/vnd.ilandcloud.api.v0.8+json")
	req.Header.Add("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	statusCode := resp.StatusCode
	responseBody := removeJSONHijackingPrefix(body)
	if statusCode >= 300 {
		var e APIError
		err = json.Unmarshal(responseBody, &e)
		if err != nil {
			return []byte{}, fmt.Errorf("Error marshaling ApiError: %s", err.Error())
		}
		var errMsg string
		if e.DetailMessage != "" {
			errMsg = e.DetailMessage
		} else {
			errMsg = e.Message
		}
		return responseBody, errors.New(errMsg)
	}
	return responseBody, nil
}

func (c *Client) RefreshTokenIfNecessary() error {
	emptyToken := Token{}
	if c == nil || c.Token == emptyToken {
		err := c.getToken()
		if err != nil {
			return fmt.Errorf("Error retrieving iland cloud API token. %s", err.Error())
		}
	}
	if c.isTokenExpiringSoon() {
		err := c.refreshToken()
		if err != nil {
			err := c.getToken()
			if err != nil {
				return fmt.Errorf("Error refreshing iland cloud API token. %s", err.Error())
			}
		}
	}
	return nil
}

func (c *Client) isTokenExpiringSoon() bool {
	if time.Now().After(c.tokenExpiration) {
		return true
	}
	return false
}
