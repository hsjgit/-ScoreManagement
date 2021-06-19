package score

import (
	"encoding/json"
	"net/http"

	"github.com/ScoreManagement/lib"
	"github.com/ScoreManagement/service/page"
)

func UploadScore(res http.ResponseWriter, req *http.Request) error {
	if err := req.ParseMultipartForm(1024 * 1024); err != nil {
		return err
	}
	for s := range req.MultipartForm.File {
		go page.SaveScore(req.MultipartForm.File[s])
	}
	_, err := res.Write([]byte(`{"message":"success"}`))
	if err != nil {
		return err
	}
	return nil
}

func GetStudentScore(res http.ResponseWriter, req *http.Request) error {
	condition := &lib.ReqGetStudentScore{
		UserName: req.FormValue("user_name"),
		Class:    req.FormValue("class"),
		Sort:     req.FormValue("sort"),
		Order:    req.FormValue("order"),
	}
	GetStudentScoreParam(condition)
	students := page.SelectStudentScore(*condition)
	marshal, MarshalErr := json.Marshal(students)
	if MarshalErr != nil {
		return MarshalErr
	}
	_, err := res.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}
