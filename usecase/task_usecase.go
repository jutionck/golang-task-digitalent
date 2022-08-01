package usecase

import (
	"github.com/jutionck/golang-task-digitalent/model"
	"github.com/jutionck/golang-task-digitalent/repository"
)

type TaskUseCase interface {
	RegisterNewTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	DeleteTask(id string) error
	FindAll() ([]model.Task, error)
	FindById(id string) (model.Task, error)
}

type taskUseCase struct {
	repo repository.TaskRepository
}

func (t *taskUseCase) RegisterNewTask(task *model.Task) error {
	return t.repo.Insert(task)
}

func (t *taskUseCase) UpdateTask(task *model.Task) error {
	return t.repo.Update(task)
}

func (t *taskUseCase) DeleteTask(id string) error {
	return t.repo.Delete(id)
}

func (t *taskUseCase) FindAll() ([]model.Task, error) {
	return t.repo.GetAll()
}

func (t *taskUseCase) FindById(id string) (model.Task, error) {
	return t.repo.GetById(id)
}

func NewTaskUseCase(repo repository.TaskRepository) TaskUseCase {
	uc := new(taskUseCase)
	uc.repo = repo
	return uc
}
