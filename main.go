// server.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/teru-0529/go_webapp_hands_on_2nd/config"
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

	// INFO:サーバーの起動
	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}
