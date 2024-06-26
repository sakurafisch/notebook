# NewSQL简介

[参考文档](https://docs.microsoft.com/zh-cn/dotnet/architecture/cloud-native/relational-vs-nosql-data#newsql-databases)

*NewSQL* 是一种新兴数据库技术，它结合了 NoSQL 的分布式可扩展性和关系数据库的 ACID 保证。 NewSQL 数据库对于必须在跨分布式环境中处理大量数据的业务系统非常重要，具有完整的事务性支持和 ACID 遵从性。 虽然 NoSQL 数据库可以提供巨大的可伸缩性，但它不保证数据一致性。 不一致数据的间歇问题会给开发团队带来负担。 开发人员必须在其微服务代码中构建安全措施，以管理由不一致数据导致的问题。

云本机计算基础（CNCF）具有多个 NewSQL 数据库项目。

| Project    | 特征                                                         |
| :--------- | :----------------------------------------------------------- |
| 蟑螂 DB    | 全局缩放的 ACID 兼容的关系数据库。 将新节点添加到群集，CockroachDB 负责跨实例和地理区域平衡数据。 它创建、管理和分发副本以确保可靠性。 它是开源的，免费提供。 |
| TiDB       | 支持混合事务性和分析处理（HTAP）工作负荷的开源数据库。 它是与 MySQL 兼容的功能，可提供水平伸缩性、强一致性和高可用性。 TiDB 的作用类似于 MySQL 服务器。 你可以继续使用现有 MySQL 客户端库，而无需对应用程序进行大量代码更改。 |
| YugabyteDB | 开源、高性能分布式 SQL 数据库。 它支持较低的查询延迟、针对故障的恢复能力和全局数据分布。 YugabyteDB 是 PostgressSQL 兼容的，可处理横向扩展 RDBMS 和 internet 缩放 OLTP 工作负载。 该产品还支持 NoSQL，并与 Cassandra 兼容。 |
| Vitess     | Vitess 是用于部署、缩放和管理 MySQL 实例的大型群集的数据库解决方案。 它可以在公共或私有云体系结构中运行。 Vitess 结合了许多重要的 MySQL 功能和功能，同时提供了垂直和水平分片支持。 源自 YouTube，Vitess 一直在为所有 YouTube 数据库流量提供服务，2011。 |

可从云本机计算基础获取上图中的开源项目。 其中三个产品是完整的数据库产品，其中包括 .NET Core 支持。 另一种是 Vitess，它是一种数据库群集系统，可用于水平缩放 MySQL 实例的大型群集。

NewSQL 数据库的关键设计目标是在 Kubernetes 中以本机方式工作，利用平台的复原能力和可伸缩性。

NewSQL 数据库旨在应对暂时的云环境，在该环境中，基础虚拟机随时可以重新启动或重新计划。 数据库旨在使节点出现故障，而不会丢失数据和停机。 例如，CockroachDB 可以通过在群集中的节点之间维护所有数据的三个一致副本，来经受计算机丢失。

Kubernetes 使用*服务构造*来允许客户端从单个 DNS 条目寻址一组完全相同的 NewSQL 数据库进程。 通过将数据库实例与相关联的服务地址进行分隔，可以扩展，而不会中断现有的应用程序实例。 在给定时间向任何服务发送请求将始终产生相同的结果。

在这种情况下，所有数据库实例都相等。 没有主要关系或次要关系。 CockroachDB 中的*共识复制*等技术允许任何数据库节点处理任何请求。 如果接收负载平衡请求的节点在本地具有所需的数据，则该节点会立即响应。 如果不是，则该节点将成为网关，并将请求转发到相应的节点以获得正确的答案。 从客户端的角度来看，每个数据库节点都是相同的：它们显示为单一的*逻辑*数据库，具有单计算机系统的一致性保证，尽管在幕后甚至有数百个节点在幕后工作。

有关 NewSQL 数据库背后的机制的详细信息，请参阅[Kubernetes-本机数据库的四个属性一](https://thenewstack.io/dash-four-properties-of-kubernetes-native-databases/)文。