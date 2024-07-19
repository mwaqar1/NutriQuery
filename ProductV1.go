package main

import "time"

//CHANGED ingredientstext and ingredientstexten from string to any

// creating go struct to store json repsonse data for v1 api
type ProductV1 struct {
	Count     int `json:"count,omitempty"`
	Page      any `json:"page,omitempty"`
	PageCount int `json:"page_count,omitempty"`
	PageSize  int `json:"page_size,omitempty"`
	//This field holds a slice of ProductsOfV1 structs. The omitempty tag means this field will be omitted from the JSON output if it is nil or empty
	ProductsV1Slice []ProductsV1Fields `json:"products,omitempty"`
	Skip            int                `json:"skip,omitempty"`
}

type ProductsV1Fields struct {
	ID                        string   `json:"_id,omitempty"`
	Keywords                  []string `json:"_keywords,omitempty"`
	AddedCountriesTags        []any    `json:"added_countries_tags,omitempty"`
	AdditivesDebugTags        []any    `json:"additives_debug_tags,omitempty"`
	AdditivesOldTags          []any    `json:"additives_old_tags,omitempty"`
	AdditivesOriginalTags     []any    `json:"additives_original_tags,omitempty"`
	AdditivesPrevOriginalTags []any    `json:"additives_prev_original_tags,omitempty"`
	AdditivesTags             []any    `json:"additives_tags,omitempty"`
	Allergens                 string   `json:"allergens,omitempty"`
	AllergensFromIngredients  string   `json:"allergens_from_ingredients,omitempty"`
	AllergensFromUser         string   `json:"allergens_from_user,omitempty"`
	AllergensHierarchy        []any    `json:"allergens_hierarchy,omitempty"`
	AllergensLc               string   `json:"allergens_lc,omitempty"`
	AllergensTags             []any    `json:"allergens_tags,omitempty"`
	AminoAcidsPrevTags        []any    `json:"amino_acids_prev_tags,omitempty"`
	AminoAcidsTags            []any    `json:"amino_acids_tags,omitempty"`
	Brands                    string   `json:"brands,omitempty"`
	BrandsTags                []string `json:"brands_tags,omitempty"`
	Categories                string   `json:"categories,omitempty"`
	CategoriesHierarchy       []string `json:"categories_hierarchy,omitempty"`
	CategoriesLc              string   `json:"categories_lc,omitempty"`
	CategoriesOld             string   `json:"categories_old,omitempty"`
	CategoriesProperties      struct {
		AgribalyseFoodCodeEn      string `json:"agribalyse_food_code:en,omitempty"`
		AgribalyseProxyFoodCodeEn string `json:"agribalyse_proxy_food_code:en,omitempty"`
		CiqualFoodCodeEn          string `json:"ciqual_food_code:en,omitempty"`
	} `json:"categories_properties,omitempty"`
	CategoriesPropertiesTags []string `json:"categories_properties_tags,omitempty"`
	CategoriesTags           []string `json:"categories_tags,omitempty"`
	CategoryProperties       struct {
		CiqualFoodNameEn string `json:"ciqual_food_name:en,omitempty"`
		CiqualFoodNameFr string `json:"ciqual_food_name:fr,omitempty"`
	} `json:"category_properties,omitempty"`
	CheckersTags            []any    `json:"checkers_tags,omitempty"`
	CiqualFoodNameTags      []string `json:"ciqual_food_name_tags,omitempty"`
	CitiesTags              []any    `json:"cities_tags,omitempty"`
	Code                    string   `json:"code,omitempty"`
	CodesTags               []string `json:"codes_tags,omitempty"`
	ComparedToCategory      string   `json:"compared_to_category,omitempty"`
	Complete                int      `json:"complete,omitempty"`
	Completeness            any      `json:"completeness,omitempty"`
	CorrectorsTags          []string `json:"correctors_tags,omitempty"`
	Countries               string   `json:"countries,omitempty"`
	CountriesHierarchy      []string `json:"countries_hierarchy,omitempty"`
	CountriesLc             string   `json:"countries_lc,omitempty"`
	CountriesTags           []string `json:"countries_tags,omitempty"`
	CreatedT                int      `json:"created_t,omitempty"`
	Creator                 string   `json:"creator,omitempty"`
	DataQualityBugsTags     []any    `json:"data_quality_bugs_tags,omitempty"`
	DataQualityErrorsTags   []any    `json:"data_quality_errors_tags,omitempty"`
	DataQualityInfoTags     []string `json:"data_quality_info_tags,omitempty"`
	DataQualityTags         []string `json:"data_quality_tags,omitempty"`
	DataQualityWarningsTags []string `json:"data_quality_warnings_tags,omitempty"`
	DataSources             string   `json:"data_sources,omitempty"`
	DataSourcesTags         []string `json:"data_sources_tags,omitempty"`
	DebugParamSortedLangs   []string `json:"debug_param_sorted_langs,omitempty"`
	EcoscoreData            struct {
		Adjustments struct {
			OriginsOfIngredients struct {
				AggregatedOrigins []struct {
					EpiScore            string `json:"epi_score,omitempty"`
					Origin              string `json:"origin,omitempty"`
					Percent             int    `json:"percent,omitempty"`
					TransportationScore any    `json:"transportation_score,omitempty"`
				} `json:"aggregated_origins,omitempty"`
				EpiScore                int      `json:"epi_score,omitempty"`
				EpiValue                int      `json:"epi_value,omitempty"`
				OriginsFromCategories   []string `json:"origins_from_categories,omitempty"`
				OriginsFromOriginsField []string `json:"origins_from_origins_field,omitempty"`
				TransportationScore     int      `json:"transportation_score,omitempty"`
				TransportationScores    struct {
					Ad    int `json:"ad,omitempty"`
					Al    int `json:"al,omitempty"`
					At    int `json:"at,omitempty"`
					Ax    int `json:"ax,omitempty"`
					Ba    int `json:"ba,omitempty"`
					Be    int `json:"be,omitempty"`
					Bg    int `json:"bg,omitempty"`
					Ch    int `json:"ch,omitempty"`
					Cy    int `json:"cy,omitempty"`
					Cz    int `json:"cz,omitempty"`
					De    int `json:"de,omitempty"`
					Dk    int `json:"dk,omitempty"`
					Dz    int `json:"dz,omitempty"`
					Ee    int `json:"ee,omitempty"`
					Eg    int `json:"eg,omitempty"`
					Es    int `json:"es,omitempty"`
					Fi    int `json:"fi,omitempty"`
					Fo    int `json:"fo,omitempty"`
					Fr    int `json:"fr,omitempty"`
					Gg    int `json:"gg,omitempty"`
					Gi    int `json:"gi,omitempty"`
					Gr    int `json:"gr,omitempty"`
					Hr    int `json:"hr,omitempty"`
					Hu    int `json:"hu,omitempty"`
					Ie    int `json:"ie,omitempty"`
					Il    int `json:"il,omitempty"`
					Im    int `json:"im,omitempty"`
					Is    int `json:"is,omitempty"`
					It    int `json:"it,omitempty"`
					Je    int `json:"je,omitempty"`
					Lb    int `json:"lb,omitempty"`
					Li    int `json:"li,omitempty"`
					Lt    int `json:"lt,omitempty"`
					Lu    int `json:"lu,omitempty"`
					Lv    int `json:"lv,omitempty"`
					Ly    int `json:"ly,omitempty"`
					Ma    int `json:"ma,omitempty"`
					Mc    int `json:"mc,omitempty"`
					Md    int `json:"md,omitempty"`
					Me    int `json:"me,omitempty"`
					Mk    int `json:"mk,omitempty"`
					Mt    int `json:"mt,omitempty"`
					Nl    int `json:"nl,omitempty"`
					No    int `json:"no,omitempty"`
					Pl    int `json:"pl,omitempty"`
					Ps    int `json:"ps,omitempty"`
					Pt    int `json:"pt,omitempty"`
					Ro    int `json:"ro,omitempty"`
					Rs    int `json:"rs,omitempty"`
					Se    int `json:"se,omitempty"`
					Si    int `json:"si,omitempty"`
					Sj    int `json:"sj,omitempty"`
					Sk    int `json:"sk,omitempty"`
					Sm    int `json:"sm,omitempty"`
					Sy    int `json:"sy,omitempty"`
					Tn    int `json:"tn,omitempty"`
					Tr    int `json:"tr,omitempty"`
					Ua    int `json:"ua,omitempty"`
					Uk    int `json:"uk,omitempty"`
					Us    int `json:"us,omitempty"`
					Va    int `json:"va,omitempty"`
					World int `json:"world,omitempty"`
					Xk    int `json:"xk,omitempty"`
				} `json:"transportation_scores,omitempty"`
				TransportationValue  int `json:"transportation_value,omitempty"`
				TransportationValues struct {
					Ad    int `json:"ad,omitempty"`
					Al    int `json:"al,omitempty"`
					At    int `json:"at,omitempty"`
					Ax    int `json:"ax,omitempty"`
					Ba    int `json:"ba,omitempty"`
					Be    int `json:"be,omitempty"`
					Bg    int `json:"bg,omitempty"`
					Ch    int `json:"ch,omitempty"`
					Cy    int `json:"cy,omitempty"`
					Cz    int `json:"cz,omitempty"`
					De    int `json:"de,omitempty"`
					Dk    int `json:"dk,omitempty"`
					Dz    int `json:"dz,omitempty"`
					Ee    int `json:"ee,omitempty"`
					Eg    int `json:"eg,omitempty"`
					Es    int `json:"es,omitempty"`
					Fi    int `json:"fi,omitempty"`
					Fo    int `json:"fo,omitempty"`
					Fr    int `json:"fr,omitempty"`
					Gg    int `json:"gg,omitempty"`
					Gi    int `json:"gi,omitempty"`
					Gr    int `json:"gr,omitempty"`
					Hr    int `json:"hr,omitempty"`
					Hu    int `json:"hu,omitempty"`
					Ie    int `json:"ie,omitempty"`
					Il    int `json:"il,omitempty"`
					Im    int `json:"im,omitempty"`
					Is    int `json:"is,omitempty"`
					It    int `json:"it,omitempty"`
					Je    int `json:"je,omitempty"`
					Lb    int `json:"lb,omitempty"`
					Li    int `json:"li,omitempty"`
					Lt    int `json:"lt,omitempty"`
					Lu    int `json:"lu,omitempty"`
					Lv    int `json:"lv,omitempty"`
					Ly    int `json:"ly,omitempty"`
					Ma    int `json:"ma,omitempty"`
					Mc    int `json:"mc,omitempty"`
					Md    int `json:"md,omitempty"`
					Me    int `json:"me,omitempty"`
					Mk    int `json:"mk,omitempty"`
					Mt    int `json:"mt,omitempty"`
					Nl    int `json:"nl,omitempty"`
					No    int `json:"no,omitempty"`
					Pl    int `json:"pl,omitempty"`
					Ps    int `json:"ps,omitempty"`
					Pt    int `json:"pt,omitempty"`
					Ro    int `json:"ro,omitempty"`
					Rs    int `json:"rs,omitempty"`
					Se    int `json:"se,omitempty"`
					Si    int `json:"si,omitempty"`
					Sj    int `json:"sj,omitempty"`
					Sk    int `json:"sk,omitempty"`
					Sm    int `json:"sm,omitempty"`
					Sy    int `json:"sy,omitempty"`
					Tn    int `json:"tn,omitempty"`
					Tr    int `json:"tr,omitempty"`
					Ua    int `json:"ua,omitempty"`
					Uk    int `json:"uk,omitempty"`
					Us    int `json:"us,omitempty"`
					Va    int `json:"va,omitempty"`
					World int `json:"world,omitempty"`
					Xk    int `json:"xk,omitempty"`
				} `json:"transportation_values,omitempty"`
				Value  int `json:"value,omitempty"`
				Values struct {
					Ad    int `json:"ad,omitempty"`
					Al    int `json:"al,omitempty"`
					At    int `json:"at,omitempty"`
					Ax    int `json:"ax,omitempty"`
					Ba    int `json:"ba,omitempty"`
					Be    int `json:"be,omitempty"`
					Bg    int `json:"bg,omitempty"`
					Ch    int `json:"ch,omitempty"`
					Cy    int `json:"cy,omitempty"`
					Cz    int `json:"cz,omitempty"`
					De    int `json:"de,omitempty"`
					Dk    int `json:"dk,omitempty"`
					Dz    int `json:"dz,omitempty"`
					Ee    int `json:"ee,omitempty"`
					Eg    int `json:"eg,omitempty"`
					Es    int `json:"es,omitempty"`
					Fi    int `json:"fi,omitempty"`
					Fo    int `json:"fo,omitempty"`
					Fr    int `json:"fr,omitempty"`
					Gg    int `json:"gg,omitempty"`
					Gi    int `json:"gi,omitempty"`
					Gr    int `json:"gr,omitempty"`
					Hr    int `json:"hr,omitempty"`
					Hu    int `json:"hu,omitempty"`
					Ie    int `json:"ie,omitempty"`
					Il    int `json:"il,omitempty"`
					Im    int `json:"im,omitempty"`
					Is    int `json:"is,omitempty"`
					It    int `json:"it,omitempty"`
					Je    int `json:"je,omitempty"`
					Lb    int `json:"lb,omitempty"`
					Li    int `json:"li,omitempty"`
					Lt    int `json:"lt,omitempty"`
					Lu    int `json:"lu,omitempty"`
					Lv    int `json:"lv,omitempty"`
					Ly    int `json:"ly,omitempty"`
					Ma    int `json:"ma,omitempty"`
					Mc    int `json:"mc,omitempty"`
					Md    int `json:"md,omitempty"`
					Me    int `json:"me,omitempty"`
					Mk    int `json:"mk,omitempty"`
					Mt    int `json:"mt,omitempty"`
					Nl    int `json:"nl,omitempty"`
					No    int `json:"no,omitempty"`
					Pl    int `json:"pl,omitempty"`
					Ps    int `json:"ps,omitempty"`
					Pt    int `json:"pt,omitempty"`
					Ro    int `json:"ro,omitempty"`
					Rs    int `json:"rs,omitempty"`
					Se    int `json:"se,omitempty"`
					Si    int `json:"si,omitempty"`
					Sj    int `json:"sj,omitempty"`
					Sk    int `json:"sk,omitempty"`
					Sm    int `json:"sm,omitempty"`
					Sy    int `json:"sy,omitempty"`
					Tn    int `json:"tn,omitempty"`
					Tr    int `json:"tr,omitempty"`
					Ua    int `json:"ua,omitempty"`
					Uk    int `json:"uk,omitempty"`
					Us    int `json:"us,omitempty"`
					Va    int `json:"va,omitempty"`
					World int `json:"world,omitempty"`
					Xk    int `json:"xk,omitempty"`
				} `json:"values,omitempty"`
				Warning string `json:"warning,omitempty"`
			} `json:"origins_of_ingredients,omitempty"`
			Packaging struct {
				NonRecyclableAndNonBiodegradableMaterials int    `json:"non_recyclable_and_non_biodegradable_materials,omitempty"`
				Value                                     int    `json:"value,omitempty"`
				Warning                                   string `json:"warning,omitempty"`
			} `json:"packaging,omitempty"`
			ProductionSystem struct {
				Labels  []any  `json:"labels,omitempty"`
				Value   int    `json:"value,omitempty"`
				Warning string `json:"warning,omitempty"`
			} `json:"production_system,omitempty"`
			ThreatenedSpecies struct {
				Warning string `json:"warning,omitempty"`
			} `json:"threatened_species,omitempty"`
		} `json:"adjustments,omitempty"`
		Agribalyse struct {
			AgribalyseFoodCode      string `json:"agribalyse_food_code,omitempty"`
			AgribalyseProxyFoodCode string `json:"agribalyse_proxy_food_code,omitempty"`
			Co2Agriculture          string `json:"co2_agriculture,omitempty"`
			Co2Consumption          int    `json:"co2_consumption,omitempty"`
			Co2Distribution         string `json:"co2_distribution,omitempty"`
			Co2Packaging            string `json:"co2_packaging,omitempty"`
			Co2Processing           string `json:"co2_processing,omitempty"`
			Co2Total                string `json:"co2_total,omitempty"`
			Co2Transportation       string `json:"co2_transportation,omitempty"`
			Code                    string `json:"code,omitempty"`
			Dqr                     string `json:"dqr,omitempty"`
			EfAgriculture           string `json:"ef_agriculture,omitempty"`
			EfConsumption           int    `json:"ef_consumption,omitempty"`
			EfDistribution          string `json:"ef_distribution,omitempty"`
			EfPackaging             string `json:"ef_packaging,omitempty"`
			EfProcessing            string `json:"ef_processing,omitempty"`
			EfTotal                 string `json:"ef_total,omitempty"`
			EfTransportation        string `json:"ef_transportation,omitempty"`
			IsBeverage              int    `json:"is_beverage,omitempty"`
			NameEn                  string `json:"name_en,omitempty"`
			NameFr                  string `json:"name_fr,omitempty"`
			Score                   int    `json:"score,omitempty"`
			Version                 string `json:"version,omitempty"`
		} `json:"agribalyse,omitempty"`
		Grade  string `json:"grade,omitempty"`
		Grades struct {
			Ad    string `json:"ad,omitempty"`
			Al    string `json:"al,omitempty"`
			At    string `json:"at,omitempty"`
			Ax    string `json:"ax,omitempty"`
			Ba    string `json:"ba,omitempty"`
			Be    string `json:"be,omitempty"`
			Bg    string `json:"bg,omitempty"`
			Ch    string `json:"ch,omitempty"`
			Cy    string `json:"cy,omitempty"`
			Cz    string `json:"cz,omitempty"`
			De    string `json:"de,omitempty"`
			Dk    string `json:"dk,omitempty"`
			Dz    string `json:"dz,omitempty"`
			Ee    string `json:"ee,omitempty"`
			Eg    string `json:"eg,omitempty"`
			Es    string `json:"es,omitempty"`
			Fi    string `json:"fi,omitempty"`
			Fo    string `json:"fo,omitempty"`
			Fr    string `json:"fr,omitempty"`
			Gg    string `json:"gg,omitempty"`
			Gi    string `json:"gi,omitempty"`
			Gr    string `json:"gr,omitempty"`
			Hr    string `json:"hr,omitempty"`
			Hu    string `json:"hu,omitempty"`
			Ie    string `json:"ie,omitempty"`
			Il    string `json:"il,omitempty"`
			Im    string `json:"im,omitempty"`
			Is    string `json:"is,omitempty"`
			It    string `json:"it,omitempty"`
			Je    string `json:"je,omitempty"`
			Lb    string `json:"lb,omitempty"`
			Li    string `json:"li,omitempty"`
			Lt    string `json:"lt,omitempty"`
			Lu    string `json:"lu,omitempty"`
			Lv    string `json:"lv,omitempty"`
			Ly    string `json:"ly,omitempty"`
			Ma    string `json:"ma,omitempty"`
			Mc    string `json:"mc,omitempty"`
			Md    string `json:"md,omitempty"`
			Me    string `json:"me,omitempty"`
			Mk    string `json:"mk,omitempty"`
			Mt    string `json:"mt,omitempty"`
			Nl    string `json:"nl,omitempty"`
			No    string `json:"no,omitempty"`
			Pl    string `json:"pl,omitempty"`
			Ps    string `json:"ps,omitempty"`
			Pt    string `json:"pt,omitempty"`
			Ro    string `json:"ro,omitempty"`
			Rs    string `json:"rs,omitempty"`
			Se    string `json:"se,omitempty"`
			Si    string `json:"si,omitempty"`
			Sj    string `json:"sj,omitempty"`
			Sk    string `json:"sk,omitempty"`
			Sm    string `json:"sm,omitempty"`
			Sy    string `json:"sy,omitempty"`
			Tn    string `json:"tn,omitempty"`
			Tr    string `json:"tr,omitempty"`
			Ua    string `json:"ua,omitempty"`
			Uk    string `json:"uk,omitempty"`
			Us    string `json:"us,omitempty"`
			Va    string `json:"va,omitempty"`
			World string `json:"world,omitempty"`
			Xk    string `json:"xk,omitempty"`
		} `json:"grades,omitempty"`
		Missing struct {
			Ingredients int `json:"ingredients,omitempty"`
			Labels      int `json:"labels,omitempty"`
			Origins     int `json:"origins,omitempty"`
			Packagings  int `json:"packagings,omitempty"`
		} `json:"missing,omitempty"`
		MissingDataWarning int `json:"missing_data_warning,omitempty"`
		MissingKeyData     int `json:"missing_key_data,omitempty"`
		PreviousData       struct {
			Agribalyse struct {
				Warning string `json:"warning,omitempty"`
			} `json:"agribalyse,omitempty"`
			Grade any `json:"grade,omitempty"`
			Score any `json:"score,omitempty"`
		} `json:"previous_data,omitempty"`
		Score  int `json:"score,omitempty"`
		Scores struct {
			Ad    int `json:"ad,omitempty"`
			Al    int `json:"al,omitempty"`
			At    int `json:"at,omitempty"`
			Ax    int `json:"ax,omitempty"`
			Ba    int `json:"ba,omitempty"`
			Be    int `json:"be,omitempty"`
			Bg    int `json:"bg,omitempty"`
			Ch    int `json:"ch,omitempty"`
			Cy    int `json:"cy,omitempty"`
			Cz    int `json:"cz,omitempty"`
			De    int `json:"de,omitempty"`
			Dk    int `json:"dk,omitempty"`
			Dz    int `json:"dz,omitempty"`
			Ee    int `json:"ee,omitempty"`
			Eg    int `json:"eg,omitempty"`
			Es    int `json:"es,omitempty"`
			Fi    int `json:"fi,omitempty"`
			Fo    int `json:"fo,omitempty"`
			Fr    int `json:"fr,omitempty"`
			Gg    int `json:"gg,omitempty"`
			Gi    int `json:"gi,omitempty"`
			Gr    int `json:"gr,omitempty"`
			Hr    int `json:"hr,omitempty"`
			Hu    int `json:"hu,omitempty"`
			Ie    int `json:"ie,omitempty"`
			Il    int `json:"il,omitempty"`
			Im    int `json:"im,omitempty"`
			Is    int `json:"is,omitempty"`
			It    int `json:"it,omitempty"`
			Je    int `json:"je,omitempty"`
			Lb    int `json:"lb,omitempty"`
			Li    int `json:"li,omitempty"`
			Lt    int `json:"lt,omitempty"`
			Lu    int `json:"lu,omitempty"`
			Lv    int `json:"lv,omitempty"`
			Ly    int `json:"ly,omitempty"`
			Ma    int `json:"ma,omitempty"`
			Mc    int `json:"mc,omitempty"`
			Md    int `json:"md,omitempty"`
			Me    int `json:"me,omitempty"`
			Mk    int `json:"mk,omitempty"`
			Mt    int `json:"mt,omitempty"`
			Nl    int `json:"nl,omitempty"`
			No    int `json:"no,omitempty"`
			Pl    int `json:"pl,omitempty"`
			Ps    int `json:"ps,omitempty"`
			Pt    int `json:"pt,omitempty"`
			Ro    int `json:"ro,omitempty"`
			Rs    int `json:"rs,omitempty"`
			Se    int `json:"se,omitempty"`
			Si    int `json:"si,omitempty"`
			Sj    int `json:"sj,omitempty"`
			Sk    int `json:"sk,omitempty"`
			Sm    int `json:"sm,omitempty"`
			Sy    int `json:"sy,omitempty"`
			Tn    int `json:"tn,omitempty"`
			Tr    int `json:"tr,omitempty"`
			Ua    int `json:"ua,omitempty"`
			Uk    int `json:"uk,omitempty"`
			Us    int `json:"us,omitempty"`
			Va    int `json:"va,omitempty"`
			World int `json:"world,omitempty"`
			Xk    int `json:"xk,omitempty"`
		} `json:"scores,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"ecoscore_data,omitempty"`
	EcoscoreGrade            string   `json:"ecoscore_grade,omitempty"`
	EcoscoreScore            int      `json:"ecoscore_score,omitempty"`
	EcoscoreTags             []string `json:"ecoscore_tags,omitempty"`
	EditorsTags              []string `json:"editors_tags,omitempty"`
	EmbCodes                 string   `json:"emb_codes,omitempty"`
	EmbCodesTags             []any    `json:"emb_codes_tags,omitempty"`
	EntryDatesTags           []string `json:"entry_dates_tags,omitempty"`
	ExpirationDate           string   `json:"expiration_date,omitempty"`
	FoodGroups               string   `json:"food_groups,omitempty"`
	FoodGroupsTags           []string `json:"food_groups_tags,omitempty"`
	GenericName              string   `json:"generic_name,omitempty"`
	GenericNameFr            string   `json:"generic_name_fr,omitempty"`
	ID0                      string   `json:"id,omitempty"`
	ImageFrontSmallURL       string   `json:"image_front_small_url,omitempty"`
	ImageFrontThumbURL       string   `json:"image_front_thumb_url,omitempty"`
	ImageFrontURL            string   `json:"image_front_url,omitempty"`
	ImageIngredientsSmallURL string   `json:"image_ingredients_small_url,omitempty"`
	ImageIngredientsThumbURL string   `json:"image_ingredients_thumb_url,omitempty"`
	ImageIngredientsURL      string   `json:"image_ingredients_url,omitempty"`
	ImageNutritionSmallURL   string   `json:"image_nutrition_small_url,omitempty"`
	ImageNutritionThumbURL   string   `json:"image_nutrition_thumb_url,omitempty"`
	ImageNutritionURL        string   `json:"image_nutrition_url,omitempty"`
	ImageSmallURL            string   `json:"image_small_url,omitempty"`
	ImageThumbURL            string   `json:"image_thumb_url,omitempty"`
	ImageURL                 string   `json:"image_url,omitempty"`
	Images                   struct {
		Num1 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"1,omitempty"`
		Num2 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"2,omitempty"`
		Num3 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"3,omitempty"`
		Num4 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"4,omitempty"`
		Num5 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"5,omitempty"`
		Num6 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"6,omitempty"`
		FrontFr struct {
			Angle                int    `json:"angle,omitempty"`
			CoordinatesImageSize string `json:"coordinates_image_size,omitempty"`
			Geometry             string `json:"geometry,omitempty"`
			Imgid                string `json:"imgid,omitempty"`
			Normalize            any    `json:"normalize,omitempty"`
			Rev                  string `json:"rev,omitempty"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num200 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"200,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			WhiteMagic any    `json:"white_magic,omitempty"`
			X1         string `json:"x1,omitempty"`
			X2         string `json:"x2,omitempty"`
			Y1         string `json:"y1,omitempty"`
			Y2         string `json:"y2,omitempty"`
		} `json:"front_fr,omitempty"`
		IngredientsDe struct {
			Angle                int    `json:"angle,omitempty"`
			CoordinatesImageSize string `json:"coordinates_image_size,omitempty"`
			Geometry             string `json:"geometry,omitempty"`
			Imgid                string `json:"imgid,omitempty"`
			Normalize            any    `json:"normalize,omitempty"`
			Rev                  string `json:"rev,omitempty"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num200 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"200,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			WhiteMagic any    `json:"white_magic,omitempty"`
			X1         string `json:"x1,omitempty"`
			X2         string `json:"x2,omitempty"`
			Y1         string `json:"y1,omitempty"`
			Y2         string `json:"y2,omitempty"`
		} `json:"ingredients_de,omitempty"`
		IngredientsFr struct {
			Angle       any    `json:"angle,omitempty"`
			Geometry    string `json:"geometry,omitempty"`
			Imgid       string `json:"imgid,omitempty"`
			Normalize   string `json:"normalize,omitempty"`
			Ocr         int    `json:"ocr,omitempty"`
			Orientation string `json:"orientation,omitempty"`
			Rev         string `json:"rev,omitempty"`
			Sizes       struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num200 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"200,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			WhiteMagic string `json:"white_magic,omitempty"`
			X1         any    `json:"x1,omitempty"`
			X2         any    `json:"x2,omitempty"`
			Y1         any    `json:"y1,omitempty"`
			Y2         any    `json:"y2,omitempty"`
		} `json:"ingredients_fr,omitempty"`
		NutritionFr struct {
			Angle                int    `json:"angle,omitempty"`
			CoordinatesImageSize string `json:"coordinates_image_size,omitempty"`
			Geometry             string `json:"geometry,omitempty"`
			Imgid                string `json:"imgid,omitempty"`
			Normalize            any    `json:"normalize,omitempty"`
			Rev                  string `json:"rev,omitempty"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num200 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"200,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			WhiteMagic any    `json:"white_magic,omitempty"`
			X1         string `json:"x1,omitempty"`
			X2         string `json:"x2,omitempty"`
			Y1         string `json:"y1,omitempty"`
			Y2         string `json:"y2,omitempty"`
		} `json:"nutrition_fr,omitempty"`
	} `json:"images,omitempty"`
	InformersTags                       []string `json:"informers_tags,omitempty"`
	IngredientsDebug                    []any    `json:"ingredients_debug,omitempty"`
	IngredientsFromPalmOilTags          []any    `json:"ingredients_from_palm_oil_tags,omitempty"`
	IngredientsIdsDebug                 []any    `json:"ingredients_ids_debug,omitempty"`
	IngredientsText                     *string  `json:"ingredients_text,omitempty"`
	IngredientsTextFr                   string   `json:"ingredients_text_fr,omitempty"`
	IngredientsTextWithAllergens        string   `json:"ingredients_text_with_allergens,omitempty"`
	IngredientsTextWithAllergensFr      string   `json:"ingredients_text_with_allergens_fr,omitempty"`
	IngredientsThatMayBeFromPalmOilTags []any    `json:"ingredients_that_may_be_from_palm_oil_tags,omitempty"`
	InterfaceVersionCreated             string   `json:"interface_version_created,omitempty"`
	InterfaceVersionModified            string   `json:"interface_version_modified,omitempty"`
	Labels                              string   `json:"labels,omitempty"`
	LabelsHierarchy                     []string `json:"labels_hierarchy,omitempty"`
	LabelsLc                            string   `json:"labels_lc,omitempty"`
	LabelsOld                           string   `json:"labels_old,omitempty"`
	LabelsTags                          []string `json:"labels_tags,omitempty"`
	Lang                                string   `json:"lang,omitempty"`
	Languages                           struct {
		EnFrench int `json:"en:french,omitempty"`
		EnGerman int `json:"en:german,omitempty"`
	} `json:"languages,omitempty"`
	LanguagesCodes struct {
		De int `json:"de,omitempty"`
		Fr int `json:"fr,omitempty"`
	} `json:"languages_codes,omitempty"`
	LanguagesHierarchy      []string `json:"languages_hierarchy,omitempty"`
	LanguagesTags           []string `json:"languages_tags,omitempty"`
	LastEditDatesTags       []string `json:"last_edit_dates_tags,omitempty"`
	LastEditor              string   `json:"last_editor,omitempty"`
	LastImageDatesTags      []string `json:"last_image_dates_tags,omitempty"`
	LastImageT              int      `json:"last_image_t,omitempty"`
	LastModifiedBy          string   `json:"last_modified_by,omitempty"`
	LastModifiedT           int      `json:"last_modified_t,omitempty"`
	LastUpdatedT            int      `json:"last_updated_t,omitempty"`
	Lc                      string   `json:"lc,omitempty"`
	Link                    string   `json:"link,omitempty"`
	MainCountriesTags       []any    `json:"main_countries_tags,omitempty"`
	ManufacturingPlaces     string   `json:"manufacturing_places,omitempty"`
	ManufacturingPlacesTags []any    `json:"manufacturing_places_tags,omitempty"`
	MaxImgid                any      `json:"max_imgid,omitempty"`
	MineralsPrevTags        []any    `json:"minerals_prev_tags,omitempty"`
	MineralsTags            []any    `json:"minerals_tags,omitempty"`
	MiscTags                []string `json:"misc_tags,omitempty"`
	NoNutritionData         string   `json:"no_nutrition_data,omitempty"`
	NovaGroup               int      `json:"nova_group,omitempty"`
	NovaGroupDebug          string   `json:"nova_group_debug,omitempty"`
	NovaGroups              string   `json:"nova_groups,omitempty"`
	NovaGroupsMarkers       struct {
		Num2 [][]string `json:"2,omitempty"`
	} `json:"nova_groups_markers,omitempty"`
	NovaGroupsTags      []string `json:"nova_groups_tags,omitempty"`
	NucleotidesPrevTags []any    `json:"nucleotides_prev_tags,omitempty"`
	NucleotidesTags     []any    `json:"nucleotides_tags,omitempty"`
	NutrientLevels      struct {
		Fat          string `json:"fat,omitempty"`
		Salt         string `json:"salt,omitempty"`
		SaturatedFat string `json:"saturated-fat,omitempty"`
		Sugars       string `json:"sugars,omitempty"`
	} `json:"nutrient_levels,omitempty"`
	NutrientLevelsTags []string `json:"nutrient_levels_tags,omitempty"`
	Nutriments         struct {
		Carbohydrates           int    `json:"carbohydrates,omitempty"`
		Carbohydrates100G       int    `json:"carbohydrates_100g,omitempty"`
		CarbohydratesUnit       string `json:"carbohydrates_unit,omitempty"`
		CarbohydratesValue      int    `json:"carbohydrates_value,omitempty"`
		Energy                  int    `json:"energy,omitempty"`
		EnergyKcal              int    `json:"energy-kcal,omitempty"`
		EnergyKcal100G          int    `json:"energy-kcal_100g,omitempty"`
		EnergyKcalUnit          string `json:"energy-kcal_unit,omitempty"`
		EnergyKcalValue         int    `json:"energy-kcal_value,omitempty"`
		EnergyKcalValueComputed int    `json:"energy-kcal_value_computed,omitempty"`
		Energy100G              int    `json:"energy_100g,omitempty"`
		EnergyUnit              string `json:"energy_unit,omitempty"`
		EnergyValue             int    `json:"energy_value,omitempty"`
		Fat                     int    `json:"fat,omitempty"`
		Fat100G                 int    `json:"fat_100g,omitempty"`
		FatUnit                 string `json:"fat_unit,omitempty"`
		FatValue                int    `json:"fat_value,omitempty"`
		NovaGroup               int    `json:"nova-group,omitempty"`
		NovaGroup100G           int    `json:"nova-group_100g,omitempty"`
		NovaGroupServing        int    `json:"nova-group_serving,omitempty"`
		NutritionScoreFr        int    `json:"nutrition-score-fr,omitempty"`
		NutritionScoreFr100G    int    `json:"nutrition-score-fr_100g,omitempty"`
		Proteins                int    `json:"proteins,omitempty"`
		Proteins100G            int    `json:"proteins_100g,omitempty"`
		ProteinsUnit            string `json:"proteins_unit,omitempty"`
		ProteinsValue           int    `json:"proteins_value,omitempty"`
		Salt                    int    `json:"salt,omitempty"`
		Salt100G                int    `json:"salt_100g,omitempty"`
		SaltUnit                string `json:"salt_unit,omitempty"`
		SaltValue               int    `json:"salt_value,omitempty"`
		SaturatedFat            int    `json:"saturated-fat,omitempty"`
		SaturatedFat100G        int    `json:"saturated-fat_100g,omitempty"`
		SaturatedFatUnit        string `json:"saturated-fat_unit,omitempty"`
		SaturatedFatValue       int    `json:"saturated-fat_value,omitempty"`
		Sodium                  int    `json:"sodium,omitempty"`
		Sodium100G              int    `json:"sodium_100g,omitempty"`
		SodiumUnit              string `json:"sodium_unit,omitempty"`
		SodiumValue             any    `json:"sodium_value,omitempty"`
		Sugars                  int    `json:"sugars,omitempty"`
		Sugars100G              int    `json:"sugars_100g,omitempty"`
		SugarsUnit              string `json:"sugars_unit,omitempty"`
		SugarsValue             int    `json:"sugars_value,omitempty"`
	} `json:"nutriments,omitempty"`
	Nutriscore struct {
		Num2021 struct {
			CategoryAvailable int `json:"category_available,omitempty"`
			Data              struct {
				Energy                                         any `json:"energy,omitempty"`
				EnergyPoints                                   int `json:"energy_points,omitempty"`
				EnergyValue                                    any `json:"energy_value,omitempty"`
				Fiber                                          any `json:"fiber,omitempty"`
				FiberPoints                                    int `json:"fiber_points,omitempty"`
				FiberValue                                     any `json:"fiber_value,omitempty"`
				FruitsVegetablesNutsColzaWalnutOliveOils       any `json:"fruits_vegetables_nuts_colza_walnut_olive_oils,omitempty"`
				FruitsVegetablesNutsColzaWalnutOliveOilsPoints any `json:"fruits_vegetables_nuts_colza_walnut_olive_oils_points,omitempty"`
				FruitsVegetablesNutsColzaWalnutOliveOilsValue  any `json:"fruits_vegetables_nuts_colza_walnut_olive_oils_value,omitempty"`
				IsBeverage                                     int `json:"is_beverage,omitempty"`
				IsCheese                                       any `json:"is_cheese,omitempty"`
				IsFat                                          int `json:"is_fat,omitempty"`
				IsWater                                        any `json:"is_water,omitempty"`
				NegativePoints                                 int `json:"negative_points,omitempty"`
				PositivePoints                                 int `json:"positive_points,omitempty"`
				Proteins                                       any `json:"proteins,omitempty"`
				ProteinsPoints                                 int `json:"proteins_points,omitempty"`
				ProteinsValue                                  any `json:"proteins_value,omitempty"`
				SaturatedFat                                   any `json:"saturated_fat,omitempty"`
				SaturatedFatPoints                             int `json:"saturated_fat_points,omitempty"`
				SaturatedFatValue                              any `json:"saturated_fat_value,omitempty"`
				Sodium                                         any `json:"sodium,omitempty"`
				SodiumPoints                                   int `json:"sodium_points,omitempty"`
				SodiumValue                                    any `json:"sodium_value,omitempty"`
				Sugars                                         any `json:"sugars,omitempty"`
				SugarsPoints                                   int `json:"sugars_points,omitempty"`
				SugarsValue                                    any `json:"sugars_value,omitempty"`
			} `json:"data,omitempty"`
			Grade                string `json:"grade,omitempty"`
			NutrientsAvailable   int    `json:"nutrients_available,omitempty"`
			NutriscoreApplicable int    `json:"nutriscore_applicable,omitempty"`
			NutriscoreComputed   int    `json:"nutriscore_computed,omitempty"`
			Score                int    `json:"score,omitempty"`
		} `json:"2021,omitempty"`
		Num2023 struct {
			CategoryAvailable int `json:"category_available,omitempty"`
			Data              struct {
				Components struct {
					Negative []struct {
						ID        string `json:"id,omitempty"`
						Points    int    `json:"points,omitempty"`
						PointsMax int    `json:"points_max,omitempty"`
						Unit      string `json:"unit,omitempty"`
						Value     any    `json:"value,omitempty"`
					} `json:"negative,omitempty"`
					Positive []struct {
						ID        string `json:"id,omitempty"`
						Points    int    `json:"points,omitempty"`
						PointsMax int    `json:"points_max,omitempty"`
						Unit      string `json:"unit,omitempty"`
						Value     any    `json:"value,omitempty"`
					} `json:"positive,omitempty"`
				} `json:"components,omitempty"`
				CountProteins       int      `json:"count_proteins,omitempty"`
				CountProteinsReason string   `json:"count_proteins_reason,omitempty"`
				IsBeverage          int      `json:"is_beverage,omitempty"`
				IsCheese            any      `json:"is_cheese,omitempty"`
				IsFatOilNutsSeeds   int      `json:"is_fat_oil_nuts_seeds,omitempty"`
				IsRedMeatProduct    int      `json:"is_red_meat_product,omitempty"`
				IsWater             any      `json:"is_water,omitempty"`
				NegativePoints      int      `json:"negative_points,omitempty"`
				NegativePointsMax   int      `json:"negative_points_max,omitempty"`
				PositiveNutrients   []string `json:"positive_nutrients,omitempty"`
				PositivePoints      int      `json:"positive_points,omitempty"`
				PositivePointsMax   int      `json:"positive_points_max,omitempty"`
			} `json:"data,omitempty"`
			Grade                string `json:"grade,omitempty"`
			NutrientsAvailable   int    `json:"nutrients_available,omitempty"`
			NutriscoreApplicable int    `json:"nutriscore_applicable,omitempty"`
			NutriscoreComputed   int    `json:"nutriscore_computed,omitempty"`
			Score                int    `json:"score,omitempty"`
		} `json:"2023,omitempty"`
	} `json:"nutriscore,omitempty"`
	Nutriscore2021Tags []string `json:"nutriscore_2021_tags,omitempty"`
	Nutriscore2023Tags []string `json:"nutriscore_2023_tags,omitempty"`
	NutriscoreData     struct {
		Energy                                         any    `json:"energy,omitempty"`
		EnergyPoints                                   int    `json:"energy_points,omitempty"`
		EnergyValue                                    any    `json:"energy_value,omitempty"`
		Fiber                                          any    `json:"fiber,omitempty"`
		FiberPoints                                    int    `json:"fiber_points,omitempty"`
		FiberValue                                     any    `json:"fiber_value,omitempty"`
		FruitsVegetablesNutsColzaWalnutOliveOils       any    `json:"fruits_vegetables_nuts_colza_walnut_olive_oils,omitempty"`
		FruitsVegetablesNutsColzaWalnutOliveOilsPoints any    `json:"fruits_vegetables_nuts_colza_walnut_olive_oils_points,omitempty"`
		FruitsVegetablesNutsColzaWalnutOliveOilsValue  any    `json:"fruits_vegetables_nuts_colza_walnut_olive_oils_value,omitempty"`
		Grade                                          string `json:"grade,omitempty"`
		IsBeverage                                     int    `json:"is_beverage,omitempty"`
		IsCheese                                       any    `json:"is_cheese,omitempty"`
		IsFat                                          int    `json:"is_fat,omitempty"`
		IsWater                                        any    `json:"is_water,omitempty"`
		NegativePoints                                 int    `json:"negative_points,omitempty"`
		PositivePoints                                 int    `json:"positive_points,omitempty"`
		Proteins                                       any    `json:"proteins,omitempty"`
		ProteinsPoints                                 int    `json:"proteins_points,omitempty"`
		ProteinsValue                                  any    `json:"proteins_value,omitempty"`
		SaturatedFat                                   any    `json:"saturated_fat,omitempty"`
		SaturatedFatPoints                             int    `json:"saturated_fat_points,omitempty"`
		SaturatedFatValue                              any    `json:"saturated_fat_value,omitempty"`
		Score                                          int    `json:"score,omitempty"`
		Sodium                                         any    `json:"sodium,omitempty"`
		SodiumPoints                                   int    `json:"sodium_points,omitempty"`
		SodiumValue                                    any    `json:"sodium_value,omitempty"`
		Sugars                                         any    `json:"sugars,omitempty"`
		SugarsPoints                                   int    `json:"sugars_points,omitempty"`
		SugarsValue                                    any    `json:"sugars_value,omitempty"`
	} `json:"nutriscore_data,omitempty"`
	NutriscoreGrade                             string   `json:"nutriscore_grade,omitempty"`
	NutriscoreScore                             int      `json:"nutriscore_score,omitempty"`
	NutriscoreScoreOpposite                     int      `json:"nutriscore_score_opposite,omitempty"`
	NutriscoreTags                              []string `json:"nutriscore_tags,omitempty"`
	NutriscoreVersion                           string   `json:"nutriscore_version,omitempty"`
	NutritionData                               string   `json:"nutrition_data,omitempty"`
	NutritionDataPer                            string   `json:"nutrition_data_per,omitempty"`
	NutritionDataPrepared                       string   `json:"nutrition_data_prepared,omitempty"`
	NutritionDataPreparedPer                    string   `json:"nutrition_data_prepared_per,omitempty"`
	NutritionGradeFr                            string   `json:"nutrition_grade_fr,omitempty"`
	NutritionGrades                             string   `json:"nutrition_grades,omitempty"`
	NutritionGradesTags                         []string `json:"nutrition_grades_tags,omitempty"`
	NutritionScoreBeverage                      int      `json:"nutrition_score_beverage,omitempty"`
	NutritionScoreDebug                         string   `json:"nutrition_score_debug,omitempty"`
	NutritionScoreWarningNoFiber                int      `json:"nutrition_score_warning_no_fiber,omitempty"`
	NutritionScoreWarningNoFruitsVegetablesNuts int      `json:"nutrition_score_warning_no_fruits_vegetables_nuts,omitempty"`
	Origin                                      string   `json:"origin,omitempty"`
	OriginFr                                    string   `json:"origin_fr,omitempty"`
	Origins                                     string   `json:"origins,omitempty"`
	OriginsHierarchy                            []any    `json:"origins_hierarchy,omitempty"`
	OriginsLc                                   string   `json:"origins_lc,omitempty"`
	OriginsTags                                 []any    `json:"origins_tags,omitempty"`
	OtherNutritionalSubstancesTags              []any    `json:"other_nutritional_substances_tags,omitempty"`
	PackagingMaterialsTags                      []any    `json:"packaging_materials_tags,omitempty"`
	PackagingRecyclingTags                      []any    `json:"packaging_recycling_tags,omitempty"`
	PackagingShapesTags                         []any    `json:"packaging_shapes_tags,omitempty"`
	PackagingText                               string   `json:"packaging_text,omitempty"`
	PackagingTextFr                             string   `json:"packaging_text_fr,omitempty"`
	Packagings                                  []any    `json:"packagings,omitempty"`
	PackagingsComplete                          int      `json:"packagings_complete,omitempty"`
	PackagingsMaterials                         struct {
	} `json:"packagings_materials,omitempty"`
	PhotographersTags    []string `json:"photographers_tags,omitempty"`
	PnnsGroups1          string   `json:"pnns_groups_1,omitempty"`
	PnnsGroups1Tags      []string `json:"pnns_groups_1_tags,omitempty"`
	PnnsGroups2          string   `json:"pnns_groups_2,omitempty"`
	PnnsGroups2Tags      []string `json:"pnns_groups_2_tags,omitempty"`
	PopularityKey        any      `json:"popularity_key,omitempty"`
	PopularityTags       []string `json:"popularity_tags,omitempty"`
	ProductName          string   `json:"product_name,omitempty"`
	ProductNameFr        string   `json:"product_name_fr,omitempty"`
	ProductQuantity      any      `json:"product_quantity,omitempty"`
	PurchasePlaces       string   `json:"purchase_places,omitempty"`
	PurchasePlacesTags   []any    `json:"purchase_places_tags,omitempty"`
	Quantity             string   `json:"quantity,omitempty"`
	RemovedCountriesTags []any    `json:"removed_countries_tags,omitempty"`
	Rev                  any      `json:"rev,omitempty"`
	ScansN               int      `json:"scans_n,omitempty"`
	SelectedImages       struct {
		Front struct {
			Display struct {
				Fr string `json:"fr,omitempty"`
			} `json:"display,omitempty"`
			Small struct {
				Fr string `json:"fr,omitempty"`
			} `json:"small,omitempty"`
			Thumb struct {
				Fr string `json:"fr,omitempty"`
			} `json:"thumb,omitempty"`
		} `json:"front,omitempty"`
		Ingredients struct {
			Display struct {
				De string `json:"de,omitempty"`
				Fr string `json:"fr,omitempty"`
			} `json:"display,omitempty"`
			Small struct {
				De string `json:"de,omitempty"`
				Fr string `json:"fr,omitempty"`
			} `json:"small,omitempty"`
			Thumb struct {
				De string `json:"de,omitempty"`
				Fr string `json:"fr,omitempty"`
			} `json:"thumb,omitempty"`
		} `json:"ingredients,omitempty"`
		Nutrition struct {
			Display struct {
				Fr string `json:"fr,omitempty"`
			} `json:"display,omitempty"`
			Small struct {
				Fr string `json:"fr,omitempty"`
			} `json:"small,omitempty"`
			Thumb struct {
				Fr string `json:"fr,omitempty"`
			} `json:"thumb,omitempty"`
		} `json:"nutrition,omitempty"`
	} `json:"selected_images,omitempty"`
	Sortkey               int      `json:"sortkey,omitempty"`
	States                string   `json:"states,omitempty"`
	StatesHierarchy       []string `json:"states_hierarchy,omitempty"`
	StatesTags            []string `json:"states_tags,omitempty"`
	Stores                string   `json:"stores,omitempty"`
	StoresTags            []any    `json:"stores_tags,omitempty"`
	Traces                string   `json:"traces,omitempty"`
	TracesFromIngredients string   `json:"traces_from_ingredients,omitempty"`
	TracesFromUser        string   `json:"traces_from_user,omitempty"`
	TracesHierarchy       []any    `json:"traces_hierarchy,omitempty"`
	TracesLc              string   `json:"traces_lc,omitempty"`
	TracesTags            []any    `json:"traces_tags,omitempty"`
	UniqueScansN          int      `json:"unique_scans_n,omitempty"`
	UnknownNutrientsTags  []any    `json:"unknown_nutrients_tags,omitempty"`
	UpdateKey             string   `json:"update_key,omitempty"`
	URL                   string   `json:"url,omitempty"`
	VitaminsPrevTags      []any    `json:"vitamins_prev_tags,omitempty"`
	VitaminsTags          []any    `json:"vitamins_tags,omitempty"`
	WeighersTags          []any    `json:"weighers_tags,omitempty"`
	CategoriesProperties0 struct {
	} `json:"categories_properties,omitempty"`
	EcoscoreData0 struct {
		Adjustments struct {
			OriginsOfIngredients struct {
				AggregatedOrigins []struct {
					EpiScore            string `json:"epi_score,omitempty"`
					Origin              string `json:"origin,omitempty"`
					Percent             int    `json:"percent,omitempty"`
					TransportationScore any    `json:"transportation_score,omitempty"`
				} `json:"aggregated_origins,omitempty"`
				EpiScore                int      `json:"epi_score,omitempty"`
				EpiValue                int      `json:"epi_value,omitempty"`
				OriginsFromCategories   []string `json:"origins_from_categories,omitempty"`
				OriginsFromOriginsField []string `json:"origins_from_origins_field,omitempty"`
				TransportationScore     int      `json:"transportation_score,omitempty"`
				TransportationScores    struct {
					Ad    int `json:"ad,omitempty"`
					Al    int `json:"al,omitempty"`
					At    int `json:"at,omitempty"`
					Ax    int `json:"ax,omitempty"`
					Ba    int `json:"ba,omitempty"`
					Be    int `json:"be,omitempty"`
					Bg    int `json:"bg,omitempty"`
					Ch    int `json:"ch,omitempty"`
					Cy    int `json:"cy,omitempty"`
					Cz    int `json:"cz,omitempty"`
					De    int `json:"de,omitempty"`
					Dk    int `json:"dk,omitempty"`
					Dz    int `json:"dz,omitempty"`
					Ee    int `json:"ee,omitempty"`
					Eg    int `json:"eg,omitempty"`
					Es    int `json:"es,omitempty"`
					Fi    int `json:"fi,omitempty"`
					Fo    int `json:"fo,omitempty"`
					Fr    int `json:"fr,omitempty"`
					Gg    int `json:"gg,omitempty"`
					Gi    int `json:"gi,omitempty"`
					Gr    int `json:"gr,omitempty"`
					Hr    int `json:"hr,omitempty"`
					Hu    int `json:"hu,omitempty"`
					Ie    int `json:"ie,omitempty"`
					Il    int `json:"il,omitempty"`
					Im    int `json:"im,omitempty"`
					Is    int `json:"is,omitempty"`
					It    int `json:"it,omitempty"`
					Je    int `json:"je,omitempty"`
					Lb    int `json:"lb,omitempty"`
					Li    int `json:"li,omitempty"`
					Lt    int `json:"lt,omitempty"`
					Lu    int `json:"lu,omitempty"`
					Lv    int `json:"lv,omitempty"`
					Ly    int `json:"ly,omitempty"`
					Ma    int `json:"ma,omitempty"`
					Mc    int `json:"mc,omitempty"`
					Md    int `json:"md,omitempty"`
					Me    int `json:"me,omitempty"`
					Mk    int `json:"mk,omitempty"`
					Mt    int `json:"mt,omitempty"`
					Nl    int `json:"nl,omitempty"`
					No    int `json:"no,omitempty"`
					Pl    int `json:"pl,omitempty"`
					Ps    int `json:"ps,omitempty"`
					Pt    int `json:"pt,omitempty"`
					Ro    int `json:"ro,omitempty"`
					Rs    int `json:"rs,omitempty"`
					Se    int `json:"se,omitempty"`
					Si    int `json:"si,omitempty"`
					Sj    int `json:"sj,omitempty"`
					Sk    int `json:"sk,omitempty"`
					Sm    int `json:"sm,omitempty"`
					Sy    int `json:"sy,omitempty"`
					Tn    int `json:"tn,omitempty"`
					Tr    int `json:"tr,omitempty"`
					Ua    int `json:"ua,omitempty"`
					Uk    int `json:"uk,omitempty"`
					Us    int `json:"us,omitempty"`
					Va    int `json:"va,omitempty"`
					World int `json:"world,omitempty"`
					Xk    int `json:"xk,omitempty"`
				} `json:"transportation_scores,omitempty"`
				TransportationValue  int `json:"transportation_value,omitempty"`
				TransportationValues struct {
					Ad    int `json:"ad,omitempty"`
					Al    int `json:"al,omitempty"`
					At    int `json:"at,omitempty"`
					Ax    int `json:"ax,omitempty"`
					Ba    int `json:"ba,omitempty"`
					Be    int `json:"be,omitempty"`
					Bg    int `json:"bg,omitempty"`
					Ch    int `json:"ch,omitempty"`
					Cy    int `json:"cy,omitempty"`
					Cz    int `json:"cz,omitempty"`
					De    int `json:"de,omitempty"`
					Dk    int `json:"dk,omitempty"`
					Dz    int `json:"dz,omitempty"`
					Ee    int `json:"ee,omitempty"`
					Eg    int `json:"eg,omitempty"`
					Es    int `json:"es,omitempty"`
					Fi    int `json:"fi,omitempty"`
					Fo    int `json:"fo,omitempty"`
					Fr    int `json:"fr,omitempty"`
					Gg    int `json:"gg,omitempty"`
					Gi    int `json:"gi,omitempty"`
					Gr    int `json:"gr,omitempty"`
					Hr    int `json:"hr,omitempty"`
					Hu    int `json:"hu,omitempty"`
					Ie    int `json:"ie,omitempty"`
					Il    int `json:"il,omitempty"`
					Im    int `json:"im,omitempty"`
					Is    int `json:"is,omitempty"`
					It    int `json:"it,omitempty"`
					Je    int `json:"je,omitempty"`
					Lb    int `json:"lb,omitempty"`
					Li    int `json:"li,omitempty"`
					Lt    int `json:"lt,omitempty"`
					Lu    int `json:"lu,omitempty"`
					Lv    int `json:"lv,omitempty"`
					Ly    int `json:"ly,omitempty"`
					Ma    int `json:"ma,omitempty"`
					Mc    int `json:"mc,omitempty"`
					Md    int `json:"md,omitempty"`
					Me    int `json:"me,omitempty"`
					Mk    int `json:"mk,omitempty"`
					Mt    int `json:"mt,omitempty"`
					Nl    int `json:"nl,omitempty"`
					No    int `json:"no,omitempty"`
					Pl    int `json:"pl,omitempty"`
					Ps    int `json:"ps,omitempty"`
					Pt    int `json:"pt,omitempty"`
					Ro    int `json:"ro,omitempty"`
					Rs    int `json:"rs,omitempty"`
					Se    int `json:"se,omitempty"`
					Si    int `json:"si,omitempty"`
					Sj    int `json:"sj,omitempty"`
					Sk    int `json:"sk,omitempty"`
					Sm    int `json:"sm,omitempty"`
					Sy    int `json:"sy,omitempty"`
					Tn    int `json:"tn,omitempty"`
					Tr    int `json:"tr,omitempty"`
					Ua    int `json:"ua,omitempty"`
					Uk    int `json:"uk,omitempty"`
					Us    int `json:"us,omitempty"`
					Va    int `json:"va,omitempty"`
					World int `json:"world,omitempty"`
					Xk    int `json:"xk,omitempty"`
				} `json:"transportation_values,omitempty"`
				Value  int `json:"value,omitempty"`
				Values struct {
					Ad    int `json:"ad,omitempty"`
					Al    int `json:"al,omitempty"`
					At    int `json:"at,omitempty"`
					Ax    int `json:"ax,omitempty"`
					Ba    int `json:"ba,omitempty"`
					Be    int `json:"be,omitempty"`
					Bg    int `json:"bg,omitempty"`
					Ch    int `json:"ch,omitempty"`
					Cy    int `json:"cy,omitempty"`
					Cz    int `json:"cz,omitempty"`
					De    int `json:"de,omitempty"`
					Dk    int `json:"dk,omitempty"`
					Dz    int `json:"dz,omitempty"`
					Ee    int `json:"ee,omitempty"`
					Eg    int `json:"eg,omitempty"`
					Es    int `json:"es,omitempty"`
					Fi    int `json:"fi,omitempty"`
					Fo    int `json:"fo,omitempty"`
					Fr    int `json:"fr,omitempty"`
					Gg    int `json:"gg,omitempty"`
					Gi    int `json:"gi,omitempty"`
					Gr    int `json:"gr,omitempty"`
					Hr    int `json:"hr,omitempty"`
					Hu    int `json:"hu,omitempty"`
					Ie    int `json:"ie,omitempty"`
					Il    int `json:"il,omitempty"`
					Im    int `json:"im,omitempty"`
					Is    int `json:"is,omitempty"`
					It    int `json:"it,omitempty"`
					Je    int `json:"je,omitempty"`
					Lb    int `json:"lb,omitempty"`
					Li    int `json:"li,omitempty"`
					Lt    int `json:"lt,omitempty"`
					Lu    int `json:"lu,omitempty"`
					Lv    int `json:"lv,omitempty"`
					Ly    int `json:"ly,omitempty"`
					Ma    int `json:"ma,omitempty"`
					Mc    int `json:"mc,omitempty"`
					Md    int `json:"md,omitempty"`
					Me    int `json:"me,omitempty"`
					Mk    int `json:"mk,omitempty"`
					Mt    int `json:"mt,omitempty"`
					Nl    int `json:"nl,omitempty"`
					No    int `json:"no,omitempty"`
					Pl    int `json:"pl,omitempty"`
					Ps    int `json:"ps,omitempty"`
					Pt    int `json:"pt,omitempty"`
					Ro    int `json:"ro,omitempty"`
					Rs    int `json:"rs,omitempty"`
					Se    int `json:"se,omitempty"`
					Si    int `json:"si,omitempty"`
					Sj    int `json:"sj,omitempty"`
					Sk    int `json:"sk,omitempty"`
					Sm    int `json:"sm,omitempty"`
					Sy    int `json:"sy,omitempty"`
					Tn    int `json:"tn,omitempty"`
					Tr    int `json:"tr,omitempty"`
					Ua    int `json:"ua,omitempty"`
					Uk    int `json:"uk,omitempty"`
					Us    int `json:"us,omitempty"`
					Va    int `json:"va,omitempty"`
					World int `json:"world,omitempty"`
					Xk    int `json:"xk,omitempty"`
				} `json:"values,omitempty"`
				Warning string `json:"warning,omitempty"`
			} `json:"origins_of_ingredients,omitempty"`
			Packaging struct {
				NonRecyclableAndNonBiodegradableMaterials int    `json:"non_recyclable_and_non_biodegradable_materials,omitempty"`
				Value                                     int    `json:"value,omitempty"`
				Warning                                   string `json:"warning,omitempty"`
			} `json:"packaging,omitempty"`
			ProductionSystem struct {
				Labels  []any  `json:"labels,omitempty"`
				Value   int    `json:"value,omitempty"`
				Warning string `json:"warning,omitempty"`
			} `json:"production_system,omitempty"`
			ThreatenedSpecies struct {
				Warning string `json:"warning,omitempty"`
			} `json:"threatened_species,omitempty"`
		} `json:"adjustments,omitempty"`
		Agribalyse struct {
			Warning string `json:"warning,omitempty"`
		} `json:"agribalyse,omitempty"`
		Missing struct {
			Categories  int `json:"categories,omitempty"`
			Ingredients int `json:"ingredients,omitempty"`
			Labels      int `json:"labels,omitempty"`
			Origins     int `json:"origins,omitempty"`
			Packagings  int `json:"packagings,omitempty"`
		} `json:"missing,omitempty"`
		MissingAgribalyseMatchWarning int `json:"missing_agribalyse_match_warning,omitempty"`
		MissingKeyData                int `json:"missing_key_data,omitempty"`
		Scores                        struct {
		} `json:"scores,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"ecoscore_data,omitempty"`
	Images0 struct {
		Num1 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"1,omitempty"`
		Num2 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"2,omitempty"`
		FrontEn struct {
			Angle                int    `json:"angle,omitempty"`
			CoordinatesImageSize string `json:"coordinates_image_size,omitempty"`
			Geometry             string `json:"geometry,omitempty"`
			Imgid                string `json:"imgid,omitempty"`
			Normalize            any    `json:"normalize,omitempty"`
			Rev                  string `json:"rev,omitempty"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num200 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"200,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			WhiteMagic any    `json:"white_magic,omitempty"`
			X1         string `json:"x1,omitempty"`
			X2         string `json:"x2,omitempty"`
			Y1         string `json:"y1,omitempty"`
			Y2         string `json:"y2,omitempty"`
		} `json:"front_en,omitempty"`
		NutritionEn struct {
			Angle                int    `json:"angle,omitempty"`
			CoordinatesImageSize string `json:"coordinates_image_size,omitempty"`
			Geometry             string `json:"geometry,omitempty"`
			Imgid                string `json:"imgid,omitempty"`
			Normalize            any    `json:"normalize,omitempty"`
			Rev                  string `json:"rev,omitempty"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num200 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"200,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			WhiteMagic any    `json:"white_magic,omitempty"`
			X1         string `json:"x1,omitempty"`
			X2         string `json:"x2,omitempty"`
			Y1         string `json:"y1,omitempty"`
			Y2         string `json:"y2,omitempty"`
		} `json:"nutrition_en,omitempty"`
	} `json:"images,omitempty"`
	Languages0 struct {
		EnEnglish int `json:"en:english,omitempty"`
	} `json:"languages,omitempty"`
	LanguagesCodes0 struct {
		En int `json:"en,omitempty"`
	} `json:"languages_codes,omitempty"`
	NovaGroupError  string `json:"nova_group_error,omitempty"`
	NutrientLevels0 struct {
	} `json:"nutrient_levels,omitempty"`
	Nutriments0 struct {
		Carbohydrates           int    `json:"carbohydrates,omitempty"`
		Carbohydrates100G       int    `json:"carbohydrates_100g,omitempty"`
		CarbohydratesUnit       string `json:"carbohydrates_unit,omitempty"`
		CarbohydratesValue      int    `json:"carbohydrates_value,omitempty"`
		Energy                  int    `json:"energy,omitempty"`
		EnergyKcal              string `json:"energy-kcal,omitempty"`
		EnergyKcal100G          string `json:"energy-kcal_100g,omitempty"`
		EnergyKcalUnit          string `json:"energy-kcal_unit,omitempty"`
		EnergyKcalValue         string `json:"energy-kcal_value,omitempty"`
		EnergyKcalValueComputed int    `json:"energy-kcal_value_computed,omitempty"`
		Energy100G              int    `json:"energy_100g,omitempty"`
		EnergyUnit              string `json:"energy_unit,omitempty"`
		EnergyValue             string `json:"energy_value,omitempty"`
		Fat                     int    `json:"fat,omitempty"`
		Fat100G                 int    `json:"fat_100g,omitempty"`
		FatUnit                 string `json:"fat_unit,omitempty"`
		FatValue                int    `json:"fat_value,omitempty"`
		Fiber                   int    `json:"fiber,omitempty"`
		Fiber100G               int    `json:"fiber_100g,omitempty"`
		FiberUnit               string `json:"fiber_unit,omitempty"`
		FiberValue              int    `json:"fiber_value,omitempty"`
		Proteins                int    `json:"proteins,omitempty"`
		Proteins100G            int    `json:"proteins_100g,omitempty"`
		ProteinsUnit            string `json:"proteins_unit,omitempty"`
		ProteinsValue           int    `json:"proteins_value,omitempty"`
		Salt                    string `json:"salt,omitempty"`
		Salt100G                string `json:"salt_100g,omitempty"`
		SaltUnit                string `json:"salt_unit,omitempty"`
		SaltValue               string `json:"salt_value,omitempty"`
		SaturatedFat            int    `json:"saturated-fat,omitempty"`
		SaturatedFat100G        int    `json:"saturated-fat_100g,omitempty"`
		SaturatedFatUnit        string `json:"saturated-fat_unit,omitempty"`
		SaturatedFatValue       int    `json:"saturated-fat_value,omitempty"`
		Sodium                  string `json:"sodium,omitempty"`
		Sodium100G              string `json:"sodium_100g,omitempty"`
		SodiumUnit              string `json:"sodium_unit,omitempty"`
		SodiumValue             string `json:"sodium_value,omitempty"`
		Sugars                  int    `json:"sugars,omitempty"`
		Sugars100G              int    `json:"sugars_100g,omitempty"`
		SugarsUnit              string `json:"sugars_unit,omitempty"`
		SugarsValue             int    `json:"sugars_value,omitempty"`
	} `json:"nutriments,omitempty"`
	ProductNameEn   string `json:"product_name_en,omitempty"`
	SelectedImages0 struct {
		Front struct {
			Display struct {
				En string `json:"en,omitempty"`
			} `json:"display,omitempty"`
			Small struct {
				En string `json:"en,omitempty"`
			} `json:"small,omitempty"`
			Thumb struct {
				En string `json:"en,omitempty"`
			} `json:"thumb,omitempty"`
		} `json:"front,omitempty"`
		Nutrition struct {
			Display struct {
				En string `json:"en,omitempty"`
			} `json:"display,omitempty"`
			Small struct {
				En string `json:"en,omitempty"`
			} `json:"small,omitempty"`
			Thumb struct {
				En string `json:"en,omitempty"`
			} `json:"thumb,omitempty"`
		} `json:"nutrition,omitempty"`
	} `json:"selected_images,omitempty"`
	AdditivesN            int    `json:"additives_n,omitempty"`
	AdditivesOldN         int    `json:"additives_old_n,omitempty"`
	BrandOwner            string `json:"brand_owner,omitempty"`
	BrandOwnerImported    string `json:"brand_owner_imported,omitempty"`
	CategoriesImported    string `json:"categories_imported,omitempty"`
	CategoriesProperties1 struct {
		AgribalyseProxyFoodCodeEn string `json:"agribalyse_proxy_food_code:en,omitempty"`
	} `json:"categories_properties,omitempty"`
	CategoryProperties0 struct {
	} `json:"category_properties,omitempty"`
	CountriesImported    string `json:"countries_imported,omitempty"`
	DataSourcesImported  string `json:"data_sources_imported,omitempty"`
	EcoscoreExtendedData struct {
		Impact struct {
			EfSingleScoreLogStddev any `json:"ef_single_score_log_stddev,omitempty"`
			LikeliestImpacts       struct {
				ClimateChange any `json:"Climate_change,omitempty"`
				EFSingleScore any `json:"EF_single_score,omitempty"`
			} `json:"likeliest_impacts,omitempty"`
			LikeliestRecipe struct {
				EnCaneSugarMolasses any `json:"en:cane_sugar_molasses,omitempty"`
				EnSugar             any `json:"en:sugar,omitempty"`
				EnWater             any `json:"en:water,omitempty"`
			} `json:"likeliest_recipe,omitempty"`
			MassRatioUncharacterized   any `json:"mass_ratio_uncharacterized,omitempty"`
			UncharacterizedIngredients struct {
				Impact    []any `json:"impact,omitempty"`
				Nutrition []any `json:"nutrition,omitempty"`
			} `json:"uncharacterized_ingredients,omitempty"`
			UncharacterizedIngredientsMassProportion struct {
				Impact    any `json:"impact,omitempty"`
				Nutrition any `json:"nutrition,omitempty"`
			} `json:"uncharacterized_ingredients_mass_proportion,omitempty"`
			UncharacterizedIngredientsRatio struct {
				Impact    any `json:"impact,omitempty"`
				Nutrition any `json:"nutrition,omitempty"`
			} `json:"uncharacterized_ingredients_ratio,omitempty"`
			Warnings []string `json:"warnings,omitempty"`
		} `json:"impact,omitempty"`
	} `json:"ecoscore_extended_data,omitempty"`
	EcoscoreExtendedDataVersion string `json:"ecoscore_extended_data_version,omitempty"`
	Images1                     struct {
		Num1 struct {
			Sizes struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			UploadedT int    `json:"uploaded_t,omitempty"`
			Uploader  string `json:"uploader,omitempty"`
		} `json:"1,omitempty"`
		FrontEn struct {
			Angle                int    `json:"angle,omitempty"`
			CoordinatesImageSize string `json:"coordinates_image_size,omitempty"`
			Geometry             string `json:"geometry,omitempty"`
			Imgid                string `json:"imgid,omitempty"`
			Normalize            any    `json:"normalize,omitempty"`
			Rev                  string `json:"rev,omitempty"`
			Sizes                struct {
				Num100 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"100,omitempty"`
				Num200 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"200,omitempty"`
				Num400 struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"400,omitempty"`
				Full struct {
					H int `json:"h,omitempty"`
					W int `json:"w,omitempty"`
				} `json:"full,omitempty"`
			} `json:"sizes,omitempty"`
			WhiteMagic any    `json:"white_magic,omitempty"`
			X1         string `json:"x1,omitempty"`
			X2         string `json:"x2,omitempty"`
			Y1         string `json:"y1,omitempty"`
			Y2         string `json:"y2,omitempty"`
		} `json:"front_en,omitempty"`
	} `json:"images,omitempty"`
	Ingredients []struct {
		CiqualProxyFoodCode string `json:"ciqual_proxy_food_code,omitempty"`
		ID                  string `json:"id,omitempty"`
		PercentEstimate     any    `json:"percent_estimate,omitempty"`
		PercentMax          any    `json:"percent_max,omitempty"`
		PercentMin          any    `json:"percent_min,omitempty"`
		Rank                int    `json:"rank,omitempty"`
		Text                string `json:"text,omitempty"`
		Vegan               string `json:"vegan,omitempty"`
		Vegetarian          string `json:"vegetarian,omitempty"`
		CiqualFoodCode      string `json:"ciqual_food_code,omitempty"`
	} `json:"ingredients,omitempty"`
	IngredientsAnalysis struct {
	} `json:"ingredients_analysis,omitempty"`
	IngredientsAnalysisTags                []string `json:"ingredients_analysis_tags,omitempty"`
	IngredientsFromOrThatMayBeFromPalmOilN int      `json:"ingredients_from_or_that_may_be_from_palm_oil_n,omitempty"`
	IngredientsFromPalmOilN                int      `json:"ingredients_from_palm_oil_n,omitempty"`
	IngredientsHierarchy                   []string `json:"ingredients_hierarchy,omitempty"`
	IngredientsN                           any      `json:"ingredients_n,omitempty"`
	IngredientsNTags                       []string `json:"ingredients_n_tags,omitempty"`
	IngredientsOriginalTags                []string `json:"ingredients_original_tags,omitempty"`
	IngredientsPercentAnalysis             int      `json:"ingredients_percent_analysis,omitempty"`
	IngredientsTags                        []string `json:"ingredients_tags,omitempty"`
	//IngredientsTextEn                     *any     `json:"ingredients_text_en,omitempty"`
	IngredientsTextEn *string `json:"ingredients_text_en,omitempty"`

	IngredientsTextEnImported            string `json:"ingredients_text_en_imported,omitempty"`
	IngredientsTextWithAllergensEn       string `json:"ingredients_text_with_allergens_en,omitempty"`
	IngredientsThatMayBeFromPalmOilN     int    `json:"ingredients_that_may_be_from_palm_oil_n,omitempty"`
	IngredientsWithSpecifiedPercentN     int    `json:"ingredients_with_specified_percent_n,omitempty"`
	IngredientsWithSpecifiedPercentSum   any    `json:"ingredients_with_specified_percent_sum,omitempty"`
	IngredientsWithUnspecifiedPercentN   int    `json:"ingredients_with_unspecified_percent_n,omitempty"`
	IngredientsWithUnspecifiedPercentSum any    `json:"ingredients_with_unspecified_percent_sum,omitempty"`
	IngredientsWithoutCiqualCodes        []any  `json:"ingredients_without_ciqual_codes,omitempty"`
	IngredientsWithoutCiqualCodesN       int    `json:"ingredients_without_ciqual_codes_n,omitempty"`
	KnownIngredientsN                    int    `json:"known_ingredients_n,omitempty"`
	Languages1                           struct {
		EnEnglish int `json:"en:english,omitempty"`
	} `json:"languages,omitempty"`
	LanguagesCodes1 struct {
		En int `json:"en,omitempty"`
	} `json:"languages_codes,omitempty"`
	LcImported      string `json:"lc_imported,omitempty"`
	NutrientLevels1 struct {
		Fat          string `json:"fat,omitempty"`
		SaturatedFat string `json:"saturated-fat,omitempty"`
		Sugars       string `json:"sugars,omitempty"`
	} `json:"nutrient_levels,omitempty"`
	Nutriments1 struct {
		Carbohydrates                                         int    `json:"carbohydrates,omitempty"`
		Carbohydrates100G                                     int    `json:"carbohydrates_100g,omitempty"`
		CarbohydratesServing                                  int    `json:"carbohydrates_serving,omitempty"`
		CarbohydratesUnit                                     string `json:"carbohydrates_unit,omitempty"`
		CarbohydratesValue                                    int    `json:"carbohydrates_value,omitempty"`
		Energy                                                int    `json:"energy,omitempty"`
		EnergyKcal                                            int    `json:"energy-kcal,omitempty"`
		EnergyKcal100G                                        int    `json:"energy-kcal_100g,omitempty"`
		EnergyKcalServing                                     int    `json:"energy-kcal_serving,omitempty"`
		EnergyKcalUnit                                        string `json:"energy-kcal_unit,omitempty"`
		EnergyKcalValue                                       int    `json:"energy-kcal_value,omitempty"`
		EnergyKcalValueComputed                               int    `json:"energy-kcal_value_computed,omitempty"`
		Energy100G                                            int    `json:"energy_100g,omitempty"`
		EnergyServing                                         string `json:"energy_serving,omitempty"`
		EnergyUnit                                            string `json:"energy_unit,omitempty"`
		EnergyValue                                           int    `json:"energy_value,omitempty"`
		Fat                                                   int    `json:"fat,omitempty"`
		Fat100G                                               int    `json:"fat_100g,omitempty"`
		FatServing                                            int    `json:"fat_serving,omitempty"`
		FatUnit                                               string `json:"fat_unit,omitempty"`
		FatValue                                              int    `json:"fat_value,omitempty"`
		FruitsVegetablesLegumesEstimateFromIngredients100G    int    `json:"fruits-vegetables-legumes-estimate-from-ingredients_100g,omitempty"`
		FruitsVegetablesLegumesEstimateFromIngredientsServing int    `json:"fruits-vegetables-legumes-estimate-from-ingredients_serving,omitempty"`
		FruitsVegetablesNutsEstimateFromIngredients100G       int    `json:"fruits-vegetables-nuts-estimate-from-ingredients_100g,omitempty"`
		FruitsVegetablesNutsEstimateFromIngredientsServing    int    `json:"fruits-vegetables-nuts-estimate-from-ingredients_serving,omitempty"`
		NovaGroup                                             int    `json:"nova-group,omitempty"`
		NovaGroup100G                                         int    `json:"nova-group_100g,omitempty"`
		NovaGroupServing                                      int    `json:"nova-group_serving,omitempty"`
		Proteins                                              int    `json:"proteins,omitempty"`
		Proteins100G                                          int    `json:"proteins_100g,omitempty"`
		ProteinsServing                                       int    `json:"proteins_serving,omitempty"`
		ProteinsUnit                                          string `json:"proteins_unit,omitempty"`
		ProteinsValue                                         int    `json:"proteins_value,omitempty"`
		SaturatedFat                                          int    `json:"saturated-fat,omitempty"`
		SaturatedFat100G                                      int    `json:"saturated-fat_100g,omitempty"`
		SaturatedFatServing                                   int    `json:"saturated-fat_serving,omitempty"`
		SaturatedFatUnit                                      string `json:"saturated-fat_unit,omitempty"`
		SaturatedFatValue                                     int    `json:"saturated-fat_value,omitempty"`
		Sugars                                                int    `json:"sugars,omitempty"`
		Sugars100G                                            int    `json:"sugars_100g,omitempty"`
		SugarsServing                                         int    `json:"sugars_serving,omitempty"`
		SugarsUnit                                            string `json:"sugars_unit,omitempty"`
		SugarsValue                                           int    `json:"sugars_value,omitempty"`
		TransFat                                              int    `json:"trans-fat,omitempty"`
		TransFat100G                                          int    `json:"trans-fat_100g,omitempty"`
		TransFatServing                                       int    `json:"trans-fat_serving,omitempty"`
		TransFatUnit                                          string `json:"trans-fat_unit,omitempty"`
		TransFatValue                                         int    `json:"trans-fat_value,omitempty"`
	} `json:"nutriments,omitempty"`
	NutritionDataPerImported                                                 string `json:"nutrition_data_per_imported,omitempty"`
	NutritionDataPreparedPerImported                                         string `json:"nutrition_data_prepared_per_imported,omitempty"`
	NutritionScoreWarningFruitsVegetablesLegumesEstimateFromIngredients      any    `json:"nutrition_score_warning_fruits_vegetables_legumes_estimate_from_ingredients,omitempty"`
	NutritionScoreWarningFruitsVegetablesLegumesEstimateFromIngredientsValue any    `json:"nutrition_score_warning_fruits_vegetables_legumes_estimate_from_ingredients_value,omitempty"`
	NutritionScoreWarningFruitsVegetablesNutsEstimateFromIngredients         any    `json:"nutrition_score_warning_fruits_vegetables_nuts_estimate_from_ingredients,omitempty"`
	NutritionScoreWarningFruitsVegetablesNutsEstimateFromIngredientsValue    any    `json:"nutrition_score_warning_fruits_vegetables_nuts_estimate_from_ingredients_value,omitempty"`
	ProductNameEnImported                                                    string `json:"product_name_en_imported,omitempty"`
	SelectedImages1                                                          struct {
		Front struct {
			Display struct {
				En string `json:"en,omitempty"`
			} `json:"display,omitempty"`
			Small struct {
				En string `json:"en,omitempty"`
			} `json:"small,omitempty"`
			Thumb struct {
				En string `json:"en,omitempty"`
			} `json:"thumb,omitempty"`
		} `json:"front,omitempty"`
	} `json:"selected_images,omitempty"`
	ServingQuantity     any    `json:"serving_quantity,omitempty"`
	ServingSize         string `json:"serving_size,omitempty"`
	ServingSizeImported string `json:"serving_size_imported,omitempty"`
	Sources             []struct {
		Fields       []string `json:"fields,omitempty"`
		ID           string   `json:"id,omitempty"`
		Images       []any    `json:"images,omitempty"`
		ImportT      int      `json:"import_t,omitempty"`
		Manufacturer any      `json:"manufacturer,omitempty"`
		Name         string   `json:"name,omitempty"`
		URL          any      `json:"url,omitempty"`
	} `json:"sources,omitempty"`
	SourcesFields struct {
		OrgDatabaseUsda struct {
			AvailableDate   time.Time `json:"available_date,omitempty"`
			FdcCategory     string    `json:"fdc_category,omitempty"`
			FdcDataSource   string    `json:"fdc_data_source,omitempty"`
			FdcID           string    `json:"fdc_id,omitempty"`
			ModifiedDate    time.Time `json:"modified_date,omitempty"`
			PublicationDate time.Time `json:"publication_date,omitempty"`
		} `json:"org-database-usda,omitempty"`
	} `json:"sources_fields,omitempty"`
	UnknownIngredientsN any `json:"unknown_ingredients_n,omitempty"`
}
