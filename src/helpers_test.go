package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", 5, -3, 2},
		{"with zero", 0, 5, 5},
		{"zero plus zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed signs", 5, -3, 8},
		{"from zero", 0, 5, -5},
		{"zero minus zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("subtract(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 6},
		{"negative numbers", -2, -3, 6},
		{"mixed signs positive", 5, -3, -15},
		{"mixed signs negative", -5, 3, -15},
		{"multiply by zero", 5, 0, 0},
		{"zero multiply zero", 0, 0, 0},
		{"multiply by one", 7, 1, 7},
		{"large numbers", 100, 50, 5000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("multiply(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		expected  int
		wantError bool
	}{
		{"positive numbers", 6, 2, 3, false},
		{"negative numbers", -6, -2, 3, false},
		{"mixed signs", -6, 2, -3, false},
		{"result is zero", 0, 5, 0, false},
		{"divide by one", 7, 1, 7, false},
		{"divide by zero", 5, 0, 0, true},
		{"zero divide by zero", 0, 0, 0, true},
		{"large division", 100, 5, 20, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divide(tt.a, tt.b)
			if (err != nil) != tt.wantError {
				t.Errorf("divide(%d, %d) error = %v, wantError %v", tt.a, tt.b, err, tt.wantError)
			}
			if err == nil && result != tt.expected {
				t.Errorf("divide(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestModulus(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		expected  int
		wantError bool
	}{
		{"positive numbers", 7, 3, 1, false},
		{"negative dividend", -7, 3, -1, false},
		{"negative divisor", 7, -3, 1, false},
		{"both negative", -7, -3, -1, false},
		{"result is zero", 6, 3, 0, false},
		{"zero dividend", 0, 5, 0, false},
		{"modulus by one", 7, 1, 0, false},
		{"large numbers", 100, 7, 2, false},
		{"modulus by zero", 5, 0, 0, true},
		{"zero modulus zero", 0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := modulus(tt.a, tt.b)
			if (err != nil) != tt.wantError {
				t.Errorf("modulus(%d, %d) error = %v, wantError %v", tt.a, tt.b, err, tt.wantError)
			}
			if err == nil && result != tt.expected {
				t.Errorf("modulus(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
