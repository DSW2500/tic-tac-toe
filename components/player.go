package components

//Player  :
type Player struct {
	Name string
	Mark string
}

//CreatePlayer :
func CreatePlayer(name string, mark string) *Player {
	return &Player{
		Name: name,
		Mark: mark,
	}

}
