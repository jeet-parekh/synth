package synth

import (
	"github.com/musica/synth/musicxml"
)

// MusicXML is a wrapper around the MusicXML struct from the musicxml sub-package.
type MusicXML struct {
	*musicxml.MusicXML
}

// NewMusicXML returns a MusicXML.
func NewMusicXML() *MusicXML {
	return &MusicXML{musicxml.NewMusicXML()}
}
