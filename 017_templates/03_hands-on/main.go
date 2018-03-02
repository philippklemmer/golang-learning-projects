/*
	1. Using the data structure created in the previous
	folder, modify it to hold menu information for an
	unlimited number of restaurants.
*/
package main

import (
	"log"
	"os"
	"text/template"
)

type food struct {
	Name, Description string
	Price             int
	Vegan             bool
}

type meal struct {
	Name string
	Food []food
}

type restaurant struct {
	Name  string
	Meals []meal
}

type meals []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {

	gramercyTavernMeals := meals{
		meal{
			Name: "Lunch",
			Food: []food{
				{
					Name: "Marinated Scallops",
					Description: `Green Apple, Daikon
						Domaine de la Taille aux Loups, Brut Tradition, Montlouis-sur-Loire, France 2014`,
					Price: 35,
					Vegan: false,
				},
				{
					Name: "Sea Bream",
					Description: `Spaghetti Squash, Pumpkin Seeds, Sherry
						Zierfandler, Stadlmann, Mandel-Höh, Thermenregion, Austria 2012`,
					Price: 45,
					Vegan: false,
				},
				{
					Name: "Salsify",
					Description: `Honshimeji Mushrooms, Meyer Lemon, Cilantro
						Malvasia Istriana, Vignai da Duline, Chioma Integrale, Venezia Giulia, Italy 2015`,
					Price: 25,
					Vegan: true,
				},
			},
		},
		meal{
			Name: "Dinner",
			Food: []food{
				{
					Name:        "Roasted Mushroom Soup",
					Description: `Mushroom Soup with Sweet Potato and Walnuts`,
					Price:       27,
					Vegan:       true,
				},
				{
					Name:        "Smoked Trout",
					Description: `Smoked Trout with scallions, kohlrabi and roe`,
					Price:       40,
					Vegan:       false,
				},
				{
					Name:        "Lumache",
					Description: `Lumache with Beef Bolognese, Black Olives, Pickled Peppers`,
					Price:       60,
					Vegan:       true,
				},
			},
		},
	}

	theHaroldMeals := meals{
		meal{
			Name: "Breakfast",
			Food: []food{
				food{
					Name:        "Oatmeal",
					Description: "yum yum",
					Price:       4,
				},
				food{
					Name:        "Cheerios",
					Description: "American eating food traditional now",
					Price:       3,
				},
				food{
					Name:        "Juice Orange",
					Description: "Delicious drinking in throat squeezed fresh",
					Price:       2,
				},
			},
		},
		meal{
			Name: "Lunch",
			Food: []food{
				food{
					Name:        "Hamburger",
					Description: "Delicous good eating for you",
					Price:       6,
				},
				food{
					Name:        "Cheese Melted Sandwich",
					Description: "Make cheese bread melt grease hot",
					Price:       3,
				},
				food{
					Name:        "French Fries",
					Description: "French eat potatoe fingers",
					Price:       2,
				},
			},
		},
		meal{
			Name: "Dinner",
			Food: []food{
				food{
					Name:        "Pasta Bolognese",
					Description: "From Italy delicious eating",
					Price:       7,
				},
				food{
					Name:        "Steak",
					Description: "Dead cow grilled bloody",
					Price:       13,
				},
				food{
					Name:        "Bistro Potatoe",
					Description: "Bistro bar wood American bacon",
					Price:       6,
				},
			},
		},
	}

	restaurants := []restaurant{
		restaurant{
			Name:  "Gramercy Tavern",
			Meals: gramercyTavernMeals,
		},
		restaurant{
			Name:  "The Harold",
			Meals: theHaroldMeals,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "data.gohtml", restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}
