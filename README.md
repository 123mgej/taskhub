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

## 7.中间件调用方式
### 请求进入
1.RequestID()，记录请求的RequestID,没有就生成一个
2.AccessLog()，记录开始1时间，跳到Recovery()中间件
3.Recovery()，捕获panic的话，status设置成500
4.Handle 构造数据返回

### 返回数据
1.Handle 构造response数据返回
2.Recovery()看有没有panic,有的话就捕获panic，status改为500
3.AccessLog() 接收数据，记录耗时，读取需要的字段并记录结构化日志内容
4.RequestID(),返回数据


    