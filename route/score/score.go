package score

import (
	"encoding/json"
	"log"
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
	resUpload := lib.ResUploadStudentScore{
		Code:    1,
		Message: "success",
	}
	marshal, MarshalErr := json.Marshal(resUpload)
	if MarshalErr != nil {
		return MarshalErr
	}
	_, err := res.Write(marshal)
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
	students, SelectStudentScoreErr := page.SelectStudentScore(*condition)
	if SelectStudentScoreErr != nil {
		log.Println(SelectStudentScoreErr.Error())
		return SelectStudentScoreErr
	}
	resStudents := lib.ResGetStudentScore{
		Code:  1,
		Data:  students,
		Page:  condition.Page,
		Count: 0,
	}
	marshal, MarshalErr := json.Marshal(resStudents)
	if MarshalErr != nil {
		return MarshalErr
	}
	_, err := res.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}
