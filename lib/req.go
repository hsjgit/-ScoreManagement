package lib

type ReqGetStudentScore struct {
	UserName string `json:"user_name"`
	Class    string `json:"class"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
}
