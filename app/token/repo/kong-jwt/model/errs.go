package model

type DuplicatedConsumerError struct {}

func (d *DuplicatedConsumerError) Error() string {
	return "DuplicatedConsumer"
}