package store

import (
	"context"
	"fmt"
	"student-app/internal/model"
)

type StudentNotFoundError struct {
	ID uint64
}

type StudentDuplicateError struct {
	ID uint64
}

func (err StudentDuplicateError) Error() string {
	return fmt.Sprintf("student %d already exists", err.ID)
}

func (err StudentNotFoundError) Error() string {
	return fmt.Sprintf("student %d doesn't exists", err.ID)
}

type Students interface {
	Save(context.Context, model.Student) error
	Get(context.Context, uint64) (model.Student, error)
	GetAll(context.Context) ([]model.Student, error)
	Delete(context.Context, uint64) error
}
