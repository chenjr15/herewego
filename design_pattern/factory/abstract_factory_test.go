package factory

import "testing"

func TestSameStyle(t *testing.T) {
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
