package puppy

import (
	"github.com/venkatasivasai1996/dog"
)

func Bark() string {
	return "first go"
}

func Barks() string {
	return "woof! woof! woof!"
}
func Bigbark() string {
	return dog.WhenGrownUP(Bark())
}

func Bigbarks() string {
	return dog.WhenGrownUP(Barks())
}
