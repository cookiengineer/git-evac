package schemas

type Selected map[string]string // map[identifier] = action

func (selected Selected) Copy() Selected {

	result := make(Selected)

	for key, val := range selected {
		result.Set(key, val)
	}

	return result

}

func (selected Selected) Count(search string) int {

	var result int = 0

	for _, value := range selected {

		if search == value {
			result++
		}

	}

	return result

}

func (selected Selected) FilterByKey(search string) {

	for key, _ := range selected {

		if search == key {
			// Do Nothing
		} else {
			delete(selected, key)
		}

	}

}

func (selected Selected) FilterByValue(search string) {

	for key, value := range selected {

		if search == value {
			// Do Nothing
		} else {
			delete(selected, key)
		}

	}

}

func (selected Selected) Keys(search string) []string {

	var result []string

	if search != "" {

		for key, value := range selected {

			if search == value {
				result = append(result, key)
			}

		}

	} else {

		for key, _ := range selected {
			result = append(result, key)
		}

	}

	return result

}

func (selected Selected) Length() int {

	var result int = 0

	for range selected {
		result++
	}

	return result

}

func (selected Selected) Reset() {

	for key, _ := range selected {
		delete(selected, key)
	}

}

func (selected Selected) Set(key string, value string) {
	selected[key] = value
}

func (selected Selected) Get(key string) string {

	var result string

	tmp, ok := selected[key]

	if ok == true {
		result = tmp
	}

	return result

}
