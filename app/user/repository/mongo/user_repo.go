package mongo

import (
	"context"
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

func NewMongoUserRepo(client *mongo.Client) user.Repo {
	return &mongoUserRepo{
		db: client.Database("user"),
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

func (m mongoUserRepo) UpdateUser(ctx context.Context, user *userModel.User) error {
	panic("implement me")
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
