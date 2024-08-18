package utils

// func Contains[T comparable](slice []T, element T) bool {
// 	for _, v := range slice {
// 		if v == element {
// 			return true
// 		}
// 	}
// 	return false
// }

func ContainsID(slice []string, id string) bool {
	for _, item := range slice {
		if item == id {
			return true
		}
	}
	return false
}
