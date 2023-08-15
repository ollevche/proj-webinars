package array

func HasArrayItem(source []string, items []string) bool {
	for _, s := range source {
		for _, i := range items {
			if s == i {
				return true
			}
		}
	}
	return false
}
