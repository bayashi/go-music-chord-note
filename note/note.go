package note

import (
	"fmt"
	"regexp"
	"strconv"
)

var BaseTones = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

//     1   3       6   8   10
// |  | | | |  |  | | | | | |  |
// |  |_| |_|  |  |_| |_| |_|  |
// |___|___|___|___|___|___|___|
//   0   2   4   5   7   9   11

// Relative note number mapping
var noteNameDegree = map[string]int{
	"C":  0, "B#": 0,
	"C#": 1,
	"Db": 1,
	"D":  2,
	"D#": 3,
	"Eb": 3,
	"E":  4, "Fb": 4,
	"F":  5, "E#": 5,
	"F#": 6,
	"Gb": 6,
	"G":  7,
	"G#": 8,
	"Ab": 8,
	"A":  9,
	"A#": 10,
	"Bb": 10,
	"B":  11, "Cb": 11,

	"Hb": 10,
	"H":  11,
}

const (
	MinimumNoteNumber = 0
	MaximumNoteNumber = 127
)

const ErrorInt = -2147483648

var (
	ErrorNotFoundNote = func(noteName string) error { return fmt.Errorf("Not found note. `%s`", noteName) }
	ErrorCouldNotGetOctave = func(noteName string) error { return fmt.Errorf("Could not get octave. `%s`", noteName) }
	ErrorCouldNotGetOctaveWithError = func(noteName string, err error) error { return fmt.Errorf("Could not get octave. `%s`, `%w`", noteName, err) }
	ErrorCouldNotGetDegree = func(noteName string) error { return fmt.Errorf("Could not get degree. `%s`", noteName) }

	ErrorInvalidOctave = fmt.Errorf("`octave` should be -1 to 9.")
	ErrorOutOfRange = fmt.Errorf("Out of range.")
)

// Get a note number `3` from a note name `Eb`. Or, Get a note number `60` from a note name with octave number `C4`.
func NoteNumber(noteName string) (int, error) {
	if degree, isExists := noteNameDegree[noteName]; isExists {
		return degree, nil
	}

	if !isValidNoteName(noteName) {
		return ErrorInt, ErrorNotFoundNote(noteName)
	}

	octave, err := getOctaveFromName(noteName)
	if err != nil {
		return ErrorInt, err
	}

	degree, err2 := getDegreeFromNameWithOctave(noteName)
	if err2 != nil {
		return ErrorInt, err2
	}

	noteNumber, err := actualNoteNumber(octave, degree)
	if err != nil {
		return ErrorInt, err
	}

	return noteNumber, nil
}

// As full note name i.e. `C4`, `Eb-1` or `G9`. Not consider valid name. `A9` is out of note number, but it's true to match.
var noteRegexp = regexp.MustCompile(`^[A-H][#b]?(\-1|[0-9])?$`)

func isValidNoteName(noteName string) bool {
	return noteRegexp.MatchString(noteName)
}

// get a octave number `1` from a note name `C1`.
func getOctaveFromName(noteName string) (int, error) {
	if length := len(noteName); length > 1 {
		// case "C-1"
		if last2Chars := noteName[length - 2:]; last2Chars == "-1" {
			return -1, nil
		}

		// case "C1" - "C9"
		lastChar := noteName[length - 1:];
		octave, err := strconv.Atoi(lastChar)
		if err != nil {
			// last char is "#" or "b" case.
			return ErrorInt, ErrorCouldNotGetOctaveWithError(noteName, err)
		}
		if isValidOctave(octave) {
			return octave, nil
		}
	}

	return ErrorInt, ErrorCouldNotGetOctave(noteName) // wrong noteName
}

// Get a degree `0` (`C` -> `0`) from a note name `C2`. This function is called when a noteName is including octave.
func getDegreeFromNameWithOctave(noteName string) (int, error) {
	first2Chars := noteName[0:2];
	degree, isExists := noteNameDegree[first2Chars];
	if isExists {
		return degree, nil // case "Db2" or "D#2"
	} else {
		firstChar := noteName[0:1];
		degree, isExists = noteNameDegree[firstChar];
		if isExists {
			return degree, nil // case "D2"
		}
	}

	return ErrorInt, ErrorCouldNotGetDegree(noteName) // wrong noteName
}


// Get an absolute note number `60` from note name `C` and octave `4`.
// This function is faster than `NoteNumber()`.
func NoteNumberWithOctave(noteName string, octave int) (int, error) {
	degree, isExists := noteNameDegree[noteName]
	if !isExists {
		return ErrorInt, ErrorNotFoundNote(noteName)
	}

	if !isValidOctave(octave) {
		return ErrorInt, ErrorInvalidOctave
	}

	noteNumber, err := actualNoteNumber(octave, degree)
	if err != nil {
		return ErrorInt, err
	}

	return noteNumber, nil
}

// get an actual valid note number. 4, degree:0 (C) -> 60
func actualNoteNumber(octave int, degree int) (int, error) {
	noteNumber := (octave + 1) * 12 + degree
	if noteNumber >= MinimumNoteNumber && noteNumber <= MaximumNoteNumber {
		return noteNumber, nil
	}

	return ErrorInt, ErrorOutOfRange
}

// Octave number should be between -1 to 9.
func isValidOctave(octave int) bool {
	return octave >= -1 && octave <= 9
}
