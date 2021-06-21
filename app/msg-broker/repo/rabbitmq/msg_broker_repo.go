package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	msg2 "pixstall-user/app/msg-broker/repo/rabbitmq/msg"
	model3 "pixstall-user/domain/artist/model"
	"pixstall-user/domain/commission/model"
	msg_broker "pixstall-user/domain/msg-broker"
	model2 "pixstall-user/domain/reg/model"
	userModel "pixstall-user/domain/user/model"
)

type rabbitmqMsgBrokerRepo struct {
	ch *amqp.Channel
}

func NewRabbitMQMsgBrokerRepo(conn *amqp.Connection) msg_broker.Repo {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}
	return &rabbitmqMsgBrokerRepo{
		ch: ch,
	}
}

func (r *rabbitmqMsgBrokerRepo) SendCreateArtistCmd(ctx context.Context, info *model2.RegInfo) error {
	b, err := json.Marshal(msg2.NewCreateArtistCmdMsg(*info))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"artist",
		"artist.cmd.create",
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

func (r *rabbitmqMsgBrokerRepo) SendUpdateArtistCmd(ctx context.Context, updater *model3.ArtistUpdater) error {
	b, err := json.Marshal(msg2.NewUpdateArtistCmdMsg(*updater))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"artist",
		"artist.cmd.update",
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

func (r *rabbitmqMsgBrokerRepo) SendUserUpdatedEvent(ctx context.Context, updater *userModel.UserUpdater) error {
	b, err := json.Marshal(msg2.NewUserUpdatedEventMsg(*updater))
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"user",
		"user.event.updated",
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

func (r *rabbitmqMsgBrokerRepo) SendCommissionUserValidationEvent(ctx context.Context, usersValidation model.CommissionUsersValidation) error {
	vValidation := msg2.CommissionUsersValidationEventMsg{
		CommissionUsersValidation: usersValidation,
	}
	b, err := json.Marshal(vValidation)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"commission",
		"commission.event.validation.users",
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

func (r *rabbitmqMsgBrokerRepo) SendUpdateCommissionCmd(ctx context.Context, updater model.CommissionUpdater) error {
	vUpdater := msg2.UpdateCommissionCmdMsg{
		CommissionUpdater: updater,
	}
	b, err := json.Marshal(vUpdater)
	if err != nil {
		return err
	}
	err = r.ch.Publish(
		"commission",
		"commission.cmd.update",
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