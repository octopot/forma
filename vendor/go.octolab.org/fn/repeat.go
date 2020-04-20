package fn

// Repeat repeats the action the required number of times.
//
//  func FillByValue(slice []int, value int) {
//  	fn.Repeat(
//  		func () { slice = append(slice, value) },
//  		cap(slice) - len(slice),
//  	)
//  }
//
func Repeat(action func(), times int) {
	for i := 0; i < times; i++ {
		action()
	}
}
