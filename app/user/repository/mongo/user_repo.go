package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	authModel "pixstall-user/app/domain/auth/model"
	"pixstall-user/app/domain/user"
	userModel "pixstall-user/app/domain/user/model"
	mongoModel "pixstall-user/app/user/repository/mongo/model"
)

type mongoUserRepo struct {
	db *mongo.Database
}

func NewMongoUserRepo(db *mongo.Database) user.Repo {
	return &mongoUserRepo{
		db: db,
	}
}

func (m mongoUserRepo) SaveAuthUser(ctx context.Context, authUserInfo *authModel.AuthUserInfo) (*userModel.User, error) {
	collection := m.db.Collection("User")
	result, err := collection.InsertOne(ctx, mongoModel.NewFromAuthUserInfo(authUserInfo))
	if err != nil {
		return nil, err
	}

	return &userModel.User{
		UserID:   result.InsertedID.(primitive.ObjectID).Hex(),
		AuthID:   authUserInfo.ID,
		UserName: authUserInfo.UserName,
		AuthType: authUserInfo.AuthType,
		Email:    authUserInfo.Email,
		Birthday: authUserInfo.Birthday,
		Gender:   authUserInfo.Gender,
		PhotoURL: authUserInfo.PhotoURL,
		State:    "P",
	}, nil

}

func (m mongoUserRepo) UpdateUser(ctx context.Context, userID string, updater *userModel.UserUpdater) error {
	collection := m.db.Collection("User")

	filter := bson.M{"userId": userID}
	update := bson.M{}

	if updater.UserName != "" {
		update["userName"] = updater.UserName
	}
	if updater.Email != "" {
		update["email"] = updater.Email
	}
	if updater.Birthday != "" {
		update["birthday"] = updater.Birthday
	}
	if updater.Gender != "" {
		update["gender"] = updater.Gender
	}
	if updater.PhotoURL != "" {
		update["photoUrl"] = updater.PhotoURL
	}
	if updater.State != "" {
		update["state"] = updater.State
	}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})

	if err != nil {
		return err
	}
	fmt.Printf("UpdateUser success: %v", result.UpsertedID)
	return nil
}

func (m mongoUserRepo) GetUserByAuthID(ctx context.Context, authID string) (*userModel.User, error) {
	collection := m.db.Collection("User")
	filter := bson.M{"authId": authID}
	mongoUser := mongoModel.User{}
	err := collection.FindOne(ctx, filter).Decode(&mongoUser)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, userModel.UserErrorNotFound
		default:
			return nil, userModel.UserErrorUnknown
		}
	}
	return mongoUser.ToDomainUser(), nil
}

func (m mongoUserRepo) GetUser(ctx context.Context, userID string) (*userModel.User, error) {
	collection := m.db.Collection("User")
	filter := bson.M{"_id": userID}
	mongoUser := mongoModel.User{}
	err := collection.FindOne(ctx, filter).Decode(&mongoUser)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, userModel.UserErrorNotFound
		default:
			return nil, userModel.UserErrorUnknown
		}
	}
	return mongoUser.ToDomainUser(), nil
}
