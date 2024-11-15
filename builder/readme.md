# 建造者模式

<h2>目录</h2>

<!-- TOC -->
* [建造者模式](#建造者模式)
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

__建造者模式（Builder Pattern）__ 是一种创建型设计模式，旨在将一个复杂对象的构建过程与其表示分离，使得同样的构建过程可以创建不同类型的对象。简单来说，建造者模式提供了一个指挥者来构建产品的各个部分，但不关心每个具体部分如何实现，最终只关注产品的最终表现。

## 关键概念

1. __复杂对象的构建__：如果一个对象的构建过程比较复杂，涉及多个步骤或者多个组成部分，建造者模式可以帮助简化这个过程。

2. __分离构建与表示__：建造者模式通过提供不同的建造者（Builder）来控制产品的构建，而不需要暴露具体的构建细节。客户端不需要知道如何组装每个部分，只需要通过建造者来一步步构建产品。

3. __可变的产品__：通过改变建造者的不同配置，可以生产不同的产品，所有的产品具有共同的构建过程和相同的表现方式，但具体的组成部分会有所不同。

## 适用场景

- __构建复杂对象__：当对象的构建涉及多个部分或步骤时，尤其当这些步骤的顺序或方式变化时，建造者模式可以帮助简化构建过程。
- __多个变体的对象__: 如果同一个对象需要有多个变体或表示，建造者模式可以通过构建者的不同实现来生产不同的变体。

## 优缺点

### 优点

- __解藕构建过程和表示__： 客户端不需要知道产品如何构建，只需要关注最终产品，简化了代码的使用。
- __可扩展性强__：添加新的构建步骤或新的产品变体时，只需要修改具体的建造者类，不需要修改客户端代码或产品本身。
- __提高代码清晰度__：将复杂的构建过程分解多个步骤，使得每个步骤都更抑郁理解和维护。

### 缺点

- __产品变化较多时，建造者增多__：当产品的种类或构建方式发生变化时，需要创建多个建造者，可能会增加系统的`复杂度`。

## 举例

假设你在构建一辆车，车的组成部分包括发动机、轮子、座椅和车门。每个部分的构建过程可以独立，但总体的构建过程可能很复杂。你可以有一个 `车的建造者（CarBuilder）`，它提供了设置每个部分的方法，比如设置发动机类型、轮胎种类、座椅材料等。而一个 `指挥者（CarDirector）` 控制着这个建造过程，确保按照正确的顺序来构建一辆完整的车。最终，客户端只需要通过建造者来获取不同配置的车，而不需要关心每个部分是如何构建的。

## 实现

### 接口定义

```go
package builder

// Speed 车辆行驶速度 一小时行驶 Speed 公里(英里)
type Speed float64

// Color 车辆颜色
type Color string

// Wheels 车辆轮胎
type Wheels string

const (
	// MPH 英里/小时
	MPH Speed = 1
	// KPH 千米/小时
	KPH Speed = 1.60934
)

const (
	Black Color = "black"
	White Color = "white"
	Red   Color = "red"
)

const (
	// Sports 运动轮胎
	Sports Wheels = "sports"
	Steel  Wheels = "steel"
)

// Builder 构建器
type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	// Drive 驾驶车辆
	Drive() error
	// Stop 停止车辆
	Stop() error
}
```

### 实例实现

```go
package builder

import "fmt"

// Car 车辆实例
type Car struct {
	Color    Color
	Wheel    Wheels
	TopSpeed Speed
}

type CarBuilder struct {
	car *Car
}

func (c *CarBuilder) Color(color Color) Builder {
	c.car.Color = color
	return c
}

func (c *CarBuilder) Wheels(wheels Wheels) Builder {
	c.car.Wheel = wheels
	return c
}

func (c *CarBuilder) TopSpeed(speed Speed) Builder {
	c.car.TopSpeed = speed
	return c
}

func (c *CarBuilder) Build() Interface {
	return c.car
}

func (c *Car) Drive() error {
	fmt.Printf("驾驶颜色 %s,最高时速 %f/英里,轮胎 %s的车辆\n", c.Color, c.TopSpeed, c.Wheel)
	return nil
}

func (c *Car) Stop() error {
	fmt.Printf("停止颜色 %s 最高时速 %f/英里 轮胎 %s 的车辆\n", c.Color, c.TopSpeed, c.Wheel)
	return nil
}

func NewBuilder() Builder {
	c := new(CarBuilder)
	c.car = new(Car)
	return c
}
```

## 使用

> 见 [usage.go](usage.go)

```
// 家庭用车
familyCar := NewBuilder().Color(Black).Wheels(Steel).TopSpeed(50 * MPH).Build()
_ = familyCar.Drive()

// 运动车辆
sportsCar := NewBuilder().Color(White).Wheels(Sports).TopSpeed(150 * MPH).Build()
_ = sportsCar.Drive()
```