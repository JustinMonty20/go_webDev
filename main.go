package main

import (
	"go_webDev/templating"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templating/html/menu.gohtml"))
}

func main() {
	bfast := templating.Menus{
		TimeOfDay: "Breafkast",
		Dishes: []templating.Plates{
			{"Eggs Benny", 16.00},
			{"Omlette", 10.00}},
	}

	lunch := templating.Menus{
		TimeOfDay: "Lunch",
		Dishes: []templating.Plates{
			{"Burger and Fries", 11.00},
			{"Flatbread", 6.00},
		},
	}

	dinner := templating.Menus{
		TimeOfDay: "Dinner",
		Dishes: []templating.Plates{
			{"Chicken and Risotto", 23.00},
			{"Filet Mignog", 45.00},
		},
	}

	r1 := templating.Resturant{
		"Justin's Place",
		[]templating.Menus{bfast, lunch, dinner},
	}

	r2 := templating.Resturant{
		Name: "Chessie's Place",
		Menu: []templating.Menus{
			{"Breakfast",
				[]templating.Plates{
					{"Instant Oats", 10.00},
					{"Tuna", 20.00},
				}}},
	}

	resturants := []templating.Resturant{r1, r2}

	err := tpl.ExecuteTemplate(os.Stdout, "menu.gohtml", resturants)
	if err != nil {
		log.Fatal(err)
	}

}
