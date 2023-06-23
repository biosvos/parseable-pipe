package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/biosvos/parseable-pipe/internal/parseable"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	stream := flag.String("stream", "", "stream")
	flag.Parse()
	if stream == nil || *stream == "" {
		log.Println("set name")
		return
	}

	auth := "YWRtaW46YWRtaW4K"
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
