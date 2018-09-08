package musicxml

import "errors"

func (m *MusicXML) preprocess() (*MusicXML, error) {
	m.meta = newMusicXMLMeta()
	return m.ppRootElementName()
}

func (m *MusicXML) ppRootElementName() (*MusicXML, error) {
	switch mxName := m.musicXML.Name(); mxName {
	case "score-partwise", "score-timewise":
		m.meta.rootElementName = mxName
	default:
		return m, errors.New("MusicXML root element must be either score-partwise or score-timewise")
	}
	return m, nil
}

func (m *MusicXML) ppSplit() (*MusicXML, error) {
	switch m.meta.rootElementName {
	case "score-partwise":
		return m.ppScorePartwise()
	case "score-timewise":
		return m.ppScoreTimewise()
	default:
		return m, errors.New("MusicXML root element must be either score-partwise or score-timewise")
	}
}
