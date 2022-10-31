
# 请求数据打印服务

用于打印 `Get` 和 `Post` 请求的数据, 支持以下数据的输出

- [x] ip
- [x] path
- [x] params
- [x] header
- [x] body

# 配置文件

程序运行时会读取 `./configs/default.yml` 作为配置文件, 当该文件不存在会使用默认配置. 具体配置在[这里](./configs/default.yml)

# 依赖说明

+ `github.com/zly-app/zapp` 是基础程序框架, 用于快速搭建一个app
+ `github.com/zly-app/service/api` 是一个web服务, 基于`zapp`快速创建一个web程序

# docker部署

```yaml
version: '3'
services:
  http-print:
    image: zlyuan/http-print:latest
    container_name: http-print
    restart: unless-stopped
    #command: /src/app
    ports:
      - "8080:8080"
```
