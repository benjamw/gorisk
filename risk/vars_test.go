package risk

import (
	"reflect"
	"testing"
)

// TestAdjacent makes sure that each territories Adjacent values is also adjacent
// back to the original territory. All paths are two-way.
func TestAdjacent(t *testing.T) {
	for k, v := range Territories {
		if k == 0 {
			continue
		}

		if len(v.Adjacent) == 0 {
			t.Fatalf("%s (%d) does not have any adjacent territories", v.Name, k)
		}

		for _, vv := range v.Adjacent {
			var found bool

			for _, vvv := range vv.Adjacent {
				if v.Name == vvv.Name {
					found = true
					break
				}
			}

			if !found {
				t.Fatalf("%s (%d) is adjacent to %s, but not the other way around", v.Name, k, vv.Name)
			}
		}
	}
}

// TestContinent makes sure that each territory is part of a continent, and that
// the territory is in the continents list of territories
func TestContinent(t *testing.T) {
	for k, v := range Continents {
		if k == 0 {
			continue
		}

		if len(v.Territories) == 0 {
			t.Fatalf("%s (%d) does not have a list of territories", v.Name, k)
		}

		for _, vv := range v.Territories {
			if vv.Continent.Name != v.Name {
				t.Fatalf("%s (%d) has territory %s, but %s has continent %s", v.Name, k, vv.Name, vv.Name, vv.Continent.Name)
			}
		}
	}
}

// TestTerritories makes sure that each territory has a continent, and that
// the continent has that territory in it's territory list
func TestTerritories(t *testing.T) {
	for k, v := range Territories {
		if k == 0 {
			continue
		}

		if v.Continent.Name == "" {
			t.Fatalf("%s (%d) does not have a continent", v.Name, k)
		}

		var found bool
		for _, vv := range v.Continent.Territories {
			if vv.Continent.Name == v.Continent.Name {
				found = true
			}
		}

		if !found {
			t.Fatalf("%s (%d) has continent %s, but %s does not have %s in it's list of territories", v.Name, k, v.Continent.Name, v.Continent.Name, v.Name)
		}
	}
}

// TestBlank makes sure that each of the main vars starts at index 1
func TestBlank(t *testing.T) {
	blankCont := Continent{}
	blankTerr := Territory{}
	blankCard := Card{}

	if !reflect.DeepEqual(blankCont, Continents[0]) {
		t.Fatal("The first continent is not empty")
	}

	if !reflect.DeepEqual(blankTerr, Territories[0]) {
		t.Fatal("The first territory is not empty")
	}

	if !reflect.DeepEqual(blankCard, Cards[0]) {
		t.Fatal("The first card is not empty")
	}
}
