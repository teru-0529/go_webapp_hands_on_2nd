// server.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/teru-0529/go_webapp_hands_on_2nd/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	// INFO:HTTPサーバーを起動
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %+v", err)
		os.Exit(1)
	}
}

// INFO:HTTPサーバーを起動
func run(ctx context.Context) error {
	// INFO:外部からのシグナルを受け付ける
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// INFO:環境変数の読込み
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// INFO:ポート番号が利用可能なことを確認
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("fail to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	// INFO:HTTPサーバーの定義
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// FIXME:実験用
			time.Sleep(5 * time.Second)
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
