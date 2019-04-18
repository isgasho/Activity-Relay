package activity

func contains(entries interface{}, finder string) bool {
	switch entry := entries.(type) {
	case string:
		return entry == finder
	case []string:
		for i := 0; i < len(entry); i++ {
			if entry[i] == finder {
				return true
			}
		}
		return false
	}
	return false
}
