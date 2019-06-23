package utils

import "github.com/360EntSecGroup-Skylar/excelize/v2"

func SetCell(val interface{}, index, sheet string, file *excelize.File) {
	switch val.(type) {
	case int:
		if v, ok := val.(int); ok {
			err := file.SetCellInt(sheet, index, v)
			PrintErr(err)
		}
	case string:
		if v, ok := val.(string); ok {
			err := file.SetCellStr(sheet, index, v)
			PrintErr(err)
		}
	case float64:
		if v, ok := val.(float64); ok {
			err := file.SetCellFloat(sheet, index, v, 2, 64)
			PrintErr(err)
		}
	default:
		err := file.SetCellDefault(sheet, index, val.(string))
		PrintErr(err)
	}
}

func SetFloatCell(val float64, index, sheet string, prec, bitSize int, file *excelize.File) {
	err := file.SetCellFloat(sheet, index, val, prec, bitSize)
	PrintErr(err)

}
