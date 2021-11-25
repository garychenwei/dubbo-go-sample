# GOLANGE 客户端实现

## 使用框架依赖

	gofiber
	dubbo-go
	

## 项目结构
	api - 对外api服务
	erp - erp模块，内部service需要token验证

## 踩坑指南

### auth/sign 配置

provider 中添加

        # 本服务开启auth
        auth: true
        # 启用auth filter，对请求进行验签
        # filter:  "auth,tracing,accesslog"
        filter:  "auth,myAttachmentFiler"
        # 默认实现通过配置文件配置AK、SK
        params:
          .accessKeyId: "SYD8-23DF"
          .secretAccessKey: "BSDY-FDF1"

官方的 auth filter 需要使用auth: true配置打开。

consumer 配置

        consumer:
          filter: "mySignFiler"

只能添加signfilter中，例子中的mySignFiler为了方便看数据。需要在references中添加params。秘钥需要跟服务者同样。官方的filter中会通过注册中心发送accessKeyId，secretAccessKey。

### 踩坑官方default_authenticator.go

其中Attachments不能强制转换为string，例子中获取的数据为[]string。需要重写为：MyDefaultAuthenticator

### next step
为防止秘钥在注册中心以明文出现，下一步使用redis管理。参考官方的例子使用redis 保存秘钥。不同分组或不consumer使用不同的秘钥。
客户端也是对不同的服务使用不同的秘钥。

## 参考
auth filter 例子 - https://developer.aliyun.com/article/762837
dubbo go 3 配置 - https://dubbogo.github.io/zh-cn/docs/user/concept/configuration.html