

## cola 
#### 适配层（Adapter Layer：负责对前端展示（web，wireless，wap）的路由和适配，对于传统B/S系统而言，adapter就相当于MVC中的controller；

#### 应用层（Application Layer）：主要负责获取输入，组装上下文，参数校验，调用领域层做业务处理，如果需要的话，发送消息通知等。层次是开放的，应用层也可以绕过领域层，直接访问基础实施层；

#### 领域层（Domain Layer）：主要是封装了核心业务逻辑，并通过领域服务（Domain Service）和领域对象（Domain Entity）的方法对App层提供业务实体和业务逻辑计算。领域是应用的核心，不依赖任何其他层次；

#### 基础实施层（Infrastructure Layer）：主要负责技术细节问题的处理，比如数据库的CRUD、搜索引擎、文件系统、分布式服务的RPC等。此外，领域防腐的重任也落在这里，外部依赖需要通过gateway的转义处理，才能被上面的App层和Domain层使用。


### ==和equals的区别，一般来说，equals里面自定义判断==

###　实例变量必须创建对象后才可以通过这个对象来使用，静态变量则可以直接使用类名来引用

### 子类覆盖父类的方法时，只能比父类抛出更少的异常

### 只有记住抽象类与普通类的唯一区别就是不能创建实例对象和允许有abstract方法。


jdk17
1. 密封类 指定继承
2. switch可以做到instanceof 的事情

jdk11
1. String新增工具：repeat 返回重复
strip、stripLeading、stripTrailing trim增强版
isblank 
lines string转stream
2. 增强文件读写 readString和writeString

四次握手： 结束连接
 WebSocket（持久化的协议） 与 Socket（一套标准IP 地址 + 端口 + 协议） 的区别
 
 人群打标

 入参后缀，如示例代码 DTO（Data Transfer Object）

@Transactional 
 P.S. 如果 try-catch 语句在 finally 语句块中进行了 return 操作，那么 catch 语句块中手动抛出的异常也会被覆盖，同样不会自动回滚。
 
 jdk8
 1.hashmap--->拉链法--->1.8尾插法--->红黑树 简单点说，旋转的目的是让树保持红黑树的特性

 O(n) > O(logn)。

ConcurrentHashMap

java的两个锁
synchronized-jvm层面
ReentrantLock-代码层面

HTTPS(SSL/TLS)
TCP三次握手
SSL/TLS四次握手

关于es8的搭建过程

关于spring初始化读取nacos配置中心的自定义配置文件

文件上传
1.秒传：MD5 redis
2.分片上传

RandomAccessFile

    位置指示器

    操作模式

MappedByteBuffer

javax.servlet

UnionFS
多个目录(也叫分支)内容联合挂载到同一个目录下，而目录的物理位置是分开的

Init层

它是一个以“-init”结尾的层，夹在只读层和读写层之间。Init层是Docker项目单独生成的一个内部层，专门用来存放/etc/hosts、/etc/resolv.conf等信息。需要这样一层的原因是，这些文件本来属于只读的系统镜像层的一部分，但是用户往往需要在启动容器时写入一些指定的值比如hostname，所以就需要在可读写层对它们进行修改。可是，这些修改往往只对当前的容器有效，我们并不希望执行docker commit时，把这些信息连同可读写层一起提交掉。所以，Docker做法是，在修改了这些文件之后，以一个单独的层挂载了出来。而用户执行docker commit只会提交可读写层，所以是不包含这些内容的。

AUFS的whiteout的实现是通过在上层的可写的目录下建立对应的whiteout隐藏文件来实现的。


Namespace分别对进程的 Cgroup root、进程间通信、网络、文件系统挂载点、进程ID、用户和组、主机名域名等进行隔离。

 在每个容器内也要有一个父进程来管理其下属的子进程
Docker就会在这个PID=3702的进程启动时给他施一个“障眼法”，让他永远看不到不属于它这个namespace中的进程。这种机制，其实就是对被隔离应用的进程空间做了手脚，使得这些进程只能看到重新计算过的进程编号，比如PID=1。可实际上，他们在宿主机的操作系统里，还是原来的第3702号进程

 
 命名空间建立系统的不同视图， 对于每一个命名空间，从用户看起来，应该像一台单独的Linux计算机一样，有自己的init进程(PID为0)，其他进程的PID依次递增，A和B空间都有PID为0的init进程，子容器的进程映射到父容器的进程上，父容器可以知道每一个子容器的运行状态，而子容器与子容器之间是隔离的。

 Cgroups
 Linux Cgroups的全称是Linux Control Groups。它最主要的作用，就是限制一个进程组能够使用的资源上限，包括CPU、内存、磁盘、网络带宽等等

 docker通过cgroup来配置容器里的各种配置参数，会生成一个控制组（一个容器一个）,最后会把容器在宿主机外面对应的pid塞到tasks文件中，标识该进程持有我这个控制组
一个容器如果不对其做任何资源限制,则宿主机会允许其占用无限大的内存空间,有时候会因为代码bug程序会一直申请内存,直到把宿主机内存占完。

　　为了避免此类的问题出现，宿主机有必要对容器进行资源分配限制，比如CPU，内存等，Linux Cgroups的全称是Linux Control Groups，它最主要的作用就是限制一个进程组能够使用的资源上限,包括CPU,内存,磁盘,网络带宽等等。

　　此外，Linux Cgroups还能够对进程优先级设置，以及将进程挂起和恢复等操作。


复合索引有最左匹配限制
or语句对于mysql来说还是没有顺序的

( 注：一个索引最多可以包含 16 列。)


覆盖索引
即普通索引查询方式，则需要先搜索name索引树，得到id的值为3，再到id索引树搜索一次。这个过程称为回表


行数据库擅长增删改，列数据库擅长查询

redis的大key问题
1.redis大key是什么 单个keyvalues很大
2.redis大key带来的影响 指令单线程操作IO阻塞
3.redis大key的解决方法 拆分KV


redis单线程？主要是指 Redis 在网络 IO和键值对读写是采用一个线程来完成的
redis内部操作？
redis为什么那么快？

Redis 在设计上采用将网络数据读写和协议解析通过多线程的方式来处理，对于命令执行来说，仍然使用单线程操作。


如何知道有消息丢失？
哪些环节可能丢消息？
如何确保消息不丢失？


JVMTI可以支持第三方工具程序以代理的方式连接和访问JVM，并利用JVMTI提供的丰富的编程接口，完成很多跟JVM相关的功能；
在JVM启动时，通过JVM参数-javaagent，传入agent jar，Instrument Agent被加载
也可以在JVM启动后，attach agent包

 Kafka 中，一个 Topic 可以配置多个 Partition（分区），数据会被写入到多个分区中，但在消费的时候，Kafka 约定一个分区只能被一个消费者消费，Topic 的分区数量决定了消费的能力




 上下文切换
 直接消耗：指的是CPU寄存器需要保存和加载
 间接消耗：指的是多核的cache之间得共享数据

 自定义的Decoder需要实现feign.codec.Decoder接口


stringbuilder和stringbuffer的区别

每次对 String 类型改变=new String 对象，指针指向新的 String 对象，所以经常进行字符串连接操作的字符串最好不要用 String 

public StringBuilder(String str) {
    super(str.length() + 16);
    append(str);
}

@Override
public synchronized StringBuffer append(String str) {
    toStringCache = null;
    super.append(str);
    return this;
}


bytebuffer是什么


什么是零拷贝
1.


DDD
与以往的仅有getter、setter的业务对象不同，领域对象具有了行为，对象更加丰满。同时，比起将这些逻辑写在服务内（例如**Service），领域功能的内聚性更强，职责更加明确。


## 关于oracle？
### 与mysql的对比
#### 主键：
mysql有自增主键
oracle需要序列来实现自增
分页
mysql limit
oracle 没有limit， 需要rownum实现分页

模糊查询
oracle模糊查询无法使用索引》
并发
mysql innodb行级锁，依赖于表索引
oracle 行级锁不依赖索引
隔离级别
mysql 4
oracle 2
数据持久性
mysql 重启丢失
oracle 重启不丢失

归档日志undo.log使用


序列化

序列化最终的目的是为了对象可以跨平台存储和进行网络传输，而我们进行跨平台存储和网络传输的方式就是 IO，而 IO 支持的数据格式就是字节数组。
常见的序列化方式有：JDK 原生、JSON、ProtoBuf、Hessian、Kryo等。
JSON 序列化常见的框架有：fastJSON、Jackson、Gson 等。


concurrentHashMap？
CAS:操作机制，Unsafe 
unsafe.compareAndSwapInt(this, valueOffset, expect, update);
它需要三个参数，分别是内存位置 V，旧的预期值 A 和新的值 B。操作时，先从内存位置读取到值，然后和预期值A比较。如果相等，则将此内存位置的值改为新值 B，返回 true。如果不相等，说明和其他线程冲突了，则不做任何改变，返回 false。

缓冲区不会立即将数据刷新到主内存中。

ABA问题？加版本号或者加更新时间


seata的原理？
cap
2pc：两阶段提交

    事务管理器发一个 prepare 指令给 A 和 B 两个服务器
    A 和 B 两个服务器收到消息后，根据自身情况，判断自己是否可以提交事务
    将处理结果记录到资源管理器中
    将处理结果返回给事务管理器
缺点：
单点故障：事务管理器
同步阻塞：单个服务因为网络故障会影响其他服务的结果返回时间
数据不一致？：协调者发出commit后，与参与者同时呆鸡

3pc：在2pc一二阶段中加了预提交阶段，并且加入了超时机制
tcc
本地消息表
- 非mq
- mq
saga
- 向后恢复
- 向前恢复

在第一阶段的 undo_log 中存放了数据修改前和修改后的值，为事务回滚做好准备，所以第一阶段就已经将分支事务提交，也就释放了锁资源！

apache的HttpClient ×

jmm
    高速缓存（Cache）
    缓存一致性协议
    当CPU写数据时，如果发现操作的变量是共享变量，即在其他CPU中也存在该变量的副本，会发出信号通知其他CPU将该变量的缓存行置为无效状态，因此当其他CPU需要读取这个变量时，发现自己缓存中缓存该变量的缓存行是无效的，那么它就会从内存重新读取。

    为了实现volatile的内存语义，JMM会限制特定类型的编译器和处理器重排序
    java编译器会在生成指令系列时在适当的位置会插入内存屏障指令来禁止特定类型的处理器重排序。

    双重检查
    可能发生指令重排序
    对象是个半成品，那就有空指针异常了。

    volatile轻量级同步。没有提供原子性和互斥性
    compilers就不会对这个属性做指令重排序。

    volatile属性不会被线程缓存，始终从主 存中读取。
    指令重排序导致的多个线程操作之间的不可见性。

    happens-before

    volatile保证了内存可见性，所以每次获取时都是最新值，获取时不用加锁，



afterPropertiesSet是在容器把依赖都设置好之后调用的，这时对象的字段都有了容器设置进来的值，调用这个方法可以执行进一步的初始化，参数校验之类的逻辑

当队列中存放的任务达到队列容量的时候，当前可以同时运行的线程数量变为最大线程数。


api规范
对非资源URL使用动词
只返回一个操作，可以使用动词，不是我们的CRUD操作

JSON属性使用camelCase

NGINX:
    初始化好这个监控的Socket
    再 fork出多个子进程出来
    子进程会竞争 accept
    异步非阻塞工作方式
    在需要等待的时候，这些进程就空闲出来待命了

    跟nio差不多
    跟 netty差不多Nginx 基于事件模型。

    异步，非阻塞，使用 epoll ，和大量细节处的优化。也正是 Nginx 之所以然的技术基石。

    正向代理和反向代理的区别：方向不一样

    最核心的区别在于apache是同步多进程模型，一个连接对应一个进程，nginx是异步的，多个连接可以对应一个进程。
    算法：
    轮询
    权重
    ip_hash
    url_hash

    ngx_http_limit_req_module

private 属性可以通过反射来获取，那么private意义何在？
    Java的private修饰符不是为了绝对安全设计的，而是对用户常规使用Java的一种约束。就好比饭店厨房门口挂着“闲人免进”的牌子，但是你还是能够通过其他方法进去。

Spring 推荐我们使用构造器的方式注入 Bean
@RequiredArgsConstructor代替@Autowired（免写构造方法）

不要返回 null
例如：Collections.emptySet();Collections.emptyMap();Collections.emptyList();
map遍历：
        for (Map.Entry<String, String> entry : map.entrySet()) {  
Optional判空
Optional.ofNullable(node)  
                .map(CatalogueTreeNode::getChild)  
                .orElse(Collections.emptyList());  


有序数组用二分法还是不错的

mysql：
    普通索引找主键索引找数据


    执行：
        客户端-连机器-查询缓存-分析器-优化器-执行器-存储引擎
    当对表中的数据进行增删改查是，索引也要动态的维护，这样就降低了数据的更新效率
    组合索引 最左匹配原则
    为什么选择innodb
    InnoDB 一页 16K 每个节点存储 二叉树改造成了多叉树
    不选择平衡二叉树：
        时间复杂度与树高相关
    不选择B树：
        三次磁盘 IO（比平衡二叉树好）
        叶子节点无指针相连，所以范围查询增加了磁盘 IO 次数，降低了查询效率。
    为什么选择B+树
    B+ 树 双向指针相连：范围查询只走了 4 次磁盘 IO
    B+ 的非叶子节点不再存储数据
    降低树高
    explain 查看是否覆盖缩影，多使用覆盖缩影，避免回表
    oracle：
    只简单说一下：mysql和oracle的区别

    日志：
    redo log 回滚日志
    先写 redo log 日志，后写磁盘
    binlog 归档日志

    redo log 是 InnoDB 引擎特有的；binlog 是 MySQL 的 Server 层实现的，所有引擎都可以使用。
    redo log 是物理日志，记录的是在某个数据页上做了什么修改；binlog 是逻辑日志，记录的是这个语句的原始逻辑，比如 **"给 ID=2 这一行的 age 字段加 1"**。
    redo log 是循环写的，空间固定会用完；binlog 是可以追加写入的。追加写是指 binlog 文件写到一定大小后会切换到下一个，并不会覆盖以前的日志。

    count：
    count (字段)<count (主键 id)<count (1)≈count ()

    order：
    给每个线程分配一块内存用于排序，称为 sort_buffer,变大后，需要临时文件辅助排序

    内存足够大，MySQL 会优先选择全字段排序，因为全字段排序不用回表

    自增：
    有主键，报主键冲突
    无主键，row_id 达到最大值后会从 0 重新开始算；前面插入的数据就会被后插入的数据覆盖，且不会报错。

    锁：

    事务：
    四大特性
    隔离级别
    三大并发问题
    Oracle 数据库的默认隔离级别是读提交，而 MySQL 是可重复读

    mvcc
    借助 undo log 通过回滚可以回到上一个版本状态

今日一坑
    @builder和@Data
    一起使用会使得无参构造失效

java8新特性：
Stream
    
ROCKETMQ刷盘和读取都使用到了林拷贝
而rocketmq使用mmap实现的林拷贝

webService能够让我们可以获取网上别人发布出来的服务。我们只要调用它，就可以获取相关的数据了

Http协议需要Socket支持
Socket其实就是对TCP/IP协议的一个封装

SOAP就是Http+XML** 支持HTTP协议->WebService需要给不同语言都能够使用，因此它使用XML

Java自带的WsImport来实现本地代理,WebService翻译成Java类

自定义服务端webservice
对服务类上加上注解。通过EndPoint(端点服务)就能够把我们webService服务类发布出去了

CXF框架spring->不用我们自己Endpoint,且有拦截(肯定又是动态代理处理类实现了拦截)


spring aop 使用的是Cglib

cglib其实子类代理


    如果加入容器的目标对象有实现接口，就使用JDK代理
    如果目标对象没有实现接口，就使用Cglib代理。


mysql  filesort
如果sort buffer不够，则在临时表 temporary table 中存储排序结果。
优化：
一次扫描算法：空间换时间，增大max_length_for_sort_data 一次取出所有字段
避免select * 
避免扩充到使用临时区


BeanFactory getBean 在获取 Bean 时把构造函数的入参信息传递进去


oracle
1.动态sql和静态sql的区别？编译和执行何时进行
2.游标：用来存储多条查询数据的一种数据结构（'结果集'），
 它有一个 '指针'，从上往下移动（'fetch'），从而能够 '遍历每条记录'

 配合loop 对数据遍历，何时停止

 游标又分为隐式和显式
rpad()函数

aware
想获得 Spring 框架提供的 BeanFactory、ApplicationContext、BeanClassLoader
谁实现了这样的一个接口，就可以获取接口入参中的各类能力

单例模式和原型模式的区别就在于是否存放到内存中

MyBatis 就是实现了一个 MapperFactoryBean 类，在 getObject 方法中提供 SqlSession 对执行 CRUD 方法的操作


非 Singleton 类型的 Bean 不需要执行销毁方法。-是因为会gc吗

FactoryBean 

事件的设计本身就是一种观察者模式的实现
观察者模式的实现过程，主要包括了事件的定义、事件的监听和发布事件，发布完成后根据匹配策略，监听器就会收到属于自己的事件内容，并做相应的处理动作

aop
代理方法主要是使用到方法拦截器类处理方法的调用 MethodInterceptor#invoke，而不是直接使用 invoke 方法中的入参 Method method 进行 method.invoke(targetObj, args) 这块是整个使用时的差异。

没有与spring整合的aop，这只是单个切面，没有细分为前还是后

beanfactory->defaultlistbeanfactory里面有beandefintion


Java和C++的区别
Oracle JDK 和 OpenJDK 的对比

1
2
4
8
4
8
2
1

强制转换类型
逻辑与跟短路与的差别
this().和super()
static

如何跳出当前的多重嵌套循环：用带有标号的break 语句

多态：编程时并不确定，而是在程序运行期间才确定

抽象类和接口的对比
抽象类是对类的抽象，是一种模板设计，接口是行为的抽象，是一种行为的规范。

在Java中定义一个不做事且没有参数的构造方法的作用：注意父子类情况
构造方法与static的顺序
构造方法默认与替换


匿名内部类：
    匿名内部类必须继承一个抽象类或者实现一个接口。
    匿名内部类不能定义任何静态成员和静态方法。
    当所在的方法的形参需要被匿名内部类使用时，必须声明为 final。
    匿名内部类不能是抽象的，它必须要实现继承的类或者实现的接口的所有抽象方法。

回调
访问外部类私有


hashset
hashcode->equal

值传递和引用传递

字符串常量池

包装类

string的不可变和常量池

String str="i"与 String str=new String(“i”)一样吗？
是不一样的


集合
set(集）、list(列表）和map(映射)。
放数据对象引用
集合的底层数据结构？
集合哪些是线程安全？

fail-fast
ConcurrentModificationException 
改变modCount的值

运行以上错误代码会报 ConcurrentModificationException 异常。这是因为当使用 foreach(for(Integer i : list)) 语句时，会自动生成一个iterator 来遍历该 list，但同时该 list 正在被 Iterator.remove() 修改。Java 一般不允许一个线程在遍历 Collection 时另一个线程修改它、

RandomAccess 接口 按位置读取元素的平均时间复杂度为 O(1)，如ArrayList。

transient 修饰 
序列化 ArrayList 中的非 transient 元素

set是无序的

关于hashmap在jdk17和18的区别

ey.hashCode()与key.hashCode()>>>16进行异或操作，高16bit补0，一个数和0异或不变，所以 hash 函数大概的作用就是：高16bit不变，低16bit和高16bit做了一个异或，目的是减少碰撞。


Spring不能对一个prototype bean的整个生命周期负责，容器在初始化、配置、装饰或者是装配完一个prototype实例后，将它交给客户端，随后就对该prototype实例不闻不问了
