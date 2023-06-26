package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"github.com/biosvos/parseable-pipe/internal/parseable"
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

	parser := parseable.NewParseable("http://127.0.0.1:8000", *user, *password)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := parser.CreateTopic(ctx, *topic)
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

		err = parser.Publish(ctx, *topic, string(line))
		if err != nil {
			log.Panicf("%+v", err)
		}
	}
}

func isSet(stream *string) bool {
	return stream != nil && *stream != ""
}
