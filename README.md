# go-music-chord-note

`go-music-chord-note` provides utilities for chords and notes.

## Usage

    package main

    import (
        "github.com/bayashi/go-music-chord-note/note"
        "github.com/bayashi/go-music-chord-note/chord"
        "github.com/bayashi/go-music-chord-note/scale"
    )

    func main() {
        n, _ := note.NoteNumber("Db") // "Db"
        println(n) // 1

        n2, _ := note.NoteNumber("C9") // "C" on octave 9
        println(n2) // 120

        chordName, _ := chord.GetChord("Csus4")
        println(chordName[0]) // "C"
        println(chordName[1]) // "F"
        println(chordName[2]) // "G"

        chordNumber, _ := chord.GetChordAsNumberList("sus4")
        println(chordNumber[0]) // "0"
        println(chordNumber[1]) // "5"
        println(chordNumber[2]) // "7"

        scaleNumber, _ := scale.GetScale("Ionian")
        println(scaleNumber[0]) // "0"
        println(scaleNumber[1]) // "2"
        println(scaleNumber[2]) // "4"
        println(scaleNumber[3]) // "5"
        println(scaleNumber[4]) // "7"
        println(scaleNumber[5]) // "9"
        println(scaleNumber[6]) // "11"
    }

See tests for more functions.

## TODO

* Add functions to return notes in YAMAHA style.
* Add a feature to treat original chords or scales.

## Installation

    go get github.com/bayashi/go-music-chord-note

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
