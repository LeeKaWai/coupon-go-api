# coupon-go-api

一个使用 Go 语言编写的淘宝客 API，支持通过淘口令查询优惠券和生成推广链接。

## 目录

- [介绍](#介绍)
- [特性](#特性)
- [目录结构](#目录结构)
- [安装](#安装)
- [使用](#使用)


## 介绍

coupon-go-api 是一个基于 Go 语言的项目，采用了gin 框架。通过该 API，可以通过淘口令获取商品详情、生成推广链接，以及进行佣金查询等操作。

## 特性

- 获取淘宝客商品详情
- 生成淘宝客推广链接
- 查询佣金信息

## 目录结构

```plaintext
coupon-go-api
├── app/
│   ├── index.go                # 项目初始化配置
├── config/
│   ├── dev.yaml                # 测试环境配置文件
│   ├── index.go                # 配置文件
│   ├── prod.yaml               # 生产环境配置文件
├── db/
│   ├── index.go                # 数据库初始化
├── pkg/
│   ├── taobao                  # 淘宝联盟提供的SDK
├── routes/                     # 路由服务
│   ├── v1/                    
│   │   ├── taobao.go      
│   ├── index.go     
├── src/                        # API 入口模块
├── main.go                    # 应用程序入口
├── go.mod                     # Go 模块文件
├── go.sum                     # Go 依赖文件
└── README.md                  # 项目说明文件
```




## 使用

| URL                                     | 方法 | 参数                                                                                                                                      | 备注                     |
|-----------------------------------------|------|-----------------------------------------------------------------------------------------------------------------------------------------|--------------------------|
| `/v1/tb/search`                    | POST | `{"q":"淘口令"}`                                                | 根据淘口令获取商品优惠券和推广信息                 |
