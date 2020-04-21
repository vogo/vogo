// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vmath

import (
	"fmt"
	"math"
	"strconv"
)

func RoundFloat(f float64) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", f), 64)
	return v
}

func RoundValidFloat(f float64) float64 {
	if math.IsNaN(f) {
		return 0
	}

	return RoundFloat(f)
}

func RoundFloat64(f float64, precision int) float64 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", f), 64)
	return v
}
