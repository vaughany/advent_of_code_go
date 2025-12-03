package main

import (
	"fmt"
	"testing"
)

func TestDay01Mod100(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name string
		n    int
		want int
	}{
		{"zero", 0, 0},
		{"within range", 42, 42},
		{"wrap positive", 150, 50},
		{"wrap negative", -1, 99},
		{"wrap negative big", -201, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day01mod100(tt.n)
			if got != tt.want {
				t.Fatalf("day01mod100(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestDay01FloorDiv(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		n, d int
		want int
	}{
		{250, 100, 2},
		{50, 100, 0},
		{-1, 100, -1},
		{-101, 100, -2},
		{-100, 100, -1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d/%d", tt.n, tt.d), func(t *testing.T) {
			got := cfg.day01floorDiv(tt.n, tt.d)
			if got != tt.want {
				t.Fatalf("day01floorDiv(%d, %d) = %d, want %d", tt.n, tt.d, got, tt.want)
			}
		})
	}
}

func TestDay01CountPasses(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name   string
		lo, hi int
		want   int
	}{
		// 50 -> 300: crosses 100, 200 (300 is excluded by strict (lo, hi))
		{"50_to_300", 50, 300, 2},
		// 300 -> 50: symmetric
		{"300_to_50", 300, 50, 2},
		// 40 -> 0: does not cross 0 in the open interval (0, 40)
		{"40_to_0", 40, 0, 0},
		// 0 -> 10: no passes either
		{"0_to_10", 0, 10, 0},
		// -250 -> 50: crosses -200, -100, 0
		{"-250_to_50", -250, 50, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day01passesBetween(tt.lo, tt.hi)
			if got != tt.want {
				t.Fatalf("day01passesBetween(%d, %d) = %d, want %d", tt.lo, tt.hi, got, tt.want)
			}
		})
	}
}

func TestDay01Part1(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name:         "no_zero_hits",
			instructions: []string{"R10", "R10"},
			want:         0,
		},
		{
			name:         "land_on_zero_once",
			instructions: []string{"L10", "L40"},
			// 50 -> 40 -> 0, only final position is 0
			want: 1,
		},
		{
			name:         "wrap_around_to_zero",
			instructions: []string{"R50"},
			// 50 -> 100 -> wrapped 0
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day01part1(tt.instructions)
			if got != tt.want {
				t.Fatalf("day01part1(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func TestDay01Part2(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name:         "simple_land_only",
			instructions: []string{"L10", "L40"},
			// 50 -> 40 (no hit), 40 -> 0: passes=0, land=1 => total 1
			want: 1,
		},
		{
			name:         "cross_multiple_times_positive",
			instructions: []string{"R250"},
			// 50 -> 300: passes at 100, 200, land at 300 => 3
			want: 3,
		},
		{
			name:         "cross_multiple_times_negative",
			instructions: []string{"L250"},
			// 50 -> -200: passes at 0, -100, land at -200 => 3
			want: 3,
		},
		{
			name:         "mixed_moves",
			instructions: []string{"R10", "L70", "R10"},
			// 50 -> 60 (no hit)
			// 60 -> -10: crosses 0 once (passes=1, land=0)
			// -10 -> 0: land on 0 once (passes=0, land=1)
			// total hits = 2
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day01part2(tt.instructions)
			if got != tt.want {
				t.Fatalf("day01part2(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func BenchmarkDay01Mod100(b *testing.B) {
	cfg := &config{}
	inputs := []int{0, 42, 150, -1, -201, 250, -250, 999, -999}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day01mod100(inputs[i%len(inputs)])
	}
}

func BenchmarkDay01FloorDiv(b *testing.B) {
	cfg := &config{}
	inputs := []struct{ n, d int }{
		{250, 100}, {50, 100}, {-1, 100}, {-101, 100}, {-100, 100},
		{500, 100}, {-500, 100}, {0, 100}, {99, 100}, {-99, 100},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := inputs[i%len(inputs)]
		_ = cfg.day01floorDiv(tt.n, tt.d)
	}
}

func BenchmarkDay01PassesBetween(b *testing.B) {
	cfg := &config{}
	inputs := []struct{ lo, hi int }{
		{50, 300}, {300, 50}, {40, 0}, {0, 10}, {-250, 50},
		{100, 200}, {-100, 100}, {0, 0}, {-500, 500}, {10, 20},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := inputs[i%len(inputs)]
		_ = cfg.day01passesBetween(tt.lo, tt.hi)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"R10", "R10", "L10", "L40", "R50", "R100", "L200", "R25", "L75", "R150",
		"L50", "R200", "L100", "R75", "L25", "R300", "L150", "R50", "L200", "R100",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day01part1(instructions)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"R10", "L70", "R10", "R250", "L250", "R100", "L100", "R50", "L50", "R200",
		"L200", "R150", "L150", "R75", "L75", "R300", "L300", "R25", "L25", "R500",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day01part2(instructions)
	}
}
