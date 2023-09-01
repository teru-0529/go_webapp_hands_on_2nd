package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// INFO:HTTPサーバー
type Server struct {
	srv *http.Server
	l   net.Listener
}

// INFO:コンストラクタ
func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

// INFO:HTTPサーバーを起動
func (s *Server) Run(ctx context.Context) error {
	// INFO:外部からのシグナルを受け付ける
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// INFO:動機制御用のエラーグループを準備
	eg, ctx := errgroup.WithContext(ctx)

	// INFO:別ゴルーチンでHTTPサーバーを起動
	eg.Go(func() error {
		// http.ErrServerClosed 以外のエラーを異常とみなす
		// (http.ErrServerClosed はhttp.Server.Shutdown()が正常に終了したことを示す)
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Panicf("faild to close: %+v", err)
			return err
		}
		return nil
	})

	// INFO:チャネルからの終了通知を待機
	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("fail to shutdown: %+v", err)
	}

	// INFO:グレースフルシャットダウンの終了を待つ
	return eg.Wait()
}
