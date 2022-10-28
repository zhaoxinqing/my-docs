# 消息队列

## Kafka

消息队列应用场景：解耦、异步、削峰。

Kafka是消息队列中间件的代表产品，它与RocketMQ和RabbitMQ最大的区别在于：在某些场景，可以弃用Flink、Spark这样的计算引擎，借助Kafka Stream轻松实现数据处理。也即，Kafka不仅是消息引擎系统，也是分布式流处理平台。
Kafka是由Apache软件基金会开发的一个开源流处理平台，由Scala和Java编写。该项目的目标是为处理实时数据提供一个统一、高吞吐、低延迟的平台。其持久化层本质上是一个“按照分布式事务日志架构的大规模发布/订阅消息队列”，这使它作为企业级基础设施来处理流式数据非常有价值。此外，Kafka可以通过Kafka Connect连接到外部系统（用于数据输入/输出），并提供了Kafka Streams——一个Java流式处理库"库 (计算机)")。

网址：<https://zhuanlan.zhihu.com/p/74063251>

解决的问题：解耦、异步、削峰
架构中引入MQ后存在的问题：
降低系统可用性。如果MQ挂掉，导致整个系统崩溃；
系统复杂性变高。可能发重复消息，导致插入重复数据；消息丢了；消息顺序乱了；系统 B，C，D 挂了，导致 MQ 消息积累，磁盘满了；
一致性问题。本来应该A，B，C，D 都执行成功了再返回，结果A，B，C 执行成功 D 失败。
建议：中小型公司 RabbitMQ，大公司：RocketMQ，大数据实时计算：Kafka。
解决方案：
事务机制：（一般不采用，同步的，生产者发送消息会同步阻塞卡住等待你是成功还是失败。会导致生产者发送消息的吞吐量降下来）；
持久化到磁盘；
知识点：<https://wenku.baidu.com/view/36c47b5a32b765ce0508763231126edb6f1a76a9.html>
链接：<https://zhuanlan.zhihu.com/p/356235333>
Go语言系列之RabbitMQ消息队列：<https://blog.51cto.com/u_13919520/3153035>
消息队列对比详解：<https://baijiahao.baidu.com/s?id=1708572554459077699&wfr=spider&for=pc>
