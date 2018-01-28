package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"192.168.3.139:2181", "192.168.3.139:2182", "192.168.3.139:2183"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	for {
		c.Create("/global_lock", "", zk.FlagEphemeral, acl)
		children, stat, ch, err := c.ChildrenW("/")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v %+v\n", children, stat)
		e := <-ch
		fmt.Printf("%+v\n", e)
	}

}
