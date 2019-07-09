package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//定义的此结构体为全局map的value值，包括每一个用户的姓名，ip地址和私人管道
type client struct {
	name string
	addr string
	C    chan string
}

/*这个函数是将私人管道中的内容发送给用户，配合全局管道Message使用可以实现广播的功能，
单独使用可以实现私聊的功能*/
func writemsg2client(clinet client, conn net.Conn) {
	for m := range clinet.C {
		conn.Write([]byte(m + "\n"))
	}
}

//这只是一个封装好用来统一(发送信息格式)的小函数，不用在意
func makemsg(name string, addr string, s string) string {
	return "[" + addr + "]" + name + s
}

//每一个进入聊天室的用户都将启动一个handleconn的go程来处理事件
func handleconn(conn net.Conn) {
	defer conn.Close()
	/*用户连接进来以后要初始化全局map，把自己的信息加入到字典里，相当于进到聊天室里之前要登
	记一下个人信息，注意姓名初始为ip地址。*/
	addr := conn.RemoteAddr().String()
	fmt.Printf("用户%s进入了房间\n", addr)
	client := client{addr, addr, make(chan string)}
	//在这里启动子go程，功能上面已经提及
	go writemsg2client(client, conn)
	onlinemap[addr] = client
	//登录进来一切准备就绪后就给所有人广播上线信息啦
	Message <- makemsg(client.name, addr, "login")
	//下面这三个变量服务于下面一些小功能
	var haschat = make(chan bool)
	var ifquit = make(chan bool)
	var flag bool
	//从这单独开启一个go程来读取用户输入的信息
	go func() {
		buf := make([]byte, 4096)
		for {
			fmt.Println("Please enter your message:")
			n, _ := conn.Read(buf)
			if n == 0 {
				fmt.Printf("%s离开了房间\n", client.name)
				ifquit <- true
				return
			}
			//改名功能的实现
			if string(buf[:7]) == "Rename|" {
				client.name = strings.Split(string(buf[:n-1]), "|")[1]
				onlinemap[addr] = client
				conn.Write([]byte("rename success\n"))
			} else if string(buf[:n-1]) == "/who" {
				//查询在线用户信息的功能
				for _, s := range onlinemap {
					conn.Write([]byte(s.name + "online\n"))
				}
			} else if string(buf[:2]) == "m|" && strings.Count(string(buf[:n]), "|") == 2 {
				/*私聊功能的实现，其实私聊功能就是跳过了往全局Message里传输信息，
				改为直接向私人管道里传输信息*/
				flag = false
				slice := strings.Split(string(buf[:n-1]), "|")
				for _, a := range onlinemap {
					//遍历所有在线用户，向指定的用户管道中发送信息
					if a.name == slice[1] {
						flag = true
						a.C <- makemsg(client.name, addr, slice[2])
						conn.Write([]byte("send success"))
					}
				}
				if flag {
					conn.Write([]byte("no such man or not online"))
				}
			} else {
				Message <- makemsg(client.name, addr, string(buf[:n-1]))
			}
			haschat <- true
		}
	}()
	for {
		select {
		case <-haschat:
			//超时强踢
		case <-time.After(time.Minute * 3):
			delete(onlinemap, addr)
			Message <- makemsg(client.name, addr, "out time to leave")
			close(client.C)
			return
		case <-ifquit:
			//退出处理
			delete(onlinemap, addr)
			Message <- makemsg(client.name, addr, "out time to leave")
			close(client.C)
			return
		}
	}
}

//这个函数用来将全局Message中的内容全部塞到私人管道C里，实现上下线广播和群聊的功能
func Manager() {
	for {
		msg := <-Message
		for _, s := range onlinemap {
			s.C <- msg
		}
	}
}

var Message = make(chan string)
var onlinemap map[string]client = make(map[string]client)

//主函数
func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:6666")
	defer listener.Close()
	//提前开启全局Message的go程，防止被阻塞
	go Manager()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			continue
		}
		//每一个连接进来的用户都会被分配进入一个子go程，用来处理上面我们提到的各种功能
		go handleconn(conn)
	}
}
