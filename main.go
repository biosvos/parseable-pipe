package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"github.com/biosvos/parseable-pipe/internal/mqtt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	topic := flag.String("topic", "", "topic")
	user := flag.String("user", "", "user")
	password := flag.String("password", "", "password")

	flag.Parse()
	if !isSet(topic) {
		log.Println("set stream")
		return
	}
	if !isSet(user) {
		log.Println("set user")
		return
	}
	if !isSet(password) {
		log.Println("set password")
		return
	}

	broker, err := mqtt.NewMqtt()
	if err != nil {
		log.Panicf("%+v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = broker.CreateTopic(ctx, *topic)
	if err != nil {
		log.Panicf("%+v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, prefix, err := reader.ReadLine()
		if prefix {
			break
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Panicf("%+v", err)
		}

		err = broker.Publish(ctx, *topic, string(line))
		if err != nil {
			log.Panicf("%+v", err)
		}
	}
}

func isSet(stream *string) bool {
	return stream != nil && *stream != ""
}
