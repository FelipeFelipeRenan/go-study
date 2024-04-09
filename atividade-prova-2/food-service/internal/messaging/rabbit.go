package messaging

import "github.com/streadway/amqp"


func InitRabbitMQ()(*amqp.Connection, error)  {
	conn , err := amqp.Dial("ampq://guest:guest@localhost:5726/")
	if err != nil {
		return nil, err
	}
	return conn, nil
}