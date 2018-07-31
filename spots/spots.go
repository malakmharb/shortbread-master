package spots

// maps a floor number to the relative URL that contains the floormap
var FloorPlans map[string]string = map[string]string{
	"1": "/assets/images/floor1.png",
	"2": "/assets/images/fllor2.png",
}

type Spot struct {
	Space    string
	X1       float64
	Y1       float64
	Width    float64
	Height   float64
	IsDouble bool
	Status   string
	floor    string
}

// maps a floor number to the list of available spots
var All map[string][]*Spot = map[string][]*Spot{
	"1": []*Spot{
		&Spot{"0", 108, 236, 31, 72, false, "Available", "1"},
		&Spot{"1", 140, 236, 31, 72, false, "Occupied", "1"},
		&Spot{"2", 180, 236, 31, 72, false, "Exceeded", "1"},
		&Spot{"3", 212, 236, 31, 72, false, "Available", "1"},
		&Spot{"4", 244, 236, 31, 72, false, "Available", "1"},
		&Spot{"5", 284, 236, 31, 72, false, "Available", "1"},
		&Spot{"6", 316, 236, 31, 72, false, "Available", "1"},
		&Spot{"7", 348, 236, 31, 72, false, "Available", "1"},
		&Spot{"9", 388, 146, 31, 162, true, "Available", "1"},
		&Spot{"11", 420, 146, 31, 162, true, "Available", "1"},
		&Spot{"13", 452, 146, 31, 162, true, "Available", "1"},
		&Spot{"15", 492, 146, 31, 162, true, "Available", "1"},
		&Spot{"17", 524, 146, 31, 162, true, "Available", "1"},
		&Spot{"19", 556, 146, 31, 162, true, "Available", "1"},
		&Spot{"21", 596, 146, 31, 162, true, "Available", "1"},
		&Spot{"23", 628, 146, 31, 162, true, "Available", "1"},
		&Spot{"25", 660, 146, 31, 162, true, "Available", "1"},
		&Spot{"26", 700, 146, 31, 162, true, "Available", "1"},
		&Spot{"27", 732, 236, 31, 72, false, "Available", "1"},
		&Spot{"29", 764, 236, 31, 72, false, "Available", "1"},
		&Spot{"29A", 804, 236, 31, 72, false, "Available", "1"},
		&Spot{"30A", 869, 236, 31, 72, false, "Available", "1"},
		&Spot{"30", 908, 236, 31, 72, false, "Available", "1"},
		&Spot{"31", 940, 236, 31, 72, false, "Available", "1"},
		&Spot{"32", 972, 236, 31, 72, false, "Available", "1"},
		&Spot{"33", 1012, 146, 32, 162, true, "Available", "1"},
		&Spot{"33C", 1048, 146, 32, 162, true, "Available", "1"},

		&Spot{"45", 628, 405, 31, 72, false, "Available", "1"},
		&Spot{"44", 660, 405, 31, 72, false, "Available", "1"},
		&Spot{"43", 700, 405, 31, 72, false, "Available", "1"},
		&Spot{"42", 732, 405, 31, 72, false, "Available", "1"},
		&Spot{"41", 764, 405, 31, 72, false, "Available", "1"},
		&Spot{"40", 804, 405, 31, 72, false, "Available", "1"},
		&Spot{"39", 836, 405, 31, 72, false, "Available", "1"},
		&Spot{"38", 869, 405, 31, 72, false, "Available", "1"},
		&Spot{"37", 908, 405, 31, 72, false, "Available", "1"},
		&Spot{"36", 940, 405, 31, 72, false, "Available", "1"},
		&Spot{"35", 972, 405, 31, 72, false, "Available", "1"},
		&Spot{"34", 1012, 405, 33, 72, true, "Available", "1"},
		&Spot{"34A", 1048, 405, 32, 72, true, "Available", "1"},

		&Spot{"50", 85, 396, 80, 31, false, "Available", "1"},
		&Spot{"49", 190, 396, 80, 31, true, "Available", "1"},
		&Spot{"48", 294, 396, 80, 31, true, "Available", "1"},
		&Spot{"47", 398, 396, 80, 31, false, "Available", "1"},
		&Spot{"46", 502, 396, 80, 31, false, "Available", "1"},
	},
	"2": []*Spot{
		&Spot{"0", 108, 236, 31, 72, false, "Available", "2"},
		&Spot{"1", 140, 236, 31, 72, false, "Occupied", "2"},
	},
}
