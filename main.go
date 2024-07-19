//TODO:
//make it a docker image
//seperate branches of trie vs contains

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// Retrieve the baseUrl from the environment variable or use a default value
var baseUrl = func() string {
	url := os.Getenv("BASEURL")
	return url
}()

const (
	searchAllEndPointsV2Api  = "api/v2/search"
	searchByProductTermV1api = "cgi/"
	queryParamsV1Api         = "search.pl?"
)

// Function to create and return an empty ProductsV1Slice from a ProductV1 struct
// When declaring a slice field in a struct, if it's not explicitly initialized it will be nil since its a reference type in go lang
func createEmptyProductV1() ProductV1 {
	return ProductV1{
		ProductsV1Slice: []ProductsV1Fields{}, // Initialize as an empty slice so that the ProductsV1Slice field  to prevent potential nil pointer dereference errors
	}
}

// creating a method on the Product struct to take the json repsonse object and store is as a go struct to process in this go code
func (productv1 *ProductV1) productStructMethodV1(data []byte) error {
	//TODO: test if & is necessary, productv1 is already a pointer
	// err := json.Unmarshal(data, &productv1)
	err := json.Unmarshal(data, productv1)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}
	return nil
}

func getRequest(endpointToSearchFor string) ([]byte, error) {
	res, err := http.Get(endpointToSearchFor)

	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", res.Status)
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}

func buildSearchProductEndpointUrl(searchProductInput *string) string {

	// Building the query parameters

	//This creates empty map data structure for storing URL parameters. The url.Values type comes from the net/url package designed  for handling URL-encoded parameters
	paramsBuildSearchProductUrl := url.Values{}

	//This adds a key-value pair to the map. The key "search_terms" specifies what the API should search for, and searchProduct holds the passed search product string.
	paramsBuildSearchProductUrl.Set("search_terms", *searchProductInput)

	//"search_simple" is a flag to simplify the search, and "1" indicates this mode is enabled
	paramsBuildSearchProductUrl.Set("search_simple", "1")

	//json" is a key of the response format data, and "1" specifies that the response should be in JSON format.
	paramsBuildSearchProductUrl.Set("json", "1")

	//return the complete URL with the complete paramters
	endpointToSearchFor := baseUrl + searchByProductTermV1api + queryParamsV1Api + paramsBuildSearchProductUrl.Encode()

	// DEBUG
	// fmt.Println(endpointToSearchFor)

	return endpointToSearchFor
}

func populateNutriScoreDiplay(nutriscoreGrade *string) string {
	if nutriscoreGrade == nil {
		log.Println("nutriscoreGrade is nil")
		return "unknown"
	}

	// Convert to lowercase
	*nutriscoreGrade = strings.ToLower(*nutriscoreGrade)
	// Remove leading and trailing whitespaces
	*nutriscoreGrade = strings.TrimSpace(*nutriscoreGrade)

	var descript string

	switch *nutriscoreGrade {
	case "a":
		descript = "Nutrition Grade: a\nNutriscore Grade Description: Excellent Choice\nNutriscore Grade Details: Products graded 'A' are among the healthiest options. They typically contain a beneficial balance of essential nutrients, with ample vitamins, minerals, and fiber, and minimal sugar, salt, and saturated fats."
	case "b":
		descript = "Nutrition Grade: b\nNutriscore Grade Description: Good Choice\nNutriscore Grade Details: These products are generally good for your health. They have more favorable nutritional profiles than average, often with lower sugar and fat content and higher essential nutrients compared to Grade C products."
	case "c":
		descript = "Nutrition Grade: c\nNutriscore Grade Description: Fair Choice\nNutriscore Grade Details: Products with a 'C' grade have a moderate nutritional value. They might contain a higher amount of sugars, salts, or saturated fats than preferable, or may lack some important nutrients."
	case "d":
		descript = "Nutrition Grade: d\nNutriscore Grade Description: Poor Choice\nNutriscore Grade Details: These items are below average in nutritional quality. They often have high levels of sugars, sodium, or unhealthy fats, and provide fewer beneficial nutrients. Consider consuming them sparingly."
	case "e":
		descript = "Nutrition Grade: e\nNutriscore Grade Description: Unhealthy Choice\nNutriscore Grade Details: Products graded 'E' have the least healthy nutritional profiles. They are typically high in negative components like sugar, salt, and saturated fats, and low in positive nutrients. Consumption should be limited due to potential health risks."
	default:
		descript = "unknown"
	}
	return descript
}

// displayProductInfo prints the product information in a user-friendly format
func (product *processedProductData) displayProductInfo(productName *string) string {
	if product == nil {
		log.Println("Product is nil when trying to display product info.")
		return ""
	}
	if productName == nil {
		log.Println("Product name is nil when trying to display product info.")
		return ""
	}

	var builder strings.Builder
	builder.WriteString("=== Detailed Product Review ===\n\n")

	var productFound string
	if product.ProductFound {
		productFound = fmt.Sprintf("Product found: Yes, the \"%s\" product entered was found in our database.\n\n", *productName)
	} else {
		productFound = fmt.Sprintf("Product found: No, the \"%s\" product entered was not found in our database.\n\n", *productName)
	}
	builder.WriteString("[Product Availability]\n" + productFound)

	// Bad ingredients
	builder.WriteString("[Health Warning]\n")
	if len(product.BadIngredients) > 0 {
		builder.WriteString("This product contains the following potentially harmful ingredients to consider avoiding:\n")
		for _, ingredient := range product.BadIngredients {
			builder.WriteString(fmt.Sprintf("- %s\n", ingredient))
		}
		builder.WriteString("\n")
	} else {
		builder.WriteString("We did not find any harmful ingredients in the product.\n\n")
	}

	// All ingredients
	builder.WriteString("[Full Ingredient List]\n")
	if product.AllIngredients == "" {
		builder.WriteString("No ingredient data is available for this product.\n")
	}
	builder.WriteString(strings.ReplaceAll(product.AllIngredients, ",", ", ") + "\n\n")

	//NutriScore information
	builder.WriteString("[Nutrition Score Information]\n")
	if product.NutriscoreDescript == "unknown" || product.NutriscoreGrade == "" {
		builder.WriteString("NutriScore information is not available for this product.\n")
	} else {
		builder.WriteString(product.NutriscoreDescript)
		builder.WriteString("\n")
	}

	var status string
	if product.Approved {
		status = "This product is recommended for consumption.\n\n"
	} else {
		status = "This product is not recommended for consumption.\n\n"
	}
	builder.WriteString("\n[Approval Status]\n" + status)

	return builder.String()
}

func searchByProductLogic(searchProductInput *string) processedProductData {

	if searchProductInput == nil {
		//create and return an empty structToPassToFrontEnd with the default values because there was an error when making a get request to the requested endpoint
		return processedProductData{}
	}

	endpointToSearchFor := buildSearchProductEndpointUrl(searchProductInput)

	//calls the getDataResponseInJsonObjectforAllSearchEndPoints which calls a get method on a search endpoint and returns the response in converted json string
	responseJson, error := getRequest(endpointToSearchFor)
	//fmt.Println(responseJson)

	if error != nil {
		log.Println("Error: ", error)
		//create and return an empty structToPassToFrontEnd with the default values because there was an error when making a get request to the requested endpoint
		return processedProductData{}
	}

	//TODO: modularize
	//create an empty productv1
	productv1 := createEmptyProductV1()
	//populate the productsV1Slice with the responseJson string using the productStructMethodV1 on the empty productsV1Slice
	err := productv1.productStructMethodV1(responseJson)

	if err != nil {
		//log.Fatal(err)
		//create and return an empty structToPassToFrontEnd because there was an error when populating productv1
		return processedProductData{}

	}

	structtopasstofrontend := productv1.checkBadIngredients()

	for _, e := range structtopasstofrontend.BadIngredients {
		fmt.Println("Bad ingredient: ", e)
	}
	return structtopasstofrontend

}

// NOTE: If data is not showing, this is why.
func (productv1 *ProductV1) checkBadIngredients() processedProductData {
	fmt.Println("Contains")

	product := processedProductData{}

	for _, prodv1 := range productv1.ProductsV1Slice {
		if prodv1.IngredientsTextEn == nil || *(prodv1.IngredientsTextEn) == " " {
			continue
		}

		product.ProductFound = true

		//populate the nutriscore display strings
		product.NutriscoreDescript = populateNutriScoreDiplay(&prodv1.NutriscoreGrade)
		product.NutriscoreGrade = prodv1.NutriscoreGrade
		product.AllIngredients = strings.ToLower(*(prodv1.IngredientsTextEn))

		for _, badIngredient := range badIngredientsSlice {
			if strings.Contains(product.AllIngredients, badIngredient) {
				product.BadIngredients = append(product.BadIngredients, badIngredient)
				//DEBUG
				// fmt.Println("bad ingredients: ", product.BadIngredients)
			}
		}
		break //break out after getting the first product with english ingredients
	}

	//assign approved field
	if (len(product.BadIngredients) == 0 && product.ProductFound) && (product.NutriscoreGrade == "a" || product.NutriscoreGrade == "b" || product.NutriscoreDescript == "unknown") {
		product.Approved = true
	}
	return product
}

func getProduct(write http.ResponseWriter, read *http.Request) {
	// Get the path parameter using mux.Vars
	vars := mux.Vars(read)
	//Retrieve the dynamic search parameter
	productName := vars["product"]

	processedData := searchByProductLogic(&productName)

	returnedString := processedData.displayProductInfo(&productName)

	fmt.Fprintln(write, returnedString)
}

func main() {
	//create a new router
	const port = "8692"
	router := mux.NewRouter()

	//register routes to the subrouter
	router.HandleFunc("/{product}", getProduct)

	//start server
	fmt.Println("Server is running on port: ", port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		//  Log the error if the server failed to start.
		fmt.Println("Failed to start server:", err)
	}
}
