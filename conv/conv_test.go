package conv

import (
	"math"
	"reflect"
	"testing"
)

/*
*

	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 32, want: 0},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if math.Floor(got) == got {
			// Hvis "got" er et heltall formater uten desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.0f, got: %.0f", tc.want, got)
			}
		} else {
			// Hvis "got" ikke er et heltall så formaterer den med 2 desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
			}
		}
	}
}
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 255.37},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if math.Floor(got) == got {
			// Hvis "got" er et heltall formater uten desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.0f, got: %.0f", tc.want, got)
			}
		} else {
			// Hvis "got" ikke er et heltall så formaterer den med 2 desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
			}
		}
	}
}
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 89.4, want: 128.92},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if math.Floor(got) == got {
			// Hvis "got" er et heltall formater uten desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.0f, got: %.0f", tc.want, got)
			}
		} else {
			// Hvis "got" ikke er et heltall så formaterer den med 2 desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
			}
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 89.4, want: 183.75},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Floor(got) == got {
			// Hvis "got" er et heltall formater uten desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.0f, got: %.0f", tc.want, got)
			}
		} else {
			// Hvis "got" ikke er et heltall så formaterer den med 2 desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
			}
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Floor(got) == got {
			// Hvis "got" er et heltall formater uten desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.0f, got: %.0f", tc.want, got)
			}
		} else {
			// Hvis "got" ikke er et heltall så formaterer den med 2 desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
			}
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 300, want: 80.33},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if math.Floor(got) == got {
			// Hvis "got" er et heltall formater uten desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.0f, got: %.0f", tc.want, got)
			}
		} else {
			// Hvis "got" ikke er et heltall så formaterer den med 2 desimaler
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %.2f, got: %.2f", tc.want, got)
			}
		}
	}
}

// De andre testfunksjonene implementeres her
// ...
