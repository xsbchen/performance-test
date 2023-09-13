package handler

import (
	"net/http"
)

const (
	helloWorldStr    = "Hello, World!"
	contentTypePlain = "text/plain; charset=utf-8"
	contentTypeHtml  = "text/html; charset=utf-8"
	contentTypeJson  = "application/json"
)

var (
	helloWorldBytes = []byte(helloWorldStr)
)

// Plaintext . Test 7: Plaintext.
func Plaintext(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", contentTypePlain)
	_, _ = w.Write(helloWorldBytes)
}
