package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// 入力値（in）と期待値（out）を定義する構造体を作成（mapなどで代替しても良い）
type calsTest struct {
	moneyArr []Money
	mes      string
}

func TestA(t *testing.T) {
	// 投入するお金
	moneyArr1 := []Money{{1000, false, "円"}, {5000, false, "円"}, {10000, false, "円"}}
	moneyArr2 := []Money{{500, false, "円"}, {100, false, "円"}, {50, false, "円"}, {10, false, "円"}, {5, false, "円"}, {1, false, "円"}}
	moneyArr3 := []Money{{500, false, "円"}, {500, false, "ドル"}}
	moneyArr4 := []Money{{500, true, "円"}, {500, false, "円"}}

	expected1 := "合計:16000 紙幣の枚数:3 硬貨の枚数:0 対象外の枚数:0"
	expected2 := "合計:666 紙幣の枚数:0 硬貨の枚数:6 対象外の枚数:0"
	expected3 := "合計:500 紙幣の枚数:0 硬貨の枚数:1 対象外の枚数:1"
	expected4 := "合計:500 紙幣の枚数:0 硬貨の枚数:1 対象外の枚数:1"

	var calsTests = []calsTest{
		{moneyArr1, expected1},
		{moneyArr2, expected2},
		{moneyArr3, expected3},
		{moneyArr4, expected4},
	}

	// 入力値と期待値を1件ずつテストする.
	for _, ct := range calsTests {
		t.Helper()

		// 既存のStdoutを退避する
		orgStdout := os.Stdout
		// パイプを定義
		pr, pw, _ := os.Pipe()
		// Stdoutの出力先をパイプのwriterに変更する
		os.Stdout = pw

		calc(ct.moneyArr)

		// Writerをクローズする。Writerオブジェクトはクローズするまで処理をブロックするので注意
		pw.Close()
		// 出力先を元に戻す。※ Close()の後に書く
		os.Stdout = orgStdout

		// Bufferに書き込こまれた内容を読み出す
		buf := bytes.Buffer{}
		io.Copy(&buf, pr)
		// buf.String()には改行コードが入っているので消す
		output := strings.TrimRight(buf.String(), "\n")

		// 期待したメッセージが出力されていることを確認する
		if output != ct.mes {
			t.Errorf("期待するメッセージと違います:%v", output)
		}

	}

}
