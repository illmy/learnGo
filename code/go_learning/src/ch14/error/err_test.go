package error_test

import (
	"errors"
	"testing"
)

var LessThanTwoError error = errors.New("n must be greater than 2")
var GreaterThanHundredError error = errors.New("n must be less than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, GreaterThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(10); err != nil {
		if err == LessThanTwoError {

		}
		if err == GreaterThanHundredError {

		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
