package lib

type ReqGetStudentScore struct {
	UserName string `json:"user_name"`
	Class    string `json:"class"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
}
