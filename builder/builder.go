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
