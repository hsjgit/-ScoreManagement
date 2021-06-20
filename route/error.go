package route

import (
	"net/http"
)

// 统一的错误处理
func errWrapper(handler appHandel) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		err := handler(res, req)
		if err != nil {
			code := http.StatusInternalServerError
			switch {
			}
			http.Error(res, err.Error(), code)
		}
	}
}
