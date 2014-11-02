// Special thanks to Johanna for providing us with the name and idea for this package.
package betofahl

import (
	"os"
	"os/signal"
	"time"
	"syscall"
	"fmt"
	"math/rand"
)

var (
	Texts = []string{"Schokodrops", "Gummibärchen", "Maoam", "Schokolade", "Lolli", "Colafläschchen"}
	c = make(chan os.Signal)
)

func init() {
	go func() {
		// Set up termination detection
		signal.Notify(c, syscall.SIGKILL, syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM)

		// Wait for termination
		<-c

		// Randomize choice
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Generate random text
		text := Texts[r.Intn(len(Texts))]

		// Voll fair
		fmt.Println(text)
		os.Exit(0)
	}()
}

func Defer() {
	// Randomize choice
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate random text
	text := Texts[r.Intn(len(Texts))]

	// Voll fair
	fmt.Println(text)
}
