package templating

type Resturant struct {
	Name string
	Menu []Menus
}

type Menus struct {
	TimeOfDay string
	Dishes    []Plates
}

type Plates struct {
	Name  string
	Price float64
}
