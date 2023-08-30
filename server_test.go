package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	// INFO:キャンセル可能なコンテキストを準備
	ctx, cancel := context.WithCancel(context.Background())

	// INFO:動機制御用のエラーグループを準備
	eg, ctx := errgroup.WithContext(ctx)

	// INFO:HTTPサーバーの起動
	eg.Go(func() error { return run(ctx) })

	// INFO:Getリクエストの送信
	in := "message"
	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("fail to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Errorf("fail to read body: %v", err)
	}
	// TODO:HTTPサーバーの戻り値を検証
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// INFO:終了通知の送信
	cancel()
	// TODO:run関数の戻り値を検証
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
