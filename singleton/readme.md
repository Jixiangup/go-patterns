<h1>单例模式</h1>

__单例模式（Singleton Pattern）__ 确保一个对象在系统中只有一个实例，并提供全局访问点。单例模式的主要目的是避免在系统中产生多个对象实例，从而节省内存和资源，确保一致性。

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
    * [实例实现](#实例实现)
  * [使用](#使用)
<!-- TOC -->

## 关键概念

1. __唯一性__: 全局只存在一个实例。
2. __全局访问__: 提供一个全局访问点，可以在任何地方访问该实例。
3. __延迟初始化__: 在第一次访问时创建实例。

## 适用场景

- 需要控制资源的共享(例如数据库连接池、日志管理器等)。
- 确保系统中某个实例只会存在一个并且能全局访问。
- 在大规模系统中，控制实例化的数量和资源消耗。

## 优缺点

### 优点

- __节省内存__: 由于只有一个实例，节省了内存。
- __全局访问__: 可以在任何地方访问该实例。
- __延迟初始化__: 在第一次访问时创建实例，避免了资源浪费。

### 缺点

- __隐藏的依赖关系__: 单例模式会隐藏类之间的依赖关系，可能会导致代码的可维护性降低。
- __难以测试__: 单例模式会使代码难以测试，因为单例对象会在整个系统中共享，可能会导致测试用例之间相互影响。
- __全局状态__: 单例模式会引入全局状态，可能会导致系统的复杂性增加。

## 举例

描述

## 实现

### 接口定义

```go
type User struct {
	Nickname string
	Email    string
}
```

### 实例实现

```go
var (
    user *User
    once sync.Once
)

func NewUser() *User {
    once.Do(func() {
        user = new(User)
        user.Nickname = "guest"
        user.Email = "guest@gmail.com"
    })
    return user
}
```

## 使用

> 见 [usage.go](usage.go)

```go
func Usage() {
	user := NewUser()
	log.Println(user)
}
```