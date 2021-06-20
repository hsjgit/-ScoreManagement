package route

import (
	"log"
	"net/http"

	"github.com/ScoreManagement/route/score"
)

type appHandel func(res http.ResponseWriter, req *http.Request) error

func RegisterRoutes() {
	http.HandleFunc("/upload", errWrapper(score.UploadScore))
	log.Println("/upload")
	http.HandleFunc("/get", errWrapper(score.GetStudentScore))
	log.Println("/get")
}
