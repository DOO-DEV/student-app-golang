package store

import (
	"context"
	"student-app/internal/model"
)

type StudentInMemory struct {
	students map[uint64]model.Student
}

func NewStudentInMemory() *StudentInMemory {
	return &StudentInMemory{
		students: make(map[uint64]model.Student),
	}
}

func (m *StudentInMemory) Save(_ context.Context, s model.Student) error {
	if _, ok := m.students[s.ID]; ok {
		return StudentDuplicateError{ID: s.ID}
	}

	m.students[s.ID] = s

	return nil
}

func (m *StudentInMemory) Get(_ context.Context, ID uint64) (model.Student, error) {
	s, ok := m.students[ID]
	if ok {
		return s, nil
	}

	return s, StudentNotFoundError{ID: ID}
}

func (m *StudentInMemory) GetAll(_ context.Context) ([]model.Student, error) {
	ss := make([]model.Student, 0)
	for _, s := range m.students {
		ss = append(ss, s)
	}

	return ss, nil
}

func (m *StudentInMemory) Delete(_ context.Context, id uint64) error {
	if _, ok := m.students[id]; !ok {
		return StudentNotFoundError{ID: id}
	}

	delete(m.students, id)

	return nil
}
