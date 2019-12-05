package vmath

import (
	"fmt"
	"strconv"
)

func RoundFloat(f float64) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", f), 64)
	return v
}

func RoundFloat64(f float64, precision int) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", f), 64)
	return v
}
