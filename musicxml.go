package synth

import (
	"io"

	"github.com/musica/musicxml"
)

// MusicXML is a container for MusicXML data parsed into a MXML struct.
type MusicXML struct {
	musicXML *musicxml.MXML
}

// NewMusicXML returns a new empty MusicXML struct.
func NewMusicXML() *MusicXML {
	return &MusicXML{}
}

// MXML returns the musicxml tree as MXML.
func (m *MusicXML) MXML() *musicxml.MXML {
	return m.musicXML
}

// SetMXML sets a MXML to the MusicXML struct.
func (m *MusicXML) SetMXML(mx *musicxml.MXML) *MusicXML {
	m.musicXML = mx
	return m
}

func (m *MusicXML) setMXML(mx *musicxml.MXML, err error) (*MusicXML, error) {
	m.musicXML = mx
	return m, err
}

// ParseXMLBuffer parses musicxml data from a io.Reader into a MusicXML struct.
func (m *MusicXML) ParseXMLBuffer(r io.Reader) (*MusicXML, error) {
	return m.setMXML(musicxml.ParseXMLBuffer(r))
}

// ParseXMLBytes parses musicxml data from []byte into a MusicXML struct.
func (m *MusicXML) ParseXMLBytes(data []byte) (*MusicXML, error) {
	return m.setMXML(musicxml.ParseXMLBytes(data))
}

// ParseXMLFile parses musicxml data from a file into a MusicXML struct.
func (m *MusicXML) ParseXMLFile(filePath string) (*MusicXML, error) {
	return m.setMXML(musicxml.ParseXMLFile(filePath))
}

// ParseMXLFile parses a musicxml data from a mxl file into a MusicXML struct.
func (m *MusicXML) ParseMXLFile(filePath string) (*MusicXML, error) {
	return m.setMXML(musicxml.ParseMXLFile(filePath))
}
