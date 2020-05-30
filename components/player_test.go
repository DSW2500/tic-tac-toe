package components

import (
	"fmt"
	"testing"
)

func TestPlayer(t *testing.T) {

	expected := *&Player{
		Name: "Devesh",
		Mark: "X",
	}

	actual := CreatePlayer("Devesh", "X")
	if actual.Mark != expected.Mark {
		t.Error("not working")
	}
	if actual.Name != expected.Name {
		t.Error("not working")
	}
	fmt.Println("Working properly!")
}
