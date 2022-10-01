package shopping

import (
	"fmt"
	"testing"
)

func TestChangeStrategy(t *testing.T) {
	cake := &Item{
		Name:  "月饼",
		Price: 100,
	}
	fmt.Println("中秋前一个月")
	fmt.Println(cake)
	fmt.Println("中秋前三天")
	cake.SetStrategy(DiscountStrategy{0.8})
	fmt.Println(cake)

	fmt.Println("中秋后")
	cake.SetStrategy(ReductionStrategy{80, 40, false})
	fmt.Println(cake)
}

func TestDiscountStrategy(t *testing.T) {
	type fields struct {
		Discount float64
	}
	type args struct {
		price float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{name: "半价", fields: fields{Discount: 0.5}, args: args{price: 100.0}, want: 50.0},
		{name: "八折", fields: fields{Discount: 0.8}, args: args{price: 100.0}, want: 80.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DiscountStrategy{
				Discount: tt.fields.Discount,
			}
			if got := s.GetPrice(tt.args.price); got != tt.want {
				t.Errorf("GetPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReductionStrategy(t *testing.T) {
	type fields struct {
		Threshold float64
		Reduction float64
		Repeat    bool
	}
	type args struct {
		price float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{name: "未达满减", fields: fields{200, 100, false}, args: args{price: 100.0}, want: 100},
		{name: "满200减100,300", fields: fields{200, 100, false}, args: args{price: 300.0}, want: 200},
		{name: "满200减100,500", fields: fields{200, 100, false}, args: args{price: 500.0}, want: 400},
		{name: "每满200减100,500", fields: fields{200, 100, true}, args: args{price: 500.0}, want: 300},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ReductionStrategy{
				Threshold: tt.fields.Threshold,
				Reduction: tt.fields.Reduction,
				Repeat:    tt.fields.Repeat,
			}
			if got := s.GetPrice(tt.args.price); got != tt.want {
				t.Errorf("GetPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
