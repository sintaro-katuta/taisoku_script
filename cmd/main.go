package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/atotto/clipboard"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error")
	}

	spreadsheetID := os.Getenv("SPREAD_SHEET_ID")
	sheetName := os.Getenv("SHEET_NAME")
	text, err := clipboard.ReadAll()

	if err != nil {
		log.Fatalf("クリップボードの読み込みに失敗: %v", err)
	}

	// 正規表現で必要な情報を抽出（(?s) フラグで '.' が改行にもマッチ）
	re := regexp.MustCompile(`(?s)(\d{4}/\d{1,2}/\d{1,2}).*打鍵数：(\d+).*平均入力打鍵：([\d.]+).*瞬間最高打鍵：(\d+).*ミス数：(\d+)`)
	matches := re.FindStringSubmatch(text)
	if len(matches) != 6 {
		log.Fatalf("データの解析に失敗しました: %v", text)
	}

	// 各値を取得
	date := matches[1]
	hits := matches[2]
	avgHits := matches[3]
	maxHits := matches[4]
	miss := matches[5]

	// Sheets API 初期化
	ctx := context.Background()
	b, err := os.ReadFile("config/credentials.json")
	if err != nil {
		log.Fatalf("認証情報の読み込みエラー: %v", err)
	}
	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("JWT構成エラー: %v", err)
	}
	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Sheets APIの初期化に失敗: %v", err)
	}

	// 送信するデータ
	values := []interface{}{date, hits, avgHits, maxHits, miss}
	vr := &sheets.ValueRange{
		Values: [][]interface{}{values},
	}

	// シートの末尾に追記
	_, err = srv.Spreadsheets.Values.Append(spreadsheetID, sheetName+"!A1", vr).
		ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("データの追加に失敗: %v", err)
	}

	fmt.Println("スプレッドシートに記録しました。")
}
