package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
**/
import (
	"fmt"
	zk "github.com/samuel/go-zookeeper/zk"
	"net"
	"os"
	"time"
)

var (
	hosts = []string{"192.168.184.164:2181", "192.168.184.164:2182", "192.168.184.164:2183"}
)

func main() {
	go starServer("127.0.0.1:8897")
	go starServer("127.0.0.1:8898")
	go starServer("127.0.0.1:8899")

	a := make(chan bool, 1)
	<-a
}

func checkError(err error) {
	fmt.Println("----- err:", err)
	os.Exit(-1)
}
func GetConnect() (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect(hosts, 5*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}

func starServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	fmt.Println(tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	//注册zk节点q
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s ", err)
	}
	defer conn.Close()
	err = RegistServer(conn, port)
	if err != nil {
		fmt.Printf(" regist node error: %s ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			continue
		}
		go handleCient(conn, port)
	}

	fmt.Println("aaaaaa")
}

func handleCient(conn net.Conn, port string) {
	defer conn.Close()

	daytime := time.Now().String()
	conn.Write([]byte(port + ": " + daytime))
}
