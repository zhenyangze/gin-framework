# gin-framework

### 启动方式
```golang
go run main.go -t [web,cron,rps,all]
```

### 目录说明

```shell
├── README.md
├── app
├── boot
├── configs
├── go.mod
├── go.sum
├── helpers
├── main.go
├── resources
└── routes
```

### 测试url

```shell
air
```

```shell
[GIN-debug] GET    /v1/test                  --> gitee.com/zhenyangze/gin-framework/app/http/controllers.TestHandle (6 handlers)
[GIN-debug] GET    /v1/view/:name/*action    --> gitee.com/zhenyangze/gin-framework/app/http/controllers.ViewHandle (6 handlers)
[GIN-debug] GET    /v1/view                  --> gitee.com/zhenyangze/gin-framework/app/http/controllers.ViewHandle (6 handlers)
[GIN-debug] POST   /v1/valid                 --> gitee.com/zhenyangze/gin-framework/app/http/controllers.ValidHandle (6 handlers)
[GIN-debug] GET    /v1/ormmiddle             --> gitee.com/zhenyangze/gin-framework/app/http/middleware.CachePage.func1 (6 handlers)
[GIN-debug] GET    /v1/ormcache              --> gitee.com/zhenyangze/gin-framework/app/http/controllers.OrmWithCacheHandle (6 handlers)
[GIN-debug] GET    /v1/orm                   --> gitee.com/zhenyangze/gin-framework/app/http/controllers.OrmHandle (6 handlers)
[GIN-debug] GET    /v1/redis                 --> gitee.com/zhenyangze/gin-framework/app/http/controllers.RedisHandle (6 handlers)
[GIN-debug] GET    /my                       --> gitee.com/zhenyangze/gin-framework/app/http/controllers.MyHandle (4 handlers)
[GIN-debug] GET    /ws                       --> gitee.com/zhenyangze/gin-framework/app/http/controllers.WebSocketHandle (4 handlers)
```



### jwt

#### Login api

```shell
http -v --json POST localhost:8000/login username=admin password=admin
```

```shell
POST /login HTTP/1.1
Accept: application/json, */*;q=0.5
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 42
Content-Type: application/json
Host: localhost:8080
User-Agent: HTTPie/2.2.0

{
    "password": "admin",
    "username": "admin"
}

HTTP/1.1 200 OK
Content-Length: 212
Content-Type: application/json; charset=utf-8
Date: Wed, 21 Oct 2020 06:36:46 GMT
X-Request-Id: 8b70b5ac-324a-4750-bd87-1a3bf63bd851

{
    "code": 200,
    "expire": "2020-10-21T15:36:46+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDMyNjU4MDYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTYwMzI2MjIwNn0.meMbQK1RL_O1AmRT6Rm3fIfCnlbnkKexO-picpjLH4c"
}
```

#### Authorization

```shell
http -f GET localhost:8000/auth/hello "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDMyNjU4MDYsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTYwMzI2MjIwNn0.meMbQK1RL_O1AmRT6Rm3fIfCnlbnkKexO-picpjLH4c"  "Content-Type: application/json"
```

```shell
HTTP/1.1 200 OK
Content-Length: 59
Content-Type: application/json; charset=utf-8
Date: Wed, 21 Oct 2020 06:37:32 GMT
X-Request-Id: ed3e2c12-c71f-4de6-9bf7-08b09049fd85

{
    "text": "Hello World.",
    "userID": "admin",
    "userName": "admin"
}
```



### 功能组件 

- [x] 缓存

- [x] 事件

- [x] 邮件

- [x] GORM

- [x] Redis

- [x] 中间件(限流，header，pageCache，RequestId) 

- [x] Cron

- [x] JWT