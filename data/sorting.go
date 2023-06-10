package data

import (
	"sort"
)

// Takes Batter in pointer of datatype batters, returns sorted batters.
// uses the sorting function to sort the batters.
func SortBatters(batters *Batters) {
	sort.Slice(batters.Batter[:], func(i, j int) bool {
		return batters.Batter[i].ID < batters.Batter[j].ID
	})
}
