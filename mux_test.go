package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	// INFO:HTTPサーバーに渡す引数モックを生成
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	// INFO:ルーティング/ハンドラーの定義
	sut := NewMux()

	// INFO:HTTP関数の実行、レスポンスの取得
	sut.ServeHTTP(w, r)
	resp := w.Result()
	t.Cleanup(func() { _ = resp.Body.Close() })

	// TODO:ステータスコードを検証
	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but", resp.StatusCode)
	}

	// TODO:レスポンス内容を検証
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("faited to read body: %v", err)
	}
	want := `{"status": " ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
