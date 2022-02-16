package mathfunc

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

//RandomInt return a number from min to max - 1
func RandInt(min, max int) (int, error) {
	i, err := Random0ToInt(max - min)
	if err != nil {
		return max, nil
	}
	i += min
	return i, nil
}

//Random0ToInt return a number from 0 to max - 1, return 0 if max == 0 and return error if max's negative
func Random0ToInt(max int) (int, error) {
	if max == 0 {
		return 0, nil
	}
	preRand, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return -1, err
	}
	return int(preRand.Int64()), nil
}

func Rand0ToFloat64(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return (min + rand.Float64()*(max-min))
}
