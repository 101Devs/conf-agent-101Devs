# 配置说明

- 配置 使用 `toml` 数据格式
- 配置分为3部分：
    - Logger：日志相关，必填，将按照配置初始化文件日志对象
    - Basic：基础配置，为Reloader配置的缺省配置，当Reloader没有配置时，会使用Basic配置作为Reloader配置。建议配置Basic配置，Reloader配置只在需要的时候进行个性化配置
    - Reloaders: reload列表。


## 1 Logger配置
| Key | 数据类型 | 含义  | 必填 | 默认值 | 说明 | 
| - | - | - | - | - | - |
| LogDir | string | 日志文件目录 | Y | - | |
| LogName | string | 日志文件名 | Y | - | |
| LogLevel | string | 日志等价 | Y | - |  可选： DEBUG TRACE INFO WARNING ERROR CRITICAL|
| RotateWhen | string | 日志文件切割策略 | Y | - | 可选：M：每分钟 H：每小时 D：每天 MIDNIGHT：午夜切割 |
| BackupCount | int | 日志文件保留格式 | Y | - | |
| Format | string | 日志消息格式 | Y | - | |
| StdOut | bool | 日志内容是否控制台输出 | N | - | |


## 2 Basic配置
| Key | 数据类型 | 含义  | 必填 | 默认值 | 说明 | 
| - | - | - | - | - | - |
| BFECluster              | string | 当前所在的BFE集群名 | Y |  |  |
| BFEConfDir              | string | bfe配置目录位置 | N | /home/work/bfe/conf |  |
| BFEMonitorPort          | int | B