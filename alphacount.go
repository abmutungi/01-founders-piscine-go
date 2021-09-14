package piscine

func AlphaCount(s string) int {
	count := 0
	result := []rune(s)

for i:= 0 ; i <= len(result) ; count++{
	if (result[i] >= 'a' && result[i] <= 'z') || (result[i] >= 'A' && result [i] <= 'Z')
	} 
	return count
}
//capital letters 65 - 90
//lowercase 97-122