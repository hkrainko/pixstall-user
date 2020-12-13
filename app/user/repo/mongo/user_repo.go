package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	mongoModel "pixstall-user/app/user/repo/mongo/model"
	authModel "pixstall-user/domain/auth/model"
	"pixstall-user/domain/user"
	userModel "pixstall-user/domain/user/model"
)

type mongoUserRepo struct {
	db *mongo.Database
}

const (
	UserCollection = "Users"
)

func NewMongoUserRepo(db *mongo.Database) user.Repo {
	return &mongoUserRepo{
		db: db,
	}
}

func (m mongoUserRepo) SaveAuthUser(ctx context.Context, authUserInfo *authModel.AuthUserInfo) (*userModel.User, error) {
	collection := m.db.Collection(UserCollection)
	result, err := collection.InsertOne(ctx, mongoModel.NewFromAuthUserInfo(authUserInfo))
	if err != nil {
		return nil, err
	}

	fmt.Printf("SaveAuthUser %v success", result.InsertedID.(primitive.ObjectID).Hex())
	return &userModel.User{
		AuthID:   authUserInfo.ID,
		UserName: authUserInfo.UserName,
		AuthType: authUserInfo.AuthType,
		Email:    authUserInfo.Email,
		Birthday: authUserInfo.Birthday,
		Gender:   authUserInfo.Gender,
		State:    userModel.UserStatePending,
	}, nil

}

func (m mongoUserRepo) UpdateUser(ctx context.Context, userID string, updater *userModel.UserUpdater) error {
	collection := m.db.Collection(UserCollection)

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
	if updater.ProfilePath != "" {
		update["profilePath"] = updater.ProfilePath
	}
	if updater.State != "" {
		update["state"] = updater.State
	}
	if updater.IsArtist != nil {
		update["isArtist"] = updater.IsArtist
	}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})

	if err != nil {
		return err
	}
	fmt.Printf("UpdateUser success: %v", result.UpsertedID)
	return nil
}

func (m mongoUserRepo) UpdateUserByAuthID(ctx context.Context, authID string, updater *userModel.UserUpdater) error {
	collection := m.db.Collection(UserCollection)

	filter := bson.M{"authId": authID}
	update := bson.M{}

	if updater.UserID != "" {
		update["userId"] = updater.UserID
	} else {
		//TODO: return error
	}
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
	if updater.ProfilePath != "" {
		update["profilePath"] = updater.ProfilePath
	}
	if updater.State != "" {
		update["state"] = updater.State
	}
	if updater.IsArtist != nil {
		update["isArtist"] = updater.IsArtist
	}

	result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update})

	if err != nil {
		return err
	}
	fmt.Printf("UpdateUser success: %v", result.UpsertedID)
	return nil
}

func (m mongoUserRepo) GetUserByAuthID(ctx context.Context, authID string) (*userModel.User, error) {
	collection := m.db.Collection(UserCollection)
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
	collection := m.db.Collection(UserCollection)
	filter := bson.M{"userId": userID}
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

func (m mongoUserRepo) IsUserExist(ctx context.Context, userID string) (*bool, error) {
	var exist bool
	collection := m.db.Collection(UserCollection)
	filter := bson.M{"userId": userID}
	mongoUser := mongoModel.User{}
	err := collection.FindOne(ctx, filter).Decode(&mongoUser)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			exist = false
			return &exist, nil
		default:
			return nil, userModel.UserErrorUnknown
		}
	}
	exist = true
	return &exist, nil
}
