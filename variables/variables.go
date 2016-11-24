package main

import "fmt"



func main() {
	//Impresión de expresiones
	fmt.Println("########## Imprime algunas expresiones ############")
	fmt.Println("Hola " + "Mundo")
	fmt.Println("2+3 ->", 2+3)
	fmt.Println("11.0/7.0 ->", 11.0/7.0)
	fmt.Println("true && false ->", false)
	fmt.Println("true || false ->", true)
	fmt.Println("!true ->", !true)
	fmt.Println("###################################################")
	fmt.Println()

	//Declaraciones de variables
	fmt.Println("########## Declaraciones de variables ############")

	var a string = "Esta es una variable string"
	fmt.Println(a)

	var b, c int = 3, 4
	fmt.Println(b, c)

	var e bool = true
	fmt.Println(e)

	f := "Esta es otra forma de declarar una variable"
	fmt.Println(f)

	fmt.Println("###################################################")
	fmt.Println()

	//Declaraciones de variables
	fmt.Println("######### Uso de constantes ############")

	const G string = "Esta es una constante string"
	const H = 5000000
	const I = 3e20/H

	fmt.Println(G)
	fmt.Println(I)
	fmt.Println(int64(I))
	fmt.Println("###################################################")


}
