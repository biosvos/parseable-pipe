package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/biosvos/parseable-pipe/internal/http"
	"io"
	"log"
	"os"
)

type Record struct {
	Logs string `json:"logs"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	name := flag.String("name", "", "name")
	flag.Parse()
	if name == nil || *name == "" {
		log.Println("set name")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	auth := "YWRtaW46YWRtaW4K"
	code, result := http.Put(ctx, fmt.Sprintf("http://127.0.0.1:8000/api/v1/logstream/%v", *name), map[string]string{
		"Authorization": fmt.Sprintf("Basic %v", auth),
	}, nil)
	log.Println(result)
	switch code {
	case 200: // 성공
	case 400: // 이미 존재
	default:
		log.Panicf(result)
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

		record := Record{Logs: string(line)}
		var buffer bytes.Buffer
		err = json.NewEncoder(&buffer).Encode(&record)
		if err != nil {
			log.Panicf("%+v", err)
		}

		code, post := http.Post(ctx, fmt.Sprintf("http://127.0.0.1:8000/api/v1/logstream/%v", *name), map[string]string{
			"Authorization": fmt.Sprintf("Basic %v", auth),
			"Content-Type":  "application/json",
		}, &buffer)
		log.Println(code, post)
	}
}
