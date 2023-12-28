package handlers

func GetFirst[Type any](slice []Type) Type{
	if len(slice) > 0 {
		return slice[0]
	}
	var t Type
	return t
}