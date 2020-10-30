package templating

type Hotel struct {
	Name string
	City string
	Zip  int
}

type Region struct {
	Region string
	Hotels []Hotel
}
