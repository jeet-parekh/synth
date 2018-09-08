package musicxml

// import "errors"

// func (m *MusicXML) TimeSignature() (int, int, error) {
// 	var n []*mxml.MXML
// 	var ok bool

// 	if m.MXML() == nil {
// 		return 0, 0, errorEmptyMusicXML
// 	}

// 	n = m.MXML()
// 	if n.Name() == "score-partwise" {
// 		n, ok = n.Children()["part"]
// 		if !ok || len(n) == 0 {
// 			return 0, 0, errorTimeSignaturePath
// 		}

// 		n, ok = n.Children()["measure"]
// 		if !ok || len(n) == 0 {
// 			return 0, 0, errorTimeSignaturePath
// 		}

// 		return 0, 0, errors.New("GAAH")
// 	}

// 	if m.MXML().Name() == "score-timewise" {
// 		n, ok = m.MXML().Children()["measure"]
// 		if !ok || len(n) == 0 {
// 			return 0, 0, errorTimeSignaturePath
// 		}
// 		return 0, 0, errors.New("GAAH")
// 	}

// 	return 0, 0, nil
// }
