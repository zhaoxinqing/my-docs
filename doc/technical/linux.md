# linux

## Linux

### IO

### 内核空间、用户空间（内核态、用户态）

- 内核空间就是操作系统内核代码运行使用的内存空间
- 用户空间就是用户程序代码运行使用的内存空间

- 因为操作系统太高级，我们用户组所有的应用，是运行在操作系统之上的，一旦操作系统不能稳定运行，那就完了。而保证操作系统运行稳定的基础，就是操作系统所使用的内存不能被随意更改，所以内存被虚拟的分成了用户内存和内核内存。

Linux
Linux总结：<https://wenku.baidu.com/view/f4aecad1b90d4a7302768e9951e79b8968026836.html>
Linux基本组件：内核、shell和GUI、系统实用程序和应用程序；

### Linux开机启动过程

1. 主机加电自检，加载BIOS硬件信息；
2. 读取MBR的引导文件（GRUB、LILO）；（LILO：是Linux的引导加载程序，主要用于将Linux操作系统加载到主内存中，以便它可以开始运行）
3. 引导Linux内核；
4. 运行第一个进程init（进程号永远是1）；
5. 进入相应的运行级别；
6. 运行终端，输入用户名和密码；

### Linux使用的进程间通信方式

- 管道（pipe）、流管道（s_pipe）、有名管道（FIFO）；
- 信号（signal）；
- 消息队列；
- 共享内存；
- 信号量；
- 套接字（socket）；

### Linux有那些系统日志文件

比较重要的是 /var/log/messages 日志文件，该日志文件是许多进程日志的汇总，从该文件可以看出任何企图或成功的入侵；

### 交换空间

交换空间是Linux使用的一定空间，用于临时保存一些并发运行的程序。当RAM没有足够的内存来容纳正在执行的所有程序时，就会发生这种情况；

### CLI和GUI

CLI（command-line interface）：命令行界面，不支持鼠标，通过键盘输入指令；
GUI（Graphical User Interface）：图形用户界面；

### 硬链接和软连接

### 常使用命令

<https://zhuanlan.zhihu.com/p/455887180>



其它：
孤儿进程，僵尸进程
父进程回收，内部引用进程没有关闭，导致孤儿进程和僵尸进程；

死锁条件，如何避免
互相引用会导致死锁现象发生，避免互相引用情况发生；

## 学习网址资源

### 视频

2021韩顺平一周学会Linux：<https://www.bilibili.com/video/BV1Sv411r7vd（基于CentOS7.6>版本较新，视频长度刚刚好，也比较完整）
玩转Vim从放弃到爱不释手：<https://www.imooc.com/learn/1129>（好评很多）
阿里云Linux运维学习路线：<https://edu.aliyun.com/roadmap/linux>

### 书籍

《鸟哥的Linux私房菜——基础篇》：<http://cn.linux.vbird.org/linux_basic/linux_basic.php>（经典）

### 文档

Linux教程（菜鸟教程）：<https://www.runoob.com/linux/linux-tutorial。html>
Linux教程（W3CSchool）：<https://www.w3cschool.cn/linux/>
Linux工具快速教程：<https://linuxtools-rst.readthedocs.io>（基础、工具进阶、工具参考）

### 合集

Linux内核学习资料：<https://github.com/0voice/linux_kernel_wiki>
GitHubLinux专区：<https://github.com/topics/linux>（很多好项目）
GitHubLinux合集：<https://github.com/inputsh/awesome-linux（Linux>系列技术）
StackOverflow：<https://stackoverflow.com/questions/tagged/linux>（解决问题必备）
掘金Linux专区：<https://juejin.cn/tag/Linux>（技术文章）

### 社区（国内倒的差不多了）

开源中国：<https://www.oschina.net/>（综合的开源社区）
红帽官网：<https://www.redhat.com/zh>

### 题

牛客网Linux专项练习：<https://www.nowcoder.com/intelligentTest>
牛客网Linux面试题：<https://www.nowcoder.com/search?query=linux%E9%9D%A2%E8%AF%95%E9%A2%98&type=question>
Linux常见面试题整理：<https://zhuanlan.zhihu.com/p/376749877>
Linux常见面试题整理：<https://github.com/0voice/linux_kernel_wiki#-%E9%9D%A2%E8%AF%95%E9%A2%98>

## Nginx（一个集**静态资源**、**负载均衡**于一身的**Web**服务器）

Nginx是当下最流行的Web[服务器](https://cloud.tencent.com/product/cvm?from=10680)，通过官方以及第三方C模块，以及在Nginx上构建出的Openresty，或者在Openresty上构建出的Kong，你可以使用Nginx生态满足任何复杂Web场景下的需求。Nginx的性能也极其优秀，它可以轻松支持百万、千万级的并发连接，也可以高效的处理磁盘IO，因而通过静态资源或者缓存，能够为Tomcat、Django等性能不佳的Web应用扛住绝大部分外部流量。

## 都是事件驱动，为什么Nginx的性能远高于Redis？

谈到Redis缓存，我们描述其性能时会这么说：支持1万并发连接，几万QPS。而我们描述Nginx的高性能时，则会宣示：支持C10M（1千万并发连接），百万级QPS。Nginx用C语言开发，而Redis是用同一家族的C++语言开发的，C与C++在性能上是同一级数的。Redis与Nginx同样使用了事件驱动、异步调用、Epoll这些机制，为什么Nginx的并发连接会高出那么多呢？

这其实是由**进程架构**决定的。为了让进程占用CPU的全部计算力，Nginx充分利用了分时操作系统的特点，比如增加CPU时间片、提高CPU二级缓存命中率、用异步IO和线程池的方式回避磁盘的阻塞读操作等等，只有清楚了Nginx的这些招数，你才能将Nginx的性能最大化发挥出来。

材料、散热这些基础科技没有获得重大突破前，CPU频率很难增长，类似Redis、NodeJS这样的单进程、单线程高并发服务，只能向分布式集群方向发展，才能继续提升性能。Nginx通过Master/Worker多进程架构，可以充分使用服务器上百个CPU核心，实现C10M。

为了榨干多核CPU的价值，Nginx无所不用其极：通过绑定CPU提升二级缓存的命中率，通过静态优先级扩大时间片，通过多种手段均衡Worker进程之间的负载，在独立线程池中隔离阻塞的IO操作，等等。可见，高性能既来自于架构，更来自于细节。

## Etcd-服务发现

[etcd](https://etcd.io/)（读作 et-see-dee）是一种[开源的](https://www.redhat.com/zh/topics/open-source/what-is-open-source)分布式统一键值存储，用于分布式系统或计算机集群的共享配置、服务发现和的调度协调。etcd 有助于促进更加安全的自动更新，协调向主机调度的工作，并帮助设置[容器](https://www.redhat.com/zh/topics/containers/whats-a-linux-container)的覆盖网络。

etcd 是许多其他项目的核心组件。最值得注意的是，它是 [Kubernetes](https://www.redhat.com/zh/topics/containers/what-is-kubernetes) 的首要数据存储，也是[容器编排](https://www.redhat.com/zh/topics/containers/what-is-container-orchestration)的实际标准系统。使用 etcd， [云原生应用](https://www.redhat.com/zh/topics/cloud-native-apps)可以保持更为一致的运行时间，而且在个别服务器发生故障时也能正常工作。应用从 etcd 读取数据并写入到其中；通过分散配置数据，为节点配置提供冗余和弹性。

服务发现（Service discovery）是自动检测一个[计算机网络](https://zh.wikipedia.org/wiki/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C "计算机网络")内的设备及其提供的服务。服务发现议定(service discovery protocol)帮助发现服务的[网络传输协议](https://zh.wikipedia.org/wiki/%E7%BD%91%E7%BB%9C%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE "网络传输协议")。服务发现的目的在于为用户自行配置而减负。服务发现通常需要一个共同的语言, 帮助用户通过软件代理器（software agents）调用另一个服务, 而不需要用户自行调用。

## Nacos-动态服务发现、配置管理和服务管理的平台

1、Nacos 是一个更易于构建云原生应用的动态服务发现、配置管理和服务管理的平台；

## 微服务

微服务是架构设计方式，分布式是系统部署方式，两者概念不同
是什么：
微服务是一种软件开发技术，是面向服务的体系结构（SOA）架构样式的一种变体，它提倡将单一应用程序划分成一组小的服务，服务之间互相协调、互相配合，为用户提供最终价值；
每个服务运行在其独立的进程中，服务与服务间采用轻量级的通信机制互相沟通（通常是基于HTTP的RESTful API）；

## SOA架构

面向服务的架构（Service Oriented Architecture）；
它是一种设计方法，其中包含多个服务， 服务之间通过相互依赖最终提供一系列的功能；

## ESB（企业服务总线）、微服务API网关

简单来说 ESB 就是一根管道，用来连接各个服务节点。为了集成不同系统，不同协议的服务，ESB 做了消息的转化解释和路由工作，让不同的服务互联互通；
API网关是一个服务器，是系统的唯一入口。API网关封装了系统内部架构，为每个客户端提供一个定制的API。它可能还具有其它职责，如身份验证、监控、负载均衡、缓存、请求分片与管理、静态响应处理。API网关方式的核心要点是，所有的客户端和消费端都通过统一的网关接入微服务，在网关层处理所有的非业务功能。通常，网关也是提供REST/HTTP的访问API。服务端通过API-GW注册和管理服务。
链接：<https://www.cnblogs.com/xuwc/p/13989081.html>

## 数据库

### 索引

索引是数据库表中一列或者多列的值进行排序后的一种结构，其作用就是提高表中数据的查询速度。类似于字典中的音序表；
索引类型：
1.主键索引（primary key）：保证数据的唯一性，不能为空，唯一标识数据库表中的每条记录（alter table table_name add primary key（column）;）
2.唯一索引（unique）：唯一索引，防止数据出现重复（alter table table_name add unique（column）;）
3.普通索引（index）：仅仅是为了提高查询的速度（alter table table_name add index_name（column）;）
4.联合索引（）：
5.全文索引（）：

索引简单来讲就是查询数据的优化方案，通过提取数据的特征信息，通过特征信息快速定位到原始数据。
类似的设计方案：
1.数组：通过下标计算元素的具体地址，然后通过地址获取数据，通过下标计算真实的地址其实也算是一种索引技术；
2.MySQL的B+树索引，通过索引减少查询磁盘的次数，提高查询效率，这是索引的经典场景；
3.ES的倒排索引；
4.Java中的Map其实也是索引技术的经典实现；

### 联合索引

链接：<https://blog.csdn.net/w13346019869/article/details/123413536>

### MySQL全文索引

全文索引与普通的索引不是一回事，在查找上方面其效率是普通模糊（like）查询和N倍，是MySQL专门提供用作搜索引擎的；
只有字段的数据类型为 char、varchar、text 及其系列才可以建全文索引；
全文索引有专用的查询语法：match(字段) against(关键字)。如：select * from answer where match(content) against("测试表")；
连接：<https://zhuanlan.zhihu.com/p/417229576>

### 哈希索引和B + 树索引

针对查询效率而言，哈希索引最快，如果不存在哈希碰撞，效率为O（1）；
但是B+树索引可以范围查询，而哈希索引不行；
使用场景：
1.大多数场景下，都会有组合查询，范围查询、排序、分组、模糊查询等查询特征，Hash 索引无法满足要求，建议数据库使用B+树索引。
2.在离散型高，数据基数大，且等值查询时候，Hash索引有优势。

### SQL查询

示例：<https://www.w3school.com.cn/sql/sql_groupby.asp>
左查询和右查询：<https://baijiahao.baidu.com/s?id=1728618861730351587&wfr=spider&for=pc>

## 分布式（技术）

一种基于网络的计算机处理技术，与集中式相对应。
分布式网络中，数据的存储和处理都是在本地工作站上进行的。因为每台计算机都能够存储和处理数据，所以不要求服务器功能十分强大，其价格也就不必过于昂贵。同时允许他们共享网络的数据、资源和服务。在分布式网络中使用的计算机既能够作为独立的系统使用，也可以把它们连接在一起得到更强的网络功能。
分布式技术优点：
快速访问、多用户使用，每台计算机都可以访问系统内其他计算机的信息文件；
系统设计上具有更大的灵活性，既可为独立的计算机的地区用户的特殊需求服务，也可为联网的企业需求服务，实现系统内不同计算机之间的通信；
分布式技术缺点：
对病毒比较敏感，任何用户都可能引入被病毒感染的文件，并将病毒扩散到整个网络；
备份困难，如果用户将数据存储在各自的系统上，而不是将他们存储在中央系统中，难于制定一项有效的备份计划。这种情况还可能导致用户使用同一文件的不同版本。
设备必须要互相兼容。
管理和维护比较复杂。
百度百科：<https://baike.baidu.com/item/分布式技术/4110919>

分布式计算：hadoop、HBase
MapReduce原理简介：<https://zhuanlan.zhihu.com/p/62135686>
Hbase简介：<https://blog.csdn.net/u012485099/article/details/110941341>
Hadoop百科：<https://baike.baidu.com/item/Hadoop/3526507?fr=aladdin>
Hive百科：<https://baike.baidu.com/item/hive/67986?fr=aladdin>

### Hadoop

Hadoop是一个分布式系统基础架构。可以使得用户在不用了解分布式底层细节的情况下，开发分布式程序。充分利用集群的威力进行高速运算和存储；

### HDFS

HDFS是指被设计成适合运行在通用硬件上的分布式文件系统；
具有高容错性特点（多备份机制），并且提供高吞吐量来访问应用程序的数据（流式访问），适合那些超大数据集的应用程序；
HDFS采用了主从（Master/Slave）结构模型，一个HDFS集群是由一个NameNode和若干个DataNode组成的。其中NameNode作为主服务器，管理文件系统的命名空间和客户端对文件的访问操作；集群中的DataNode管理存储的数据；

### Hbase

Hbase是一种分布式、可扩展、支持海量数据存储的NoSQL数据库；

### Hive

hive是基于Hadoop的一个数据仓库工具，用来进行数据提取、转化、加载，这是一种可以存储、查询和分析存储在Hadoop中的大规模数据的机制；
Hive 的本质其实就相当于将HDFS 中已经存储的文件在Mysql中做了一个双射关系，以方便使用 HQL 去管理查询；

### MapReduce

MapReduce是一种分布式计算框架 ，以一种可靠的，具有容错能力的方式并行地处理TB级别的海量数据集。主要用于搜索领域，解决海量数据的计算问题。
YARN：
是一种新的 Hadoop 资源管理器，它是一个通用资源管理系统，可为上层应用提供统一的资源管理和调度，它的引入为集群在利用率、资源统一管理和数据共享等方面带来了巨大好处。

### 分布式系统

建立在网络之上的软件系统
链接：<https://baike。baidu。com/item/分布式系统/4905336>
分布式相关知识点：
1.CAP原则：一致性、可用性、分区容忍性（只能满足两个）；
2.分布式一致性协议raft，两大核心：选主和复制日志；
3.一致性哈希算法，保证负载均衡，可添加虚拟节点；
4.大型网站系统架构：用户-负载均衡器-N台服务器-redis缓存集群-mysql集群
(1)前端限流（例如一个用户10秒内只能点击一次，异步处理，消息队列）；
(2)负载均衡一般采用NGINX反向代理；
(3)mysql读写分离，主库写，从库读，分库分表。
5.准点秒杀设计：承受高并发的压力，采用前端限流，削峰（异步处理，消息队列）的方法；
6.分布式锁ZooKeeper：在分布式系统中，常常需要协调他们的动作。如果不同的系统或是同一个系统的不同主机之间共享了一个或一组资源，那么访问这些资源的时候，往往需要互斥来防止彼此干扰来保证一致性，在这种情况下，便需要使用到分布式锁。
7.分布式缓存：
8.微服务：优点：松耦合，开发容易，职责单一，可以技术异构；缺点:可能会重复开发，各个微服务之间维护复杂，通讯耗时长；
9.消息队列：
10.数据库缓存一致性：保证redis和mysql的一致性；
(1)先删缓存，再更新数据库；
(2)一般采用延时双删策略，即先删缓存，更新数据库，休眠1秒，再删缓存。第二次删除缓存为异步。如果第二次删缓存失败，还会有缓存和数据库不一致的问题。可以引入保障机制，删除失败重新删除，直到成功。
11.数据库主从一致性；
12.缓存雪崩，缓存穿透；
链接：<https://zhuanlan.zhihu.com/p/140272240>

## 高并发、高性能、高可用

### 集群（技术）

集群（cluster）技术是一种较新的技术，通过集群技术，可以在付出较低成本的情况下获得在性能、可靠性、灵活性方面的相对较高的收益，其任务调度则是集群系统中的核心技术。
集群是一组相互独立的、通过高速网络互联的计算机，它们构成了一个组，并以单一系统的模式加以管理。一个客户与集群相互作用时，集群像是一个独立的服务器。集群配置是用于提高可用性和可缩放性。
百度百科：<https://baike.baidu.com/item/集群技术/9774443>

集群分类：
科学集群：是并行计算的基础，以集群开发的并行应用程序，可以解决复杂的科学问题。科学集群对外好像一个超级计算机，这种计算机内部由十至上万个独立处理器组成，并且在公共消息传递层上进行通信以运行并行应用程序。
负载均衡集群：可以使用特定的算法，使得应用程序处理负载或网络流量负载尽可能平均分摊处理，以实现负载均衡。每个节点都是运行单独软件的独立系统，都可以承担一定的处理负载，并且可以实现处理负载在节点之间的动态分配；
高可用性能集群：考虑到计算机硬件和软件的易错性，高可用性能集群的主要目的是为了使集权的服务尽可能可用。当主节点发生了故障，那么这段时间内将有次节点代替它，次节点主要是主节点的镜像。当它代替主节点时，它可以完全接管其身份，因此使系统环境对用户是一致的。

### 如何设计一个秒杀系统（完整版）

链接：<https://blog.csdn.net/weixin_45953989/article/details/120016954>

### C10K问题及解决方案

C10k问题是优化网络套接字以同时处理大量客户端的问题；
链接：<https://cloud.tencent.com/developer/article/1031629>
