package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type HttpCode int

func Put(url string, header map[string]string, reader io.Reader) (HttpCode, string) {
	request, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return doRequest(request)
}

func Post(url string, header map[string]string, reader io.Reader) (HttpCode, string) {
	request, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return doRequest(request)
}

func doRequest(request *http.Request) (HttpCode, string) {
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
		return HttpCode(rsp.StatusCode), ""
	}

	return HttpCode(rsp.StatusCode), string(all)
}

func isSuccess(code int) bool {
	return code == 200
}

type Record struct {
	Logs string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	name := flag.String("name", "", "name")
	flag.Parse()
	if name == nil || *name == "" {
		log.Println("set name")
		return
	}

	auth := "YWRtaW46YWRtaW4K"
	code, result := Put(fmt.Sprintf("http://127.0.0.1:8000/api/v1/logstream/%v", *name), map[string]string{
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
			if err == io.EOF {
				break
			}
			log.Panicf("%+v", err)
		}
		log.Println(string(line), prefix, err)

		record := Record{Logs: string(line)}
		var buffer bytes.Buffer
		err = json.NewEncoder(&buffer).Encode(&record)
		if err != nil {
			log.Panicf("%+v", err)
		}

		code, post := Post(fmt.Sprintf("http://127.0.0.1:8000/api/v1/logstream/%v", *name), map[string]string{
			"Authorization": fmt.Sprintf("Basic %v", auth),
			"Content-Type":  "application/json",
		}, &buffer)
		log.Println(code, post)
	}
}
