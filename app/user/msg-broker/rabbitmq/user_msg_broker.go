package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"pixstall-user/app/domain/reg/model"
	"pixstall-user/app/domain/user"
	domainUserModel "pixstall-user/app/domain/user/model"
	msgBrokerModel "pixstall-user/app/user/msg-broker/model"
)

type rabbitMQUserMsgBroker struct {
	ch *amqp.Channel
}

func NewRabbitMQUserMsgBroker(ch *amqp.Channel) user.MsgBroker {
	return &rabbitMQUserMsgBroker{
		ch: ch,
	}
}

func (r *rabbitMQUserMsgBroker) SendRegisterArtistMsg(ctx context.Context, info *model.RegInfo) error {
	b, err := json.Marshal(msgBrokerModel.NewRegInfoFromDomainRegInfo(info))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"user",
		"user.new.isArtist",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitMQUserMsgBroker) SendArtistUpdateMsg(ctx context.Context, updater *domainUserModel.UserUpdater) error {
	b, err := json.Marshal(msgBrokerModel.NewUserUpdaterFromDomainUserUpdater(updater))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"user",
		"user.update.isArtist",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitMQUserMsgBroker) SendRegisterUserMsg(ctx context.Context, info *model.RegInfo) error {
	b, err := json.Marshal(msgBrokerModel.NewRegInfoFromDomainRegInfo(info))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"user",
		"user.new",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitMQUserMsgBroker) SendUserUpdateMsg(ctx context.Context, updater *domainUserModel.UserUpdater) error {
	b, err := json.Marshal(msgBrokerModel.NewUserUpdaterFromDomainUserUpdater(updater))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"user",
		"user.update",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
