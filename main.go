package main

import "fmt"

// お金
type Money struct {
	kind   int    // 金額
	damage bool   // 傷の有無
	unit   string // 単位
}

func main() {
	// 投入するお金
	moneyArr := []Money{}
	money1 := Money{500, false, "円"}
	money2 := Money{1000, false, "円"}
	moneyArr = append(moneyArr, money1)
	moneyArr = append(moneyArr, money2)

	calc(moneyArr)
}

func calc(moneyArr []Money) {
	// 合計金額を計算する処理を呼び出す
	amount, bill, coin, na := calcAmount(moneyArr)

	// 標準出力に結果を出力
	fmt.Printf("合計:%v 紙幣の枚数:%v 硬貨の枚数:%v 対象外の枚数:%v", amount, bill, coin, na)
}

// 合計金額を計算する処理
func calcAmount(money []Money) (amount int, bill int, coin int, na int) {
	// Todo 合計金額の計算処理
	for _, m := range money {
		if m.damage || m.unit != "円" {
			na++
			continue
		}

		if m.kind >= 1000 {
			bill++
		} else {
			coin++
		}
		amount += m.kind
	}

	return amount, bill, coin, na

}
