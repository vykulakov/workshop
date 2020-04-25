package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"workshop/internal/handler"

	"github.com/stretchr/testify/require"

	"workshop/internal/api"
	"workshop/internal/api/mocks"
)

func TestHandler_Hello(t *testing.T) {
	tests := []struct {
		name     string
		joke     *api.JokeResponse
		err      error
		codeWant int
		bodyWant string
	}{{
		name:     "simple test",
		joke:     &api.JokeResponse{"test joke"},
		codeWant: 200,
		bodyWant: "test joke",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/hello", nil)

			apiMock := &mocks.Client{}
			apiMock.On("GetJoke").Return(tt.joke, tt.err)

			h := handler.NewHandler(apiMock)
			h.Hello(rr, req)

			require.Equal(t, tt.bodyWant, string(rr.Body.Bytes()))
			require.Equal(t, tt.codeWant, rr.Result().StatusCode)
		})
	}
}
