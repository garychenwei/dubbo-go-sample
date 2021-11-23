# GOLANGE 客户端实现

## 使用框架依赖

	gofiber
	dubbo-go
	

## 项目结构
	api - 对外api服务
	erp - erp模块，内部service需要token验证
	
## 项目搭建

### 定proto协议文件
	proto/erp

### 配置providr

### 配置consumer

## 踩坑指南

### nacos
使用JAVA 11需要把开始脚本中的变量`JAVA_OPT_EXT_FIX` 删除。

### provider
provider 中services 名字需要是golange中provider中实现功能的struct的名字。interface如果不需要跟JAVA通讯可以自定义，不跟protobuffer中的 rpc 结构有关。

### consumer
consumer中的配置，references的名字必须是pb中clientimpl中的名字统一。interface与provider相同。

## 参考
seata 不同协议中传递xid - http://blog.itpub.net/69953029/viewspace-2757142/
dubbo golange 例子 - https://github.com/apache/dubbo-go-samples
java seata tcc 详解 - https://blog.csdn.net/w1014074794/article/details/116131139
dubbo go auth sample - https://blog.csdn.net/weixin_39860915/article/details/106270958
dubbo go triple sample - https://github.com/dubbogo/triple