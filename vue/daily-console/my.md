#一、项目


#二、碰到的难点
2.1 redis内存达到90%
0.现象redis 16g的内存 占用了14g触发告警
1.拉到rdb文件
2.使用工具分析(rdr) 得出key占用分布图（某个前缀的key占用多少百分比 和 大小）
发现有一类key占用很大，没有设置超时时间
3.分析key的业务含义
是审批任务建立后发送代办给审批人，审批后删除此数据
4.从业务上分析合理性，超过5天的代办系统不做自动删除，由用户手动删除，对应技术上将这批key设置5天超时
5.结果 5天后内存下降至20%

2.2 日志打印导致接口慢
1.收到业务方反馈接口变慢，也收到接口变慢告警
2.根据api网关的报错日志，定位到报错节点集中到某两台机器上
3.对机器进行dump+jstack保留现场后，在api网关上摘除节点，恢复正常
4.对dump和jstack进行分析，dump没有问题
5.jstack分析发现有大量线程dubbo provider线程是block
6.根据block线程里查找等待的锁是由哪个线程持有
7.发现持有的线程是logback的线程，线程状态为runable，正在使用调用file.getLength方法
8.ll发现写日志的文件夹挂在一个nas盘上
9.使用dd命令做性能测试，发现写1g的文件要1min,正常的机器只要0.8秒
10.联系iaas同事查看发现nas盘正在做坏道修复

2.3 FULLGC问题排查(survivor过小)
1.每隔几个小时就FULLGC 5 6次 产生告警
2.通过观察gc日志是老年代被打满 但是没有配置自动打出内存dump 所以只通过fullgc日志目前得不到什么信息
3.发现fullgc前伴随着好几次younggc
4.观察younggc的日志 发现eden区清零 survivor区没有增加反而减少 但是older区却增加，fullgc前几次的younggc都是这样，知道把older区打满触发fullgc
5.有观察了jvm大小只有100多m,xmx设的是1g,所以定位为jvm自己管理的内存太小
6.增加xms参数解决问题

2.4 FULLGC问题(全表扫描)
1.收到业务方反馈接口缓慢
2.监控产生告警接口缓慢 fullgc次数过高
3.对产生fullgc的应用紧急重启
4.问题恢复
事后分析:
1.jvm老年代并没有缓慢上涨，说明不是已有问题，有可能TPS突然升高导致大流量
2.查看机器网卡入方向监控，问题时间点有突刺，但是排查入口apache tps监控并没有突刺，所以怀疑是查外部数据源MySQL redis返回大量数据
3.redis出口流量正常，mysql有一台从库出口流量飙升，基本判定某条sql返回大量数据将java进程内存打满导致fullgc进而服务不可用
4.通过排查apm链路监控，按照出问题的时间点找到耗时长的接口，最后定位到时外部掉我们接口参数穿了空串，导致mybatisplus框架生成sql时没有待where条件，查询全表数据大概800m

2.5 接口变慢
