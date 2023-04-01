package store_test

import (
	"context"
	"student-app/internal/model"
	"student-app/internal/store"
	"testing"
)

func TestStudentInMemory_Save(t *testing.T) {
	st := store.NewStudentInMemory()
	ctx := context.Background()

	st.Save(ctx, model.Student{
		ID:        9813100,
		FirstName: "Hossein",
		LastName:  "Testy",
	})

	m, err := st.Get(ctx, 9813100)
	if err != nil {
		t.Fatal(err)
	}

	if m.FirstName != "Hossein" {
		t.Fatal("first name should be Hossein")
	}

	if m.LastName != "Testy" {
		t.Fatal("first name should be Testy")
	}
}
