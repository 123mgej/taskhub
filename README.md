# TaskHub 已经完成内容

 ## 1.项目目标

- 搭建一个后端服务的最小可运行的骨架；
- 后续逐步扩展为任务系统；



## 2. 技术栈（当前已经使用）

- Golang
- Gin（Web框架）



## 3.当前功能

- [x] 启动http服务
- [x] 健康检查接口：
  - GET /healthz
  - 返回 {“code”:0,"message":ok}



## 4.项目结构

```bash
taskhub/

​    cmd/api/main.go      # 程序入口
     internal/router/router.go # 路由与中间件注册
     internal/handle/health.go # /healthz 处理逻辑
     go.mod / go.sum           # Go module 依赖管理
```



## 5.运行方式

- 启动：

  - PORT=9000 ENV=prod go run ./cmd/api

      

- 测试

  -   curl -s http://127.0.0.1:8080/healthz
  

## 6. 明天第一步
1. 配置（env 读取配置）
2. 结构化的日志（写控制台，文件等，可以控制）
3. request_id()

    