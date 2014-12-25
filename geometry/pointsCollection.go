package geometry

type PointsCollection []Points

func NewPointsCollectoin(psc ...Points) PointsCollection {
	var npc PointsCollection
	count := len(psc)
	npc = make([]Points, count, 10)
	for i := 0; i < count; i++ {
		pcount := len(psc[i])
		npc[i] = make([]Point, pcount, pcount)
		for j := 0; j < pcount; j++ {
			npc[i][j] = psc[i][j].Clone()
		}
	}
	return npc
}

func (p PointsCollection) Clone() PointsCollection {
	var npc PointsCollection
	count := len(p)
	npc = make([]Points, count, 10)
	for i := 0; i < count; i++ {
		pcount := len(p[i])
		npc[i] = make([]Point, pcount, pcount)
		for j := 0; j < pcount; j++ {
			npc[i][j] = p[i][j].Clone()
		}
	}
	return npc
}
