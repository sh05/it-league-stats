package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// 存在確認してなければ作る
func TouchExcelFile(filePath string) {
	// filenameで存在確認
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		file = excelize.NewFile()
		if err := file.SaveAs(filePath); err != nil {
			log.Fatal(err)
		}
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

func WriteExcelSheet(filePath string, sheetName string, data [][]interface{}) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.NewSheet(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	sw, err := file.NewStreamWriter(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range data {
		cell, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			log.Fatal(err)
		}
		err = sw.SetRow(cell, row)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := sw.Flush(); err != nil {
		log.Fatal(err)
	}
	if err := file.SaveAs(filePath); err != nil {
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
