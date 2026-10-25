func isAnagram(s string, t string) bool {

if len(s) != len(t){
		return false
}

 sMap := createHash (s)
 tMap := createHash (t)
 for key, value := range sMap {
	val, ok := tMap[key]
	if !ok || val != value {
		return false
	}
 }
	return true
}


func createHash (s string) map[string]int {
	sMap := map[string]int{}

	for _,value := range s {
		sMap[string(value)] = sMap[string(value)] + 1
	}
	return sMap
}