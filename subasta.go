package main

import "fmt"

func main() {
	fmt.Println("INGRESA EL NOMBRE DEL PRODUCTO A SUBASTAR")
	var nombre string
	fmt.Scanf("%s", &nombre)

	fmt.Println("INGRESA EL VALOR INICIAL")
	var valor_inicial int
	_, err := fmt.Scanf("%d", &valor_inicial)
	if err != nil {
		panic(err)
	}

	fmt.Println("QUE INICIE LA SUBASTA!!!")

	valor_actual := valor_inicial
	valor_siguiente := 0

	for {
		fmt.Println("EL VALOR ACTUAL DEL PRODUCTO ES DE ", valor_actual)

		valor_siguiente = valor_actual + 50

		fmt.Println("¿ALGUIEN DA MÁS POR ", valor_siguiente, "?")
		fmt.Println("¿DESEA SEGUIR?")
		fmt.Println("DIGITE SI PARA CONTINUAR", "DIGITE NO PARA FINALIZAR")
		var sino string
		fmt.Scanf("%s", &sino)
		if sino == "SI" {

			valor_actual = valor_siguiente
			fmt.Println("¿ALGUIEN DA MÁS POR ", valor_actual, "?")

		} else if sino == "NO" {
			fmt.Println("VENDIDO POR ", valor_actual)
			break
		}
	}

}
