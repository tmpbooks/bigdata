zkServer.sh start-foreground cluster/z1/zoo.cfg
zkServer.sh start-foreground cluster/z2/zoo.cfg
zkServer.sh start-foreground cluster/z3/zoo.cfg

zkCli.sh -server localhost:2181
zkCli.sh -server localhost:2182
zkCli.sh -server localhost:2183
ls /
create /workers ""
ls /