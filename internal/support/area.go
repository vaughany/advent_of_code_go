package support

type area struct {
	x, y, z, xy, yz, zx, area, volume, smallestSide, smallestPerimeter int
}

func (a *area) GetAreaOfBox() {
	a.xy = a.x * a.y
	a.yz = a.y * a.z
	a.zx = a.z * a.x

	a.area = (2 * a.xy) + (2 * a.yz) + (2 * a.zx)
}

func (a *area) GetVolumeOfBox() {
	a.volume = a.x * a.y * a.z
}

func (a *area) GetSmallestSide() {
	var n [3]int

	n[0] = a.xy
	n[1] = a.yz
	n[2] = a.zx

	smallest := n[0]
	for _, num := range n[1:] {
		if num < smallest {
			smallest = num
		}
	}

	a.smallestSide = smallest
}

func (a *area) GetSmallestPerimeter() {
	var n [3]int

	n[0] = (a.x * 2) + (a.y * 2)
	n[1] = (a.y * 2) + (a.z * 2)
	n[2] = (a.z * 2) + (a.x * 2)

	smallest := n[0]
	for _, num := range n[1:] {
		if num < smallest {
			smallest = num
		}
	}

	a.smallestPerimeter = smallest
}
