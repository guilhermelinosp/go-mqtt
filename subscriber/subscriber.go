package main

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Define the MQTT broker URL
	broker := "tcp://localhost:1883"

	// Create an MQTT client options
	opts := MQTT.NewClientOptions().AddBroker(broker).SetClientID("mqtt_client")

	// Define the message handler
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})

	// Create and start the MQTT client
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := client.Subscribe("test/topic", 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Wait indefinitely to receive messages
	select {}
}
