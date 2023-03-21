package note

import "testing"

func TestNoteNumber(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "C",
			want: 0,
		},
		{
			name: "C-1",
			want: 0,
		},
		{
			name: "C0",
			want: 12,
		},
		{
			name: "C1",
			want: 24,
		},
		{
			name: "C9",
			want: 120,
		},
		{
			name: "B",
			want: 11,
		},
		{
			name: "B-1",
			want: 11,
		},
		{
			name: "G9",
			want: 127,
		},
		{
			name: "C#",
			want: 1,
		},
		{
			name: "Db",
			want: 1,
		},
		{
			name: "C#-1",
			want: 1,
		},
		{
			name: "C#0",
			want: 13,
		},
		{
			name: "H",
			want: 11,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := NoteNumber(test.name);
			if actual != test.want {
				t.Errorf(`NoteNumber("%v"), actual:"%v", want:"%v"`, test.name, actual, test.want)
			}
		})
	}
}

func TestNoteNumberError(t *testing.T) {
	tests := []struct {
		name string
		want error
		reason string
	}{
		{
			name: "C-5",
			want: ErrorNotFoundNote("C-5"),
			reason: "Out of range.",
		},
		{
			name: "C-2",
			want: ErrorNotFoundNote("C-2"),
			reason: "Out of range.",
		},
		{
			name: "C10",
			want: ErrorNotFoundNote("C10"),
			reason: "Out of range.",
		},
		{
			name: "G#9",
			want: ErrorOutOfRange,
			reason: "Out of range",
		},
		{
			name: "A9",
			want: ErrorOutOfRange,
			reason: "Out of range",
		},
		{
			name: "I9",
			want: ErrorNotFoundNote("I9"),
			reason: "Wrong note name",
		},
		{
			name: "C-",
			want: ErrorNotFoundNote("C-"),
			reason: "Wrong note name. Use `Cb` instead",
		},
		{
			name: "C+",
			want: ErrorNotFoundNote("C+"),
			reason: "Wrong note name. Use `C#` instead",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := NoteNumber(test.name);
			if actual != ErrorInt {
				t.Errorf(`NoteNumber("%v") got "%v"`, test.name, actual)
			}
			if err.Error() != test.want.Error() {
				t.Errorf(`NoteNumber("%v"), actual err:"%v", want:"%v"`, test.name, err, test.want)
			}
		})
	}
}

func TestNoteNumberWithOctave(t *testing.T) {
	tests := []struct {
		name string
		octave int
		want int
	}{
		{
			name: "C", octave: -1,
			want: 0,
		},
		{
			name: "C", octave: 0,
			want: 12,
		},
		{
			name: "C", octave: 4,
			want: 60,
		},
		{
			name: "D", octave: 8,
			want: 110,
		},
		{
			name: "G", octave: 9,
			want: 127,
		},
		{
			name: "H", octave: 1,
			want: 35,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := NoteNumberWithOctave(test.name, test.octave);
			if actual != test.want {
				t.Errorf(`NoteNumber("%v", %v), actual:"%v", want:"%v"`, test.name, test.octave, actual, test.want)
			}
		})
	}
}

func TestNoteNumberErrorWithOctave(t *testing.T) {
	tests := []struct {
		name string
		octave int
		want error
	}{
		{
			name: "C",
			octave: -2,
			want: ErrorInvalidOctave,
		},
		{
			name: "A",
			octave: 9,
			want: ErrorOutOfRange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := NoteNumberWithOctave(test.name, test.octave);
			if actual != ErrorInt {
				t.Errorf(`NoteNumber("%v") got "%v"`, test.name, actual)
			}
			if err.Error() != test.want.Error() {
				t.Errorf(`NoteNumber("%v", %v), actual err:"%v", want:"%v"`, test.name, test.octave, err, test.want)
			}
		})
	}
}