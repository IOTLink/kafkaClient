package main

//http://blog.csdn.net/pkueecser/article/details/50613989


//kafka zookeeper 原理
//http://www.cnblogs.com/yangxiaoyi/p/7359236.html


/*
多个broker协同合作，producer、consumer和broker三者之间通过zookeeper来协调请求和转发。

kafka 特殊配置：http://blog.csdn.net/hadas_wang/article/details/50056381


zk：
Kafka依赖Zookeeper管理自身集群（Broker、Offset、Producer、Consumer等），所以先要安装 Zookeeper。自然，为了达到高可用的目的，Zookeeper自身也不能是单点，接下来就介绍如何搭建一个最小的Zookeeper集群（3个 zk节点）
    此处选用Zookeeper的版本是3.4.6，此为Kafka0.9中推荐的Zookeeper版本

	http://blog.csdn.net/maomao5987370/article/details/51384694
http://www.cnblogs.com/5iTech/articles/6043224.html

*/