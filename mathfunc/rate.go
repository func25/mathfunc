package mathfunc

import (
	"fmt"
	"reflect"
)

type RateUnit struct {
	Value interface{}
	Rate  float64
}

type RarityConfig struct {
	Rates []RateUnit
}

func RandomWithRarities(rateUnits RarityConfig, result interface{}) error {

	var totalRate float64 = 0
	for _, rate := range rateUnits.Rates {
		totalRate += rate.Rate
	}

	randNum := Rand0ToFloat64(0, totalRate)

	var lastValue interface{} = nil
	for _, v := range rateUnits.Rates {
		var percent float64 = v.Rate
		if percent > 0 {
			lastValue = v.Value
			if randNum > percent {
				randNum -= percent
			} else {
				break
			}
		}
	}

	resultsVal := reflect.ValueOf(result)
	if resultsVal.Kind() != reflect.Ptr {
		return fmt.Errorf("results argument must be a pointer to a slice, but was a %s", resultsVal.Kind())
	}

	resultsVal.Elem().Set(reflect.ValueOf(lastValue))
	return nil
}
