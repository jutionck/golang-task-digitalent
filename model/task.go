package model

import "time"

type Task struct {
	Id           string
	TaskDetail   string
	EmployeeName string
	TaskDeadline time.Time
}
