package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	// INFO:利用可能なポート番号を確保
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("fail to listen port %v", err)
	}

	// INFO:キャンセル可能なコンテキストを準備
	ctx, cancel := context.WithCancel(context.Background())

	// INFO:ハンドラー実装を定義
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	// INFO:動機制御用のエラーグループを準備
	eg, ctx := errgroup.WithContext(ctx)

	// INFO:HTTPサーバーの起動
	eg.Go(func() error {
		s := NewServer(l, mux)
		return s.Run(ctx)
	})

	// INFO:Getリクエストの送信
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	t.Logf("try request to [%q]", url)
	rsp, err := http.Get(url)
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
