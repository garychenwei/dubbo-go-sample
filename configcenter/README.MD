# GOLANGE 客户端实现

## 使用框架依赖

	gofiber
	dubbo-go
	

## 项目结构
	api - 对外api服务
	erp - erp模块，内部service需要token验证
	
## 项目搭建

### nacos 上添加配置  
  config 文件中的data-id 匹配即可。原来basic项目的config复制到nacos的配置表中。

## 参考
dubbo go sample 中的configcenter