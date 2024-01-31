package tests

import (
	"test/domain"
	"test/repository"
	"testing"
)

func EfficientSum(a, b int) int {
	return a + b
}

func InefficientSum(a, b int) int {
	res := make(chan int, 1)
	res <- a + b
	return <-res
}
func BenchmarkEfficientSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EfficientSum(i, i+1)
	}

}
func BenchmarkInEfficientSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		InefficientSum(i, i+1)
	}
}

func BenchmarkGetByName(b *testing.B) {
	b.ReportAllocs()
	testUser := domain.User{
		FirstName: "Aruna",
	}
	for i := 0; i < b.N; i++ {
		err := repository.GetByName(testUser)
		if err.Error() != "username already exists" {
			b.Errorf("benchmark test failed %d: %s ", i, err)
		}
	}
}
