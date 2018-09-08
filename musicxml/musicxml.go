package musicxml

import (
	"errors"
	"io"

	"github.com/musica/musicxml"
)

var (
	errorEmptyMusicXML     = errors.New("synth: no musicxml data found")
	errorTimeSignaturePath = errors.New("synth: xml path for time signature not found")
)

// MusicXML is a container for MusicXML data parsed into a MXML struct.
type MusicXML struct {
	musicXML *musicxml.MXML
	meta     *musicXMLMeta
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
func (m *MusicXML) SetMXML(mx *musicxml.MXML) (*MusicXML, error) {
	if root := mx.Name(); root == "score-partwise" || root == "score-timewise" {
		m.musicXML = mx
		return m.preprocess()
	}
	return m, errors.New("MusicXML root element must be either score-partwise or score-timewise")
}

func (m *MusicXML) setMXML(mx *musicxml.MXML, err error) (*MusicXML, error) {
	m.musicXML = mx
	if err != nil {
		return m, err
	}
	return m.preprocess()
}

// ParseXMLBuffer parses musicxml data from a io.Reader into a MusicXML struct.
func (m *MusicXML) ParseXMLBuffer(r io.Reader) (*MusicXML, error) {
	return m.setMXML(musicxml.ParseXMLBuffer(r))
}

// ParseXMLFile parses musicxml data from a file into a MusicXML struct.
func (m *MusicXML) ParseXMLFile(filePath string) (*MusicXML, error) {
	return m.setMXML(musicxml.ParseXMLFile(filePath))
}

// ParseMXLFile parses a musicxml data from a mxl file into a MusicXML struct.
func (m *MusicXML) ParseMXLFile(filePath string) (*MusicXML, error) {
	return m.setMXML(musicxml.ParseMXLFile(filePath))
}
