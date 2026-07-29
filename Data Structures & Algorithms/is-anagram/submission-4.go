func isAnagram(s string, t string) bool {
	lens := len(s)
	lent := len(t)
	if lens == 0 && lent == 0 {
		return true
	}

	if lens <= 0 || lent <= 0{
		return false
	}

	if lens != lent {
		return false
	}

	var seenSymbols map[string]int = make(map[string]int)
	for i:=0; i < lens; i++ { 
		seenSymbols[string(s[i])] = seenSymbols[string(s[i])] + 1
		seenSymbols[string(t[i])] = seenSymbols[string(t[i])] - 1 
	}

	for _, val := range seenSymbols {
		if val != 0 {
			return false
		}
	}

	return true
}
