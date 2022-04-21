package mathfunc

//BitmaskUInt is just for convenient, if you want to safely use, please use "math/bits"
type BitmaskUInt uint

func (b BitmaskUInt) TurnOn(other BitmaskUInt) BitmaskUInt {
	return b | other
}

func (b BitmaskUInt) TurnOff(other BitmaskUInt) BitmaskUInt {
	return b &^ other
}

func (b BitmaskUInt) Has(other BitmaskUInt) bool {
	return (b & other) == other
}
