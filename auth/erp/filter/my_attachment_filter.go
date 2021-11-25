package filter


import (
	"context"
	"sync"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

var (
	signOnce sync.Once
	sign     *myAttachmentFiler
)

func init() {
	extension.SetFilter("myAttachmentFiler", newMyAttachmentFiler)
}

// signFilter signs the request on consumer side
type myAttachmentFiler struct{}

func newMyAttachmentFiler() filter.Filter {
	if sign == nil {
		signOnce.Do(func() {
			sign = &myAttachmentFiler{}
		})
	}
	return sign
}

// Invoke retrieves the configured Authenticator to add signature to invocation
func (sf *myAttachmentFiler) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	logger.Info("Attachments:", invocation.Attachments())
	return invoker.Invoke(ctx, invocation)
}

// OnResponse dummy process, returns the result directly
func (sf *myAttachmentFiler) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	return result
}
