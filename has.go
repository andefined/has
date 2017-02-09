package has

// Email ...
func Email(str string) []string {
	elemetns := ReEmail().FindStringSubmatch(str)
	return removeDuplicates(elemetns)
}

// Bitcoin ...
func Bitcoin(str string) []string {
	elemetns := ReBitcoin().FindStringSubmatch(str)
	return removeDuplicates(elemetns)
}

// Helper Function
// removeDuplicates : Remove Duplicates
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}

	return result
}
