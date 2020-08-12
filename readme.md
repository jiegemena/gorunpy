## 由于windows 上部署 python 应用困难，开发了一个能够批量启动 python程序的应用

- 自动重启 python 异常退出应用
- 定时启动 python 应用
- 循环定时启动 python 应用


### config.json

config 配置说明

- path 运行目录
- cmd 运行命令
- name 服务名（不允许重名）
- state 0：不启动，1：正常启动
- runtype 任务类型 
    - 1 普通任务
    - 2 定时任务(未实现)
    - 3 循环任务(未实现)

### 例子：
```
{
    "cmds":[
        {
            "path":"D:\\codehome\\pyhome\\ptest",
            "cmd":"ping 127.0.0.1",
            "name":"a1",
            "state" : 0,
            "runtype" : 1
        },
        {
            "path":"D:\\codehome\\pyhome\\ptest",
            "cmd":"python -u index.py",
            "name":"a2",
            "state" : 1,
            "runtype" : 1
        }
    ]
}
```


src\github.com\jiegemena
git clone https://github.com/jiegemena/gorunpy.git
go build -o dkproxy.exe