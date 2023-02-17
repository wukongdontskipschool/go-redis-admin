# go-redis-admin

基于gin的redis mvc架构管理工具  
布局参考 https://github.com/golang-standards/project-layout  
遵循RESTFUL网络请求  
遵循单一设计模式原则  

主要扩展有：
- 权限管理 casbin
- 数据库orm gorm
- redis go-redis
- 前端模板基于layui的WeAdmin

### 主要目录结构
cmd  
| - admin // main.go为入口文件  
| - internal  
　　| - accessControl  // 权限验证  
　　| - api // 路由执行控制器  
　　| - configs // 配置文件  
　　| - consts // 常量  
　　| - databases // 数据库操作  
　　| - jobs // 实际业务  
　　| - redisPool // redis池  
　　| - router // 路由  
　　| - services // 服务  
　　| - web // 静态页面  

### 配置
目前监听12345端口  
internal/configs/env/env.yaml.example 复制为internal/configs/env/env.yaml  
internal/configs/databases/redisAdmin.yaml 复制为internal/configs/databases/redisAdmin.yaml 配置数据库（需手动建库）  
准备redis 127.0.0.1 端口6379  

http://127.0.0.1:12356/v1/admin/migrate 请求初始化表  
show.jpg　管理页面初步展示  

### 基本业务流程

路由 -> api控制器 -> 具体业务  
internal/router/router.go -> internal/api -> internal/job  

![](show1.jpg)
