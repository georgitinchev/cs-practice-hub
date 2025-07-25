package gross

// stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"each":               1,
		"dozen":              12,
		"gross":              144,
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"small_gross":        120,
		"great_gross":        1728,
	}
}

// creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}
	bill[item] += units[unit]
	return true
}

// removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, ok := units[unit]
	if !ok {
		return false
	}
	quantity, exists := bill[item]
	if !exists || quantity < value {
		return false
	}
	bill[item] -= value
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// returns the quantity of an item that the customer has in their bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
	return quantity, ok
}
