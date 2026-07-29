func isAnagram(s string, t string) bool {
	var seenSymbols map[string]int = make(map[string]int)
	if len(s) == 0 && len(t) == 0 {
		return true
	}

	if len(s) <= 0 || len(t) <= 0{
		return false
	}

	if len(s) != len(t) {
		return false
	}

	for i:=0; i < len(s); i++ { 
		firstWordSymbol := s[i]
		secondWordSymbol := t[i]
		seenSymbols[string(firstWordSymbol)] = seenSymbols[string(firstWordSymbol)] + 1
		seenSymbols[string(secondWordSymbol)] = seenSymbols[string(secondWordSymbol)] - 1 
	}

	for _, val := range seenSymbols {
		if val != 0 {
			return false
		}
	}

	return true
}
