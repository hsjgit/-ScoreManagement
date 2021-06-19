package route

import (
	"net/http"

	"github.com/ScoreManagement/route/score"
)

type appHandel func(res http.ResponseWriter, req *http.Request) error

func RegisterRoutes() {
	http.HandleFunc("/upload", errWrapper(score.UploadScore))
}
