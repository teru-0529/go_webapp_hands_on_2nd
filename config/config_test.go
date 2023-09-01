package config

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	// INFO:ポート番号の環境変数設定
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))

	// TODO:設定した環境変数値が期待通り取得できることを検証
	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}

	// TODO:デフォルト値が期待通り取得できることを検証
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
}
