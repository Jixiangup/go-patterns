package object_pool

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Usage() {
	pool := NewConnectionPool(5)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Press enter to run a query: ")
	query := make(chan string)

	for {
		// 逐行读取输入
		scanner.Scan()
		queryText := scanner.Text()
		if queryText == "exit" {
			log.Println("Exiting...")
			// 理论上应该先吧连接池所有连接都关掉
			os.Exit(1)
		}
		// 模拟并发请求
		go func(query chan string) {
			for {
				select {
				case queryText := <-query:
					// 借用连接
					conn, err := pool.Borrow()
					if err != nil {
						log.Println(err.Error())
						fmt.Printf("Press enter to run a query: ")
						continue
					}
					conn.Query(queryText)

					// 等待一下模拟并发连接不够的情况

					split := strings.Split(queryText, " ")
					if split != nil {
						// 模拟查询耗时
						sleepTime := split[len(split)-1]
						if sleepTime == "1" {
							log.Println("检测到需要休眠 程序将休眠20秒")
							fmt.Printf("Press enter to run a query: ")
							time.Sleep(20 * time.Second)
							pool.Return(conn)
							continue
						}

					}

					pool.Return(conn)
					fmt.Printf("Press enter to run a query: ")
				}
			}
		}(query)
		query <- queryText
	}
}
