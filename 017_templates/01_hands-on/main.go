/*
	1. Create a data structure to pass to a template which
	* contains information about California hotels including Name, Address, City, Zip, Region
	* region can be: Southern, Central, Northern
	* can hold an unlimited number of hotels
*/
package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City, Zip, Region string
}

type region struct {
	Region string
	Hotel  []hotel
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	regionsHotelData :=
		[]region{
			region{
				Region: "Central",
				Hotel: []hotel{
					{
						Name:    "Best Western Califonia City Inn & Suites",
						Address: "10386 California City Boulevard",
						City:    "California City",
						Zip:     "93505",
					},
					{
						Name:    "Ayres Lodge Alpine",
						Address: "1251 Tavern Rd",
						City:    "Alpine",
						Zip:     "991901",
					},
					{
						Name:    "Fairfiled Inn & Suites by Marriott",
						Address: "525 North Sepulveda Boulevard",
						City:    "El Segundo",
						Zip:     "90245",
					},
				},
			},
			region{
				Region: "Southern",
				Hotel: []hotel{
					{
						Name:    "Southern",
						Address: "10386 California City Boulevard",
						City:    "California City",
						Zip:     "93505",
					},
					{
						Name:    "Southern",
						Address: "1251 Tavern Rd",
						City:    "Alpine",
						Zip:     "991901",
					},
					{
						Name:    "Southern",
						Address: "525 North Sepulveda Boulevard",
						City:    "El Segundo",
						Zip:     "90245",
					},
				},
			},
			region{
				Region: "Northern",
				Hotel: []hotel{
					{
						Name:    "Northern",
						Address: "10386 California City Boulevard",
						City:    "California City",
						Zip:     "93505",
					},
					{
						Name:    "Northern",
						Address: "1251 Tavern Rd",
						City:    "Alpine",
						Zip:     "991901",
					},
					{
						Name:    "Northern",
						Address: "525 North Sepulveda Boulevard",
						City:    "El Segundo",
						Zip:     "90245",
					},
				},
			},
		}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", regionsHotelData)
	if err != nil {
		log.Fatalln(err)
	}
}
