package vector

import "testing"

//	func TestPointAdd(t *testing.T) {
//		tests := []struct {
//			name     string
//			p1, p2   *Point
//			expected *Point
//		}{
//			{name: "Zero points", p1: &Point{0, 0}, p2: &Point{0, 0}, expected: &Point{0, 0}},
//			{name: "Positive points", p1: &Point{1.2, 3.4}, p2: &Point{4.5, 6.7}, expected: &Point{5.7, 10.1}},
//			{name: "Negative points", p1: &Point{-1.1, -2.2}, p2: &Point{-3.3, -4.4}, expected: &Point{-4.4, -6.6}},
//			{name: "Mixed positive and negative points", p1: &Point{-1.5, 2.7}, p2: &Point{3.6, -1.4}, expected: &Point{2.1, 1.3}},
//			{name: "Large values", p1: &Point{1e6, 1e6}, p2: &Point{-1e6, -1e6}, expected: &Point{0, 0}},
//			{name: "Small values", p1: &Point{1e-6, 1e-6}, p2: &Point{-1e-6, -1e-6}, expected: &Point{0, 0}},
//		}
//
//		for _, tt := range tests {
//			t.Run(tt.name, func(t *testing.T) {
//
//				tt.p1.Add(tt.p2)
//				if tt.p1.X() != tt.expected.X() || tt.p1.Y() != tt.expected.Y() {
//					t.Errorf("%v.Add(%v) = %v, want %v", tt.p1, tt.p2, result, tt.expected)
//				}
//			})
//		}
//	}
func TestNewPoint(t *testing.T) {
	tests := []struct {
		name     string
		x, y     float32
		expected *Point
	}{
		{name: "Zero point", x: 0, y: 0, expected: &Point{0, 0}},
		{name: "Positive values", x: 3.5, y: 4.2, expected: &Point{3.5, 4.2}},
		{name: "Negative values", x: -3.5, y: -4.2, expected: &Point{-3.5, -4.2}},
		{name: "Mixed values", x: -2.0, y: 4.7, expected: &Point{-2.0, 4.7}},
		{name: "Large values", x: 1e6, y: -1e6, expected: &Point{1e6, -1e6}},
		{name: "Small values", x: 1e-6, y: -1e-6, expected: &Point{1e-6, -1e-6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewPoint(tt.x, tt.y)
			if result.X() != tt.expected.X() || result.Y() != tt.expected.Y() {
				t.Errorf("NewPoint(%v, %v) = %v, want %v", tt.x, tt.y, result, tt.expected)
			}
		})
	}
}
