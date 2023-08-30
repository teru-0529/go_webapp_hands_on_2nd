// server.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %+v", err)
	}
}

// INFO:HTTPサーバーの起動
func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
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
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
