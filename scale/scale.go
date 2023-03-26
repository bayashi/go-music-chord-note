package scale

import (
	"fmt"
	"strings"

	"github.com/bayashi/go-music-chord-note/note"
)

//    1   3     6   8   10       Db  Eb    Gb  Ab  Bb       C#  D#    F#  G#  A#
// | | | | | | | | | | | | |  | | | | | | | | | | | | |  | | | | | | | | | | | | |
// | |_| |_| | |_| |_| |_| |  | |_| |_| | |_| |_| |_| |  | |_| |_| | |_| |_| |_| |
// |__|___|__|__|___|___|__|  |__|___|__|__|___|___|__|  |__|___|__|__|___|___|__|
//  0   2  4   5  7   9  11    C   D  E   F  G   A   B    C   D  E   F  G   A   B

const allScale = 29

var allKindOfScales = map[string][]int{
	// Major
	"ionian":     {0, 2, 4, 5, 7, 9, 11}, // C D E F G A B
	// Gregorian mode https://en.wikipedia.org/wiki/Gregorian_mode
	"dorian":     {0, 2, 3, 5, 7, 9, 10}, // C D Eb F G A Bb
	"phrigian":   {0, 1, 3, 5, 7, 8, 10}, // C Db Eb F G Ab Bb
	"lydian":     {0, 2, 4, 6, 7, 9, 11}, // C D E F# G A B
	"mixolydian": {0, 2, 4, 5, 7, 9, 10}, // C D E F G A Bb
	"aeolian":    {0, 2, 3, 5, 7, 8, 10}, // C D Eb F G Ab Bb
	"locrian":    {0, 1, 3, 5, 6, 8, 10}, // C Db Eb F Gb Ab Bb

	// Harmonic minor https://en.wikipedia.org/wiki/Minor_scale
	"harmonic-minor":  {0, 2, 3, 5, 7, 8, 11}, // C D Eb F G Ab B
	"locrian#6":       {0, 1, 3, 5, 6, 9, 10}, // C Db Eb F Gb A Bb
	"ionian#5":        {0, 2, 4, 5, 8, 9, 11}, // C D E F G# A B
	"dorian#4":        {0, 2, 3, 6, 7, 9, 10}, // C D Eb F# G A Bb
	"phrigian-major":  {0, 1, 4, 5, 7, 8, 10}, // C Db E F G Ab Bb
	"lydian#2":        {0, 3, 4, 6, 7, 9, 11}, // C D# E F# G A B
	"super-locrianb7": {0, 1, 3, 4, 6, 8, 9}, // C Db Eb Fb Gb Ab Bbb

	// Melodic minor
	"super-ionian":     {0, 2, 3, 5, 7, 9, 11}, // C D Eb F G A B
	"super-dorian":     {0, 1, 3, 5, 7, 9, 10}, // C Db Eb F G A Bb
	"super-phrigian":   {0, 2, 4, 6, 8, 9, 11}, // C D E F# G# A B
	"super-lydian":     {0, 2, 4, 6, 7, 9, 10}, // C D E F# G A Bb
	"super-mixolydian": {0, 2, 4, 5, 7, 8, 10}, // C D E F G Ab Bb
	"super-aeolian":    {0, 2, 3, 5, 6, 8, 10}, // C D Eb F Gb Ab Bb
	"super-locrian":    {0, 1, 3, 4, 6, 8, 10}, // C Db Eb Fb Gb Ab Bb

	// Whole Tone (Symmetrical Scale)
	"whole-tone": {0, 2, 4, 6, 8, 10}, // C D E F# Ab Bb

	// Diminished
	"diminished": {0, 2, 3, 5, 6, 8, 9, 11}, // C D Eb F Gb G# A B

	// Chromatic
	"chromatic": {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, // C C# D D# E F F# G G# A A# B

	// Pentatonic
	"pentatonic-minor": {0, 3, 5, 7, 10}, // C Eb F G Bb
	"pentatonic-major": {0, 3, 4, 7, 9},  // C D  E G A
	// Blues
	"blues-minor": {0, 3, 5, 6, 7, 10}, // C Eb F Gb G Bb
	"blues-major": {0, 2, 3, 4, 7, 9},  // C D Eb E G A
	// Blue Note
	"blue-note": {0, 2, 3, 4, 5, 6, 7, 9, 10, 11}, // C D Eb E F Gb G A Bb B
}

func allScales() [allScale]string {
	var list [allScale]string
	i := 0
	for k := range allKindOfScales {
		list[i] = k
		i++
	}

	return list
}

// All scales
var AllScales = allScales()

var (
	ErrorNotFoundScale = func(scaleName string) error { return fmt.Errorf("Not found scale. `%s`", scaleName) }
)

// Get scale notes from scale name: `ionian` -> `{0, 2, 4, 5, 7, 9, 11}`
func GetScale(scaleName string) ([]int, error) {
	sc, isExists := allKindOfScales[strings.ToLower(scaleName)]

	if !isExists {
		return nil, ErrorNotFoundScale(scaleName)
	}

	return sc, nil
}

// Get scale notes as note number from the root note: `ionian, D4` -> `{62, 64, 66, 67, 69, 71, 73}`
func GetScaleFromRoot(scaleName string, rootNote string) ([]int, error) {
	sc, err := GetScale(scaleName)
	if err != nil {
		return nil, err
	}

	rootNoteNumber, err := note.NoteNumber(rootNote)
	if err != nil {
		return nil, err
	}

	var actualScale []int
	for _, v := range sc {
		actualScale = append(actualScale, rootNoteNumber + v)
	}

	return actualScale, nil
}
