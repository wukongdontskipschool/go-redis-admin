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
bin // 64位 二进制软件包  
|  
cmd  
| - admin // main.go为入口文件  
internal  
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

### 启动配置
要求安装mysql  
1. 执行rdadmin_linux -init 生成配置  
2. 配置数据库配置internal/configs/databases/redisAdmin.yaml  
3. 再次执行rdadmin_linux -init 初始化数据表  
4. 执行rdadmin_linux -port=12367 指定端口  

### 基本业务流程

路由 -> api控制器 -> 具体业务  
internal/router/router.go -> internal/api -> internal/job  

### 暂未完成
1. 暂时只支持读redis  
2. 参数未作校验  
3. 前端界面未完善  

![](show1.jpg)
