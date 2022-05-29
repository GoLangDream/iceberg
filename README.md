# iceberg

[![Build](https://github.com/GoLangDream/iceberg/actions/workflows/build.yml/badge.svg)](https://github.com/GoLangDream/iceberg/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/GoLangDream/iceberg/badge.svg?branch=main)](https://coveralls.io/github/GoLangDream/iceberg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoLangDream/iceberg)](https://goreportcard.com/report/github.com/GoLangDream/iceberg)

一个方便快速的web开发框架

## 开始

参考 https://github.com/GoLangDream/iceberg-cli 使用命令行工具开始

### 新建一个项目

```shell
iceberg new test_project
```

### 其他一些命令

```shell
iceberg g m create_user                    # 生成一个migration
iceberg g model user name:string age:uint  # 生成一个模型
iceberg g controller user index show       # 生成一个controller 和 对应的 action
```

### 运行

```shell
iceberg run
```

### 访问

http://localhost:3000

## 示例

参考 https://github.com/GoLangDream/iceberg_example
