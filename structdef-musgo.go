package goserbench

type MusgoA struct {
	Name     string
	BirthDay int64 `mus:"#raw"`
	Phone    string
	Siblings int32
	Spouse   bool
	Money    float64 `mus:"#raw"`
}
