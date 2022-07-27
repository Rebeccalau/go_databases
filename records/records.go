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

func DefaultRecord2() Record {
	features := []Feature{{Name: "Flag1", Enabled: true}, {Name: "Flag2", Enabled: false}}

	testing := Record{
		Id:        "12345",
		Primary:   "1234-12324",
		Secondary: nil,
		Code:      "Code-1234",
		Features:  features,
	}
	return testing
}

func DefaultRecord1() Record {
	features := []Feature{{Name: "Flag1", Enabled: true}, {Name: "Flag2", Enabled: false}}

	testing := Record{
		Id:        "11111",
		Primary:   "2222-2222",
		Secondary: nil,
		Code:      "Code-3333",
		Features:  features,
	}
	return testing
}
