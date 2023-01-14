package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubBookStore struct {
	titles map[string]string
}

func (s StubBookStore) GetISBNTitle(isbn string) string {
	title := s.titles[isbn]
	return title
}

func TestGetTitlesfromISBN(t *testing.T) {
	store := StubBookStore{
		map[string]string{
			"1234": "Title1",
			"1235": "Title2",
		},
	}
	server := &BookServer{&store}

	t.Run("return First ISBN title ", func(t *testing.T) {
		request := newGetScoreRequest("1234")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "Title1")
	})

	t.Run("return Second ISBN title ", func(t *testing.T) {
		request := newGetScoreRequest("1235")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		fmt.Println("Second Test ISBN 1235 returns ", response.Body.String()) // Just for TS

		assertResponseBody(t, response.Body.String(), "Title2")
	})

}

func newGetScoreRequest(isbn string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/books/%s", isbn), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Response body is wrong got %q wanted %q", got, want)
	}

}
