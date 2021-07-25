package adding_credits

type Solution interface {
	Insert(clientID string, c int)
	Remove(clientID string) bool
	Lookup(clientID string) int
	AddAll(c int)
	Max() string
}

type ClientsCreditsInfo struct {
	// TODO - Add your code here
}

func NewClientsCreditsInfo() Solution {
	// TODO - Add your code here
	return &ClientsCreditsInfo{}
}

func (cc *ClientsCreditsInfo) Insert(clientID string, c int) {
	// TODO - Add your code here
}

func (cc *ClientsCreditsInfo) Remove(clientID string) bool {
	// TODO - Add your code here
	return false
}

func (cc *ClientsCreditsInfo) Lookup(clientID string) int {
	// TODO - Add your code here
	return 0
}

func (cc *ClientsCreditsInfo) AddAll(c int) {
	// TODO - Add your code here
}

func (cc *ClientsCreditsInfo) Max() string {
	// TODO - Add your code here
	return ""
}
