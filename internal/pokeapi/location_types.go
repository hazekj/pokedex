package pokeapi

type pokeLocation struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
}
