package handlers

func Hello(s string) string {
	return "Hey, Myself " + s
}
func Add(a, b int) int {
	return a + b
}
func Repeat(char string) string {
	var repeat string
	for i := 0; i < 5; i++ {
		repeat += char
	}
	return repeat
}
