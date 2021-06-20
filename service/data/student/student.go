package student

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/ScoreManagement/lib/db"
)

type Student struct {
	Id       int64   `json:"id"`
	Class    string  `json:"class"`
	UserName string  `json:"user_name"`
	Subject  string  `json:"subject"`
	Score    float32 `json:"score"`
}

type StudentDB struct {
	Student
	DB *sql.DB
}

func NewStudentDB() *StudentDB {
	return &StudentDB{
		DB:      db.DB,
		Student: Student{},
	}
}

func (s *StudentDB) SaveOneStudentScore() error {
	sql := "INSERT INTO student(`user_name`, `class`, `subject`, `score`) VALUES (?, ?, ?, ?)"
	prepare, PrepareErr := s.DB.Prepare(sql)
	if PrepareErr != nil {
		return PrepareErr
	}
	_, ExecErr := prepare.Exec(s.UserName, s.Class, s.Subject, s.Score)
	if ExecErr != nil {
		return ExecErr
	}
	return nil
}

func (s *StudentDB) SaveStudentsScore(sql string, students []Student) error {
	value := make([]interface{}, 0)
	for i := range students {
		value = append(value, students[i].UserName, students[i].Class, students[i].Subject, students[i].Score)
	}

	prepare, PrepareErr := s.DB.Prepare(sql)
	if PrepareErr != nil {
		return PrepareErr
	}
	_, ExecErr := prepare.Exec(value...)
	if ExecErr != nil {
		return ExecErr
	}
	return nil

}

func (s *StudentDB) SelectStudentsScoreByName(sort, order string, page, pageSize int) ([]Student, error) {
	sql := "select * from student where user_name = ? ORDER BY " +
		sort + " " + order +
		" LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa((page-1)*pageSize)
	return s.singleCondition(sql, s.UserName)
}

func (s *StudentDB) SelectStudentsScoreByClass(sort, order string, page, pageSize int) ([]Student, error) {

	sql := "select * from student where class = ? ORDER BY " +
		sort + " " + order +
		" LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa((page-1)*pageSize)
	return s.singleCondition(sql, s.Class)
}

func (s *StudentDB) singleCondition(sql string, condition interface{}) ([]Student, error) {
	prepare, PrepareErr := s.DB.Prepare(sql)
	if PrepareErr != nil {
		return nil, PrepareErr
	}
	query, QueryErr := prepare.Query(condition)

	if QueryErr != nil {
		return nil, QueryErr
	}
	students := make([]Student, 0)
	for query.Next() {
		student := Student{}
		if err := query.Scan(&student.Id, &student.UserName, &student.Class, &student.Subject, &student.Score); err != nil {
			log.Println(err.Error())
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *StudentDB) SelectStudentsScoreByClassAndName(sort, order string, page, pageSize int) ([]Student, error) {
	sql := "select * from student where class = ? and user_name = ? ORDER BY " +
		sort + " " + order +
		" LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa((page-1)*pageSize)
	prepare, PrepareErr := s.DB.Prepare(sql)
	if PrepareErr != nil {
		return nil, PrepareErr
	}
	query, QueryErr := prepare.Query(s.Class, s.UserName)
	if QueryErr != nil {
		return nil, QueryErr
	}
	students := make([]Student, 0)
	for query.Next() {
		student := Student{}
		if err := query.Scan(&student.Id, &student.UserName, &student.Class, &student.Subject, &student.Score); err != nil {
			log.Println(err.Error())
		}
		students = append(students, student)
	}
	return students, nil
}

func (s StudentDB) Count() (int64, error) {
	sqlCount := ""
	var err error
	var count int64 = 0
	var prepare *sql.Stmt
	var rows *sql.Rows
	switch {
	case s.UserName != "" && s.Class == "":
		sqlCount = "select count(*) as count from student where user_name = ?"
		count, err = s.getCount(prepare, sqlCount, rows, []interface{}{s.UserName})

	case s.UserName == "" && s.Class != "":
		sqlCount = "select count(*) as count from student where class = ?"
		count, err = s.getCount(prepare, sqlCount, rows, []interface{}{s.Class})

	default:
		sqlCount = "select count(*) as count from student where user_name = ? and class = ?"
		count, err = s.getCount(prepare, sqlCount, rows, []interface{}{s.UserName, s.Class})
	}
	return count, err
}

func (s StudentDB) getCount(prepare *sql.Stmt, sqlCount string, rows *sql.Rows, condition []interface{}) (int64, error) {
	var err error
	prepare, err = s.DB.Prepare(sqlCount)
	if err != nil {
		return 0, err
	}
	rows, err = prepare.Query(condition...)
	if err != nil {
		return 0, err
	}
	count := int64(0)
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return count, err
}
