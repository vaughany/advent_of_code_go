package main

import (
	"testing"
)

func TestDay03DigitsToInt(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name   string
		digits []int
		want   int64
	}{
		{"empty slice", []int{}, 0},
		{"single digit", []int{5}, 5},
		{"two digits", []int{1, 2}, 12},
		{"three digits", []int{1, 2, 3}, 123},
		{"twelve digits", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, 123456789012},
		{"all zeros", []int{0, 0, 0}, 0},
		{"leading zero", []int{0, 1, 2}, 12},
		{"max digits", []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 999999999999},
		{"mixed", []int{4, 2, 0}, 420},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day03DigitsToInt(tt.digits)
			if got != tt.want {
				t.Fatalf("day03DigitsToInt(%v) = %d, want %d", tt.digits, got, tt.want)
			}
		})
	}
}

func TestDay03Part1(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name:         "empty",
			instructions: []string{},
			want:         0,
		},
		{
			name:         "single digit",
			instructions: []string{"5"},
			want:         0, // Can't form a 2-digit pair from a single digit
		},
		{
			name:         "two digits ascending",
			instructions: []string{"12"},
			want:         12,
		},
		{
			name:         "two digits descending",
			instructions: []string{"21"},
			want:         21,
		},
		{
			name:         "simple max pair",
			instructions: []string{"123"},
			want:         23, // 2*10 + 3 = 23
		},
		{
			name:         "max pair in middle",
			instructions: []string{"12345"},
			want:         45, // 4*10 + 5 = 45
		},
		{
			name:         "max pair not at end",
			instructions: []string{"54321"},
			want:         54, // 5*10 + 4 = 54 (5 before 4)
		},
		{
			name:         "multiple instructions",
			instructions: []string{"12", "34"},
			want:         12 + 34,
		},
		{
			name:         "complex example",
			instructions: []string{"987654321"},
			want:         98, // 9*10 + 8 = 98
		},
		{
			name:         "all same digits",
			instructions: []string{"111"},
			want:         11, // 1*10 + 1 = 11
		},
		{
			name:         "zero in string",
			instructions: []string{"102"},
			want:         12, // Scanning right to left: 2 (best=2), 0 (pair=02=2, best=2), 1 (pair=12, best=2). Max = 12
		},
		{
			name:         "zero first",
			instructions: []string{"012"},
			want:         12, // 1*10 + 2 = 12
		},
		{
			name:         "long string",
			instructions: []string{"123456789"},
			want:         89, // 8*10 + 9 = 89
		},
		{
			name:         "pair requires scanning",
			instructions: []string{"192837465"},
			want:         98, // Scanning right to left finds 9*10+8=98 as max pair
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day03part1(tt.instructions)
			if got != tt.want {
				t.Fatalf("day03part1(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func TestDay03Part2(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int64
	}{
		{
			name:         "empty",
			instructions: []string{},
			want:         0,
		},
		{
			name:         "exactly 12 digits",
			instructions: []string{"123456789012"},
			want:         123456789012,
		},
		{
			name:         "more than 12 digits",
			instructions: []string{"98765432109876543210"},
			want:         989876543210, // Greedily picks digits: 9,8,9,8,7,6,5,4,3,2,1,0
		},
		{
			name:         "greedy selection",
			instructions: []string{"192837465012"},
			want:         192837465012, // Exactly 12 digits, picks all
		},
		{
			name:         "all same digits",
			instructions: []string{"111111111111"},
			want:         111111111111,
		},
		{
			name:         "ascending digits",
			instructions: []string{"01234567890123456789"},
			want:         890123456789, // Picks digits starting from position 8: 8,9,0,1,2,3,4,5,6,7,8,9
		},
		{
			name:         "descending digits",
			instructions: []string{"98765432109876543210"},
			want:         989876543210, // Picks digits: 9,8,9,8,7,6,5,4,3,2,1,0
		},
		{
			name:         "multiple instructions",
			instructions: []string{"123456789012", "345678901234"},
			want:         123456789012 + 345678901234,
		},
		{
			name:         "with zeros",
			instructions: []string{"01234567890123456789"},
			want:         890123456789,
		},
		{
			name:         "complex greedy",
			instructions: []string{"54321098765432109876"},
			want:         987654329876, // Greedily picks highest digits available
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day03part2(tt.instructions)
			if got != tt.want {
				t.Fatalf("day03part2(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func BenchmarkDay03DigitsToInt(b *testing.B) {
	cfg := &config{}
	inputs := [][]int{
		{1, 2, 3},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		{0, 0, 0},
		{4, 2, 0},
		{1},
		{1, 2, 3, 4, 5},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day03DigitsToInt(inputs[i%len(inputs)])
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"123456789",
		"987654321",
		"192837465",
		"54321098765432109876",
		"01234567890123456789",
		"11111111111111111111",
		"12345678901234567890",
		"98765432109876543210",
		"123",
		"45678901234567890123",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day03part1(instructions)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"12345678901234567890",
		"98765432109876543210",
		"19283746501234567890",
		"54321098765432109876",
		"01234567890123456789",
		"11111111111111111111",
		"12345678901234567890",
		"98765432109876543210",
		"123456789012",
		"99999999999999999999",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day03part2(instructions)
	}
}
