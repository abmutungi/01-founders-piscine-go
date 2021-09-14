/*package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}
package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
*/
package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3))
	z01.PrintRune(piscine.NRune("Salut!", 2))
	z01.PrintRune(piscine.NRune("Bye!", -1))
	z01.PrintRune(piscine.NRune("Bye!", 5))
	z01.PrintRune(piscine.NRune("Ola!", 4))
	z01.PrintRune('\n')
}
