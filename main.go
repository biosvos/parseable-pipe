package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"github.com/biosvos/parseable-pipe/internal/parseable"
	"io"
	"log"
	"os"
)

func encodeAuth(user, password string) string {
	form := fmt.Sprintf("%v:%v", user, password)
	return base64.StdEncoding.EncodeToString([]byte(form))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	stream := flag.String("stream", "", "stream")
	user := flag.String("user", "", "user")
	password := flag.String("password", "", "password")

	flag.Parse()
	if !isSet(stream) {
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

	auth := encodeAuth(*user, *password)
	parser := parseable.NewParseable("http://127.0.0.1:8000", fmt.Sprintf("Basic %v", auth))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := parser.CreateStream(ctx, *stream)
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

		err = parser.SendLog(ctx, *stream, string(line))
		if err != nil {
			log.Panicf("%+v", err)
		}
	}
}

func isSet(stream *string) bool {
	return stream != nil && *stream != ""
}
