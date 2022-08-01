package model

type Task struct {
	Id           string `json:"id,omitempty"`
	TaskDetail   string `json:"taskDetail"`
	EmployeeName string `json:"employeeName"`
	TaskDeadline string `json:"taskDeadline"`
}
