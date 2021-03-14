package ocp

type color int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) filterByColor(
	products []Product, color Color) []*Product {
		result := make([]*Product, 0)

		
	}

	return result
}