func isAnagram(s string, t string) bool {

if len(s) != len(t){
		return false
}
var counts [26]int

for _, value := range s {
	counts[value - 'a']++
}

for _,value := range t {
	counts[value-'a']--
	if counts[value-'a']<0{
		return false
	}
}
return true
}


