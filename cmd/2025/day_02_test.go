package main

import (
	"testing"
)

func TestDay02CheckInvalid(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name string
		in   int
		want int
	}{
		{"valid single digit", 1, 0},
		{"valid two different digits", 12, 0},
		{"invalid two same digits", 55, 55},
		{"invalid repeated twice", 6464, 6464},
		{"invalid repeated twice longer", 123123, 123123},
		{"valid odd length", 123, 0},
		{"valid odd length longer", 12345, 0},
		{"valid even but not repeated", 1234, 0},
		{"invalid four digits", 1111, 1111},
		{"invalid six digits", 123456123456, 123456123456},
		{"valid zero", 0, 0},
		{"valid single digit zero", 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day02CheckInvalid(tt.in)
			if got != tt.want {
				t.Fatalf("day02CheckInvalid(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestDay02CheckInvalid2(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name string
		in   int
		want int
	}{
		{"valid single digit", 1, 0},
		{"valid two different digits", 12, 0},
		{"invalid all same", 1111111, 1111111},
		{"invalid all same two", 22, 22},
		{"invalid repeated twice", 12341234, 12341234},
		{"invalid repeated three times", 123123123, 123123123},
		{"invalid repeated five times", 1212121212, 1212121212},
		{"valid odd length", 123, 0},
		{"valid not repeated", 1234, 0},
		{"invalid length 2", 11, 11},
		{"invalid length 3", 111, 111},
		{"invalid length 4 block 2", 1212, 1212},
		{"invalid length 4 block 4", 1111, 1111},
		{"invalid length 6 block 2", 123123, 123123},
		{"invalid length 6 block 3", 121212, 121212},
		{"invalid length 8 block 2", 12341234, 12341234},
		{"invalid length 8 block 4", 12121212, 12121212},
		{"invalid length 8 block 8", 11111111, 11111111},
		{"invalid length 9 block 3", 123123123, 123123123},
		{"invalid length 9 block 9", 111111111, 111111111},
		{"invalid length 10 block 2", 1234512345, 1234512345},
		{"invalid length 10 block 5", 1212121212, 1212121212},
		{"invalid length 10 block 10", 1111111111, 1111111111},
		{"valid length 11", 12345678901, 0},
		{"valid length 1", 5, 0},
		{"valid zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day02CheckInvalid2(tt.in)
			if got != tt.want {
				t.Fatalf("day02CheckInvalid2(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestAllSame(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"empty string", "", false},
		{"single char", "a", true},
		{"all same", "aaa", true},
		{"all same digits", "1111111", true},
		{"different chars", "abc", false},
		{"different digits", "123", false},
		{"mixed", "aab", false},
		{"two same", "aa", true},
		{"two different", "ab", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := allSame(tt.in)
			if got != tt.want {
				t.Fatalf("allSame(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestRepeatedBlocks(t *testing.T) {
	tests := []struct {
		name   string
		in     string
		blocks int
		want   bool
	}{
		{"empty string", "", 2, true}, // Empty string can be split into empty chunks
		{"two blocks same", "abab", 2, true},
		{"two blocks different", "abcd", 2, false},
		{"three blocks same", "abcabcabc", 3, true},
		{"three blocks different", "abcabcabd", 3, false},
		{"five blocks same", "1212121212", 5, true},
		{"five blocks different", "1212121213", 5, false},
		{"invalid blocks zero", "ab", 0, false},
		{"invalid blocks negative", "ab", -1, false},
		{"not divisible", "abc", 2, false},
		{"single block", "abc", 1, true},
		{"all same two blocks", "1111", 2, true},
		{"all same four blocks", "1111", 4, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repeatedBlocks(tt.in, tt.blocks)
			if got != tt.want {
				t.Fatalf("repeatedBlocks(%q, %d) = %v, want %v", tt.in, tt.blocks, got, tt.want)
			}
		})
	}
}

func TestDay02Part1(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions string
		want         int
	}{
		{
			name:         "single range no invalid",
			instructions: "1-5",
			want:         0,
		},
		{
			name:         "single range with invalid",
			instructions: "50-55",
			want:         55, // 55 is invalid (repeated twice)
		},
		{
			name:         "multiple ranges",
			instructions: "50-55,100-105",
			want:         55, // Only 55 is invalid in these ranges
		},
		{
			name:         "range with invalid",
			instructions: "10-15",
			want:         11, // 11 is invalid (repeated twice)
		},
		{
			name:         "single number range",
			instructions: "55-55",
			want:         55,
		},
		{
			name:         "range with 6464",
			instructions: "6460-6470",
			want:         6464, // 6464 is invalid (repeated twice)
		},
		{
			name:         "range with 123123",
			instructions: "123120-123125",
			want:         123123, // 123123 is invalid (repeated twice)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day02part1(tt.instructions)
			if got != tt.want {
				t.Fatalf("day02part1(%q) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func TestDay02Part2(t *testing.T) {
	cfg := &config{}

	tests := []struct {
		name         string
		instructions string
		want         int
	}{
		{
			name:         "single range no invalid",
			instructions: "1-5",
			want:         0,
		},
		{
			name:         "single range with invalid",
			instructions: "10-15",
			want:         11, // 11 is invalid (all same)
		},
		{
			name:         "range with all same",
			instructions: "111-111",
			want:         111, // 111 is invalid (all same)
		},
		{
			name:         "range with repeated blocks",
			instructions: "1210-1215",
			want:         1212, // 1212 is invalid (repeated twice: "12"+"12")
		},
		{
			name:         "multiple ranges",
			instructions: "10-15,111-111",
			want:         11 + 111, // 11 (all same) and 111 (all same)
		},
		{
			name:         "single number range",
			instructions: "111-111",
			want:         111,
		},
		{
			name:         "range with 12341234",
			instructions: "12341230-12341240",
			want:         12341234, // 12341234 is invalid (repeated twice)
		},
		{
			name:         "range with 123123123",
			instructions: "123123120-123123125",
			want:         123123123, // 123123123 is invalid (repeated three times)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cfg.day02part2(tt.instructions)
			if got != tt.want {
				t.Fatalf("day02part2(%q) = %d, want %d", tt.instructions, got, tt.want)
			}
		})
	}
}

func BenchmarkDay02CheckInvalid(b *testing.B) {
	cfg := &config{}
	inputs := []int{0, 1, 12, 55, 123, 6464, 123123, 1234, 1111, 123456, 999999, 12345678}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day02CheckInvalid(inputs[i%len(inputs)])
	}
}

func BenchmarkDay02CheckInvalid2(b *testing.B) {
	cfg := &config{}
	inputs := []int{0, 1, 12, 55, 111, 123, 1212, 1234, 1111, 123123, 12341234, 1212121212, 1111111111}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day02CheckInvalid2(inputs[i%len(inputs)])
	}
}

func BenchmarkAllSame(b *testing.B) {
	inputs := []string{"", "a", "aa", "aaa", "abc", "1111111", "1234567", "aaaaaaaa"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = allSame(inputs[i%len(inputs)])
	}
}

func BenchmarkRepeatedBlocks(b *testing.B) {
	inputs := []struct {
		s      string
		blocks int
	}{
		{"abab", 2},
		{"abcabcabc", 3},
		{"1212121212", 5},
		{"12341234", 2},
		{"abcd", 2},
		{"11111111", 4},
		{"123123123", 3},
		{"abcdefgh", 2},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := inputs[i%len(inputs)]
		_ = repeatedBlocks(tt.s, tt.blocks)
	}
}

func BenchmarkDay02Part1(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"1-100",
		"50-150",
		"100-200",
		"1-1000",
		"500-600",
		"1000-2000",
		"1-100,200-300",
		"50-150,250-350",
		"100-200,300-400,500-600",
		"1-1000,2000-3000",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day02part1(instructions[i%len(instructions)])
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	cfg := &config{}
	instructions := []string{
		"1-100",
		"50-150",
		"100-200",
		"1-1000",
		"500-600",
		"1000-2000",
		"1-100,200-300",
		"50-150,250-350",
		"100-200,300-400,500-600",
		"1-1000,2000-3000",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cfg.day02part2(instructions[i%len(instructions)])
	}
}
