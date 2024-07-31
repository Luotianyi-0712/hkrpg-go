# 准备环境
1. golang >= 1.22.1
2. mysql
3. redis
4. bash(使用build.sh时才需要)

## 编译
1. 安装依赖
`go mod tidy`
2. 开始编译
#### 脚本编译
在bash环境下直接执行bulid.sh即可，脚本执行结束后可在文件夹 ./build 里看到各平台的可执行文件
#### 自行编译
> 此处不交流交叉编译

进入./ cmd文件夹中，可看到各个服务的文件夹，进入每一个文件夹执行 `go build xxxx.go` 即可
### 不想编译
前往[Build-dev](https://github.com/gucooing/hkrpg-go-public/actions/workflows/HKRPG-Go-Publish.yml)下载

## 运行
### 1.准备资源：
data resources，data使用仓库的data即可，但资源文件夹需要给予读写权限。

resources的准备:
1. 下载[StarRailData](https://github.com/Dimbreath/StarRailData)
2. 下载补充文件（任务文件）[StarRailData-missioninfo](https://github.com/oureveryday/StarRailData-missioninfo)
3. 先将StarRailData解压到resources中，然后用StarRailData-missioninfo覆盖一次(只覆盖更新Config即可，不要覆盖ExcelOutput不兼容)
### 2.运行，
运行时需要携带启动参数 -i appid ， 其中appid格式为ipv4格式，如：9001.1.1.1 其中含义：
```bash
9001: 区服id;
1:    服务id; 
1:    主机id;
1:    本次启动服务id;
```
了解到了appid的组成含义后你可以先不携带参数启动一次，使其生成各个服务的配置文件，生成的配置文件在conf文件夹里，然后根据你自己定义的appid更改默认配置文件中的appid(虽然服务采用发现形式添加新服务，但是还是推荐每一个配置文件中的appid配置表都相同)，然后根据自己的想法更改配置文件中的其他参数。
### 3.数据库的准备，
安装mysql，mysql中新建数据库：hkrpg-go-account && hkrpg-go-user && hkrpg-go-player && hkrpg-go-conf (utf8mb4),然后更改配置文件中的账户和密码，安装redis，更改配置文件中的密码（本服务可采用分表分库形式，但同一张表一定要是同一个数据库）
### 4.启动，
前期的准备工作已经全部完成了到了启动的时候了，推荐的启动顺序为：
> 下面示例的启动方法为默认配置文件的启动参数
```bash
./nodeserver -i 9001.3.1.1
./gameserver -i 9001.2.1.1
./gateserver -i 9001.1.1.1
./dispatch -i 9001.4.1.1
./multiserver -i 9001.5.1.1
./muipserver -i 9001.6.1.1
```

## 各服务功能
### nodeserver 节点服务器（有状态，不可集群），服务发现，服务管理
### dispatch 登录服务器（无状态，可集群）
### gateserver 网关服务器（有状态，可集群），内部网络与外界唯一交互口
### gameserver 逻辑服务器（有状态，可集群），处理业务逻辑
### multiserver 多人服务器（有状态，不可集群）没有用的服务
### muipserver 目前仅负责api

## 进阶操作
### 多gateserver、多gameserver部署
以gateserver为例，默认只有一个9001.1.1.1的配置，可添加一个9001.1.1.2的配置，启动时可以使用同一个可执行文件，第一个gateserver使用-i 9001.1.1.1启动，第二个使用 -i 9001.1.1.2 启动即可，如果在同一台机器注意两个配置的端口不要冲突了

等.........

## 注意事项
请处理好内外网，不要让外网可随意访问到集群内部网络
如果你的外网带宽不足 1Gpbs/s 延迟不低于10ms 请不要使用外网数据库

## 想测试但不想配置复杂的环境

1.前往[Build-dev](https://github.com/gucooing/hkrpg-go-public/actions/workflows/HKRPG-Go-Publish.yml)下载hkrpg-pe-go文件