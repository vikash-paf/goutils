package str

// Jaro returns the Jaro similarity between two strings, ranging from 0.0 (no match) to 1.0 (exact match).
// It accounts for matching characters and transpositions within a dynamic window.
func Jaro(s1, s2 string) float64 {
	r1, r2 := []rune(s1), []rune(s2)
	len1, len2 := len(r1), len(r2)

	if len1 == 0 && len2 == 0 {
		return 1.0
	}
	if len1 == 0 || len2 == 0 {
		return 0.0
	}

	matchDistance := max(len1, len2)/2 - 1
	if matchDistance < 0 {
		matchDistance = 0
	}

	s1Matches := make([]bool, len1)
	s2Matches := make([]bool, len2)

	matches := 0
	for i := 0; i < len1; i++ {
		start := max(0, i-matchDistance)
		end := min(i+matchDistance+1, len2)

		for j := start; j < end; j++ {
			if s2Matches[j] {
				continue
			}
			if r1[i] == r2[j] {
				s1Matches[i] = true
				s2Matches[j] = true
				matches++
				break
			}
		}
	}

	if matches == 0 {
		return 0.0
	}

	transpositions := 0
	k := 0
	for i := 0; i < len1; i++ {
		if !s1Matches[i] {
			continue
		}
		for j := k; j < len2; j++ {
			k++
			if s2Matches[j] {
				if r1[i] != r2[j] {
					transpositions++
				}
				break
			}
		}
	}

	m := float64(matches)
	t := float64(transpositions) / 2.0

	return (m/float64(len1) + m/float64(len2) + (m-t)/m) / 3.0
}

// JaroWinkler returns the Jaro-Winkler similarity between two strings.
// It builds on the Jaro similarity by giving a higher score to strings that match
// from the beginning for up to 4 characters.
func JaroWinkler(s1, s2 string) float64 {
	jaroDist := Jaro(s1, s2)

	// The Winkler modification is typically only applied if the Jaro similarity is above a certain threshold (0.7).
	if jaroDist < 0.7 {
		return jaroDist
	}

	r1, r2 := []rune(s1), []rune(s2)
	prefixLen := 0
	maxPrefix := min(min(len(r1), len(r2)), 4)

	for i := 0; i < maxPrefix; i++ {
		if r1[i] == r2[i] {
			prefixLen++
		} else {
			break
		}
	}

	// 0.1 is the standard scaling factor.
	return jaroDist + (float64(prefixLen) * 0.1 * (1.0 - jaroDist))
}
