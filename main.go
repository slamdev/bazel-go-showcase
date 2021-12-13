package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	cm := &kafka.ConfigMap{}
	fmt.Printf("%+v", cm)
}
