package score

import (
	"net/http"

	"github.com/ScoreManagement/service/page"
)

func UploadScore(res http.ResponseWriter, req *http.Request) error {
	if err := req.ParseMultipartForm(1024 * 1024); err != nil {
		return err
	}
	for s := range req.MultipartForm.File {
		for i := range req.MultipartForm.File[s] {
			file, _ := req.MultipartForm.File[s][i].Open()
			//go page.SaveScore(file)
			page.SaveOneStudentScore(file)
		}
	}
	_, err := res.Write([]byte(`{"message":"success"}`))
	if err != nil {
		return err
	}
	return nil
}
