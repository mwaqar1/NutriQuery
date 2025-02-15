package main

var (
	badIngredientsSlice = []string{
		"carrageenan", "sugar", "fluoride", "aluminum",
		"sodium aluminum phosphate", "sodium aluminum sulfate",
		"arsenic", "rice syrup", "cadmium", "lead",
		"mercury", "monosodium glutamate(msg)", "alachlor",
		"acetochlor", "metolachlor", "atrazine",
		"simazine", "cyanazine", "processed free glutamic acid",
		"organophosphates", "malathion", "diazinon",
		"dichlorvos", "fenthion", "chlorpyrif",
		"glyphosate", "polyethoxylated tallowamine", "poea",
		"phosphates", "bpa", "bisphenol a", "1-bromopropane", "1-bp", "1 bromopropane",
		"di(2-ethylhexyl)adipate", "bis(2-ethylhexyl)adipate", "dioctyladipate", "deha", "dioxins", "dioctyl adipate",
		" hexanedioic acid", "ethylbenzene", "glycol ethers", "methylene chloride",
		"p-dichlorobenzene", "phthalates", "styrene",
		"white refined flour", "white refined sugar",
		"artificial color", "yellow 5", "4-methylimidazole", "4 methylimidazole",
		"4-mei", "butylated hydroxyanisole", "bht", "butylated hydroxytoluene",
		"nitrites", "propionates", "parabens", "methylparaben",
		"ethylparaben", "propylparaben", "potassium sorbate",
		"sodium benzoate", "potassium benzoate", "calcium benzoate",
		"benzene", "artificial sweetener", "aspartame",
		"sucralose", "saccharin", "recombinant bovine growth hormone", "rbgh",
		"high fructose corn syrup", "hfcs", "maize syrup", "glucose syrup",
		"tapioca syrup", "fruit fructose", "crystalline fructose",
		"fructose", "hydrogenated oil", "pfcs", "perfluorinated compounds",
		"perfluorooctanoic acid", "pfoa", "pfos", "perfluorooctane sulfonic acid",
		"ptfe", "polytetrafluoroetheylene", "artificial flavors",
		"hydrolyzed vegetable protein", "natural flavors",
		"tvp", "textured vegetable protein", "perfluorinated chemicals", "perfluorochemicals", "pfc",
		"fluorotelomer", "teflon", "c8", "c7", "c9",
		"c10", "c11", "c12", "dioxin",
		"2 tetrachlorodibenzo para dioxin", "tcdd", "3 tetrachlorodibenzo para dioxin",
		"7 tetrachlorodibenzo para dioxin", "8 tetrachlorodibenzo para dioxin",
		"polychlorinated dibenzo para dioxins", "pcdds", "polychlorinated dibenzofurans", "pcdfs",
		"polychlorinated biphenyls", "pcbs", "canola oil", "sunflower oil", "safflower oil",
		"vegetable oil", "palm oil", "partially hydrogenated oil", "corn oil", "rapeseed oil", "soy oil",
		"rice bran oil", "grapeseed oil", "red 40", "yellow 6", "blue 1", "blue 2 ", "green 3", "red 3",
		"citrus red", "artifical color", "natural smoke flavor", "soy bean oil", "corn syrup", "cottonseed oil",
		"malted cereal syrup", "sodium nitrite", "potassium nitrate", "celery powder", "celery juice", "cured",
		"curing", "nitrate", "nitrite", "niacin", "nicotinic acid", "potassium bromate",
		"heterocyclic amines", "polycyclic aromatic hydrocarbons", "titanium dioxide", "artificial", "tbhq",
		"tert butylhydroquinone", "azodicarbonamide", "ada",
		"brominated vegetable oil", "bvo", "propyl parabens", "propyl gallate",
		"sodium benzoate", "caramel color", "edta", "calcium disodium", "emulsifier lecithine", "Soy Lecithin", "emulsifier emulsifier lecithin",
		"flavour enhancers", "monosodium glutamate", "inosinato", "guanylate disódicos", "smoke flavouring", "colouring", "soybean oil",
		"skimmed milk powder", "sunflower lecithin", "e101a", "riboflavin 5 phosphate", "e338", "phosphoric acid", "e339", "sodium phosphate",
		"e340", "potassium phosphate", "e341", "calcium phosphate", "e343", "magnesium phosphate",
		"e442", "ammonium phosphatide", "e450", "diphosphate", "e451", "triphosphate", "e452",
		"polyphosphate", "e540", "dicalcium diphosphate", "e541",
		"aluminium phosphate", "e543", "calcium sodium polyphosphate",
		"e544", "calcium polyphosphate", "e1410", "phosphate starch",
		"e1412", "distarch phosphate", "e1413", "phosphated distarch phosphate",
		"e1414", "acetylated distarch phosphate", "e1442", "hydroxypropyl distarch phosphate",
		"datem", "diacetyl tartaric esters of monoglycerides", "diacetyl tartaric acid esters",
		"diacetyl tartaric acid esters of diglycerides", "corn malt extract", "malt extract", "malt syrup", "e471", "e250",
		"sodium tripolyphosphate", "e316", "sodium erythorbate", "e155", "brown ht", "caramellsed milk powder", "natural flavouring",
		"colour", "modified palm oil", "emulsifiers: lecithin", "sweetened condensed whole milk ", "sweetened condensed milk", "maltodextrin", "corn maltodextrin",
		"disodium guanylate", "interesterified soybean oil", "e627", "refined sunflower oil", "vegetable oils", "vegetables oil", "natural flavor", "liquid sugar",
		"flavourings",
	}
)
