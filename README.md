## flinkx
基于kitex RPC微服务 + Hertz HTTP服务完成的第四届字节跳动青训营 - 极简流处理引擎项目

## 项目特点

1. 项目基于Go语言1.18编写

2. 采用字节跳动开源RPC框架（Kitex），提供**RPC微服务**（算子之间的通信）
3. 采用字节跳动开源HTTP框架（Hertz），提供**HTTP服务**（DAG任务提交）
4. 使用**Kafka**存放需要处理的流数据，并进行消费

5. 各个算子实现基于Kitex的脚手架生成的代码进行开发，项目**结构清晰**，代码**符合规范**

6. 使用**Nacos**进行服务注册和服务发现（Kitex框架扩展）

7. 使用Kitex的**熔断器**，当下游服务出现故障时，主动断流

## 项目地址

https://github.com/BaiZe1998/flinkx

## 项目说明

### 1. 项目模块介绍

| 服务名称 |             模块介绍              | 技术框架 | 传输协议 | 注册中心 |  日志  |     数据存取      |
| :------: | :-------------------------------: | :------: | :------: | :------: | :----: | :---------------: |
|   api    | 接受HTTP请求提交流处理任务（DAG） | `hertz`  |  `http`  | `nacos`  | `klog` |      `hertz`      |
|   data   |         数据生产导入kafka         | `kitex`  | `thrift` |          |        | `kafka`、`sarama` |
|  source  |         从kafka中消费数据         |          |          |          |        | `kitex`、`sarama` |
|   map    |            map算子服务            |          |          |          |        |                   |
|  keyby   |           keyby算子服务           |          |          |          |        |                   |
|  reduce  |          reduce算子服务           |          |          |          |        |                   |
|   sink   |           sink算子服务            |          |          |          |        |      `file`       |

### 2. 服务调用关系

![image-20220824164723742](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220824164723742.png)

### 3. 代码介绍

#### 3.1 代码目录结构介绍

|   目录    |  子目录   |               说明                |
| :-------: | :-------: | :-------------------------------: |
|    cmd    |    api    |         api服务的业务代码         |
|           |   data    |       kafka的生产者业务代码       |
|           |    map    |         map服务的业务代码         |
|           |   keyby   |        keyby服务的业务代码        |
|           |  reduce   |       reduce服务的业务代码        |
|           |   sink    |        sink服务的业务代码         |
|           |  source   | kafka消费者业务代码（source服务） |
|  config   |           |    flinkx的配置文件以及DAG文件    |
|    idl    |           |        thrift接口定义文件         |
| kitex_gen |           |        Kitex自动生成的代码        |
|    pkg    | constants |            系统常量包             |
|           |   errno   |              错误码               |

#### 3.2 代码运行

1. 提前修改config目录的相关配置
2. 运行api服务（接收DAG）

```bash
cd word-count
go run word-count/cmd/api
```

3. 运行消息生产服务

```bash
cd word-count
go run word-count/cmd/data/producer
```

4. 运行source服务（kafka消息消费者）

```bash
cd word-count
go run word-count/cmd/source
```

5. 运行map服务

```bash
cd word-count
go run word-count/cmd/map
```

6. 运行keyby服务

```bash
cd word-count
go run word-count/cmd/keyby
```

7. 运行reduce服务

```bash
cd word-count
go run word-count/cmd/reduce
```

8. 运行sink服务

```bash
cd word-count
go run word-count/cmd/sink
```

## 存在的问题

1. 伪分布式：当前每类算子占用一个进程，而相同算子的并发通过该进程内的协程并发实现，并没有实现JobManager和TaskManager的调度模型
2. DataStream结构没有清晰定义
3. 引擎的功能与DAG任务描述之间存在耦合性
4. 算子的负载均衡规则较为简单

## 下一步计划

1. Kitex开启opentelemetry链路追踪扩展，选择Jaeger渲染链路
2. 故障恢复（服务拉起）

3. 一致性语义（at least once || exactly once）
4. 项目重构：实现JobManager和TaskManager模块，算子的创建符合Flink的模式
5. 算子shuffle倾斜优化

