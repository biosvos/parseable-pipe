package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type HTTPCode int

func Put(ctx context.Context, url string, header map[string]string, reader io.Reader) (HTTPCode, string) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, url, reader)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return doRequest(request)
}

func Post(ctx context.Context, url string, header map[string]string, reader io.Reader) (HTTPCode, string) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return doRequest(request)
}

func doRequest(request *http.Request) (HTTPCode, string) {
	rsp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := rsp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	all, err := io.ReadAll(rsp.Body)
	if err != nil {
		return HTTPCode(rsp.StatusCode), ""
	}

	return HTTPCode(rsp.StatusCode), string(all)
}

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
	code, result := Put(ctx, fmt.Sprintf("http://127.0.0.1:8000/api/v1/logstream/%v", *name), map[string]string{
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

		code, post := Post(ctx, fmt.Sprintf("http://127.0.0.1:8000/api/v1/logstream/%v", *name), map[string]string{
			"Authorization": fmt.Sprintf("Basic %v", auth),
			"Content-Type":  "application/json",
		}, &buffer)
		log.Println(code, post)
	}
}
