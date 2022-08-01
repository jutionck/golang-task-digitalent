package repository

import (
	"database/sql"

	"github.com/jutionck/golang-task-digitalent/model"
)

type TaskRepository interface {
	Insert(task *model.Task) error
	Update(task *model.Task) error
	Delete(id string) error
	GetAll() ([]model.Task, error)
	GetById(id string) (model.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func (t *taskRepository) Insert(task *model.Task) error {
	_, err := t.db.Exec("insert into m_task (task,employee,deadline) values (?, ?, ?)", task.TaskDetail, task.EmployeeName, task.TaskDeadline)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Update(task *model.Task) error {
	_, err := t.db.Exec("update m_task set task=?, employee=?, deadline=? where id=?", task.TaskDetail, task.EmployeeName, task.TaskDeadline, task.Id)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Delete(id string) error {
	_, err := t.db.Exec("delete from m_task where id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetAll() ([]model.Task, error) {
	rows, err := t.db.Query("select id, task, employee, deadline from m_task")
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	for rows.Next() {
		var task = model.Task{}
		var err = rows.Scan(&task.Id, &task.TaskDetail, &task.EmployeeName, &task.TaskDeadline)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) GetById(id string) (model.Task, error) {
	var task = model.Task{}
	err := t.db.QueryRow("select id, task, employee, deadline from m_task where id = ?", id).Scan(&task.Id, &task.TaskDetail, &task.EmployeeName, &task.TaskDeadline)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	repo := new(taskRepository)
	repo.db = db
	return repo
}
