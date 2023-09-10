package domain

type Pack struct {
	Size int
}

type Packs []Pack

func (p Packs) TotalSize() int {
	var result int
	for _, pack := range p {
		result += pack.Size
	}
	return result
}
