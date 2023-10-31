package utils

import (
	"context"
	"fmt"
)

func Logger(ctx context.Context, log func(format string, v ...interface{}), method string, service string, requestID string, v ...interface{}) {
	if ctx.Err() != context.Canceled {
		logTxt := fmt.Sprintf("%s -> %s ::: %s --> %v", method, service, requestID, v)
		log("%s", logTxt)
	}
}
