package vquery

import (
	"net/http"
	"strconv"
)

func Int(r *http.Request, name string) (int, bool) {
	param, ok := String(r, name)
	if !ok {
		return 0, false
	}

	i, err := strconv.Atoi(param)
	if err != nil {
		return 0, false
	}

	return i, true
}

func String(r *http.Request, name string) (string, bool) {
	query := r.URL.Query()
	if len(query) == 0 {
		return "", false
	}
	param := query.Get(name)
	if param == "" {
		return "", false
	}
	return param, true
}

func Float(r *http.Request, name string) (float64, bool) {
	param, ok := String(r, name)
	if !ok {
		return 0, false
	}
	f, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

func Bool(r *http.Request, name string) (bool, bool) {
	param, ok := String(r, name)
	if !ok {
		return false, false
	}
	b, err := strconv.ParseBool(param)
	if err != nil {
		return false, false
	}

	return b, true
}
