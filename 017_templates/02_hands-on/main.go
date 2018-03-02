/*
	1. Create a data structure to pass to a template which
	* contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
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

type meals []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {

	m := meals{
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

	err := tpl.ExecuteTemplate(os.Stdout, "data.gohtml", m)
	if err != nil {
		log.Fatalln(err)
	}
}
