package main

import (
	"flag"
	"fmt"
)

var kleur string

const (
	Blauw string = "blauw"
	Rood  string = "rood"
	Geel  string = "geel"
	Groen string = "groen"
	Bruin string = "bruin"
	Wit   string = "wit"
	Zwart string = "zwart"
)

func init() {

	flag.StringVar(&kleur, "k", "blauw", "ingevoerde kleur")
	flag.Parse()
}

func main() {
	if kleur == Blauw {
		fmt.Println(kleur, "Blauw zoals de lucht")
	}
	if kleur == Rood {
		fmt.Println("Rood met Passie")
	}
	if kleur == Geel {
		fmt.Println("Geel als de stralen van de zon")
	}
	if kleur == Groen {
		fmt.Println("Groen van de natuur")
	}
	if kleur == Bruin {
		fmt.Println("Bruin als de inhoud van een pot Nutella")
	}
	if kleur == Wit {
		fmt.Println("Wit als een gamer in het midden van de zomer")
	}
	if kleur == Zwart {
		fmt.Println("Zwart als de krochten van mijn ziel")
	}
}
