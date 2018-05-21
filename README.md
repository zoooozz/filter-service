
# filter-service [敏感词服务]

## 接入规范说明
    端口6511 -> internal 为内部接口为管理后台提供 
    端口6510 -> outter 为内部接口为业务提供接口 
    URL http:://www.xxxxxx.co  替换真实slb地址
    
## 错误码

状态码  | 返回值
---|---
0    | 请求成功
-400 | 参数错误

## 功能点

 1. <a href="#新增敏感词">新增敏感词   (internal)   POST</a>
 2. <a href="#获取敏感词 ">获取敏感词  (internal) GET</a>
 3. <a href="#获取业务方 ">获取业务方  (internal) GET</a>
 4. <a href="#更新敏感词状态 ">更新敏感词状态  (internal) POST</a>
 4. <a href="#更新敏感词级别 ">更新敏感词级别  (internal) POST</a>
 5. <a href="#敏感词匹配 ">敏感词匹配  (outter) POST</a>
 
 


### 新增敏感词
    
    URL:  /x/internal/filter/keyword/add [POST]
    参数:
    
参数 | 类型 | 必须 | 说明
---|---|---|---
content | string | true | 关键词内容
flag | string | false | 用户组ID 多个逗号分隔
level | int | false | 0 普通 1 一级 2 二级
state | int | false | 默认关闭 1开启

#### 返回值

```json
{
    "code": 0,
    "message": "ok"
}
```

### 获取敏感词 

    URL:  URL:  /x/internal/filter/keyword/list [GET]
    参数:

参数 | 类型 | 必须 | 说明
---|---|---|---
flag | string | true | 用户组 
content | string | flase | 敏感词搜索昵称


#### 返回值

```json
{
    "code": 0,
    "data": {
        "bs": [
            {
                "id": 15,
                "content": "我是敏感词2233",
                "flag": "username",
                "state": 0,
                "level": 0
            }
        ],
        "count": 2
    },
    "message": "ok"
}
```

## 获取业务方

    URL:  /x/internal/filter/business/list [GET]
    参数

参数 | 类型 | 必须 | 说明
---|---|---|---

#### 返回值

```json
{
    "code": 0,
    "data": {
        "bs": [
            {
                "id": 7,
                "name": "其他",
                "flag": "all",
                "state": "0"
            }
        ]
    },
    "message": "ok"
}
```


## 更新敏感词状态

    URL:  /x/internal/filter/keyword/state [POST]
    参数

参数 | 类型 | 必须 | 说明
---|---|---|---
state|int|true| 状态
id|int|true|等级

#### 返回值

```json
{
    "code": 0,
    "message": "ok"
}
```


## 更新敏感词级别

    URL:  /x/internal/filter/keyword/edit [POST]
    参数

参数 | 类型 | 必须 | 说明
---|---|---|---
level|int|true| 状态
id|int|true|等级

#### 返回值

```json
{
    "code": 0,
    "message": "ok"
}
```


## 敏感词匹配

    URL:  /x/outter/filter/list [POST]
    参数

参数 | 类型 | 必须 | 说明
---|---|---|---
content|string|true| 内容
flag|string|true| 业务方

#### 返回值

```json
{
    "code": 0,
    "data": {
        "content": "*****323123",
        "keyword": [
            "我是敏感词"
        ]
    },
    "message": "ok"
}
```




