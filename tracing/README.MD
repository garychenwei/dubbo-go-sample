# GOLANGE 客户端实现

## 使用框架依赖

	gofiber
	dubbo-go
	

## 项目结构
	api - 对外api服务
	erp - erp模块，内部service需要token验证
	
## 项目搭建

### zipkin服务启用
  https://www.cnblogs.com/jmcui/p/10940372.html

### 代码中添加GlobalTracer

	不同的tracing平台添加不同的tracer

### 启用tracing filter配置

  
## 参考
dubbo go samples tracing - https://github.com/apache/dubbo-go-samples