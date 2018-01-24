zkServer.sh start-foreground cluster/z1/zoo.cfg
zkServer.sh start-foreground cluster/z2/zoo.cfg
zkServer.sh start-foreground cluster/z3/zoo.cfg

//连接单个服务器
zkCli.sh -server localhost:2181
zkCli.sh -server localhost:2182
zkCli.sh -server localhost:2183

//连接服务器列表，逗号分隔，不能有空格
zkCli.sh -server localhost:2181,localhost:2182,localhost:2183
ls /
create /workers ""
ls /