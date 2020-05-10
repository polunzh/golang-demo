// map_drinks.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{
		"beer":   "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé"}

	names := make([]string, len(drinks))
	idx := 0
	for k := range drinks {
		names[idx] = k
		idx++
	}

	sort.Strings(names)

	for _, v := range names {
		fmt.Printf("%s: %s\n", v, drinks[v])
	}
}

/* Output:
The following drinks are available:
wine
beer
water
coffee
thea

The french for wine is vin
The french for beer is bière
The french for water is eau
The french for coffee is café
The french for thea is thé

Now the sorted output:
The following sorted drinks are available:
beer
coffee
thea
water
wine

The french for beer is bière
The french for coffee is café
The french for thea is thé
The french for water is eau
The french for wine is vin
*/
