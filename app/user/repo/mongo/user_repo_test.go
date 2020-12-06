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
	model2 "pixstall-user/app/domain/artist/model"
	"pixstall-user/app/domain/user"
	"pixstall-user/app/domain/user/model"
	mongoModel "pixstall-user/app/user/repo/mongo/model"
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

func TestMongoUserRepo_UpdateUser(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "test_user_id", model.UserStatePending)
	updater := model.UserUpdater{
		UserName:   "new_UserName",
		Email:      "new_Email",
		Birthday:   "20201230",
		Gender:     "F",
		PhotoURL:   "new_PhotoURL",
		State:      model.UserStateActive,
		ArtistInfo: nil,
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
	assert.Equal(t, "new_PhotoURL", mongoUser.PhotoURL)
	assert.Equal(t, model.UserStateActive, mongoUser.State)
}

func TestMongoUserRepo_UpdateUser_userNameOnly(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "test_user_id", model.UserStatePending)
	updater := model.UserUpdater{
		UserName:   "new_UserName",
		State:      model.UserStateActive,
		ArtistInfo: nil,
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
	assert.Equal(t, "Dummy_PhotoURL", mongoUser.PhotoURL)
	assert.Equal(t, model.UserStateActive, mongoUser.State)
}

func TestMongoUserRepo_UpdateUserByAuthID_userNameOnly(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "test_user_id", model.UserStatePending)
	updater := model.UserUpdater{
		UserID:     "new_UserID",
		UserName:   "new_UserName",
		State:      model.UserStateActive,
		ArtistInfo: nil,
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
	assert.Equal(t, "Dummy_PhotoURL", mongoUser.PhotoURL)
	assert.Equal(t, model.UserStateActive, mongoUser.State)
}

func TestMongoUserRepo_UpdateUserByAuthID_BeArtist(t *testing.T) {
	cleanAll()
	objectID := insertDummyUser(ctx, "Dummy_AuthID", model.UserStatePending)
	isArtist := true
	yearOfDrawing := 5
	SelfIntro := "Hello"
	updater := model.UserUpdater{
		UserID:   "new_UserID",
		UserName: "new_UserName",
		State:    model.UserStateActive,
		IsArtist: &isArtist,
		ArtistInfo: &model2.ArtistIntro{
			YearOfDrawing: &yearOfDrawing,
			ArtTypes:      &[]string{"A", "B", "C", "D", "E"},
			SelfIntro:     &SelfIntro,
		},
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
	assert.Equal(t, "Dummy_PhotoURL", mongoUser.PhotoURL)
	assert.True(t, mongoUser.IsArtist)
	assert.NotNil(t, mongoUser.ArtistInfo)
	assert.Equal(t, 5, *mongoUser.ArtistInfo.YearOfDrawing)
	assert.Equal(t, 5, len(*mongoUser.ArtistInfo.ArtTypes))
	assert.Equal(t, "Hello", *mongoUser.ArtistInfo.SelfIntro)
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
		UserID:   userId,
		AuthID:   "Dummy_AuthID",
		UserName: "Dummy_UserName",
		AuthType: "Dummy_AuthType",
		Token:    "Dummy_Token",
		Email:    "Dummy_Email",
		Birthday: "20200101",
		Gender:   "M",
		PhotoURL: "Dummy_PhotoURL",
		State:    state,
	}
	result, err := c.InsertOne(ctx, &user)
	if err != nil {
		panic(err)
	}
	return result.InsertedID.(primitive.ObjectID)
}
