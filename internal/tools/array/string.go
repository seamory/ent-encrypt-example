package array

type Strings []string

func (x Strings) Includes(element string) bool {
	for _, el := range x {
		if el == element {
			return true
		}
	}
	return false
}

func (x Strings) Contains(elements []string) Strings {
	var data []string
	for _, s := range x {
		for _, role := range elements {
			if s == role {
				data = append(data, s)
			}
		}
	}
	return data
}
