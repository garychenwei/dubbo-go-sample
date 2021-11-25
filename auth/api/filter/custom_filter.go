package filter

import (
	"context"
	"fmt"
	"sync"

	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

var (
	signOnce sync.Once
	sign     *mySignFilter
)

func init() {
	extension.SetFilter("mySignFiler", newMySignFilter)
}

// signFilter signs the request on consumer side
type mySignFilter struct{}

func newMySignFilter() filter.Filter {
	if sign == nil {
		signOnce.Do(func() {
			sign = &mySignFilter{}
		})
	}
	return sign
}

// Invoke retrieves the configured Authenticator to add signature to invocation
func (sf *mySignFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	logger.Info("call the mySignFilter and call the auth")
	url := invoker.GetURL()

	err := myDoAuthWork(url, func(authenticator filter.Authenticator) error {
		tempErr := authenticator.Sign(invocation, url)
		// pring the attament
		logger.Info("Attachments:", invocation.Attachments())
		return tempErr
	})
	if err != nil {
		panic(fmt.Sprintf("Sign for invocation %s # %s failed msg %s", url.ServiceKey(), invocation.MethodName(), err.Error()))
	}
	return invoker.Invoke(ctx, invocation)
}

// OnResponse dummy process, returns the result directly
func (sf *mySignFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	return result
}

func myDoAuthWork(url *common.URL, do func(filter.Authenticator) error) error {
	shouldAuth := url.GetParamBool(constant.ServiceAuthKey, false)
	if shouldAuth {
		authenticator := extension.GetAuthenticator(url.GetParam(constant.AuthenticatorKey, constant.DefaultAuthenticator))
		return do(authenticator)
	}
	return nil
}