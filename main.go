package main

func cafe(menu string, money float64) (string, float64) {
	price := map[string]float64{
		"tea": 50.0,
		"greentea": 60.0,
		"coffee": 70.0,
	}

	change := money - price[menu]
	return menu, change
}

