package shopping

import "fmt"

// PromotionStrategy 促销策略
type PromotionStrategy interface {
	GetPrice(price float64) float64
}

type Item struct {
	Name     string
	Price    float64
	strategy PromotionStrategy
}

func (item *Item) SetStrategy(strategy PromotionStrategy) {
	item.strategy = strategy
}

func (item *Item) GetSellPrice() float64 {
	if item.strategy == nil {
		return item.Price
	}
	return item.strategy.GetPrice(item.Price)
}
func (item *Item) String() string {
	return fmt.Sprintf("[%s 原价:%.2f 现价: %.2f]", item.Name, item.Price, item.GetSellPrice())
}

// DiscountStrategy  打折策略，直接乘上系数
type DiscountStrategy struct {
	Discount float64
}

func (s DiscountStrategy) GetPrice(price float64) float64 {
	return price * s.Discount
}

// ReductionStrategy  满减策略，满指定金额就减
type ReductionStrategy struct {
	Threshold float64
	Reduction float64
	Repeat    bool
}

func (s ReductionStrategy) GetPrice(price float64) float64 {
	if price < s.Threshold {
		return price
	}
	if s.Repeat {
		return price - s.Reduction*float64(int32(price/s.Threshold))
	}
	return price - s.Reduction
}
