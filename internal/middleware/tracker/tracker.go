package tracker

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

const startTimeCtxKey string = "XXX-ctx-key-tracker-start-time"

func MiddlewareTracker() app.HandlerFunc {
	return func(ctx context.Context, rCtx *app.RequestContext) {
		rCtx.Set(startTimeCtxKey, time.Now().Unix())
		rCtx.Next(ctx)
	}
}

func GetStartTimeFromRequestContext(rCtx *app.RequestContext) time.Time {
	rawStartTime, ok := rCtx.Get(startTimeCtxKey)
	if !ok {
		return time.Time{}
	}
	startTime, ok := rawStartTime.(int64)
	if !ok {
		return time.Time{}
	}
	return time.Unix(startTime, 0)
}
