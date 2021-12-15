/*zk virtual device*/
package zvdev

type Vessel struct {
	devices *zmap.IntMap
}

func NewVessel() *Vessel {
	return &Vessel{devices: zmap.NewIntMap()}
}
