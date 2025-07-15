package services

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"todo-app/backend/models"
)

type TaskService struct {
	Collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{Collection: collection}
}

func (s *TaskService) Create(title string) (*models.Task, error) {
	task := &models.Task{
		ID:        primitive.NewObjectID(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := s.Collection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) List() ([]models.Task, error) {
	cur, err := s.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var tasks []models.Task
	for cur.Next(context.TODO()) {
		var task models.Task
		if err := cur.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *TaskService) Update(id string, title string, completed *bool) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id")
	}
	update := bson.M{"updated_at": time.Now()}
	if title != "" {
		update["title"] = title
	}
	if completed != nil {
		update["completed"] = *completed
	}
	res := s.Collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": update},
		// Return the updated document
	)
	var updatedTask models.Task
	if err := res.Decode(&updatedTask); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &updatedTask, nil
}

func (s *TaskService) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id")
	}
	res, err := s.Collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("not found")
	}
	return nil
}
