package store

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"student-app/internal/model"
)

const StudentCollection = "students"

type StudentSMongoDB struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewStudentMongoDB(db *mongo.Database, logger *zap.Logger) *StudentSMongoDB {
	return &StudentSMongoDB{
		collection: db.Collection(StudentCollection),
		logger:     logger,
	}
}

func (store *StudentSMongoDB) Save(ctx context.Context, m model.Student) error {
	if _, err := store.collection.InsertOne(ctx, m); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return StudentDuplicateError{
				ID: m.ID,
			}
		}

		return fmt.Errorf("document creation on mongodb failed %w", err)
	}

	return nil

}
func (store *StudentSMongoDB) Get(ctx context.Context, id uint64) (model.Student, error) {
	var student model.Student

	res := store.collection.FindOne(ctx, bson.M{
		"id": id,
	})

	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return student, StudentNotFoundError{
				ID: id,
			}
		}

		return student, fmt.Errorf("cannot read from collection %w", err)
	}

	if err := res.Decode(&student); err != nil {
		return student, fmt.Errorf("cannot decode result into student %w", err)
	}

	return student, nil
}
func (store *StudentSMongoDB) GetAll(ctx context.Context) ([]model.Student, error) {
	cursor, err := store.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("cannot read from collection %w", err)
	}

	students := make([]model.Student, 0)
	for cursor.Next(ctx) {
		var student model.Student
		if err := cursor.Err(); err != nil {
			return nil, fmt.Errorf("cannot read current curr from collection %w", err)
		}

		if err := cursor.Decode(&student); err != nil {
			return nil, fmt.Errorf("cannot decode current cursor result into student %w", err)
		}

		students = append(students, student)
	}

	return students, nil
}
