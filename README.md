# synth

Music package for Go.

---

## MusicXML

- You can use `synth` to parse MusicXML files. Internally, `synth` uses `github.com/musica/musicxml`.

- The `MusicXML` struct *would provide* helper functions to handle the music aspects of the MusicXML.

- To parse a MusicXML file, call `NewMusicXML()`, and then use any of these methods.

  - `ParseXMLBuffer(XMLFileReader io.Reader)`.

  - `ParseXMLFile(XMLFilePath string)`.

  - `ParseMXLFile(MXLFilePath string)`
