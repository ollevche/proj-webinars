package order

type User struct {
	Name     string
	Email    string
	Password string
}

type Order struct {
	User        User
	ProductName string
	TotalPrice  float64
}

func ProcessOrder(o Order) error {
	// processing logic
	return nil
}
