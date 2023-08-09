package repository

import (
	"context"
	"project-p-back/internal/entity"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Collection
}

type IUserRepository interface {
	GetAllUser() ([]entity.User, error)
	GetUserById(string) (*entity.User, error)
	GetUserByUsername(string) (*entity.User, error)
	CreateUser(*entity.User) (*entity.User, error)
	UpdateUser(string, *entity.User) (*entity.User, error)
	IsUserExists(string) bool
}

func NewUserRepository(client *mongo.Client) *userRepository {
	collectionName := viper.GetString("MongoDb.Database.Collections.UserCollection")
	databaseName := viper.GetString("MongoDb.Database.Name")
	db := client.Database(databaseName).Collection(collectionName)

	repository := userRepository{}
	repository.db = db
	return &repository
}

func (repo *userRepository) GetAllUser() ([]entity.User, error) {
	var result []entity.User
	cursor, err := repo.db.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var i entity.User
		err := cursor.Decode(&i)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func (repo *userRepository) GetUserById(id string) (*entity.User, error) {
	var user entity.User
	filter := bson.D{{Key: "_id", Value: id}}
	err := repo.db.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User

	filter := bson.D{{Key: "username", Value: username}}
	err := repo.db.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	_, err := repo.db.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) UpdateUser(id string, user *entity.User) (*entity.User, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: user}}

	_, err := repo.db.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepository) IsUserExists(username string) bool {
	filter := bson.D{{Key: "username", Value: username}}

	count, err := repo.db.CountDocuments(context.Background(), filter)
	if err != nil {
		return false
	}

	return count > 0
}
