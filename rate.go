package mathfunc

import (
	"errors"
	"fmt"
	"reflect"
)

type RateUnit struct {
	Value interface{}
	Rate  float64
}

func RandomWithRarities(rateUnits []RateUnit, result interface{}) error {
	if len(rateUnits) <= 0 {
		return errors.New("cannot rate pick the empty array")
	}

	var totalRate float64 = 0
	for _, rate := range rateUnits {
		totalRate += rate.Rate
	}

	randNum := Rand0ToFloat64(0, totalRate)

	var lastValue interface{} = nil
	for _, v := range rateUnits {
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

	if lastValue == nil {
		return errors.New("something wrong with rateUnits array")
	}

	resultsVal := reflect.ValueOf(result)
	if resultsVal.Kind() != reflect.Ptr {
		return fmt.Errorf("results argument must be a pointer to a slice, but was a %s", resultsVal.Kind())
	}

	resultsVal.Elem().Set(reflect.ValueOf(lastValue))
	return nil
}
