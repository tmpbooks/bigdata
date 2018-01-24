安装Zookeeper出现Unable to start AdminServer，existing abnormally问题解决方法
http://www.cnblogs.com/byfcumt/p/8176903.html
出现这个问题主要是由于8080端口占用，可在zoo.cfg中增加admin.serverPort=没有被占用的端口号解决问题。