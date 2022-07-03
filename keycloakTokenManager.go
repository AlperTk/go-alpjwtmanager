package tokenmanager

import (
	"encoding/json"
	"errors"
	"github.com/Masterminds/log-go"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewTokenManager(openIdConnectTokenUrl string, clientId string, secretKey string) AlpTokenManager {
	return &tokenManagerKeycloak{
		openIdConnectTokenUrl: openIdConnectTokenUrl,
		clientId:              clientId,
		secretKey:             secretKey,
		responseModel:         &TokenResponseModel{},
		client:                &http.Client{},
	}
}

func NewTokenManagerWithCustomClient(openIdConnectTokenUrl string, clientId string, secretKey string, client *http.Client) AlpTokenManager {
	return &tokenManagerKeycloak{
		openIdConnectTokenUrl: openIdConnectTokenUrl,
		clientId:              clientId,
		secretKey:             secretKey,
		responseModel:         &TokenResponseModel{},
		client:                client,
	}
}

type tokenManagerKeycloak struct {
	openIdConnectTokenUrl string
	clientId              string
	secretKey             string
	responseModel         *TokenResponseModel
	client                *http.Client
}

var grantTypeCredentials = "client_credentials"
var grantTypeRefreshToken = "refresh_token"
var expireThreshold int64 = 30

func (t tokenManagerKeycloak) GetBearerToken() (string, error) {

	if t.responseModel.AccessToken == "" {
		err := t.getToken()
		if err != nil {
			log.Error("error while getting token. ", err)
			return "", err
		}
	} else if t.responseModel.ExpiredAt < (time.Now().Unix() + expireThreshold) {
		err := t.RefreshToken()
		if err != nil {
			log.Error("error while getting token. ", err)
			return "Bearer " + t.responseModel.AccessToken, err
		}
	}

	return "Bearer " + t.responseModel.AccessToken, nil
}

func (t tokenManagerKeycloak) RefreshToken() error {
	var err error
	if t.responseModel.RefreshToken == "" || t.responseModel.RefreshExpiredAt < (time.Now().Unix()+expireThreshold) {
		err = t.getToken()
	} else {
		err = t.refreshToken()
	}

	if err != nil {
		log.Error("error while getting token. ", err)
	}

	return nil
}

func (t *tokenManagerKeycloak) ClearToken() {
	t.responseModel = &TokenResponseModel{}
	log.Info("Token cleared")
}

func (t *tokenManagerKeycloak) refreshToken() error {
	data := url.Values{}
	data.Set("client_id", t.clientId)
	data.Set("client_secret", t.secretKey)
	data.Set("grant_type", grantTypeRefreshToken)
	data.Set("refresh_token", t.responseModel.RefreshToken)
	err := t.requestToken(data)
	if err != nil {
		log.Error("Token refresh request error", err)
	} else {
		log.Info("Token refreshed")
	}
	return err
}

func (t *tokenManagerKeycloak) getToken() error {
	data := url.Values{}
	data.Set("client_id", t.clientId)
	data.Set("client_secret", t.secretKey)
	data.Set("grant_type", grantTypeCredentials)
	err := t.requestToken(data)
	if err != nil {
		log.Error("Token request error", err)
	} else {
		log.Info("Token received")
	}
	return err
}

func (t *tokenManagerKeycloak) requestToken(data url.Values) error {

	r, _ := http.NewRequest(http.MethodPost, t.openIdConnectTokenUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := t.client.Do(r)
	if err != nil {
		return err
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return errors.New("response not successful. statusCode: " + resp.Status + " response: " + string(bodyBytes))
		}

		_ = json.Unmarshal(bodyBytes, t.responseModel)
		t.responseModel.ExpiredAt = time.Now().Unix() + t.responseModel.ExpiresIn
		t.responseModel.RefreshExpiredAt = time.Now().Unix() + t.responseModel.RefreshExpiresIn
		return nil
	}
}
