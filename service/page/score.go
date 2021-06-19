package page

import (
	"bufio"
	"mime/multipart"
	"strconv"
	"strings"

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
		studentDB.SaveOneStudentScore()
	}

}
