package schemas

type Selected map[string]string // map[identifier] = action

func (selected Selected) Count(search string) int {

	var result int = 0

	for _, value := range selected {

		if search == value {
			result++
		}

	}

	return result

}

func (selected Selected) Filter(search string) {

	for id, value := range selected {

		if search == value {
			// Do Nothing
		} else {
			delete(selected, id)
		}

	}

}

func (selected Selected) Keys(search string) []string {

	var result []string

	for id, value := range selected {

		if search == value {
			result = append(result, id)
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
