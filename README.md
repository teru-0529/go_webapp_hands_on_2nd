# go_webapp_hands_on_2nd
Go言語 Webアプリケーション開発　ハンズオンrev2.0


## 写経元リポジトリ（参考）
https://github.com/budougumi0617/go_todo_app


## セントラルリポジトリ
https://github.com/teru-0529/go_webapp_hands_on_2nd


## 変更履歴

2023/08/30

### SECTION-053 プロジェクトの初期化

* リポジトリの作成、クローン
* VS-CODEワークスペースの保存
* Goプロジェクトの初期化
  > [【参考】go mod完全に理解した](https://zenn.dev/optimisuke/articles/105feac3f8e726830f8c)

---

### SECTION-054 Webサーバーを起動する

* 動くだけのWebサーバーを構築
  * リクエストのパスを使ってレスポンスメッセージの組み立て
  * ポート番号固定
  > [【参考】【2023年最新版】VSCodeをGo言語超特化型にする、最高の拡張機能10選まとめ。](https://yurupro.cloud/2531/)
* Webサーバーの起動
  ```
  REM main関数の実行
  go run .
  ```
* APIの確認
  > http://localhost:18080/from_browser

---

### SECTION-055 リファクタリングとテストコード

* テスト容易性を高めるため`main`から`run`関数へ処理を分離する。
  * 出力結果を検証しやすくする。
  * 外部からの中断操作、異常状態を検知できるようにする。
* 外部からの中断操作を受け付けるため、`Shutdown`メソッドが実装されている`*http.Server`の`ListenAndServe`メソッドを利用してHTTPサーバーを起動する。
  * `*http.Server.ListenAndServe`メソッドを実行してHTTPリクエストを受け付ける。
  * 引数で渡された`context.Context`を通じて処理の中断命令を検知し、`*http.Server.Shutdown`メソッドでサーバー機能を終了する。
  * 戻り値として`*http.Server.ListenAndServe`の戻り値のエラーを返す。
* `*http.Server.ListenAndServe`メソッドを実行しつつ、`context.Context`から伝播される終了通知を待機する。
  * `[golang.org/x/sync/errgroup]`パッケージを利用して終了通知を待機する。
  * `errgroup.WithContext`関数を使い、取得した`*errgroup.Group`型の値の`Go`メソッドを利用することで、`func() error`というシグネチャの関数を別ゴルーチンで起動する。
  * `run`関数はHTTPリクエストを待機しつつ、引数で受け取った`context.Context`型の値の`Done`メソッドの戻り値として得られるチャネルからの通知を待つ。
  > [【参考】sync.ErrGroupで複数のgoroutineを制御する](https://deeeet.com/writing/2016/10/12/errgroup/)
  ```
  REM パッケージの取得
  go get -u golang.org/x/sync
  ```
  ```
  REM go.mod ファイル/go.sum ファイルの更新
  go mod tidy
  ```
* `run`関数のテスト
  * 期待通りにHTTPサーバーが起動しているか(HTTPサーバーの戻り値の検証)
  * 意図通りに終了するか(run関数の終了通知処理検証)
  ```
  REM テスト実行
  go test -v ./...
  ```

---

### SECTION-056 ポート番号を変更できるようにする

* 任意のポートでHTTPサーバーを起動できるようにする。
  * `run`関数外部から動的に選択したポート番号のリッスンを開始した`net.Listener`インターフェースを満たす型の値を渡す。
  * `net/http`パッケージではポート番号に`0`を指定すると、利用可能なポート番号を選択してくれることを利用。
  * `main`関数では、実行時の引数でポート番号を指定する。
  ```
  REM ポート番号を指定してmain関数を実行する
  go run . %port_no%
  ```

---

### SECTION-057 Dockerを利用した実行環境

* マルチステージビルドを実施し、ビルド環境/実行環境を分割する。
* dev環境では`air`コマンドを実行しファイルの更新を検知した際に`go build`コマンドを再実行、プログラム再起動を行う。ローカルマシンのディレクトリをマウントしておくことでホットリロードを実現。
* ビルド後コンテナを起動することで、ローカルマシンからリクエストが送信できる
  ```
  docker compose up
  ```
  [【参考】Goでリリースビルドするときに最低限付けておきたいオプション](https://qiita.com/ssc-ynakamura/items/da37856f7f217d708a07)

---

### SECTION-058 Makefileを追加する

* windwos環境のため、help(grep)/testの一部(-raceオプション)以外を作成

---
