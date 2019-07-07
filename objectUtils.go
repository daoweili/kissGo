package kissGo

import "strconv"

func FloatToStr(v float64, prec int) string {
	return strconv.FormatFloat(v, 'f', prec, 64)
}

func StrToFloat(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func IntToStr(v int) string {
	return strconv.Itoa(v)
}

func StrToInt(v string) (int, error) {
	return strconv.Atoi(v)
}
