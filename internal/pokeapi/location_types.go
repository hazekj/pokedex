package pokeapi

type pokeLocation struct {
	Count    int
	Next     string
	Previous any
	Results  []struct {
		Name string
		URL  string
	}
}
