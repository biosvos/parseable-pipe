package parseable

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/biosvos/parseable-pipe/internal/http"
	"github.com/pkg/errors"
)

type Parseable struct {
	url  string
	auth string
}

func NewParseable(url string, auth string) *Parseable {
	return &Parseable{url: url, auth: auth}
}

func (p *Parseable) CreateStream(ctx context.Context, name string) error {
	code, result := http.Put(ctx, fmt.Sprintf("%v/api/v1/logstream/%v", p.url, name), map[string]string{
		"Authorization": p.auth,
	}, nil)
	switch code {
	case 200: // 성공
		return nil
	case 400: // 이미 존재
		return nil
	default:
		return errors.Errorf("%v: %v", code, result)
	}
}

type Record struct {
	Logs string `json:"logs"`
}

func (p *Parseable) SendLog(ctx context.Context, name string, log string) error {
	record := Record{Logs: log}
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(&record)
	if err != nil {
		return errors.WithStack(err)
	}

	code, post := http.Post(ctx, fmt.Sprintf("%v/api/v1/logstream/%v", p.url, name), map[string]string{
		"Authorization": p.auth,
		"Content-Type":  "application/json",
	}, &buffer)
	if code != 200 {
		return errors.Errorf("%v: %v", code, post)
	}
	return nil
}