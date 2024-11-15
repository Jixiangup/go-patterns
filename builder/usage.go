package builder

// Usage 建造者使用方式示例
func Usage() {
	// 家庭用车
	familyCar := NewBuilder().Color(Black).Wheels(Steel).TopSpeed(50 * MPH).Build()
	_ = familyCar.Drive()

	// 运动车辆
	sportsCar := NewBuilder().Color(White).Wheels(Sports).TopSpeed(150 * MPH).Build()
	_ = sportsCar.Drive()

}
