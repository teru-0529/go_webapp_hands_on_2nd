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
  go run .
  ```
* APIの確認
  > http://localhost:18080/from_browser
