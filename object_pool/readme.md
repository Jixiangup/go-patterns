<h1>对象池模式</h1>

__对象池模式（Object Pool Pattern）__ 是一种创建型设计模式，它通过缓存已经创建好的对象，避免频繁的创建和销毁对象，尤其在对象创建成本较高或对象数量有限的情况下。对象池模式的核心思想是：当系统需要某些资源的时候，从池中获取一个对象，而不是每次都创建一个新对象；在资源使用完毕后，将其放回到池中，而不是销毁。这样可以在其他地方重复使用这些对象。

主要目的

- __节省资源__: 通过对象池，可以减少对象的创建和销毁次数以及避免频繁创建和销毁对象，提高系统性能。
- __管理资源__: 通过对象池，可以更好地管理对象的生命周期，避免资源泄漏。
- __控制并发__: 通过对象池，可以控制通过限制池中最大可用对象的数量，可以有效的控制并发访问资源的数量，避免资源竞争和超负荷。

--- 

<h2>目录</h2>

<!-- TOC -->
  * [关键概念](#关键概念)
  * [适用场景](#适用场景)
  * [优缺点](#优缺点)
    * [优点](#优点)
    * [缺点](#缺点)
  * [举例](#举例)
  * [实现](#实现)
    * [接口定义](#接口定义)
  * [使用](#使用)
<!-- TOC -->

## 关键概念

1. __对象池__: 池管理多个可重用的对象，负责管理对象的创建、借用、归还、销毁等操作。
2. __对象__: 池中存放的对象，通常具有某些特殊资源的实例，比如数据库连接、网络连接、线程等。
3. __借用__: 从对象池中获取一个对象，通常会先检查池中是否有空闲对象，如果池中没有可用对象，如果没有则根据一定的策略创建一个新对象或者等待以及报错。
4. __归还__: 将对象放回到对象池中，以便其他地方重复使用。
5. __最大池大小__: 对象池中最大可用对象的数量，当池中对象数量达到最大值时，新的借用请求会被阻塞或者报错。

创建对象池可以将其分为以下几点

- __对象池__: 管理池中所有对象，提供借用和归还接口，负责对象的创建、销毁和资源管理。
- __对象__: 实际资源对象，如数据库连接、网络连接等。
- __工厂__: 负责对象的创建，可以是简单工厂、工厂方法、抽象工厂等。
- __对象回收__: 对象池中对象的回收策略，如对象的空闲时间、最大存活时间等。
- __策略__: 对象池中对象的策略，如对象的获取方式、等待超时时间等，最大池大小等。

## 适用场景

- __数据库连接池__: 在应用程序中，连接数据库的操作可能非常昂贵。通过使用对象池，可以避免频繁创建和销毁数据库连接，提高性能。
- __线程池__: 线程的创建和销毁也是非常耗费资源的，线程池允许程序使用固定数量的线程，并复用这些线程处理多个任务。
- __网络连接池__: 在网络编程中，网络连接的创建和销毁也是非常耗费资源的，通过使用对象池，可以避免频繁创建和销毁网络连接，提高性能。
- __大型游戏服务器__: 在大型游戏服务器中，可能会有大量的玩家连接，通过使用对象池，可以避免频繁创建和销毁玩家对象，提高性能。

## 优缺点

### 优点

- __减少资源消耗__: 避免了频繁创建和销毁对象，减少了内存分配和垃圾回收的压力。
- __提高性能__: 通过对象池，可以重复使用对象，避免了频繁创建和销毁对象，提高了系统性能。
- __管理资源__: 对象池可以设置最大池大小，控制系统的并发访问亮，避免资源耗尽或过度使用。
- __便于管理__: 通过对象池，可以更好地管理对象的生命周期， 可以集中管理资源的借用、归还、清理等操作，减少了代码的复杂性和错误的可能性。

### 缺点

- __对象池管理复杂性__: 需要确保对象池的管理方式合理，避免内存泄漏或资源占用过多。
- __初始化开销__: 当池中对象的创建非常昂贵时，池的初始化可能会变得耗时。
- __池满的处理__: 如果池的对象数量达到最大限制，后续请求可能会被阻塞或失败，因此需要合理设计最大池大小和阻塞策略。
- __过期对象__: 池中对象可能会因为长时间未使用而变得不可用，因此需要定期清理过期或无效对象(这里我们可以考虑采用心跳做保活，然后根据策略比如最小连接池大小来清理连接知道当前池中数量和最小连接池一致之后，为这些连接做保活)。

## 举例

当Web服务器等场景下会出现并发请求后台接口的情况，如果每次请求都创建新的数据库连接性能和`GC`压力时比较大的所以可以选择提前先创建一部份连接放在数据库连接池，如果有请求过来自动请求即可。

## 实现

### 接口定义

```go
package object_pool

import (
	"fmt"
	"log"
	"sync"
)

// Connection 连接对象
type Connection struct {
	ID int
}

// Query 运行查询
func (c Connection) Query(query string) {
	fmt.Printf("Connection %d: Running query: %s\n", c.ID, query)
}

// ConnectionPool 连接池
type ConnectionPool struct {
	// connection 连接池
	connection chan *Connection
	// lock 互斥锁
	lock   sync.Mutex
	nextID int
}

// NewConnectionPool 创建连接池
func NewConnectionPool(size int) *ConnectionPool {
	pool := &ConnectionPool{
		connection: make(chan *Connection, size),
		nextID:     0,
	}

	// 初始化连接池
	for i := 0; i < size; i++ {
		pool.connection <- &Connection{ID: pool.nextID}
		pool.nextID++
	}
	return pool

}

// Borrow 借用连接
func (p *ConnectionPool) Borrow() (*Connection, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.nextID += 1
	select {
	case connection := <-p.connection:
		if connection == nil {
			return nil, fmt.Errorf("connection pool is empty")
		} else {
			log.Println("Borrow connection", connection.ID)
		}
		return connection, nil
	default:
		// 如果没有连接可用 则返回错误
		return nil, fmt.Errorf("connection pool is empty")
	}

}

// Return 归还连接
func (p *ConnectionPool) Return(conn *Connection) {
	p.lock.Lock()
	defer p.lock.Unlock()
	log.Println("Return connection", conn.ID)
	p.nextID = conn.ID
	p.connection <- conn
}
```

## 使用

> 见 [usage.go](usage.go)

```go
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
```