# 实现原理

配置加载包括如下几个过程：
- 配置文件拉取：
    - 以 bfe 热加载 API 触发后会读取的 配置文件列表 为集合，从 API Server 拉取 一到多个 配置文件
    - 如果没有更新的配置，退出本次配置加载
- 配置文件落盘：
    - 将现有的正式的指定配置文件列表