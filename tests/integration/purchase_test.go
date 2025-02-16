package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"shop-service/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	baseURL = "http://localhost:8080"
)

func TestPurchaseFlow(t *testing.T) {
	// 1. Регистрация пользователя
	registerResp := makeRequest(t, "POST", "/auth/register", map[string]string{
		"username": "testuser5",
		"password": "password123",
	})
	assert.Equal(t, http.StatusOK, registerResp.StatusCode)

	// 2. Авторизация пользователя
	loginResp := makeRequest(t, "POST", "/auth/login", map[string]string{
		"username": "testuser5",
		"password": "password123",
	})
	assert.Equal(t, http.StatusOK, loginResp.StatusCode)

	var loginData map[string]string
	decodeJSON(t, loginResp.Body, &loginData)
	token := loginData["token"]
	require.NotEmpty(t, token)

	// // 3. Покупка мерча
	purchaseResp := makeAuthorizedRequest(t, "GET", "/api/buy/t-shirt", token, nil)
	assert.Equal(t, http.StatusOK, purchaseResp.StatusCode)

	// // 4. Проверка информации о пользователе
	infoResp := makeAuthorizedRequest(t, "GET", "/api/info", token, nil)
	assert.Equal(t, http.StatusOK, infoResp.StatusCode)

	var infoData map[string]interface{}
	logger.Log.Println(infoData)

	decodeJSON(t, infoResp.Body, &infoData)
}

// Вспомогательные функции
func makeRequest(t *testing.T, method, path string, body interface{}) *http.Response {
	reqBody, _ := json.Marshal(body)
	resp, err := http.Post(baseURL+path, "application/json", bytes.NewReader(reqBody))
	require.NoError(t, err)
	return resp
}

func makeAuthorizedRequest(t *testing.T, method, path, token string, body interface{}) *http.Response {
	reqBody, _ := json.Marshal(body)
	client := &http.Client{}
	req, err := http.NewRequest(method, baseURL+path, bytes.NewReader(reqBody))
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	require.NoError(t, err)
	return resp
}

func decodeJSON(t *testing.T, reader io.Reader, target interface{}) {
	err := json.NewDecoder(reader).Decode(target)
	require.NoError(t, err)
}
