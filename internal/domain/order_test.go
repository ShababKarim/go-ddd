package domain

import "testing"

func TestOrder(t *testing.T) {
	testDrink := &coffee{}
	var testDrinkWithOatmilk = &AddOn{
		name:  "oatmilk",
		price: 1,
		drink: testDrink,
	}

	var (
		customer = "bob"
		drinks   = testDrinkWithOatmilk
	)

	order, err := NewOrder(PlaceOrder{customer: customer, drinks: []Drink{drinks}})
	if err != nil {
		t.Fail()
	}

	if totalPrice := order.CalculateTotalPrice(); totalPrice != 4 {
		t.Fail()
	}
}

func TestOrder_WithInvalidInputs(t *testing.T) {
	testDrink := &coffee{}
	var testDrinkWithOatmilk = &AddOn{
		name:  "oatmilk",
		price: 1,
		drink: testDrink,
	}

	var (
		customer = "bob"
		drinks   = testDrinkWithOatmilk
	)

	_, err := NewOrder(PlaceOrder{customer: "", drinks: []Drink{drinks}})
	if err == nil {
		t.Fail()
	}

	_, err = NewOrder(PlaceOrder{customer: customer, drinks: nil})
	if err == nil {
		t.Fail()
	}
}

func TestUpdateOrderStatus(t *testing.T) {
	testDrink := &coffee{}
	order, _ := NewOrder(PlaceOrder{customer: "bob", drinks: []Drink{testDrink}})

	order, err := order.UpdateOrderStatus(UpdateOrderStatus{
		NewOrderStatus: Completed,
	})
	if err != nil {
		t.Fail()
	}
}

type coffee struct{}

func (c coffee) GetName() string {
	return "coffee"
}

func (c coffee) GetPrice() int {
	return 3
}
