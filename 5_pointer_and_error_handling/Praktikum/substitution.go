package pointerAndErrorHandling

import "strings"

type Siswa struct {
	Nama       string
	NamaEncode string
	Skor       int
}

type Cipher interface {
	Encode() string
	Decode() string
}

func (s *Siswa) Encode() string {
	nameEncode := ""
	text := strings.ToLower(s.Nama)
	for _, char := range text {
		nameEncode += string(rune(-13)%26 + char)
	}
	return nameEncode
}

func (s *Siswa) Decode() string {
	nameDecode := ""
	text := s.Nama
	for _, char := range text {
		nameDecode += string(rune(13)%26 + char)
	}
	return nameDecode
}
