package model

type DestinationsPool struct {
	destinations []*Destination
	current      int
	total_weight int
}

func (dp DestinationsPool) Get() *Destination {

	destination := dp.destinations[dp.current]

	if dp.current >= len(dp.destinations)-1 {
		dp.current = 0
	} else {
		dp.current++
	}

	return destination
}
