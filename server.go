// server.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	// INFO:引数でポート番号が指定されていることを確認
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}
	// INFO:ポート番号が利用可能なことを確認
	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("fail to listen port %s: %v", p, err)
	}
	// INFO:HTTPサーバーを起動
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %+v", err)
		os.Exit(1)
	}
}

// INFO:HTTPサーバーを起動
func run(ctx context.Context, l net.Listener) error {
	// INFO:HTTPサーバーの定義
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}

	// INFO:動機制御用のエラーグループを準備
	eg, ctx := errgroup.WithContext(ctx)

	// INFO:別ゴルーチンでHTTPサーバーを起動
	eg.Go(func() error {
		// http.ErrServerClosed 以外のエラーを異常とみなす
		// (http.ErrServerClosed はhttp.Server.Shutdown()が正常に終了したことを示す)
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Panicf("faild to close: %+v", err)
			return err
		}
		return nil
	})

	// INFO:チャネルからの終了通知を待機
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("fail to shutdown: %+v", err)
	}

	// INFO:別ゴルーチンの終了を待つ
	return eg.Wait()
}
