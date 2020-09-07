package client_test

import (
	"testing"

	"illmy.com/m/code/go_learning/src/ch15/series"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(10))
}
