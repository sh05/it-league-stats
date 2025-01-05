package excel

import (
	"github.com/xuri/excelize/v2"
)

func ReadExcelFile(filePath string) (*excelize.File, error) {
	return excelize.OpenFile(filePath)
}

func ReadExcelSheet(file *excelize.File, sheetName string) ([][]string, error) {
	return file.GetRows(sheetName)
}
