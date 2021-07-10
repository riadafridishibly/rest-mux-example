package api

// List of products

type Product struct {
	Name string
	ID   string
}

var products = []Product{
	{
		Name: "Shanling UA2",
		ID:   "E123XC",
	},
	{
		Name: "Moondrop Aria",
		ID:   "E3C919",
	},
	{
		Name: "Final Audio E3000c",
		ID:   "E13ED3",
	},
}
