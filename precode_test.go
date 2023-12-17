package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getResponseFromUrl(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	res := getResponseFromUrl("/cafe?city=moscow&count=7")

	assert.Len(t, strings.Split(string(res.Body.String()), ","), totalCount)
	require.Equal(t, http.StatusOK, res.Code)
	require.NotEmpty(t, res.Body)
}

func TestMainHandlerWhenCityWrong(t *testing.T) {
	res := getResponseFromUrl("/cafe?city=almaty&count=3")
	expectedTestValue := "wrong city value"

	require.Equal(t, http.StatusBadRequest, res.Code)
	require.NotEmpty(t, res.Body)
	assert.Equal(t, expectedTestValue, res.Body.String())
}

func TestMainHandlerWhenReturnsOk(t *testing.T) {
	res := getResponseFromUrl("/cafe?city=moscow&count=2")

	require.Equal(t, http.StatusOK, res.Code)
	require.NotEmpty(t, res.Body)
}
