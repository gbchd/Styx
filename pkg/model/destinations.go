package model

// Struct that define the pool of destination
// Current is where the Round Robin is
// Total_weight is the total of weight in the destination pool
type DestinationsPool struct {
	Destinations []*Destination
	Current      int
	Total_weight int
}

// Function that return the current destination in the Round Robin and select the next destination
func (dp *DestinationsPool) Get() *Destination {

	destination := dp.Destinations[dp.Current]

	if dp.Current >= len(dp.Destinations)-1 {
		dp.Current = 0
	} else {
		dp.Current++
	}

	return destination
}
