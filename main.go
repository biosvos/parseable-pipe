package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Put(url string, header map[string]string, reader io.Reader) string {
	request, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return doRequest(request)
}

func doRequest(request *http.Request) string {
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
		panic(err)
	}

	if !isSuccess(rsp.StatusCode) {
		log.Panicf("%+v %v", rsp, string(all))
	}
	return string(all)
}

func isSuccess(code int) bool {
	return code == 200
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
	result := Put(fmt.Sprintf("http://127.0.0.1:8000/api/v1/logstream/%v", *name), map[string]string{
		"Authorization": fmt.Sprintf("Basic %v", auth),
	}, nil)
	log.Println(result)

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
	}
}
