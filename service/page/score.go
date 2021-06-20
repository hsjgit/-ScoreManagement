package page

import (
	"bufio"
	"errors"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/ScoreManagement/lib"
	"github.com/ScoreManagement/service/data/student"
)

func SaveScore(files []*multipart.FileHeader) {
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			log.Println(err.Error())
		}
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
			sql, err := buildSql(students)
			if err != nil {
				log.Println(err.Error())
			}
			if err := studentDB.SaveStudentsScore(sql, students); err != nil {
				Retry(studentDB, sql, students)
			}
			students = students[:0]
		}
	}
	sql, _ := buildSql(students)
	if err := studentDB.SaveStudentsScore(sql, students); err != nil {
		log.Println(err.Error())
		Retry(studentDB, sql, students)
	}

}

func Retry(studentDB *student.StudentDB, sql string, students []student.Student) {
	count := 0
	for ; count < 3; count++ {
		if err := studentDB.SaveStudentsScore(sql, students); err == nil {
			break
		}
	}
	if count == 3 {
		for i := range students {
			studentDB.Score = students[i].Score
			studentDB.UserName = students[i].UserName
			studentDB.Class = students[i].Class
			studentDB.Subject = students[i].Subject
			if err := studentDB.SaveOneStudentScore(); err != nil {
				log.Println(err.Error())
			}

		}
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

func SelectStudentScore(condition lib.ReqGetStudentScore) ([]student.Student, error) {
	studentDB := student.NewStudentDB()
	studentDB.UserName = condition.UserName
	studentDB.Class = condition.Class

	switch {
	case studentDB.UserName != "" && studentDB.Class == "":
		students, err := studentDB.SelectStudentsScoreByName(condition.Sort, condition.Order)
		return students, err
	case studentDB.UserName == "" && studentDB.Class != "":
		students, err := studentDB.SelectStudentsScoreByClass(condition.Sort, condition.Order)
		return students, err
	default:
		students, err := studentDB.SelectStudentsScoreByClassAndName(condition.Sort, condition.Order)
		return students, err
	}

}
