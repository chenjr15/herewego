package factory

import "testing"

func TestSameStyle(t *testing.T) {

	// 可以方便地增加产品族，增加一个产品族直接增加对应的工厂即可
	t.Run("Modern Style", func(t *testing.T) {
		var factory = new(ModernFactory)

		createFurniture(t, factory)

	})

	t.Run("Retro Style", func(t *testing.T) {
		var factory = new(RetroFactory)
		createFurniture(t, factory)

	})
}

func createFurniture(t *testing.T, factory FurnitureFactory) {
	// 但是对于来自不同产品族的产品，并不一定能减少代码量
	var chair Chair = factory.CreateChair()
	t.Log(chair.DescribeChair())

	var table Table = factory.CreateTable()
	t.Log(table.DescribeTable())

	var sofa Sofa = factory.CreateSofa()
	t.Log(sofa.DescribeSofa())
}

func TestDifferentStyle(t *testing.T) {
	modernFactory := new(ModernFactory)
	retroFactory := new(RetroFactory)

	var chair Chair
	var sofa Sofa

	chair = modernFactory.CreateChair()
	sofa = retroFactory.CreateSofa()

	t.Log(chair.DescribeChair())
	t.Log(sofa.DescribeSofa())
}
