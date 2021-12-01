package filter

 import (
	 "errors"
	 "fmt"
	 "strconv"
	 "time"
 )
 
 import (
	 "dubbo.apache.org/dubbo-go/v3/common"
	 "dubbo.apache.org/dubbo-go/v3/common/constant"
	 "dubbo.apache.org/dubbo-go/v3/common/extension"
	 "dubbo.apache.org/dubbo-go/v3/filter"
	 "dubbo.apache.org/dubbo-go/v3/protocol"
	 invocation_impl "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	 "dubbo.apache.org/dubbo-go/v3/filter/auth"
	 "dubbo.apache.org/dubbo-go/v3/common/logger"
 )
 
 func init() {
	 extension.SetAuthenticator(constant.DefaultAuthenticator, func() filter.Authenticator {
		 return &MyDefaultAuthenticator{}
	 })
 }
 
 // DefaultAuthenticator is the default implementation of Authenticator
 type MyDefaultAuthenticator struct{}
 
 // Sign adds the signature to the invocation
 func (authenticator *MyDefaultAuthenticator) Sign(invocation protocol.Invocation, url *common.URL) error {
	 currentTimeMillis := strconv.Itoa(int(time.Now().Unix() * 1000))
 
	 consumer := url.GetParam(constant.ApplicationKey, "")
	 accessKeyPair, err := getAccessKeyPair(invocation, url)
	 if err != nil {
		 return errors.New("get accesskey pair failed, cause: " + err.Error())
	 }
	 inv := invocation.(*invocation_impl.RPCInvocation)
	 signature, err := getSignature(url, invocation, accessKeyPair.SecretKey, currentTimeMillis)
	 if err != nil {
		 return err
	 }
	 inv.SetAttachments(constant.RequestSignatureKey, signature)
	 inv.SetAttachments(constant.RequestTimestampKey, currentTimeMillis)
	 inv.SetAttachments(constant.AKKey, accessKeyPair.AccessKey)
	 inv.SetAttachments(constant.Consumer, consumer)
	 return nil
 }
 
 // getSignature
 // get signature by the metadata and params of the invocation
 func getSignature(url *common.URL, invocation protocol.Invocation, secrectKey string, currentTime string) (string, error) {
	 requestString := fmt.Sprintf(constant.SignatureStringFormat,
		 url.ColonSeparatedKey(), invocation.MethodName(), secrectKey, currentTime)
	 var signature string
	 if parameterEncrypt := url.GetParamBool(constant.ParameterSignatureEnableKey, false); parameterEncrypt {
		 var err error
		 if signature, err = auth.SignWithParams(invocation.Arguments(), requestString, secrectKey); err != nil {
			 // TODO
			 return "", errors.New("sign the request with params failed, cause:" + err.Error())
		 }
	 } else {
		 signature = auth.Sign(requestString, secrectKey)
	 }
 
	 return signature, nil
 }
 
 // Authenticate verifies whether the signature sent by the requester is correct
 func (authenticator *MyDefaultAuthenticator) Authenticate(invocation protocol.Invocation, url *common.URL) error {
	 attachments := invocation.Attachments()
	 accessKeyId := ""
	 requestTimestamp := ""
	 originSignature := ""
	 consumer := ""
	 for k, v := range attachments {
	 	switch k{
			case constant.AKKey:
				accessKeyId = v.([]string)[0]
			case constant.RequestTimestampKey:
				requestTimestamp = v.([]string)[0]
			case constant.RequestSignatureKey:
				originSignature = v.([]string)[0]
			case constant.Consumer:
				consumer = v.([]string)[0]
		}
	 }
	 logger.Info("comm url=", url)
	 if auth.IsEmpty(accessKeyId, false) || auth.IsEmpty(consumer, false) ||
	 		auth.IsEmpty(requestTimestamp, false) || auth.IsEmpty(originSignature, false) {
		 return errors.New("failed to authenticate your ak/sk, maybe the consumer has not enabled the auth")
	 }
 
	 accessKeyPair, err := getAccessKeyPair(invocation, url)
	 if err != nil {
		 return errors.New("failed to authenticate , can't load the accessKeyPair")
	 }
 
	 computeSignature, err := getSignature(url, invocation, accessKeyPair.SecretKey, requestTimestamp)
	 if err != nil {
		 return err
	 }
	 if success := computeSignature == originSignature; !success {
		 return errors.New("failed to authenticate, signature is not correct")
	 }
	 logger.Info("auth success ",computeSignature, originSignature)
	 return nil
 }

 
 func getAccessKeyPair(invocation protocol.Invocation, url *common.URL) (*filter.AccessKeyPair, error) {
	 accesskeyStorage := extension.GetAccessKeyStorages(url.GetParam(constant.AccessKeyStorageKey, constant.DefaultAccessKeyStorage))
	 accessKeyPair := accesskeyStorage.GetAccessKeyPair(invocation, url)
	 if accessKeyPair == nil || auth.IsEmpty(accessKeyPair.AccessKey, false) || auth.IsEmpty(accessKeyPair.SecretKey, true) {
		 return nil, errors.New("accessKeyId or secretAccessKey not found")
	 } else {
		 return accessKeyPair, nil
	 }
 }
 