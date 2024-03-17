package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"go-packages/packages/money"
)

type goods struct {
	Id    int
	Count int
	Price string
}

type Coupon struct {
	GoodsId int
	Amount  string
}

func main() {

	coupon := Coupon{GoodsId: 1, Amount: "101.2982"}
	goodsA := goods{
		Id:    1,
		Count: 2,
		Price: "129.8612",
	}
	goodsB := goods{
		Id:    2,
		Count: 13,
		Price: "0.8231",
	}

	goodsAPrice := money.NewMoneyFromString(goodsA.Price, money.MoneyScale)

	goodsBPrice := money.NewMoneyFromString(goodsB.Price, money.MoneyScale)

	couponPrice := money.NewMoneyFromString(coupon.Amount, money.MoneyScale)

	goodsAPrice = goodsAPrice.Mul(decimal.NewFromInt(int64(goodsA.Count)))

	goodsBPrice = goodsBPrice.Mul(decimal.NewFromInt(int64(goodsB.Count)))

	money.FormatMoney("0", 2)

	orderAmount := goodsAPrice.Add(goodsBPrice).Sub(couponPrice)

	fmt.Println("itemA", goodsAPrice.String())
	fmt.Println("itemB", goodsBPrice.String())
	fmt.Println("coupon", couponPrice.StringFixed(money.MoneyScale))

	fmt.Println("order", orderAmount.String())

	money.FormatMoney("11111111110", 2)

}
