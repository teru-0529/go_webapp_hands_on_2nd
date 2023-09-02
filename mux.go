package main

import "net/http"

func NewMux() http.Handler {
	mux := http.NewServeMux()

	// INFO:稼働中かを確認する`/health`エンドポイント
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset-utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	return mux
}
