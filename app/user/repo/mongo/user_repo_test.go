package mongo

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	mongoModel "pixstall-user/app/user/repo/mongo/model"
	model2 "pixstall-user/domain/commission/model"
	indexModel "pixstall-user/domain/inbox/model"
	"pixstall-user/domain/user"
	"pixstall-user/domain/user/model"
	"testing"
)

var db *mongo.Database
var dbClient *mongo.Client
var repo user.Repo
var ctx context.Context

const (
	TestDBName = "pixstall-user-test"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before all tests")
	ctx = context.TODO()
	var err error
	dbClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	db = dbClient.Database(TestDBName)
	repo = NewMongoUserRepo(db)
}

func teardown() {
	dropAll()
	fmt.Println("After all tests")
	err := dbClient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}

func TestMongoUserRepo_GetUser(t *testing.T) {
	cleanAll()
	_ = insertDummyUser(ctx, "test_user_id", model.UserStateActive)

	dUser, err := repo.GetUser(ctx, "test_user_id")
	assert.NoError(t, err)

	assert.Equal(t, "test_user_id", dUser.UserID)
	assert.Equal(t, "Dummy_UserName", dUser.UserName)
	assert.Equal(t, "Dummy_ProfilePath", dUser.ProfilePath)
	assert.Equal(t, model.UserStateActive, dUser.State)

	assert.Empty(t, dUser.AuthID)
	assert.Empty(t, dUser.Email)
	assert.Empty(t, dUser.Gender)
}

func TestMongoUserRepo_GetUserDetails(t *testing.T) {
	cleanAll()
	_ = insertDummyUser(ctx, "test_user_id", model.UserStateActive)

	dUser, err := repo.GetUserDetails(ctx, "test_user_id")
	assert.NoError(t, err)

	assert.Equal(t, "test_user_id", dUser.UserID)
	assert.Equal(t, "Dummy_UserName", dUser.UserName)
	assert.Equal(t, "Dummy_ProfilePath", dUser.ProfilePath)
	assert.Equal(t, model.UserStateActive, dUser.State)

	assert.Equal(t, "Dummy_AuthID", dUser.AuthID)
	assert.Equal(t, "Dummy_Email", dUser.Email)
	assert.Equal(t, "M", dUser.Gender)
}

func TestMongoUserRepo_UpdateUser(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "test_user_id", model.UserStatePending)

	userName := "new_UserName"
	email := "new_Email"
	birthday := "20201230"
	gender := "F"
	profilePath := "new_ProfilePath"
	state := model.UserStateActive
	updater := model.UserUpdater{
		UserName:    &userName,
		Email:       &email,
		Birthday:    &birthday,
		Gender:      &gender,
		ProfilePath: &profilePath,
		State:       &state,
	}
	err := repo.UpdateUser(ctx, "test_user_id", &updater)
	assert.NoError(t, err)

	mongoUser := mongoModel.User{}
	err = db.Collection(UserCollection).FindOne(ctx, bson.M{"_id": objectID}).Decode(&mongoUser)
	assert.NoError(t, err)

	assert.Equal(t, "new_UserName", mongoUser.UserName)
	assert.Equal(t, "new_Email", mongoUser.Email)
	assert.Equal(t, "20201230", mongoUser.Birthday)
	assert.Equal(t, "F", mongoUser.Gender)
	assert.Equal(t, "new_ProfilePath", mongoUser.ProfilePath)
	assert.Equal(t, model.UserStateActive, mongoUser.State)
}

func TestMongoUserRepo_UpdateUser_userNameOnly(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "test_user_id", model.UserStatePending)
	userName := "new_UserName"
	state := model.UserStateActive
	updater := model.UserUpdater{
		UserName: &userName,
		State:    &state,
	}
	err := repo.UpdateUser(ctx, "test_user_id", &updater)
	assert.NoError(t, err)

	mongoUser := mongoModel.User{}
	err = db.Collection(UserCollection).FindOne(ctx, bson.M{"_id": objectID}).Decode(&mongoUser)
	assert.NoError(t, err)

	assert.Equal(t, "new_UserName", mongoUser.UserName)
	assert.Equal(t, "Dummy_Email", mongoUser.Email)
	assert.Equal(t, "20200101", mongoUser.Birthday)
	assert.Equal(t, "M", mongoUser.Gender)
	assert.Equal(t, "Dummy_ProfilePath", mongoUser.ProfilePath)
	assert.Equal(t, model.UserStateActive, mongoUser.State)
}

func TestMongoUserRepo_UpdateUserByAuthID_userNameOnly(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "test_user_id", model.UserStatePending)
	userName := "new_UserName"
	state := model.UserStateActive
	updater := model.UserUpdater{
		UserID:   "new_UserID",
		UserName: &userName,
		State:    &state,
	}
	err := repo.UpdateUserByAuthID(ctx, "Dummy_AuthID", &updater)
	assert.NoError(t, err)

	mongoUser := mongoModel.User{}
	err = db.Collection(UserCollection).FindOne(ctx, bson.M{"_id": objectID}).Decode(&mongoUser)
	assert.NoError(t, err)

	assert.Equal(t, "new_UserID", mongoUser.UserID)
	assert.Equal(t, "new_UserName", mongoUser.UserName)
	assert.Equal(t, "Dummy_Email", mongoUser.Email)
	assert.Equal(t, "20200101", mongoUser.Birthday)
	assert.Equal(t, "M", mongoUser.Gender)
	assert.Equal(t, "Dummy_ProfilePath", mongoUser.ProfilePath)
	assert.Equal(t, model.UserStateActive, mongoUser.State)
}

func TestMongoUserRepo_UpdateUserByAuthID_BeArtist(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "Dummy_AuthID", model.UserStatePending)
	isArtist := true
	userName := "new_UserName"
	state := model.UserStateActive
	updater := model.UserUpdater{
		UserID:   "new_UserID",
		UserName: &userName,
		State:    &state,
		IsArtist: &isArtist,
	}
	err := repo.UpdateUserByAuthID(ctx, "Dummy_AuthID", &updater)
	assert.NoError(t, err)

	mongoUser := mongoModel.User{}
	err = db.Collection(UserCollection).FindOne(ctx, bson.M{"_id": objectID}).Decode(&mongoUser)
	assert.NoError(t, err)

	assert.Equal(t, "new_UserID", mongoUser.UserID)
	assert.Equal(t, "new_UserName", mongoUser.UserName)
	assert.Equal(t, "Dummy_Email", mongoUser.Email)
	assert.Equal(t, "20200101", mongoUser.Birthday)
	assert.Equal(t, "M", mongoUser.Gender)
	assert.Equal(t, "Dummy_ProfilePath", mongoUser.ProfilePath)
	assert.True(t, mongoUser.IsArtist)
	assert.Equal(t, model.UserStateActive, mongoUser.State)
}

//Private
func cleanAll() {
	_, err := db.Collection(UserCollection).DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}
}

func dropAll() {
	err := db.Collection(UserCollection).Drop(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
}

func insertDummyUser(ctx context.Context, userId string, state model.UserState) primitive.ObjectID {
	c := db.Collection(UserCollection)

	user := mongoModel.User{
		UserID:      userId,
		AuthID:      "Dummy_AuthID",
		UserName:    "Dummy_UserName",
		AuthType:    "Dummy_AuthType",
		Email:       "Dummy_Email",
		Birthday:    "20200101",
		Gender:      "M",
		ProfilePath: "Dummy_ProfilePath",
		State:       state,
		Inbox: indexModel.Inbox{
			UnreadMessageCount: 10,
			LastChat:           indexModel.Chat{
				ID:          "12345",
				UserIDs:     nil,
				LastMessage: "Dummy_LastMessage",
				Read:        "",
			},
		},
		Commission: model2.Commission{
			MessageCount: 10,
			LastMessage:  "Dummy_LastMessage",
		},
	}
	result, err := c.InsertOne(ctx, &user)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}
