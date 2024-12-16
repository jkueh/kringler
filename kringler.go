package kringler

type KringlerAssignment struct {
	Giver    string
	Receiver string
}

type Kringler struct {
	Kringlers   []string
	Assignments []KringlerAssignment
}

func New() *Kringler {
	return &Kringler{}
}

func NewFromList(kringlers []string) *Kringler {
	return &Kringler{Kringlers: kringlers}
}

// Sets the entire kringlers list in one go.
func (k *Kringler) SetKringlers(kringlers []string) {
	k.Kringlers = kringlers
}

// Adds a kringler to the list
func (k *Kringler) AddKringler(kringler string) {
	k.Kringlers = append(k.Kringlers, kringler)
}
