# taisoku_script

[タイピング速度測定](https://taisoku.com/)で計測した結果をGoogle スプレッドシートに記録するスクリプト

## 事前準備
- GCP
  - プロジェクトの作成
  - Google Sheets APIの有効化
  - 認証情報の作成
  - APIキーの作成
  - 認証情報ダウンロード(こちらの鍵情報が必要です)
    - `config/credentials.json`に配置
  - スプレッドシートのidとシート名を取得し`.env`に記載
    ```env
    SPREAD_SHEET_ID=""
    SHEET_NAME=""
    ```
  - devcontainerでGo環境構築
    ```shell
    go mod tidy
    ```
  - タイピング速度測定の結果画面から`結果をコピー`

```shell
go run cmd/main.go
```

## 参考
https://zenn.dev/ttsbs/articles/a8500c2a44356a
https://note.com/s_t877/n/n7ce48a6e945f
