package main 

import(
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14115

	var raio = 3.2 // tipo (flot64) inferido pelo compilador / em Go não somos obrigados a declarar 
	               // o tipo da variável se atribuirmos um valor para ela
				   
		// forma reduzida de criar uma variável 

		area := PI * math.Pow(raio, 2)
		fmt.Println("A área da circuferência é:",  area)

		const( 
			a = 1
			b = 2 
		)

		var (
			c = 3
			d = 4
		)

		fmt.Println(a, b, c, d) // importante sempre utilizar as variáveis sempre que definidas para o código rodar
								// o GO prioriza código limpo, caso não utilize a variável, ele não irá rodar

			
	






}