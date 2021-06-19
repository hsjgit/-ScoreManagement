package page

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/ScoreManagement/lib"
	"github.com/ScoreManagement/service/data/student"
)

func SaveScore(files []*multipart.FileHeader) {
	for i := range files {
		file, _ := files[i].Open()
		go SaveOneFile(file)
	}
}

func SaveOneFile(file multipart.File) {
	scanner := bufio.NewScanner(file)
	studentDB := student.NewStudentDB()
	students := make([]student.Student, 0)
	for scanner.Scan() {
		s := student.Student{}
		content := scanner.Text()
		if content == "" {
			continue
		}
		score := strings.Split(content, " ")
		s.Class = score[0]
		s.UserName = score[1]
		float, _ := strconv.ParseFloat(score[2], 32)
		s.Score = float32(float)
		s.Subject = score[3]
		students = append(students, s)
		if len(students) == 300 {
			sql, _ := buildSql(students)
			studentDB.SaveStudentsScore(sql, students)
			students = students[:0]
		}
	}
	sql, _ := buildSql(students)
	err := studentDB.SaveStudentsScore(sql, students)
	if err != nil {
		log.Println(err.Error())
	}

}

func buildSql(students []student.Student) (string, error) {
	if len(students) > 300 {
		return "", errors.New("最多一次拼接300个实体")
	}
	insert := "INSERT INTO student(`user_name`, `class`, `subject`, `score`) VALUES"
	buf := strings.Builder{}
	buf.WriteString(insert)
	count := 0
	for range students {
		buf.WriteString("(")
		buf.WriteString("?, ?, ?, ?")
		buf.WriteString("),")
		count++
	}
	sql := buf.String()
	return sql[:len(sql)-1], nil
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
