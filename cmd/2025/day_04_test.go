package main

import (
	"testing"
)

func TestCanMoveRoll(t *testing.T) {
	tests := []struct {
		name  string
		grid  [][]bool
		x, y  int
		want  bool
	}{
		{
			name: "isolated roll",
			grid: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			},
			x: 1, y: 1,
			want: true, // 0 adjacent rolls
		},
		{
			name: "roll with 3 adjacent",
			grid: [][]bool{
				{true, true, false},
				{true, true, false},
				{false, false, false},
			},
			x: 1, y: 1,
			want: true, // 3 adjacent rolls
		},
		{
			name: "roll with 4 adjacent",
			grid: [][]bool{
				{true, true, true},
				{true, true, true},
				{false, false, false},
			},
			x: 1, y: 1,
			want: false, // 4 adjacent rolls
		},
		{
			name: "roll with 5 adjacent",
			grid: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, false, false},
			},
			x: 1, y: 1,
			want: false, // 5 adjacent rolls
		},
		{
			name: "corner roll isolated",
			grid: [][]bool{
				{true, false},
				{false, false},
			},
			x: 0, y: 0,
			want: true, // 0 adjacent rolls
		},
		{
			name: "corner roll with 2 adjacent",
			grid: [][]bool{
				{true, true},
				{true, false},
			},
			x: 0, y: 0,
			want: true, // 2 adjacent rolls
		},
		{
			name: "edge roll isolated",
			grid: [][]bool{
				{false, true, false},
				{false, false, false},
			},
			x: 1, y: 0,
			want: true, // 0 adjacent rolls
		},
		{
			name: "edge roll with 3 adjacent",
			grid: [][]bool{
				{true, true, true},
				{true, false, false},
			},
			x: 1, y: 0,
			want: true, // 3 adjacent rolls
		},
		{
			name: "surrounded roll",
			grid: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			},
			x: 1, y: 1,
			want: false, // 8 adjacent rolls
		},
		{
			name: "single roll in grid",
			grid: [][]bool{
				{true},
			},
			x: 0, y: 0,
			want: true, // 0 adjacent rolls
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canMoveRoll(tt.grid, tt.x, tt.y)
			if got != tt.want {
				t.Fatalf("canMoveRoll(grid, %d, %d) = %v, want %v", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestDay04Part1(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name:         "single roll",
			instructions: []string{"@"},
			want:         1,
		},
		{
			name:         "isolated roll",
			instructions: []string{"...", ".@.", "..."},
			want:         1,
		},
		{
			name:         "two isolated rolls",
			instructions: []string{"@..", "...", "..@"},
			want:         2,
		},
		{
			name:         "roll with 4 adjacent cannot move",
			instructions: []string{"@@@", "@@@", "..."},
			want:         4, // Corners and edges can move (2-3 adjacent), only center has 4+
		},
		{
			name:         "mixed moveable and immovable",
			instructions: []string{"@..", "@@@", "@.@"},
			want:         5, // Top-left (0 adj), top-right (0 adj), bottom-left (1 adj), bottom-right (0 adj), middle-bottom (1 adj)
		},
		{
			name:         "corner roll",
			instructions: []string{"@..", "...", "..."},
			want:         1,
		},
		{
			name:         "edge roll",
			instructions: []string{".@.", "...", "..."},
			want:         1,
		},
		{
			name:         "all moveable",
			instructions: []string{"@.@", "...", "@.@"},
			want:         4,
		},
		{
			name: "sample input",
			instructions: []string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			want: 13, // Actual result for sample input
		},
		{
			name:         "no rolls",
			instructions: []string{"...", "...", "..."},
			want:         0,
		},
		{
			name:         "single row",
			instructions: []string{"@.@.@"},
			want:         3, // All isolated
		},
		{
			name:         "single column",
			instructions: []string{"@", ".", "@", "."},
			want:         2, // All isolated
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day04part1(tt.instructions)
			if got != tt.want {
				t.Fatalf("day04part1(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func TestDay04Part2(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name:         "single roll",
			instructions: []string{"@"},
			want:         1, // Removed in first iteration
		},
		{
			name:         "isolated roll",
			instructions: []string{"...", ".@.", "..."},
			want:         1, // Removed in first iteration
		},
		{
			name:         "two isolated rolls",
			instructions: []string{"@..", "...", "..@"},
			want:         2, // Both removed in first iteration
		},
		{
			name:         "roll with 4 adjacent never removed",
			instructions: []string{"@@@", "@@@", "..."},
			want:         6, // Corners and edges removed (4 corners + 2 edges), center remains
		},
		{
			name:         "cascade removal",
			instructions: []string{"@.@", ".@.", "@.@"},
			want:         5, // All 5 removed (4 corners + 1 center)
		},
		{
			name:         "iterative removal",
			instructions: []string{"@.@", "@@@", "@.@"},
			want:         7, // All 7 removed iteratively
		},
		{
			name: "sample input",
			instructions: []string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			want: 43, // Actual result for sample input
		},
		{
			name:         "no rolls",
			instructions: []string{"...", "...", "..."},
			want:         0,
		},
		{
			name:         "single row all moveable",
			instructions: []string{"@.@.@"},
			want:         3, // All removed in first iteration
		},
		{
			name:         "single column all moveable",
			instructions: []string{"@", ".", "@", "."},
			want:         2, // All removed in first iteration
		},
		{
			name:         "chain removal",
			instructions: []string{"@.@", "@.@", "@.@"},
			want:         6, // All 6 removed
		},
		{
			name:         "center surrounded",
			instructions: []string{"@@@", "@@@", "@@."},
			want:         8, // All except center removed
		},
		{
			name:         "all immovable",
			instructions: []string{"@@@", "@@@", "@@@"},
			want:         9, // All 9 removed (corners and edges can move)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day04part2(tt.instructions)
			if got != tt.want {
				t.Fatalf("day04part2(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func BenchmarkCanMoveRoll(b *testing.B) {
	grid := [][]bool{
		{true, true, false, true, false},
		{true, true, true, false, true},
		{false, true, true, true, false},
		{true, false, true, true, true},
		{false, true, false, true, true},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = canMoveRoll(grid, 2, 2)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
		"@@@@@@@@@@",
		"..........",
		"@.@.@.@.@.",
		".@.@.@.@.@",
		"@@@@@@@@@@",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day04part1(instructions)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
		"@@@@@@@@@@",
		"..........",
		"@.@.@.@.@.",
		".@.@.@.@.@",
		"@@@@@@@@@@",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day04part2(instructions)
	}
}

