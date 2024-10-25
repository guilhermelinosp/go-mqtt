package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Define the MQTT broker URL
	broker := "tcp://localhost:1883"

	// Create an MQTT client options
	opts := MQTT.NewClientOptions().AddBroker(broker).SetClientID("mqtt_server")
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Message received: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})

	// Create and start the MQTT client
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Publish messages
	for {
		token := client.Publish("test/topic", 0, false, "Hello from MQTT Server!")
		token.Wait()
		fmt.Println("Message published")
		time.Sleep(1 * time.Second) // Wait 1 second before the next message
	}
}
