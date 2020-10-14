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



### 功能组件 

[ x ] 缓存

[x] 事件

[x] 邮件

[x] GORM

[x] Redis

[x] 中间件(限流，header，pageCache，RequestId) 

[x] Cron

[ ] JWT
