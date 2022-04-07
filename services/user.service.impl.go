package services

import (
	"context"
	"errors"
	"fasta/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (service *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := service.userCollection.InsertOne(service.ctx, user)
	return err
}

func (service *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	err := service.userCollection.FindOne(service.ctx, filter).Decode(&user)
	return user, err
}

func (service *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	cursor, err := service.userCollection.Find(service.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(service.ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("collection not found")
	}
	return users, nil
}

func (service *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}
	result, _ := service.userCollection.UpdateOne(service.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (service *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := service.userCollection.DeleteOne(service.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
