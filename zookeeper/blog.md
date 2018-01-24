1.安装Zookeeper出现Unable to start AdminServer，existing abnormally问题解决方法
  http://www.cnblogs.com/byfcumt/p/8176903.html
  出现这个问题主要是由于8080端口占用，可在zoo.cfg中增加admin.serverPort=没有被占用的端口号解决问题。

2.bin/zkCli.sh -server ip:port
  bin/zkCli.sh -server 127.0.0.1:2181

3.ZooKeeper深入浅出
  https://holynull.gitbooks.io/zookeeper/content/zookeeperkafashili_md.html