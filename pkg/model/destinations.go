package model

type DestinationsPool struct {
	Destinations []*Destination
	Current      int
	Total_weight int
}

func (dp DestinationsPool) Get() *Destination {

	destination := dp.Destinations[dp.Current]

	if dp.Current >= len(dp.Destinations)-1 {
		dp.Current = 0
	} else {
		dp.Current++
	}

	return destination
}
