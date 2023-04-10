package chord

import (
	"testing"

	"github.com/bayashi/actually"
)

func TestAllChords(t *testing.T) {
	if len(allChords()) != 75 {
		t.Error("Wrong number of all chords.")
	}

	for _, v := range AllChords {
		if v == "base" {
			return
		}
	}
	t.Error("AllChords doesn't have `base`.")
}

func TestGetChordAsNumberList(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "base",
			want: []int{0, 4, 7},
		},
		{
			name: "sus4",
			want: []int{0, 5, 7},
		},
		{
			name: "13",
			want: []int{0, 4, 7, 10, 14, 17, 21},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := GetChordAsNumberList(test.name);
			actually.Got(err).FailNow().Nil(t)
			if len(actual) != len(test.want) {
				t.Errorf(`GetChordAsNumberList("%v"), actual:"%v", want:"%v"`, test.name, actual, test.want)
			}
			for i, v := range test.want {
				if actual[i] != v {
					t.Errorf(`GetChordAsNumberList("%v"), note No.%v is wrong. actual:"%v", want:"%v"`, test.name, i+1, actual, test.want)
				}
			}
		})
	}
}

func TestGetChordAsNumberListError(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "N7",
			want: ErrorNotFoundChordKind("N7"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetChordAsNumberList(test.name);
			actually.Got(err).FailNow().NotNil(t)
			if len(got) != 0 {
				t.Errorf(`GetChordAsNumberList("%v") wants empty result. But got (%v).`, test.name, got)
			}
			if err.Error() != test.want.Error() {
				t.Errorf(`GetChordAsNumberList("%v") wants Error(%v). but it's wrong. "%v"`, test.name, err, test.want)
			}
		})
	}
}

func TestGetChord(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "C",
			want: []string{"C", "E", "G"},
		},
		{
			name: "D",
			want: []string{"D", "F#", "A"},
		},
		{
			name: "BM7",
			want: []string{"B", "D#", "F#", "A#"},
		},
		{
			name: "Csus4",
			want: []string{"C", "F", "G"},
		},
		{
			name: "C13",
			want: []string{"C", "E", "G", "A#", "D", "F", "A"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := GetChord(test.name);
			actually.Got(err).FailNow().Nil(t)
			if len(actual) != len(test.want) {
				t.Errorf(`GetChord("%v"), actual:"%v", want:"%v"`, test.name, actual, test.want)
			}
			for i, v := range test.want {
				if actual[i] != v {
					t.Errorf(`GetChord("%v"), note No.%v is wrong. actual:"%v", want:"%v"`, test.name, i+1, actual, test.want)
				}
			}
		})
	}
}

func TestGetChordError(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{
			name: "X",
			want: ErrorNotFoundChord("X"),
		},
		{
			name: "CN7",
			want: ErrorNotFoundChordKind("N7"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetChord(test.name);
			actually.Got(err).FailNow().NotNil(t)
			if len(got) != 0 {
				t.Errorf(`GetChord("%v") wants empty result. But got (%v).`, test.name, got)
			}
			if err.Error() != test.want.Error() {
				t.Errorf(`GetChord("%v") wants Error(%v). but it's wrong. "%v"`, test.name, err, test.want)
			}
		})
	}
}

func TestGetChordWithOctave(t *testing.T) {
	tests := []struct {
		name string
		octave int
		want []string
	}{
		{
			name: "C",
			octave: -1,
			want: []string{"C-1", "E-1", "G-1"},
		},
		{
			name: "BM7",
			octave: -1,
			want: []string{"B-1", "D#0", "F#0", "A#0"},
		},
		{
			name: "Asus4",
			octave: 2,
			want: []string{"A2", "D3", "E3"},
		},
		{
			name: "C",
			octave: 9,
			want: []string{"C9", "E9", "G9"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := GetChordWithOctave(test.name, test.octave);
			actually.Got(err).FailNow().Nil(t)
			if len(actual) != len(test.want) {
				t.Errorf(`GetChordWithOctave("%v, %v"), actual:"%v", want:"%v"`, test.name, test.octave, actual, test.want)
			}
			for i, v := range test.want {
				if actual[i] != v {
					t.Errorf(`GetChordWithOctave("%v, %v"), note No.%v is wrong. actual:"%v", want:"%v"`,
						test.name, test.octave, i+1, actual, test.want)
				}
			}
		})
	}
}

func TestGetChordWithOctaveError(t *testing.T) {
	tests := []struct {
		name string
		octave int
		want error
	}{
		{
			name: "Xm7",
			octave: 4,
			want: ErrorNotFoundChord("Xm7"),
		},
		{
			name: "CN7",
			octave: 5,
			want: ErrorNotFoundChordKind("N7"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetChordWithOctave(test.name, test.octave);
			actually.Got(err).FailNow().NotNil(t)
			if len(got) != 0 {
				t.Errorf(`GetChord("%v") wants empty result. But got (%v).`, test.name, got)
			}
			if err.Error() != test.want.Error() {
				t.Errorf(`GetChord("%v") wants Error(%v). but it's wrong. "%v"`, test.name, err, test.want)
			}
		})
	}
}

