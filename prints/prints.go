package main 
import "fmt"

func main(){
	fmt.Print("Mesma")
	fmt.Print("linha.")  

	fmt.Println("Nova ")  // "Println" irá colocar a próxima string impressa em uma nova linha
	fmt.Println("linha")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) Esse tipo de concatenação não é perimitido em GO 

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	// método mais fácil 
	fmt.Println("O valor de x é ", x)


	fmt.Printf("O valor de x é %.2f" , x) // o número 2, significa que eu quero imprimir apenas 2 casas decimais 
										  // %f significa que o que eu vou imprimir é do tipo float

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\nd %f %.1f %t %s", a, b, b, c, d)

}