package tests

import (
	"context"
	"errors"
	"fmt"
	"test/entity"
	mock "test/mock/mockRepo"
	"testing"

	"github.com/golang/mock/gomock"
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
	var cx context.Context

	testUser := entity.CreateUserInput{
		FirstName: "Aruna",
		LastName:  "Das",
		Email:     "aa",
		Gender:    "MALE",
	}
	for i := 0; i < b.N; i++ {
		ctrl := gomock.NewController(b)
		defer ctrl.Finish()

		c := mock.NewMockUserRepo(ctrl)
		c.EXPECT().GetByName(cx, testUser).Return(errors.New("username already exists"))

		fmt.Println(c)
		err := c.GetByName(cx, testUser)
		if err.Error() != "username already exists" {
			b.Errorf("benchmark test failed %d: %s ", i, err)
		}
	}
}

func TestFactorial(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "N is 0",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "N is 1",
			args: args{n: 3},
			want: 1,
		},
		{
			name: "N is 3",
			args: args{n: 3},
			want: 6,
		},
		{
			name: "N is 10",
			args: args{n: 10},
			want: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Factorial(n uint) uint {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}
