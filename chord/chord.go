package chord

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bayashi/go-music-chord-note/note"
)

// The number of supported chord
const allChord = 75

//     1   3       6   8   10
// |  | | | |  |  | | | | | |  |
// |  |_| |_|  |  |_| |_| |_|  |
// |___|___|___|___|___|___|___|
//   0   2   4   5   7   9   11

// Mapping of chord name and note list
var allKindOfChords = map[string][]int{
	"base":      {0, 4, 7},
	"-5":        {0, 4, 6},
	"-6":        {0, 4, 7, 8},
	"6":         {0, 4, 7, 9},
	"6(9)":      {0, 4, 7, 9, 14}, "69": {0, 4, 7, 9, 14},
	"M7":        {0, 4, 7, 11},
	"M7(9)":     {0, 4, 7, 11, 14}, "M79": {0, 4, 7, 11, 14},
	"M9":        {0, 4, 7, 11, 14},
	"M11":       {0, 4, 7, 11, 14, 17},
	"M13":       {0, 4, 7, 11, 14, 17, 21},
	"7":         {0, 4, 7, 10},
	"7(b5)":     {0, 4, 6, 10},     "7b5": {0, 4, 6, 10},
	"7(-5)":     {0, 4, 6, 10},     "7-5": {0, 4, 6, 10},
	"7(#5)":     {0, 4, 7, 8, 10},  "7#5": {0, 4, 7, 8, 10},
	"7(b9)":     {0, 4, 7, 10, 13}, "7b9": {0, 4, 7, 10, 13},
	"7(-9)":     {0, 4, 7, 10, 13}, "7-9": {0, 4, 7, 10, 13},
	"-9":        {0, 4, 7, 10, 13},
	"-9(#5)":    {0, 4, 8, 10, 13},     "-9#5": {0, 4, 8, 10, 13},
	"7(b9, 13)": {0, 4, 7, 10, 13, 21}, "7(-9, 13)": {0, 4, 7, 10, 13, 21},
	"7(9, 13)":  {0, 4, 7, 10, 14, 21},
	"7(#9)":     {0, 4, 7, 10, 15},     "7#9": {0, 4, 7, 10, 15},
	"7(#11)":    {0, 4, 7, 10, 15, 18}, "7#11": {0, 4, 7, 10, 15, 18},
	"7(#13)":    {0, 4, 10, 21},        "7#13": {0, 4, 10, 21},
	"9":         {0, 4, 7, 10, 14},
	"9(b5)":     {0, 4, 6, 10, 14},     "9b5": {0, 4, 6, 10, 14},
	"9(-5)":     {0, 4, 6, 10, 14},     "9-5": {0, 4, 6, 10, 14},
	"11":        {0, 4, 7, 10, 14, 17},
	"13":        {0, 4, 7, 10, 14, 17, 21},
	"m":         {0, 3, 7},
	"madd4":     {0, 3, 5, 7},
	"m6":        {0, 3, 7, 9},
	"m6(9)":     {0, 3, 7, 9, 14},  "m69": {0, 3, 7, 9, 14},
	"mM7":       {0, 3, 7, 11},
	"m7":        {0, 3, 7, 10},
	"m7(b5)":    {0, 3, 6, 10},     "m7b5": {0, 3, 6, 10},
	"m7(-5)":    {0, 3, 6, 10},     "m7-5": {0, 3, 6, 10},
	"m7(#5)":    {0, 3, 8, 10},     "m7#5": {0, 3, 8, 10},
	"m7(9)":     {0, 3, 7, 10, 14}, "m79": {0, 3, 7, 10, 14},
	"m9":        {0, 3, 7, 10, 14},
	"m7(9, 11)": {0, 3, 7, 10, 14, 17},
	"m11":       {0, 3, 7, 10, 14, 17},
	"m13":       {0, 3, 7, 10, 14, 17, 21},
	"dim":       {0, 3, 6},
	"dim7":      {0, 3, 6, 9}, "dim6": {0, 3, 6, 9},
	"aug":       {0, 4, 8},
	"aug7":      {0, 4, 8, 10},
	"augM7":     {0, 4, 8, 11},
	"aug9":      {0, 4, 8, 10, 14},
	"sus2":      {0, 2, 7},
	"sus":       {0, 5, 7},
	"sus4":      {0, 5, 7},
	"7sus4":     {0, 5, 7, 10},
	"add2":      {0, 2, 4, 7},
	"add4":      {0, 4, 5, 7},
	"add9":      {0, 4, 7, 14},
}

func allChords() [allChord]string {
	var list [allChord]string
	i := 0
	for k := range allKindOfChords {
		list[i] = k
		i++
	}

	return list
}

// All types of chord
var AllChords = allChords()

var (
	ErrorNotFoundChord = func(chordName string) error { return fmt.Errorf("Not found chord. `%s`", chordName) }
	ErrorNotFoundChordKind = func(chordKind string) error { return fmt.Errorf("Not found chord Kind. `%s`", chordKind) }
	ErrorNoteOutOfRange = func(chordName string) error { return fmt.Errorf("Note out of range. `%s`", chordName) }
)

// Get a chord as a note number list `{0, 4, 7, 11}` from kind of chord `M7`.
// NOTE: Note number is within 0 to 24, actually.
func GetChordAsNumberList(chordKind string) ([]int, error) {
	if chordKind == "" {
		chordKind = "base"
	}

	chord, isExists := allKindOfChords[chordKind]
	if !isExists {
		return nil, ErrorNotFoundChordKind(chordKind)
	}

	return chord, nil
}

// Get a chord as a note list `{"C", "E", "G", "B"}` from full chord name `CM7`.
func GetChord(chordName string) ([]string, error) {
	scalic, chordNumbers, err := parseChordName(chordName)
	if err != nil {
		return nil, err
	}

	var notes []string
	for _, n := range chordNumbers {
		noteNumber := (n + scalic) % 12
		notes = append(notes, note.BaseTones[noteNumber])
	}

	return notes, nil
}

// get a base note number `0` and a note number list `{0, 4, 7, 11}` from full chord name `CM7`.
func parseChordName(chordName string) (int, []int, error) {
	tonic, kind, err := splitChord(chordName)
	if err != nil {
		return note.ErrorInt, nil, err
	}

	scalic, _ := note.NoteNumber(tonic)
	chordNumbers, err2 := GetChordAsNumberList(kind);
	if  err2 != nil {
		return note.ErrorInt, nil, err2
	}

	return scalic, chordNumbers, err
}

var regexpSplitChord = regexp.MustCompile("^([A-G][b#]?)(.*)$")

// split full chord name `CM7` to note name `C` and kind of chord `M7`.
func splitChord(chordName string) (string, string, error) {
	if re := regexpSplitChord.FindStringSubmatch(chordName); re != nil {
		return re[1], re[2], nil
	}

	return "", "", ErrorNotFoundChord(chordName)
}

// Get a note list `{"F4", "A4", "C#5", "E5"}` from full chord name `FM7` and octave number `4`.
func GetChordWithOctave(chordName string, octave int) ([]string, error) {
	noteList, err := GetChord(chordName)
	if err != nil {
		return nil, err
	}

	var notesWithOctave []string
	var lastPosition = -1
	for _, n := range noteList {
		position, _ := note.NoteNumber(n)
		if position < lastPosition {
			octave++
		}
		nn := n + strconv.Itoa(octave)
		if octave > 9 || (octave == 9 && position > 7) {
			return nil, ErrorNoteOutOfRange(nn)
		}
		notesWithOctave = append(notesWithOctave, nn)
		lastPosition = position
	}

	return notesWithOctave, nil
}
