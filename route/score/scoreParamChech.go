package score

import (
	"strings"

	"github.com/ScoreManagement/lib"
)

func GetStudentScoreParam(param *lib.ReqGetStudentScore) {
	if param.Order == "" {
		param.Order = "DESC"
	}
	if param.Sort == "" {
		param.Sort = "id"
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.PageSize <= 0 {
		param.PageSize = 10
	}
	param.Class = strings.Trim(param.Class, " ")
	param.UserName = strings.Trim(param.UserName, " ")
}
