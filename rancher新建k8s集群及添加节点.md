# rancher新建k8s集群及添加节点

## 1、添加k8s集群

1.1 全局-添加集群-创建新的kubernetes集群-自定义

1.2 填写集群名称

1.3 kubernetes选项：

选择版本

1.4 私有镜像仓库

启用

仓库地址：nexus.gree.com

用户名：100193685

密码：Qwer!234

这里区分一下/etc/docker/daemon.json文件中的

registry-mirror:"http:xx.xx.xx.xx"

执行docker pull时宿主机先从docker hub拉取docker image，内网不可达，则从registry-mirror配置的私有镜像拉取(允许匿名访问的镜像库)，在rancherUI中新建集群时指定nexus.gree.com并配置账号密码则宿主机执行docker pull时会从nexus.gree.com登录并拉取镜像

1.5 下一步-完成

## 2、添加节点：

**添加前的检查**

1、禁用selinux vi /etc/selinux/config	修改：SELINUX=disabled

2、禁用swap	vi /etc/fstab					 修改：#   ... swap ...

3、关闭和禁用防火墙 systemctl stop firewalld && systemctl disable firewalld

2.1 全局-选择刚才新建的集群-集群-添加节点

2.1.1 作为master节点：需要勾选：ETcd，Control Plane，Worker

2.1.2 作为worker节点：只需勾选Worker(默认)

复制生成的命令

2.1.3 在主机的 SSH 终端运行前

**先检查是否有docker镜像在运行**

docker ps

如果有

docker stop +names 停止

检查/etc/kubernetes等目录是否存有旧的etcd状态

如果有，则执行/root/clean.sh脚本清理状态数据，路由表等

如果节点机器上没有clean.sh，可以通过**scp root@10.2.30.21:/root/clean.sh /root/clean.sh**从其他节点上复制到本机的/root目录下

执行命令，将节点注册到集群中

<!--注册为全部-->

sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run  nexus.gree.com/rancher/rancher-agent:v2.5.9 --server https://xka.gree.com --token f8nbrg4px4cvpm5dznncqmd4hpvsrhxqfvw4p6knnhm2dbftnrxrnw --etcd --controlplane --worker

<!--注册为worker-->

sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run  nexus.gree.com/rancher/rancher-agent:v2.5.11 --server https://xka.gree.com --token rfwr2smhlp5tn98wjwq2j4h5xbzvjmlk4snwnfvscv867kddrkc9b6 --worker



sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run  nexus.gree.com/rancher/rancher-agent:v2.5.9 --server https://xka.gree.com --token f8nbrg4px4cvpm5dznncqmd4hpvsrhxqfvw4p6knnhm2dbftnrxrnw --worker



docker pull nexus.gree.com/rancher/hyperkube:v1.20.8-rancher1可以查看拉取和解压的进度