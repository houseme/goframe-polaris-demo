# GoFrame-Polaris-Demo
Sample code for GoFrame and polaris


Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.

## How to use

A、写入SQL文件到数据库，文件在`manifest/sql/`目录下

B、执行`gf gen pbentity`生成`Entity/DAO/DO`文件

C、执行`gf gen pb`生成`PB`文件

D、修改`manifest/config/config.yaml`文件
```yaml
# GRPC Server.
grpc:
  address: "192.168.124.12:8199"
  name:    "GoFramePolarisDemo"
  logPath: ""
  logStdout: true
  errorStack: true
  errorLogEnabled: true
  errorLogPattern: "error-{Ymd}.log"
  accessLogEnabled: true
  accessLogPattern: "access-{Ymd}.log"

# Global logging.
logger:
  level: "all"
  stdout: true

# Database.
database:
  logger:
    level: "all"
    stdout: true

  default:
    link: "mysql:root:root@tcp(127.0.0.1:3306)/test"
    debug: true
```

    1、修改IP地址`address`为本机IP

    2、修改`link`为本机数据库地址

    3、修改`client.go`中的`Passport`的值，保证唯一性

E、先运行`main.go`中的`main`函数，启动`GRPC Server`服务

F、再运行`client.go`中的`main`函数，启动`GRPC Client`服务



# License

`GoFrame-Polaris-Demo` software is licenced under the Apache License Version 2.0. See the [LICENSE](LICENSE) file for details. 