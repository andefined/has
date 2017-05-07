package has

// Helper Functions

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
