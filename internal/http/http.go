package http

import (
	"context"
	"io"
	"log"
	"net/http"
)

type Code int

func Put(ctx context.Context, url string, header map[string]string, reader io.Reader) (Code, string) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, url, reader)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return doRequest(request)
}

func Post(ctx context.Context, url string, header map[string]string, reader io.Reader) (Code, string) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
	if err != nil {
		panic(err)
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return doRequest(request)
}

func doRequest(request *http.Request) (Code, string) {
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
		return Code(rsp.StatusCode), ""
	}

	return Code(rsp.StatusCode), string(all)
}
