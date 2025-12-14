package main

import (
	"slices"
	"strconv"
	"strings"

	"vaughany.com/advent_of_code_go/internal/loaders"
)

func (cfg *config) day08() error {
	return runDayWithInput(
		cfg,
		loaders.GetStrings,
		cfg.day08part1,
		cfg.day08part2,
	)
}

type point struct {
	x, y, z int
}

type edge struct {
	i, j int   // indices into points slice
	d2   int64 // squared distance
}

func parsePoints(instructions []string) []point {
	points := make([]point, 0, len(instructions))

	for _, instruction := range instructions {
		parts := strings.Split(instruction, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		points = append(points, point{x: x, y: y, z: z})
	}

	return points
}

func buildSortedEdges(points []point) []edge {
	n := len(points)
	edges := make([]edge, 0, n*(n-1)/2)

	for i := range n {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			d2 := int64(dx*dx + dy*dy + dz*dz)

			edges = append(edges, edge{i: i, j: j, d2: d2})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		switch {
		case a.d2 < b.d2:
			return -1
		case a.d2 > b.d2:
			return 1
		default:
			return 0
		}
	})

	return edges
}

// 2025-08-1: 67488
func (cfg *config) day08part1(instructions []string) int {
	points := parsePoints(instructions)
	insLen := len(points)
	edges := buildSortedEdges(points)
	dsu := newDSU(len(points))

	// Specific details given for the sample data / main data.
	pairsToConnect := 1000
	if cfg.sample {
		pairsToConnect = 10
	}

	for k := range pairsToConnect {
		e := edges[k]
		dsu.union(e.i, e.j) // Whether or not it merges.
	}

	componentSizes := []int{}
	seen := make(map[int]bool)

	for i := range insLen {
		root := dsu.find(i)
		if !seen[root] {
			seen[root] = true
			componentSizes = append(componentSizes, dsu.size[root])
		}
	}

	slices.SortFunc(componentSizes, func(a, b int) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}

		return 0
	})

	return componentSizes[0] * componentSizes[1] * componentSizes[2]
}

// 2025-08-2: 3767453340
func (cfg *config) day08part2(instructions []string) int {
	points := parsePoints(instructions)
	insLen := len(points)
	edges := buildSortedEdges(points)
	dsu := newDSU(len(points))

	for k := range edges {
		e := edges[k]

		if dsu.union(e.i, e.j) {
			insLen--
			if insLen == 1 {
				return points[e.i].x * points[e.j].x
			}
		}
	}

	return 0
}

// https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type disjointSetUnion struct {
	parent []int
	size   []int
}

func newDSU(n int) *disjointSetUnion {
	parent := make([]int, n)
	size := make([]int, n)

	for i := range n {
		parent[i] = i
		size[i] = 1
	}

	return &disjointSetUnion{parent: parent, size: size}
}

func (d *disjointSetUnion) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x]) // Path compression.
	}

	return d.parent[x]
}

func (d *disjointSetUnion) union(a, b int) bool {
	ra := d.find(a)
	rb := d.find(b)

	if ra == rb {
		return false // Already in the same set, so there's nothing to do.
	}

	// Union by size: attach smaller tree under larger tree.
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}

	d.parent[rb] = ra
	d.size[ra] += d.size[rb]

	return true
}
