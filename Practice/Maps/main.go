package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["GA"] = "Georgia"
	states["UT"] = "Utah"
	fmt.Println(states)

	Georgia := states["GA"]
	fmt.Println(Georgia)
	delete(states, "GA")
	states["NY"] = "New York"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v %v\n", k, v)
	}

	keys := make([]string, len(states))

	i := 0

	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
