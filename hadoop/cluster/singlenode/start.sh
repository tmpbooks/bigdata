#!/bin/bash

# http://hadoop.apache.org/docs/stable/hadoop-project-dist/hadoop-common/SingleCluster.html

#config ssh
ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa
cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys
chmod 0600 ~/.ssh/authorized_keys

bin/hdfs namenode -format

#start
sbin/start-dfs.sh

# visit http://localhost:50070/

#example
bin/hdfs dfs -mkdir /user
bin/hdfs dfs -mkdir /user/<username>
bin/hdfs dfs -put etc/hadoop /input
bin/hadoop jar share/hadoop/mapreduce/hadoop-mapreduce-examples-2.9.0.jar grep /input /output 'dfs[a-z.]+'
bin/hdfs dfs -get output output
cat output/*
bin/hdfs dfs -cat /output/*

#stop
sbin/stop-dfs.sh
