package page

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/ScoreManagement/lib"
	"github.com/ScoreManagement/service/data/student"
)

func SaveScore(file []*multipart.FileHeader) {

}

func SaveOneStudentScore(file multipart.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		studentDB := student.NewStudentDB()
		content := scanner.Text()
		if content == "" {
			continue
		}
		score := strings.Split(content, " ")
		studentDB.Class = score[0]
		studentDB.UserName = score[1]
		float, _ := strconv.ParseFloat(score[2], 32)
		studentDB.Score = float32(float)
		studentDB.Subject = score[3]
		go studentDB.SaveOneStudentScore()
	}

}

func SelectStudentScore(condition lib.ReqGetStudentScore) []student.Student {
	studentDB := student.NewStudentDB()
	studentDB.UserName = condition.UserName
	studentDB.Class = condition.Class

	switch {
	case studentDB.UserName != "" && studentDB.Class == "":
		students, err := studentDB.SelectStudentsScoreByName(condition.Sort, condition.Order)
		if err != nil {
			fmt.Println(err.Error())
		}
		return students
	case studentDB.UserName == "" && studentDB.Class != "":
		students, _ := studentDB.SelectStudentsScoreByClass(condition.Sort, condition.Order)
		return students
	default:
		students, _ := studentDB.SelectStudentsScoreByClassAndName(condition.Sort, condition.Order)
		return students
	}

}
