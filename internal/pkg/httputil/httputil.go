package httputil

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/middleware/tracker"
	"github.com/bearaujus/go-warehouse-api/internal/pkg"
	"github.com/bearaujus/go-warehouse-api/internal/pkg/errwrap"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"io"
	"log"
	"time"
)

func StartHTTPServer(port string, registerServerFunc func(s *server.Hertz)) error {
	if registerServerFunc == nil {
		return errors.New("nil register server function")
	}

	srvAddress := fmt.Sprintf(":%s", port)
	s := server.Default(server.WithHostPorts(srvAddress))

	registerServerFunc(s)
	log.Print("HTTP server registered")

	sCtx, sCtxCancel := context.WithCancel(context.Background())
	defer sCtxCancel()

	attachedStatusChan := make(chan struct{}, 1)
	go func() {
		pkg.WaitShutdownSigterm(sCtx, func() {
			_ = s.Shutdown(sCtx)
		}, attachedStatusChan)
	}()
	<-attachedStatusChan
	close(attachedStatusChan)

	err := s.Run()
	if err != nil {
		return err
	}

	log.Print("HTTP server shut down gracefully")
	return nil
}

type responseBody struct {
	Header responseBodyHeader `json:"header"`
	Data   interface{}        `json:"data,omitempty"`
}

type responseBodyHeader struct {
	IsSuccess   bool    `json:"is_success"`
	ProcessTime float64 `json:"process_time"`
	Code        string  `json:"code,omitempty"`
	Reason      string  `json:"reason,omitempty"`
	Stack       string  `json:"stack_trace,omitempty"`
}

func WriteResponse(rCtx *app.RequestContext, statusCode int, data interface{}) {
	resp := responseBody{
		Header: responseBodyHeader{
			IsSuccess:   true,
			ProcessTime: time.Since(tracker.GetStartTimeFromRequestContext(rCtx)).Seconds(),
		},
		Data: data,
	}

	rCtx.JSON(statusCode, resp)
}

func WriteEmptyResponse(rCtx *app.RequestContext, statusCode int) {
	WriteResponse(rCtx, statusCode, nil)
}

func WriteErrorResponseAndAbort(rCtx *app.RequestContext, statusCode int, err error) {
	resp := responseBody{
		Header: responseBodyHeader{
			IsSuccess:   false,
			ProcessTime: time.Since(tracker.GetStartTimeFromRequestContext(rCtx)).Seconds(),
			Reason:      err.Error(),
		},
		Data: nil,
	}

	var errWrap errwrap.ErrWrap
	if errors.As(err, &errWrap) {
		resp.Header.Code = errWrap.Code()
		resp.Header.Reason = errWrap.RawError()
		resp.Header.Stack = errWrap.StackTrace()
	}

	rCtx.JSON(statusCode, resp)
	rCtx.Abort()
}

func DecodeUnmarshalResponseBody(r io.ReadCloser, v any) error {
	resp := responseBody{}
	err := json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return err
	}

	d, err := json.Marshal(resp.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, v)
}
