package main

type processedProductData struct {
	ProductName        string   `json:"productname"`
	Approved           bool     `json:"approved"`
	ProductFound       bool     `json:"productfound"`
	BadIngredients     []string `json:"badingredients"`
	AllIngredients     string   `json:"allIngredients"`
	NutriscoreGrade    string   `json:"nutriscoregrade"`
	NutriscoreDescript string   `json:"nutriscoredescript"`
}
