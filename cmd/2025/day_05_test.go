package main

import (
	"testing"
)

func TestIsFresh(t *testing.T) {
	tests := []struct {
		name       string
		ranges     []Fresh
		ingredient int
		want       bool
	}{
		{
			name:       "empty ranges",
			ranges:     []Fresh{},
			ingredient: 5,
			want:       false,
		},
		{
			name:       "single range match",
			ranges:     []Fresh{{from: 3, to: 5}},
			ingredient: 4,
			want:       true,
		},
		{
			name:       "single range match at start",
			ranges:     []Fresh{{from: 3, to: 5}},
			ingredient: 3,
			want:       true,
		},
		{
			name:       "single range match at end",
			ranges:     []Fresh{{from: 3, to: 5}},
			ingredient: 5,
			want:       true,
		},
		{
			name:       "single range no match below",
			ranges:     []Fresh{{from: 3, to: 5}},
			ingredient: 2,
			want:       false,
		},
		{
			name:       "single range no match above",
			ranges:     []Fresh{{from: 3, to: 5}},
			ingredient: 6,
			want:       false,
		},
		{
			name:       "multiple ranges match first",
			ranges:     []Fresh{{from: 3, to: 5}, {from: 10, to: 14}},
			ingredient: 4,
			want:       true,
		},
		{
			name:       "multiple ranges match second",
			ranges:     []Fresh{{from: 3, to: 5}, {from: 10, to: 14}},
			ingredient: 12,
			want:       true,
		},
		{
			name:       "multiple ranges no match",
			ranges:     []Fresh{{from: 3, to: 5}, {from: 10, to: 14}},
			ingredient: 8,
			want:       false,
		},
		{
			name:       "overlapping ranges match",
			ranges:     []Fresh{{from: 3, to: 5}, {from: 4, to: 6}},
			ingredient: 5,
			want:       true,
		},
		{
			name:       "single value range",
			ranges:     []Fresh{{from: 5, to: 5}},
			ingredient: 5,
			want:       true,
		},
		{
			name:       "single value range no match",
			ranges:     []Fresh{{from: 5, to: 5}},
			ingredient: 4,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isFresh(tt.ranges, tt.ingredient)
			if got != tt.want {
				t.Fatalf("isFresh(%v, %d) = %v, want %v", tt.ranges, tt.ingredient, got, tt.want)
			}
		})
	}
}

func TestDay05Part1(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name:         "empty input",
			instructions: []string{},
			want:         0,
		},
		{
			name:         "only ranges no ingredients",
			instructions: []string{"3-5", "10-14"},
			want:         0,
		},
		{
			name:         "single range single ingredient match",
			instructions: []string{"3-5", "4"},
			want:         1,
		},
		{
			name:         "single range single ingredient no match",
			instructions: []string{"3-5", "8"},
			want:         0,
		},
		{
			name:         "sample input",
			instructions: []string{"3-5", "10-14", "16-20", "12-18", "", "1", "5", "8", "11", "17", "32"},
			want:         3, // 5, 11, 17 are in ranges
		},
		{
			name:         "multiple ranges multiple ingredients",
			instructions: []string{"1-5", "10-15", "3", "7", "12", "20"},
			want:         2, // 3 and 12 are in ranges
		},
		{
			name:         "ingredient at range boundary",
			instructions: []string{"3-5", "5"},
			want:         1,
		},
		{
			name:         "ingredient at range start",
			instructions: []string{"3-5", "3"},
			want:         1,
		},
		{
			name:         "overlapping ranges",
			instructions: []string{"3-5", "4-6", "4"},
			want:         1,
		},
		{
			name:         "empty lines ignored",
			instructions: []string{"3-5", "", "", "4", ""},
			want:         1,
		},
		{
			name:         "whitespace trimmed",
			instructions: []string{"  3-5  ", "  4  "},
			want:         1,
		},
		{
			name:         "all ingredients match",
			instructions: []string{"1-10", "2", "5", "8"},
			want:         3,
		},
		{
			name:         "no ingredients match",
			instructions: []string{"1-5", "10", "20", "30"},
			want:         0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day05part1(tt.instructions)
			if got != tt.want {
				t.Fatalf("day05part1(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func TestDay05Part2(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions []string
		want         int
	}{
		{
			name:         "single range",
			instructions: []string{"3-5"},
			want:         3, // 3, 4, 5 = 3 numbers
		},
		{
			name:         "single value range",
			instructions: []string{"5-5"},
			want:         1, // Just 5
		},
		{
			name:         "non-overlapping ranges",
			instructions: []string{"3-5", "10-12"},
			want:         6, // (5-3+1) + (12-10+1) = 3 + 3 = 6
		},
		{
			name:         "overlapping ranges",
			instructions: []string{"3-5", "4-6"},
			want:         4, // Merged: 3-6 = 4 numbers (3,4,5,6)
		},
		{
			name:         "adjacent ranges",
			instructions: []string{"3-5", "6-8"},
			want:         6, // Merged: 3-8 = 6 numbers (3,4,5,6,7,8)
		},
		{
			name:         "sample input",
			instructions: []string{"3-5", "10-14", "16-20", "12-18"},
			want:         14, // Merged: 3-5 (3 numbers) + 10-20 (11 numbers) = 14
		},
		{
			name:         "nested ranges",
			instructions: []string{"1-10", "3-7"},
			want:         10, // Merged: 1-10 = 10 numbers
		},
		{
			name:         "completely overlapping ranges",
			instructions: []string{"5-10", "5-10"},
			want:         6, // Still just 5-10 = 6 numbers
		},
		{
			name:         "multiple merges needed",
			instructions: []string{"1-3", "2-4", "5-7", "6-8"},
			want:         8, // Merged: 1-4, 5-8 = 4 + 4 = 8
		},
		{
			name:         "unsorted ranges",
			instructions: []string{"10-15", "1-5", "20-25"},
			want:         17, // 1-5 (5 numbers) + 10-15 (6 numbers) + 20-25 (6 numbers) = 17
		},
		{
			name:         "empty lines ignored",
			instructions: []string{"3-5", "", "10-12", ""},
			want:         6,
		},
		{
			name:         "whitespace trimmed",
			instructions: []string{"  3-5  ", "  10-12  "},
			want:         6,
		},
		{
			name:         "ingredients ignored in part2",
			instructions: []string{"3-5", "10-12", "4", "11"},
			want:         6, // Ingredients should be ignored, only ranges count
		},
		{
			name:         "single large range",
			instructions: []string{"1-100"},
			want:         100,
		},
		{
			name:         "chain of overlapping ranges",
			instructions: []string{"1-3", "3-5", "5-7", "7-9"},
			want:         9, // Merged: 1-9 = 9 numbers
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day05part2(tt.instructions)
			if got != tt.want {
				t.Fatalf("day05part2(%v) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func BenchmarkIsFresh(b *testing.B) {
	ranges := []Fresh{
		{from: 3, to: 5},
		{from: 10, to: 14},
		{from: 16, to: 20},
		{from: 12, to: 18},
		{from: 25, to: 30},
		{from: 35, to: 40},
		{from: 50, to: 60},
		{from: 70, to: 80},
	}
	ingredients := []int{1, 4, 8, 12, 17, 25, 32, 50, 75, 100}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isFresh(ranges, ingredients[i%len(ingredients)])
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"25-30",
		"35-40",
		"50-60",
		"70-80",
		"",
		"1",
		"5",
		"8",
		"11",
		"17",
		"25",
		"32",
		"50",
		"75",
		"100",
		"150",
		"200",
		"250",
		"300",
		"350",
		"400",
		"450",
		"500",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day05part1(instructions)
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"3-5",
		"10-14",
		"16-20",
		"12-18",
		"25-30",
		"28-35",
		"50-60",
		"55-65",
		"70-80",
		"75-85",
		"100-110",
		"105-115",
		"150-160",
		"155-165",
		"200-210",
		"205-215",
		"250-260",
		"255-265",
		"300-310",
		"305-315",
		"350-360",
		"355-365",
		"400-410",
		"405-415",
		"450-460",
		"455-465",
		"500-510",
		"505-515",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day05part2(instructions)
	}
}

