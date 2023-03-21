# go-music-chord-note

`go-music-chord-note` provides utilities for chords and notes.

## Usage

    package music

    import (
        "github.com/bayashi/go-music-chord-note/note"
        "github.com/bayashi/go-music-chord-note/chord"
    )

    n, _ := note.NoteNumber("Db") // "Db"
    println(n) // 1

    n2, _ := note.NoteNumber("C9") // "C" on octave 9
    println(n2) // 120

    chord, _ := chord.GetChord("C")
    println(chord[0]) // "C"
    println(chord[1]) // "E"
    println(chord[2]) // "G"

    chordn, _ := chord.GetChordAsNumberList("D")
    println(chordn[0]) // "D"
    println(chordn[1]) // "F#"
    println(chordn[2]) // "A"

See tests for more functions.

## TODO

* Write more documents.
* Add functions to return notes in YAMAHA style.
* Add a utility to get some scales.

## Installation

    go get github.com/bayashi/go-music-chord-note

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
