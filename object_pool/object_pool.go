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
