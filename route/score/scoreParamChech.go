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
	param.Class = strings.Trim(param.Class, " ")
	param.UserName = strings.Trim(param.UserName, " ")
}
