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
