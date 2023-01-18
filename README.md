# OJ网站

##  项目简介

实现了类似于OJ的一些功能

##  代码架构

### 项目结构

<details>
<summary>展开查看</summary>
<pre>
<code>
    ├── app ----------------------------- (项目文件)
    	├── api ------------------------- (api层)
    		├── category ------------------ (关于题目分类的api)
    		├── problem -------------- (关于问题的api)
    		├── submit -------------------- (关于用户提交代码的api)
    		├── user -------------------- (关于用户的api)
    	├── global ---------------------- (全局组件)
    	├── internal -------------------- (内部包)
    		├── middleware -------------- (中间件)
    		├── model ------------------- (模型)
    		├── service ----------------- (服务层)
    	├── router ---------------------- (路由层)
    ├── boot ---------------------------- (项目启动包)
    ├── manifest ------------------------ (交付清单)
    	├── config ---------------------- (项目配置)
		├── sql ------------------------- (sql文件)
			├── mysql ------------------- (mysql表结构)
    ├── utils --------------------------- (工具包)
    ├── build.sh ------------------------ (项目启动shell脚本)
    ├── docker-compose.yml -------------- (docker容器)
</code>
</pre>
</details>


### 技术栈

<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="15%">

- [gin](https://github.com/gin-gonic/gin)

> `Gin`是一个用Go语言编写的web框架。它是一个类似于`martini`但拥有更好性能的API框架, 由于使用了`httprouter`，速度提高了近40倍。 如果你是性能和高效的追求者, 你会爱上`Gin`。



<img src="http://jwt.io/img/logo-asset.svg">

- jwt

> SON Web Token (JWT)是一个开放标准(RFC 7519)，它定义了一种紧凑的、自包含的方式，用于作为JSON对象在各方之间安全地传输信息。该信息可以被验证和信任，因为它是数字签名的。

- [zap](https://github.com/uber-go/zap)

> `zap`是`Uber`开发的非常快的、结构化的，分日志级别的Go日志库。根据Uber-go Zap的文档，它的性能比类似的结构化日志包更好，也比标准库更快。具体的性能测试可以去`github`上看到。

- [viper](https://github.com/spf13/viper)

> Viper是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

<img src="https://upload.wikimedia.org/wikipedia/zh/thumb/6/62/MySQL.svg/1200px-MySQL.svg.png" width="30%">

- [mysql](https://www.mysql.com/)

> 一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统关系型数据库管理系统之一，在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一

- [redis](https://redis.io/)

> 一个开源的、使用C语言编写的、支持网络交互的、可基于内存也可持久化的Key-Value数据库

<img src="https://developers.redhat.com/sites/default/files/styles/article_feature/public/blog/2014/05/homepage-docker-logo.png?itok=zx0e-vcP" width="30%">

- [docker](https://www.docker.com/)

> Google 公司推出的 Go 语言 进行开发实现，基于 Linux 内核的 cgroup，namespace，以及 AUFS 类的 Union FS 等技术的一个容器服务

​	容器用docker-compose部署

##  功能模块

### API文档

> v1.0.0

Base URLs:

* <a href="127.0.0.1:8080">测试环境: 127.0.0.1:8080</a>

# 用户

## POST 用户注册

POST /api/user/register

> Body 请求参数

```yaml
username: "1119209427"
password: 197920jc

```

### 请求参数

| 名称       | 位置 | 类型   | 必选 | 说明 |
| ---------- | ---- | ------ | ---- | ---- |
| body       | body | object | 否   | none |
| » username | body | string | 是   | none |
| » password | body | string | 是   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## POST 用户登录

POST /api/user/login

> Body 请求参数

```yaml
username: "1119209427"
password: 197920jc

```

### 请求参数

| 名称       | 位置 | 类型   | 必选 | 说明 |
| ---------- | ---- | ------ | ---- | ---- |
| body       | body | object | 否   | none |
| » username | body | string | 是   | none |
| » password | body | string | 是   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## PUT 升级管理员权限

PUT /api/user/admin

> Body 请求参数

```yaml
ids:
  - "3"

```

### 请求参数

| 名称  | 位置  | 类型   | 必选 | 说明 |
| ----- | ----- | ------ | ---- | ---- |
| token | query | string | 是   | none |
| body  | body  | object | 否   | none |
| » ids | body  | array  | 是   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## PUT 取消管理员权限

PUT /api/user/cancel%20admin

> Body 请求参数

```yaml
ids: "3"

```

### 请求参数

| 名称  | 位置  | 类型   | 必选 | 说明 |
| ----- | ----- | ------ | ---- | ---- |
| toekn | query | string | 是   | none |
| body  | body  | object | 否   | none |
| » ids | body  | string | 否   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

# 问题

## GET 获取问题列表

GET /api/problem/list

### 请求参数

| 名称    | 位置  | 类型    | 必选 | 说明 |
| ------- | ----- | ------- | ---- | ---- |
| size    | query | integer | 否   | none |
| page    | query | integer | 否   | none |
| keyword | query | string  | 否   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## POST 创建问题

POST /api/problem/create

> Body 请求参数

```yaml
title: 两数相加
content: 给你两个数a和b，请求出他们的和
textCase: 1 2
max_runtime: "100"
max_mem: "100"

```

### 请求参数

| 名称          | 位置  | 类型    | 必选 | 说明 |
| ------------- | ----- | ------- | ---- | ---- |
| token         | query | string  | 是   | none |
| body          | body  | object  | 否   | none |
| » title       | body  | string  | 是   | none |
| » content     | body  | string  | 是   | none |
| » textCase    | body  | string  | 是   | none |
| » max_runtime | body  | integer | 是   | none |
| » max_mem     | body  | integer | 是   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## PUT 修改问题

PUT /api/problem/update

> Body 请求参数

```yaml
id: "2"
title: 三数之和
content: 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
  ？请你找出所有和为 0 且不重复的三元组。注意：答案中不可以包含重复的三元组。
textCase: 给定数组 nums = [-1, 0, 1, 2, -1, -4]，满足要求的三元组集合为：[  [-1, 0, 1],  [-1, -1, 2]]
max_runtime: "100"
max_mem: "30"

```

### 请求参数

| 名称          | 位置  | 类型   | 必选 | 说明 |
| ------------- | ----- | ------ | ---- | ---- |
| token         | query | string | 是   | none |
| body          | body  | object | 否   | none |
| » id          | body  | string | 否   | none |
| » title       | body  | string | 否   | none |
| » content     | body  | string | 否   | none |
| » textCase    | body  | string | 否   | none |
| » max_runtime | body  | string | 否   | none |
| » max_mem     | body  | string | 否   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## POST 添加测试用例

POST /api/testcase/create

> Body 请求参数

```yaml
problem_id: "9"
input: 2 3
output: "5"

```

### 请求参数

| 名称         | 位置  | 类型    | 必选 | 说明 |
| ------------ | ----- | ------- | ---- | ---- |
| token        | query | string  | 是   | none |
| body         | body  | object  | 否   | none |
| » problem_id | body  | integer | 是   | none |
| » input      | body  | string  | 是   | none |
| » output     | body  | string  | 是   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## PUT 修改测试用例

PUT /api/testcase/update

> Body 请求参数

```yaml
id: "2"
problem_id: "9"
input: 100 210
output: "310"

```

### 请求参数

| 名称         | 位置  | 类型    | 必选 | 说明 |
| ------------ | ----- | ------- | ---- | ---- |
| token        | query | string  | 是   | none |
| body         | body  | object  | 否   | none |
| » id         | body  | integer | 是   | none |
| » problem_id | body  | integer | 是   | none |
| » input      | body  | string  | 是   | none |
| » output     | body  | string  | 是   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

# 提交

## GET 获取提交列表

GET /api/submit/lists

### 请求参数

| 名称  | 位置  | 类型    | 必选 | 说明 |
| ----- | ----- | ------- | ---- | ---- |
| page  | query | integer | 否   | none |
| size  | query | integer | 否   | none |
| token | query | string  | 是   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## POST 提交代码

POST /api/submit/submit

> Body 请求参数

```
string

```

### 请求参数

| 名称       | 位置  | 类型    | 必选 | 说明 |
| ---------- | ----- | ------- | ---- | ---- |
| token      | query | string  | 否   | none |
| problem_id | query | integer | 否   | none |
| body       | body  | string  | 否   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

# 分类

## POST 创建分类

POST /api/category/created

> Body 请求参数

```yaml
name: 链表
parent_id: string

```

### 请求参数

| 名称        | 位置  | 类型   | 必选 | 说明 |
| ----------- | ----- | ------ | ---- | ---- |
| token       | query | string | 否   | none |
| body        | body  | object | 否   | none |
| » name      | body  | string | 否   | none |
| » parent_id | body  | string | 否   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

## PUT 修改分类

PUT /api/category/update

> Body 请求参数

```yaml
id: "1"
name: 算法
parent_id: "0"

```

### 请求参数

| 名称        | 位置  | 类型   | 必选 | 说明 |
| ----------- | ----- | ------ | ---- | ---- |
| token       | query | string | 是   | none |
| body        | body  | object | 否   | none |
| » id        | body  | string | 是   | none |
| » name      | body  | string | 是   | none |
| » parent_id | body  | string | 否   | none |

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功 | Inline   |

### 返回数据结构

# 数据模型

