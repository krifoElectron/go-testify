package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenAllRight(t *testing.T) {
	count := 2
	city := "moscow"
	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(count)+"&city="+city, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)

	assert.NotEqual(t, responseRecorder.Body.String(), "")
}

func TestMainHandlerWhenCityWrong(t *testing.T) {
	count := 2
	city := "moscow228"
	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(count)+"&city="+city, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	count := 7
	city := "moscow"
	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(count)+"&city="+city, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)

	cafe := cafeList[city]
	answer := strings.Join(cafe[:], ",")
	assert.Equal(t, responseRecorder.Body.String(), answer)
}
