package vstrconv

import (
	"log"
	"strconv"
)

func EnsureInt(s string) int {
	return int(EnsureInt64(s))
}

func EnsureInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Panicf("parse int64 error: %v", err)
	}
	return i
}

func EnsureInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Panicf("parse int32 error: %v", err)
	}
	return int32(i)
}

func EnsureUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Panicf("parse uint error: %v", err)
	}
	return uint(i)
}

func EnsureBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Panicf("parse bool error: %v", err)
	}
	return b
}

func EnsureFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Panicf("parse float64 error: %v", err)
	}
	return f
}

func EnsureFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Panicf("parse float32 error: %v", err)
	}
	return float32(f)
}

func EnsureUint64(s string) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Panicf("parse uint64 error: %v", err)
	}
	return i
}
