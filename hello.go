package hello

import (
	"fmt"

	"github.com/pbsilvab/greetings"
)

func main() {
	message := greetings.Message("Pedro")
	fmt.Println(message)
}
