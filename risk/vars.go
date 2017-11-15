package risk

const (
	WILD = iota
	INFANTRY
	CAVALRY
	ARTILLERY
)

type Continent struct {
	Name        string
	Bonus       int
	Territories []*Territory
}

type Territory struct {
	Name      string
	ShortName string
	Continent *Continent
	Adjacent  []*Territory
	Card      Card
}

type Card struct {
	Territory *Territory
	Type      int
}

func (c *Card) TypeName() string {
	switch c.Type {
	case WILD:
		return "Wild"
	case INFANTRY:
		return "Infantry"
	case CAVALRY:
		return "Cavalry"
	case ARTILLERY:
		return "Artillery"
	}

	return "ERROR"
}

var (
	Continents = []Continent{
		{}, // 1-index for the humans

		{
			Name:        "North America",
			Bonus:       5,
			Territories: make([]*Territory, 0),
		},
		{
			Name:        "South America",
			Bonus:       2,
			Territories: make([]*Territory, 0),
		},
		{
			Name:        "Europe",
			Bonus:       5,
			Territories: make([]*Territory, 0),
		},
		{
			Name:        "Africa",
			Bonus:       3,
			Territories: make([]*Territory, 0),
		},
		{
			Name:        "Asia",
			Bonus:       7,
			Territories: make([]*Territory, 0),
		},
		{
			Name:        "Australia",
			Bonus:       2,
			Territories: make([]*Territory, 0),
		},
	}

	Territories = []Territory{
		{}, // 1-index for the humans

		// North America
		{Name: "Alaska"},
		{Name: "Alberta"},
		{Name: "Central America"},
		{Name: "Eastern United States"},
		{Name: "Greenland"},
		{Name: "Northwest Territory"},
		{Name: "Ontario"},
		{Name: "Quebec"},
		{Name: "Western United States"},

		// South America
		{Name: "Argentina"},
		{Name: "Brazil"},
		{Name: "Peru"},
		{Name: "Venezuela"},

		// Europe
		{Name: "Great Britain"}, // officially 'Great Britain and Ireland', but it's too long
		{Name: "Iceland"},
		{Name: "Northern Europe"},
		{Name: "Scandinavia"},
		{Name: "Southern Europe"},
		{Name: "Ukraine"},
		{Name: "Western Europe"},

		// Africa
		{Name: "Congo"},
		{Name: "East Africa"},
		{Name: "Egypt"},
		{Name: "Madagascar"},
		{Name: "North Africa"},
		{Name: "South Africa"},

		// Asia
		{Name: "Afghanistan"},
		{Name: "China"},
		{Name: "India"},
		{Name: "Irkutsk"},
		{Name: "Japan"},
		{Name: "Kamchatka"},
		{Name: "Middle East"},
		{Name: "Mongolia"},
		{Name: "Siam"},
		{Name: "Siberia"},
		{Name: "Ural"},
		{Name: "Yakutsk"},

		// Australia
		{Name: "Eastern Australia"},
		{Name: "Indonesia"},
		{Name: "New Guinea"},
		{Name: "Western Australia"},
	}

	Cards = []Card{
		{}, // 1-index for the humans

		{Type: INFANTRY}, // 1
		{Type: CAVALRY},
		{Type: ARTILLERY},
		{Type: ARTILLERY},
		{Type: CAVALRY}, // 5
		{Type: ARTILLERY},
		{Type: CAVALRY},
		{Type: CAVALRY},
		{Type: ARTILLERY},
		{Type: INFANTRY}, // 10
		{Type: ARTILLERY},
		{Type: INFANTRY},
		{Type: INFANTRY},
		{Type: ARTILLERY},
		{Type: INFANTRY}, // 15
		{Type: ARTILLERY},
		{Type: CAVALRY},
		{Type: ARTILLERY},
		{Type: CAVALRY},
		{Type: ARTILLERY}, // 20
		{Type: INFANTRY},
		{Type: INFANTRY},
		{Type: INFANTRY},
		{Type: CAVALRY},
		{Type: CAVALRY}, // 25
		{Type: ARTILLERY},
		{Type: CAVALRY},
		{Type: INFANTRY},
		{Type: ARTILLERY},
		{Type: CAVALRY}, // 30
		{Type: CAVALRY},
		{Type: INFANTRY},
		{Type: INFANTRY},
		{Type: INFANTRY},
		{Type: INFANTRY}, // 35
		{Type: CAVALRY},
		{Type: CAVALRY},
		{Type: CAVALRY},
		{Type: ARTILLERY},
		{Type: ARTILLERY}, // 40
		{Type: INFANTRY},
		{Type: ARTILLERY},

		{Type: WILD},
		{Type: WILD},
	}
)

// init sets all the territories to the continents and adjacent territories
// because they can't be set in the var declaration and still be references
// to each other
func init() {
	// don't use v (or the value) here, as range creates a copy and
	// causes &v to hold the memory location of v, not the original value
	for k := range Territories {
		if k == 0 {
			continue
		}

		var cont int
		if 1 <= k && k <= 9 {
			cont = 1 // North America
		} else if 10 <= k && k <= 13 {
			cont = 2 // South America
		} else if 14 <= k && k <= 20 {
			cont = 3 // Europe
		} else if 21 <= k && k <= 26 {
			cont = 4 // Africa
		} else if 27 <= k && k <= 38 {
			cont = 5 // Asia
		} else if 39 <= k && k <= 42 {
			cont = 6 // Australia
		}

		Territories[k].ShortName = shortenName(Territories[k].Name)

		// set references to each other
		Territories[k].Continent = &Continents[cont]

		Continents[cont].Territories = append(Continents[cont].Territories, &Territories[k])

		Cards[k].Territory = &Territories[k]

		Territories[k].Card = Cards[k]
	}

	// this part is a bit complicated, so do it manually...
	// North America
	Territories[1].Adjacent = []*Territory{&Territories[2], &Territories[6], &Territories[32]}
	Territories[2].Adjacent = []*Territory{&Territories[1], &Territories[6], &Territories[7], &Territories[9]}
	Territories[3].Adjacent = []*Territory{&Territories[4], &Territories[9], &Territories[13]}
	Territories[4].Adjacent = []*Territory{&Territories[3], &Territories[7], &Territories[8], &Territories[9]}
	Territories[5].Adjacent = []*Territory{&Territories[6], &Territories[7], &Territories[8], &Territories[15]}
	Territories[6].Adjacent = []*Territory{&Territories[1], &Territories[2], &Territories[5], &Territories[7]}
	Territories[7].Adjacent = []*Territory{&Territories[2], &Territories[4], &Territories[5], &Territories[6], &Territories[8], &Territories[9]}
	Territories[8].Adjacent = []*Territory{&Territories[4], &Territories[5], &Territories[7]}
	Territories[9].Adjacent = []*Territory{&Territories[2], &Territories[3], &Territories[4], &Territories[7]}

	// South America
	Territories[10].Adjacent = []*Territory{&Territories[11], &Territories[12]}
	Territories[11].Adjacent = []*Territory{&Territories[10], &Territories[12], &Territories[13], &Territories[25]}
	Territories[12].Adjacent = []*Territory{&Territories[10], &Territories[11], &Territories[13]}
	Territories[13].Adjacent = []*Territory{&Territories[3], &Territories[11], &Territories[12]}

	// Europe
	Territories[14].Adjacent = []*Territory{&Territories[15], &Territories[16], &Territories[17], &Territories[20]}
	Territories[15].Adjacent = []*Territory{&Territories[5], &Territories[14], &Territories[17]}
	Territories[16].Adjacent = []*Territory{&Territories[14], &Territories[17], &Territories[18], &Territories[19], &Territories[20]}
	Territories[17].Adjacent = []*Territory{&Territories[14], &Territories[15], &Territories[16], &Territories[19]}
	Territories[18].Adjacent = []*Territory{&Territories[16], &Territories[19], &Territories[20], &Territories[23], &Territories[25], &Territories[33]}
	Territories[19].Adjacent = []*Territory{&Territories[16], &Territories[17], &Territories[18], &Territories[27], &Territories[33], &Territories[37]}
	Territories[20].Adjacent = []*Territory{&Territories[14], &Territories[16], &Territories[18], &Territories[25]}

	// Africa
	Territories[21].Adjacent = []*Territory{&Territories[22], &Territories[25], &Territories[26]}
	Territories[22].Adjacent = []*Territory{&Territories[21], &Territories[23], &Territories[24], &Territories[25], &Territories[26], &Territories[33]}
	Territories[23].Adjacent = []*Territory{&Territories[18], &Territories[22], &Territories[25], &Territories[33]}
	Territories[24].Adjacent = []*Territory{&Territories[22], &Territories[26]}
	Territories[25].Adjacent = []*Territory{&Territories[11], &Territories[18], &Territories[20], &Territories[21], &Territories[22], &Territories[23]}
	Territories[26].Adjacent = []*Territory{&Territories[21], &Territories[22], &Territories[24]}

	// Asia
	Territories[27].Adjacent = []*Territory{&Territories[19], &Territories[28], &Territories[29], &Territories[33], &Territories[37]}
	Territories[28].Adjacent = []*Territory{&Territories[27], &Territories[29], &Territories[34], &Territories[35], &Territories[36], &Territories[37]}
	Territories[29].Adjacent = []*Territory{&Territories[27], &Territories[28], &Territories[33], &Territories[35]}
	Territories[30].Adjacent = []*Territory{&Territories[32], &Territories[34], &Territories[36], &Territories[38]}
	Territories[31].Adjacent = []*Territory{&Territories[32], &Territories[34]}
	Territories[32].Adjacent = []*Territory{&Territories[1], &Territories[30], &Territories[31], &Territories[34], &Territories[38]}
	Territories[33].Adjacent = []*Territory{&Territories[18], &Territories[19], &Territories[22], &Territories[23], &Territories[27], &Territories[29]}
	Territories[34].Adjacent = []*Territory{&Territories[28], &Territories[30], &Territories[31], &Territories[32], &Territories[36]}
	Territories[35].Adjacent = []*Territory{&Territories[28], &Territories[29], &Territories[40]}
	Territories[36].Adjacent = []*Territory{&Territories[28], &Territories[30], &Territories[34], &Territories[37], &Territories[38]}
	Territories[37].Adjacent = []*Territory{&Territories[19], &Territories[27], &Territories[28], &Territories[36]}
	Territories[38].Adjacent = []*Territory{&Territories[30], &Territories[32], &Territories[36]}

	// Australia
	Territories[39].Adjacent = []*Territory{&Territories[41], &Territories[42]}
	Territories[40].Adjacent = []*Territory{&Territories[35], &Territories[41], &Territories[42]}
	Territories[41].Adjacent = []*Territory{&Territories[39], &Territories[40], &Territories[42]}
	Territories[42].Adjacent = []*Territory{&Territories[39], &Territories[40], &Territories[41]}
}

func shortenName(s string) string {
	switch s {
	case "Central America":
		return "Cent. America"
	case "Eastern United States":
		return "Eastern U.S."
	case "Northwest Territory":
		return "N.W. Territory"
	case "Western United States":
		return "Western U.S."
	case "Northern Europe":
		return "N. Europe"
	case "Southern Europe":
		return "S. Europe"
	case "Western Europe":
		return "W. Europe"
	case "East Africa":
		return "E. Africa"
	case "North Africa":
		return "N. Africa"
	case "South Africa":
		return "S. Africa"
	case "Middle East":
		return "Mid. East"
	case "Eastern Australia":
		return "E. Australia"
	case "Western Australia":
		return "W. Australia"
	}

	return s
}
