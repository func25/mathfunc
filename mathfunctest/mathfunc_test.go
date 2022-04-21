package mathfunctest

import (
	"fmt"
	"testing"

	"github.com/func25/mathfunc"
)

func TestRandRates(t *testing.T) {
	var x string = ""
	quantity := map[string]int{}
	n := 1000
	for i := 0; i < n; i++ {
		err := mathfunc.RandomWithRarities([]mathfunc.RateUnit{
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
		}, &x)
		if err != nil {
			t.Error(err)
			return
		}
		quantity[x]++
	}

	fmt.Println(quantity)
}

type abc struct {
	a, b, c mathfunc.BitmaskUInt
}

func TestMe(t *testing.T) {
	fmt.Println(1 & -2)
}

func TestBitwise(t *testing.T) {
	turnOnTable := []abc{
		{2, 1, 3},
		{89, 20, 93},
	}
	t.Run("turn on", func(t *testing.T) {
		for _, v := range turnOnTable {
			if v.a.TurnOn(v.b) != v.c {
				t.Error(fmt.Sprintf("turn on failed case: %v", v))
			}
		}
	})

	turnOffTable := []abc{
		{2, 1, 2},
		{2, 2, 0},
		{89, 20, 73},
	}
	t.Run("turn off", func(t *testing.T) {
		for _, v := range turnOffTable {
			if v.a.TurnOff(v.b) != v.c {
				t.Error(fmt.Sprintf("turn on failed case: %v", v))
			}
		}
	})

	hasTable := []abc{
		{2, 1, 0},
		{3, 1, 1},
		{3, 2, 1},
		{45, 36, 1},
		{45, 35, 0},
	}
	t.Run("turn off", func(t *testing.T) {
		for _, v := range hasTable {
			res := true
			if v.c == 0 {
				res = false
			}
			if v.a.Has(v.b) != res {
				t.Error(fmt.Sprintf("turn on failed case: %v", v))
			}
		}
	})
}
