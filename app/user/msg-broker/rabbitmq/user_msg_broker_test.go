package rabbitmq

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	model2 "pixstall-user/domain/artist/model"
	"pixstall-user/domain/reg/model"
	"pixstall-user/domain/user"
	"testing"
)

var conn *amqp.Connection
var userMsgBroker user.MsgBroker
var ctx context.Context

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Before all tests")
	ctx = context.Background()
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	err = ch.ExchangeDeclare(
		"user",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create exchange %v", err)
	}

	userMsgBroker = NewRabbitMQUserMsgBroker(ch)
}

func teardown() {
	fmt.Println("After all tests")
	conn.Close()
}

func TestRabbitMQUserMsgBroker_SendRegisterArtistMsg(t *testing.T) {
	err := userMsgBroker.SendRegisterArtistMsg(ctx, &model.RegInfo{
		AuthID:      "123",
		UserID:      "123",
		DisplayName: "1231",
		Email:       "helloTest",
		Birthday:    "20102020",
		Gender:      "M",
		ProfilePath: "test/path",
		RegAsArtist: true,
		RegArtistIntro: model2.ArtistIntro{
			YearOfDrawing: nil,
			ArtTypes:      nil,
		},
	})
	assert.NoError(t, err)
}

func TestRabbitMQUserMsgBroker_SendRegisterArtistMsg_repeat(t *testing.T) {

	for i := 0; i < 10; i++ {
		err := userMsgBroker.SendRegisterArtistMsg(ctx, &model.RegInfo{
			AuthID:      "123",
			UserID:      "123",
			DisplayName: "1231",
			Email:       "helloTest",
			Birthday:    "20102020",
			Gender:      "M",
			ProfilePath: "test/path",
			RegAsArtist: true,
			RegArtistIntro: model2.ArtistIntro{
				YearOfDrawing: nil,
				ArtTypes:      nil,
			},
		})
		assert.NoError(t, err)
	}
}
