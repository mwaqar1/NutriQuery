package main

type processedProductData struct {
	//TODO:
	//1. return a json with two keys
	//key1 approved boolean
	//key2 badingredients array[] string
	//key3 allingredients array [s]string"`
	//ORIGINAL
	//AllIngredients   []string `json:"allIngredients"`
	ProductName        string   `json:"productname"`
	Approved           bool     `json:"approved"`
	ProductFound       bool     `json:"productfound"`
	BadIngredients     []string `json:"badingredients"`
	AllIngredients     string   `json:"allIngredients"`
	NutriscoreGrade    string   `json:"nutriscoregrade"`
	NutriscoreDescript string   `json:"nutriscoredescript"`

	//1b. when no product found return no product found
	//when bad ingreidient found, not approved
	//when no bad ingredient, approved

}
