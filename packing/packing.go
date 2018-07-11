package packing

type Category struct {
	ListName    string
	Suggestions []string
}

var (
	Categories = []Category{
		Category{
			ListName: "Clothing",
			Suggestions: []string{
				"Tops",
				"Bottoms",
				"Jacket",
				// "Sweatshirt",
				// "Pajamas",
				// "Socks",
				// "Shoes",
				// "Sandals",
				// "Undershirts",
				// "Belt",
			},
		},
		// Category{
		// 	ListName: "Documents",
		// 	Suggestions: []string{
		// 		"Passport",
		// 		"License",
		// 		"Tickets",
		// 		"Reservation Info",
		// 		"Itinerary",
		// 	},
		// },
		// Category{
		// 	ListName: "Toiletries",
		// 	Suggestions: []string{
		// 		"Shampoo",
		// 		"Conditioner",
		// 		"Saving Items",
		// 		"Soap",
		// 		"Towel",
		// 		"Lotion",
		// 		"Makeup",
		// 		"Deoderant",
		// 		"Cologne",
		// 	},
		// },
		// Category{
		// 	ListName: "Equipment",
		// 	Suggestions: []string{
		// 		"Camera",
		// 		"Camera Battery",
		// 		"SD Card",
		// 		"Laptop",
		// 		"Charger",
		// 		"Phone",
		// 		"Phone wall charger",
		// 		"Phone car charger",
		// 		"Battery Pack",
		// 	},
		// },
		// Category{
		// 	ListName: "Misc.",
		// 	Suggestions: []string{
		// 		"Purse",
		// 		"Snacks",
		// 		"Sunglasses",
		// 		"Glasses",
		// 		"Medications",
		// 	},
		// },
	}
)

type Item struct {
	Name       string
	BeenPacked bool
}

var packingList []*Item

func AddToPackingList(name string, amount int) {
	for i := 0; i < amount; i++ {
		item := &Item{
			Name: name,
		}
		packingList = append(packingList, item)

	}

}

func FinalList() []*Item {
	return packingList

}
