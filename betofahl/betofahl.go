// Special thanks to Johanna for providing us with the name and idea for this package.
package betofahl

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

var (
	Texts = []string{"Schokodrops", "Gummibärchen", "Maoam", "Schokolade", "Lolli", "Colafläschchen"}
)

func init() {
	go func() {
		// Set up termination detection
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGKILL, syscall.SIGQUIT)

		// Wait for termination
		<-c

		// Generate random text
		text := Texts[int(time.Now().Unix()) % len(Texts)]

		// Voll fair
		fmt.Println(text)
	}()
}
