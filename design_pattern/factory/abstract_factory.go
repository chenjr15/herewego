package factory

/**
抽象工厂模式，尝试将产品按产品族(系列)对工厂进行分类，不同产品族生产的产品类型相同。如下面来自不同工厂的产品。
*/

/* ===== 抽象产品   =====*/

// Chair  抽象椅子
type Chair interface {
	DescribeChair() string
}

// Table 抽象桌子
type Table interface {
	DescribeTable() string
}

// Sofa 抽象沙发
type Sofa interface {
	DescribeSofa() string
}

// FurnitureFactory 抽象的工厂
type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
	CreateSofa() Sofa
}

/* ==== 具体产品 ====*/

type ModernChair struct{}

func (ca ModernChair) DescribeChair() string {
	return "[Modern] Chair"
}

type ModernTable struct{}

func (ca ModernTable) DescribeTable() string {
	return "[Modern] Table"
}

type ModernSofa struct{}

func (ca ModernSofa) DescribeSofa() string {
	return "[Modern] Sofa"
}

/* === 具体工厂 - Modern === */

type ModernFactory struct{}

func (cf *ModernFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (cf *ModernFactory) CreateTable() Table {
	return &ModernTable{}
}

func (cf *ModernFactory) CreateSofa() Sofa {
	return &ModernSofa{}
}

type RetroChair struct{}

func (ca RetroChair) DescribeChair() string {
	return "[Retro] Chair"
}

type RetroTable struct{}

func (ca RetroTable) DescribeTable() string {
	return "[Retro] Table"
}

type RetroSofa struct{}

func (ca RetroSofa) DescribeSofa() string {
	return "[Retro] Sofa "
}

/* === 具体工厂 - Retro === */

type RetroFactory struct{}

func (cf *RetroFactory) CreateChair() Chair {
	return &RetroChair{}
}

func (cf *RetroFactory) CreateTable() Table {
	return &RetroTable{}
}

func (cf *RetroFactory) CreateSofa() Sofa {
	return &RetroSofa{}
}
