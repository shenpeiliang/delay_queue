# 功能说明
基于Go实现的任务队列
- 使用Redis接收外部数据，定时扫描Redis队列加入处理
- 可通过接口地址传递任务，或直接使用Redis键值缓存任务
- 后台提供日志管理页面，查看任务执行状态


### 文件结构说明
#### main
已编译好的可执行文件，用于Web服务

使用Redis存储Session

Redis、Mysql均使用连接池配置

采用Go简易的开发框架Gin

Gin 是使用 Go/golang 语言实现的 HTTP Web 框架。接口简洁，性能极高。截止 1.4.0 版本，包含测试代码，仅14K，其中测试代码 9K 左右，也就是说框架源码仅 5K 左右

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


# 实现原理
![image](https://github.com/shenpeiliang/delay_queue/blob/master/data/1.jpg)

结构属性
```go
//环形队列结构
type Queue struct {
	currentIndex int                //当前扫到的索引号
	slots        []map[string]*Task //队列中每个槽位的元素

	//信道
	closed        chan bool //关闭
	popTaskClose  chan bool //任务出队关闭
	pushTaskClose chan bool //任务入队关闭
	timeClose     chan bool //时间关闭
}
```

currentIndex当前扫到的索引号，如当前配置的槽位数为60，相当于走完一圈就是1分钟，当索引号为59时，下一秒就会变成
0以达到闭环的效果

slots []map[string]*Task 队列中每个槽位的元素存储着任务集合，也就是说每一秒要执行的任务应该是多个的，会循环检查哪个
任务达到可执行的条件


```go
//队列结构元素
type Task struct {
	cycleNum int //执行任务的循环次数
	retryNum int //重试次数

	queueIndex string
	queueSlot  int

	notifyUrl    string    //请求通知地址
	planTime     time.Time //时间戳，指定要执行任务的计划时间
	methodName   string    //请求方法 post、get
	notifyParams string    //请求通知数据，请求查询字符串
}
```
Task队列结构元素里cycleNum标识执行任务的循环次数，每走完一圈的时间就会减去1，直到为0的时候就会被执行任务，
类似一个离心圆，距离圆心越近那么就更早地被执行

queueIndex作为任务的索引值，因为使用异步处理，不需要等待数据库返回来的ID，所以使用uuid生成

retryNum重试次数标识执行任务失败的次数，执行失败的规则在配置文件中已经有说明，是一个延时队列，有最大
错误重试次数，每次失败都会延迟多少秒后再执行，类似支付通知，原理就是未接收到成功返回就重新计算要执行任务的时间，然后
放回队列中等待下一次执行


# Supervisor独立部署
```shell script
[root@host5 MyApp]# yum install supervisor
[root@host5 MyApp]cd /etc/supervisord.d
[root@host5 MyApp]vim gin.ini
```

配置内容
```shell script
[program:web]
# 运行用户身份
user=root
# 执行的命令
command=/data/wwwwroot/gin/main
# 日志输出
stdout_logfile=/data/wwwwlog/gin/stdout.log
stderr_logfile=/data/wwwwlog/gin/stderr.log
# supervisor启动的时候是否随着同时启动，默认True
autostart=true
# 当程序exit的时候是否重启
autorestart=true

[program:queue]
# 运行用户身份
user=root
# 执行的命令
command=/data/wwwwroot/gin/queue
# 日志输出
stdout_logfile=/data/wwwwlog/queue/stdout.log
stderr_logfile=/data/wwwwlog/queue/stderr.log
# supervisor启动的时候是否随着同时启动，默认True
autostart=true
# 当程序exit的时候是否重启
autorestart=true
```

命令说明：
```shell script
supervisorctl status            //查看所有进程的状态
supervisorctl stop app1         //停止app1
supervisorctl start app1        //启动app1
supervisorctl restart app1      //重启app1
supervisorctl update            //配置文件修改后使用该命令加载新的配置
supervisorctl reload            //重新启动配置中的所有程序
```


# nginx代理部署
需要的文件：
```go
config 配置文件
static 静态文件
view 模板文件
main 可执行文件
queue 可执行文件
```

配置文件：
```shell script
server {
	# 监听的端口
    listen       8080;
    # 域名 
    server_name  _ ;
    # 访问日志
    access_log   /data/wwwlogs/gin-access.log;
    # 错误日志
    error_log    /data/wwwlogs/gin-error.log;
    
	# 静态文件交给nginx处理,这里是采用文件后缀来区分的
    location ~ .*\.(gif|jpg|jpeg|png|js|css|eot|ttf|woff|svg|otf)$ {
        access_log off;
        expires    1d;
        root       /data/wwwroot/gin/static;
        try_files  $uri @go_dispose;
    }
 #  也可根据文件夹目录区分,指定目录的访问交给Nginx处理(将public目录交给nginx处理)
 #  location ^~ /public {
 #      access_log off;
 #      expires    1d;
 #      root       /data/wwwroot/gin;
 #      try_files  $uri @go_dispose;
 #  }
 
	# 将其他程序交给后端Go处理
    location / {
        try_files $uri @go_dispose;
    }
	
    location @go_dispose {
    	# Go程序处理的地址
        proxy_pass                 http://127.0.0.1:8080;
        proxy_redirect             off;
        proxy_set_header           Host             $host;
        proxy_set_header           X-Real-IP        $remote_addr;
        proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}
```
重启nginx

nginx -t

service nginx restart 


# 开发说明
待完善..
