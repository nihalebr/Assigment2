package data

// Takes Batter in pointer of datatype batters and value in string,
// returns bool that whether the value is present in batters are not.
func HasValueBatterType(batters *Batters, value string) bool {
	for _, values := range batters.Batter {
		if values.Type == value {
			return true
		}
	}
	return false
}
