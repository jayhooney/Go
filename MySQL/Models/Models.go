package models

type Topic struct {
	DOCIDs int
	TERMS  string
}

type Emotion struct {
	Seq     int
	CmtID   string
	Emotion string
}
