# 功能说明
基于Go实现的任务队列
- 使用Redis接收外部数据，定时扫描Redis队列加入处理
- 可通过接口地址传递任务，或直接使用Redis键值缓存任务
- 后台提供日志管理页面，查看任务执行状态


### 文件结构说明
#### main
已编译好的可执行文件，用于Web服务

#### queue
已编译好的可执行文件，用于队列任务服务

#### 服务配置说明
使用yaml格式配置，路径 config/config.yaml
```go
server:
  host: 0.0.0.0
  port: 8080
  read_timeout: 10 #超时秒数
  write_timeout: 10
  max_header_bytes: 0 #为0使用默认 1<<20
  views_pattern: view/*/*.html
  left_delims:  #模板渲染分隔符 - 左
  right_delims: #模板渲染分隔符 - 右
  env: release #环境模式 release/debug/test
static_version: #静态文件版本号
  js: 0.0.0
  css: 0.0.0
master_db:
  host: 127.0.0.1
  port: 3306
  db_name: gin
  db_user: root
  db_pwd: root
  prefix: hs_
  db_charset: utf8
  db_max_open_conns: 20 #连接池最大连接数
  db_max_idle_conns: 10 #连接池最大空闲数
  db_max_lifetime_conns: 7200 #连接池链接最长生命周期s
slave_db:
  host: 127.0.0.1
  port: 3306
  db_name: gin
  db_user: root
  db_pwd: root
  prefix: hs_
  db_charset: utf8
  db_max_open_conns: 20
  db_max_idle_conns: 10
  db_max_lifetime_conns: 7200
redis:
  host: 127.0.0.1
  port: 6379
  #db_name: 0 #数据库名整数
  #db_pwd:  #验证密码
  db_max_open_conns: 20 #连接池最大连接数
  db_max_idle_conns: 0 #连接池最大空闲数
  db_max_lifetime_conns: 0 #连接池链接最长生命周期s
session:
  key_pairs: secret
  name: GOSESSIONID
queue:
  key: delay_queue_task #队列键名
  slot: 60 #队列槽位数
  time_interval: 10 #如果加入队列时的计划时间小于当前时间，设置任务计划时间为当前时间之后的配置秒数
  max_retry_num: 6 #最大重试次数，超过该次数之后将从定时任务中删除，不再处理
  retry_default_time: 60 #重试次数默认间隔时间，如果retry_time_config配置没有就使用默认间隔
  retry_time_config: [1_5,2_15,3_30,4_180,5_3600,6_7200] #重试延迟队列配置 重试第几次_当前之后的多少秒
```
### 如何使用
#### 需要的文件
```go
config 配置文件
static 静态文件
view 模板文件
main 可执行文件
queue 可执行文件
```

#### 环境要求
不需要安装Go环境

需要初始化数据库data/db.sql

Redis服务需要开启

#### 启动服务
可执行文件需要执行权限
```shell script
chmod +x ./main
chmod +x ./queue
```

然后直接启动服务./main和./queue即可

#### 后台管理
地址：http://127.0.0.1:8080/admin/login/index

账号： 管理员 
 
密码：zxc123


#### 添加任务
```shell script
curl -d 'notify_url=http://sksystem.sk0.com/queue/vip&plan_time=1593317127' http://127.0.0.1:8080/api/task/save
```
参数说明：
```shell script
notify_url 任务通知地址
plan_time 任务预计执行时间（时间戳）
method_name 请求方法get/post
notify_param 通知参数
```

#### 效果图
Web服务：
![image](https://github.com/shenpeiliang/delay_queue/blob/master/data/img_2.png)

Queue队列服务：
![image](https://github.com/shenpeiliang/delay_queue/blob/master/data/img_1.png)

![image](https://github.com/shenpeiliang/delay_queue/blob/master/data/1.gif)
