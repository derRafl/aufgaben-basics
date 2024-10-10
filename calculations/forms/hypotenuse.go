package forms

import "math"

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	// SOLUTION
	return math.Sqrt(a*a + b*b)
	// SOLUTION_END
}

// HINT
// Benutzen Sie die Funktion math.Sqrt().