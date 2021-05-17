package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	// 新規ファイルを作成 --- (*1)
	f := excelize.NewFile()
	// 新規シートを追加 --- (*2)
	index := f.NewSheet("Sheet2")
	// セルに値をセット --- (*3)
	f.SetCellValue("Sheet2", "C2", "こんにちは")
	// アクティブシートを変更 --- (*4)
	f.SetActiveSheet(index)
	// ファイルに保存 --- (*5)
	f.SaveAs("hello.xlsx")
}
