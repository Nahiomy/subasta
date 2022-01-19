package main

import "fmt"

func main() {
	fmt.Println("INGRESA EL NOMBRE DEL PRODUCTO A SUBASTAR")
	var nombre string
	fmt.Scanf("%s", &nombre)

	fmt.Println("INGRESA EL VALOR INICIAL")
	var valor_prod int
	fmt.Scanf("%d", &valor_prod)

	fmt.Println("QUE INICIE LA SUBASTA!!!")

	for {
		fmt.Println("EL VALOR ACTUAL DEL PRODUCTO ES DE ", valor_prod)

		valor_prod += 50

		fmt.Println("¿ALGUIEN DA MÁS POR ", valor_prod, "?")
		fmt.Println("¿DESEA SEGUIR?")
		fmt.Println("DIGITE SI PARA CONTINUAR", "DIGITE NO PARA FINALIZAR")
		var sino string
		fmt.Scanf("%s", &sino)
		if sino == "SI" {
			fmt.Println("EL VALOR ACTUAL DEL PRODUCTO ES DE ", valor_prod)
			valor_prod += 50
			fmt.Println("¿ALGUIEN DA MÁS POR ", valor_prod, "?")

		} else if sino == "NO" {
			fmt.Println("VENDIDO POR ", valor_prod)
			break
		}
	}

}
