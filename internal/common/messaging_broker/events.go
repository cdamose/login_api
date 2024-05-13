package messagingbroker

type Events interface {
	Publish(topic string, mesage string)
	Subscribe(topic string)
}
