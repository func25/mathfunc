package mathfunctest

import (
	"fmt"
	"testing"

	"github.com/func25/mathfunc/mathfunc"
)

func TestRandRates(t *testing.T) {
	var x string = ""
	quantity := map[string]int{}
	n := 1000
	for i := 0; i < n; i++ {
		err := mathfunc.RandomWithRarities(mathfunc.RarityConfig{
			Rates: []mathfunc.RateUnit{
				{
					Value: "a",
					Rate:  0.1,
				},
				{
					Value: "ab",
					Rate:  0.1,
				},
				{
					Value: "abc",
					Rate:  0.1,
				},
				{
					Value: "abcd",
					Rate:  0.1,
				},
				{
					Value: "abcde",
					Rate:  0.6,
				},
			},
		}, &x)
		if err != nil {
			t.Error(err)
			return
		}
		quantity[x]++
	}

	fmt.Println(quantity)
}
