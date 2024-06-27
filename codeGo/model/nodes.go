package model

type Employee struct {
	Name        string
	MailAddress string
	Domain      string
	Password    string
	PhoneNumber string
	Role        string
	Location    string
}

type Tool struct {
	Name         string
	ApprovalType string
}

type Department struct {
	DepartmentName string
}

type Community struct {
	Name        string
	Description string
	AccessType  string
}

type Communication struct {
	Type    string
	Content string `json:"content"`
}

type Skills struct {
	SkillName string
}

