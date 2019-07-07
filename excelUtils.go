package kissGo

import "github.com/360EntSecGroup-Skylar/excelize/v2"

func SetCell(val interface{}, index, sheet string, file *excelize.File) {
	switch val.(type) {
	case int:
		if v, ok := val.(int); ok {
			_ = file.SetCellInt(sheet, index, v)
		}
	case string:
		if v, ok := val.(string); ok {
			_ = file.SetCellStr(sheet, index, v)
		}
	case float64:
		if v, ok := val.(float64); ok {
			_ = file.SetCellFloat(sheet, index, v, 2, 64)
		}
	default:
		_ = file.SetCellDefault(sheet, index, val.(string))
	}
}

func SetFloatCell(val float64, index, sheet string, prec, bitSize int, file *excelize.File) {
	_ = file.SetCellFloat(sheet, index, val, prec, bitSize)
}
