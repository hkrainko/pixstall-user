package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pixstall-user/domain/commission"
	"pixstall-user/domain/commission/model"
	"time"
)

type CommissionMessageBroker struct {
	commUseCase commission.UseCase
	ch          *amqp.Channel
}

func NewRabbitMQCommissionMessageBroker(commUseCase commission.UseCase, conn *amqp.Connection) CommissionMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return CommissionMessageBroker{
		commUseCase: commUseCase,
		ch:          ch,
	}
}

func (c CommissionMessageBroker) StartCommUsersValidateQueue() {
	//TODO
	q, err := c.ch.QueueDeclare(
		"commission-users-validate", // name
		true,                        // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		nil,                         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = c.ch.QueueBind(
		q.Name,
		"commission.event.created",
		"commission",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue %v", err)
	}

	msgs, err := c.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			go func() {
				for {
					select {
					case <-ctx.Done():
						switch ctx.Err() {
						case context.DeadlineExceeded:
							log.Println("context.DeadlineExceeded")
						case context.Canceled:
							log.Println("context.Canceled")
						default:
							log.Println("default")
						}
						return // returning not to leak the goroutine
					}
				}
			}()

			switch d.RoutingKey {
			case "commission.event.created":
				err := c.newCommissionCreated(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			default:
				cancel()
			}

		}
	}()

	<-forever
}

func (c CommissionMessageBroker) StopAllQueue() {
	err := c.ch.Close()
	if err != nil {
		log.Printf("StopCommissionQueue err %v", err)
	}
	log.Printf("StopCommissionQueue success")
}

func (c CommissionMessageBroker) newCommissionCreated(ctx context.Context, body []byte) error {
	comm := model.Commission{}
	err := json.Unmarshal(body, &comm)
	if err != nil {
		return err
	}

	return c.commUseCase.HandleNewCreatedCommission(ctx, comm)
}
