# unknown

## Service
- gate: 网关服务(tcp/udp/websocket),保持与客户端的长链接，根据并发承载，动态扩容
- web: http接口服务，支持登录/支付/资源版本等功能，可分布式部署
- center: 中心服务，负责产品全局相关的逻辑与数据存储
- log: 日志服务，负责游戏行为日志汇总
- game: 游戏服务，负责处理游戏内部逻辑，支持滚服模式开服
- cross: 跨服节点
- admin: 管理面板，负责管理所有的节点和游戏业务

## Middleware
- etcd: 注册中心
- kafka: 消息中间件
- mongo: 数据库
- redis: 缓存中间件
- elasticsearch: 日志存储中间件