package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

type testSerer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testSerer {
	ts := httptest.NewTLSServer(h)
	return &testSerer{ts}
}

func (ts *testSerer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get()

	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}
