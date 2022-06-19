package session

const (
	DirectionNorth Direction = iota
	DirectionSouth
	DirectionEast
	DirectionWest
	DirectionNorthWest
	DirectionNorthEast
	DirectionSouthWest
	DirectionSouthEast
)

type Direction int

func (d Direction) Invert() Direction {
	switch d {
	case DirectionNorth:
		return DirectionSouth
	case DirectionSouth:
		return DirectionNorth
	case DirectionWest:
		return DirectionEast
	case DirectionEast:
		return DirectionWest
	}
	return d
}
