package main

import (
	"fmt"
)

func main() {

	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(0),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(),
		Sodium:              SodiumMilligram(),
		Fruits:              FruitsPercent(),
		Fibre:               FibreGram(),
		Protein:             ProteinGram(),
	}, Food)
	fmt.Printf("Nurtitional Score %d\n", ns.Value)
}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints()
		fibrePoints := n.Fibre.GetPoints()

		negative = n.Energy.GetPoints() + n.Sugars.GetPoints() + n.SaturatedFattyAcids.GetPoints() + n.Sodium.GetPoints()
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints()
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}
