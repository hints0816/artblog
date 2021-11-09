组件：
    brog：
    etcd：v2 memory，v3 database

    1.apiservice：所有服务访问统一入口
    2.crontrollerManager：维持副本期望数目
    3.scheduler：负责介绍任务，选择合适的节点进行分配任务
    4.etcd：键值对数据库，存储k8s集群的所有重要信息（如：flannel的ip）
    5.kubelet：直接跟容器引擎交互实现容器的生命周期管理
    6.kube-proxy：负责写入规则至 IPtables ipvs 实现服务映射访问
    7.coredns：为集群中svc创建域名ip的对应关系（其实就是帮你修改nginx）
    8.dashboard：k8s ui
    9.ingress controller：七层代理
pod：
    Lamp LAMP(Linux + Apache + MySQL/MariaDB + PHP)

    pod内容器共用pause的网络盏，共用存储卷
    localhost
    同pod port冲突

    控制器pod：
        1.RC：期望的副本数
        2.RS：集合式
        3.Deployment：滚动更新

        Deployment->RS->pods

        RS<-HPA[cpu>80 max=10 min=2]

        statefulset：
        1.稳定的持久化，pvc
        2.稳定的网络标记 headless service
        3.有序部署
        4.有序收缩

        daemonset
        1.node的守护进程
        日志

        job
        cronjob


服务发现
rs，rc，同一组标签
service 

squid

网络通讯：
1.同一pod(直接回返网卡)
2.pod1->pod2:同一主机由docker0网桥转发

网段
etcd 和 flannel
1存储管理flannel可分配的ip地址段资源
2监控etcd中每个pod的实际地址，并在内存中建立维护pod节点路由表

pod-->service lvs


1.安装依赖包
2.设置防火墙
3.关闭selinux
4.关闭swap
为什么k8s要关闭swap分区，关闭虚拟内存，虚拟内存会影响k8s
5.调整内核参数
6.设置系统时区
timedatectl set-timezone xxx
7.关闭不需要服务
8.关闭一个日志
9.升级内核4.44

资源清单--剧本
    资源类型：

    资源清单yaml：

pod的生命周期:
kuebctl->请求api->kubelet->pod生命周期
                    pause->initC(线性)(多个)
                                -就绪检测(readiness)->生存检测(liveness)

initC(初始化容器)- 会导致pod重启
1.成功
2.线性成功执行

binary
source
pap
sha512