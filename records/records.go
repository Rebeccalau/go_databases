package records

type Feature struct {
	Name    string
	Enabled bool
}

type Record struct {
	Id        string
	Primary   string
	Secondary []string
	Code      string
	Features  []Feature
}
