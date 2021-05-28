package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Función principal del programa
func main() {

	// Variable para leer el texto ingresado
	scanner := bufio.NewScanner(os.Stdin)

	// Variable para controlar las opciones de ejecución del programa
	var opcion string

	// Ciclo que se repite hasta que el usuario seleccione la opcion de salir en el menú
	for {
		// Muestra las opciones del menú
		fmt.Print("Operación secreta Fuego de Quasar\n\n")
		fmt.Println("1) Obtener localización del emisor")
		fmt.Println("2) Obtener mensaje transmitido")
		fmt.Println("3) Salir del sistema")

		// Captura la opción seleccionada por el usuario
		fmt.Print("\nPor favor ingresa la opción para el procesamiento del sistema: ")
		fmt.Scanln(&opcion)

		/*
			** Si la opción es 3 entonces se cierra el programa, de lo contrario según la opción seleccionada
				el programa sigue el flujo normal
		*/
		if opcion == "3" {
			fmt.Println("Gracias por salvar el universo!")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			break
		} else {
			switch opcion {
			// Si se selecciona la opción para calcular la distancia del emisor entonces procedemos a capturar la distancia a cada satelite
			case "1":
				var distanceKenobi float64
				var distanceSkywalker float64
				var distanceSato float64
				fmt.Print("Ingresar la distancia desde el satelite Kenobi: ")
				fmt.Scanln(&distanceKenobi)
				fmt.Print("Ingresar la distancia desde el satelite Skywalker: ")
				fmt.Scanln(&distanceSkywalker)
				fmt.Print("Ingresar la distancia desde el satelite Sato: ")
				fmt.Scanln(&distanceSato)

				// Guardar variables en arreglo
				var distances [3]float64
				distances[0] = distanceKenobi
				distances[1] = distanceSkywalker
				distances[2] = distanceSato

				// Invocar función GetLocation
				fmt.Println(GetLocation(distances[:]))
				bufio.NewReader(os.Stdin).ReadBytes('\n')

			// Si se selecciona la opcion de obtener el mensaje se procede a solicitar el mensaje recibido en cada satelite
			case "2":
				var messages [3]string
				fmt.Println("Para ingresar el mensaje recibido en cada satelite debe diligenciar el texto normalmente, en caso de que una palabra no este determinada por favor ingrese lo siguiente ''")
				fmt.Print("Ingresar el mensaje recibido en el satelite Kenobi: ")
				scanner.Scan()
				messages[0] = scanner.Text()
				fmt.Print("Ingresar el mensaje recibido en el satelite Skywalker: ")
				scanner.Scan()
				messages[1] = scanner.Text()
				fmt.Print("Ingresar el mensaje recibido en el satelite Sato: ")
				scanner.Scan()
				messages[2] = scanner.Text()

				// Invocar función GetMessage
				fmt.Println(GetMessage(messages))
				bufio.NewReader(os.Stdin).ReadBytes('\n')
			default:
				fmt.Println("La opción seleccionada no está disponible!")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
			}
		}

	}
}

// Función para obtener la localización del emisor
func GetLocation(distances []float64) string {
	// Posiciones actuales de los satelites
	kenobi := [2]float64{-500, -200}
	skywalker := [2]float64{100, -100}
	sato := [2]float64{500, 100}

	// Variables a utilizar para uso del algoritmo
	d1 := distances[0]
	d2 := distances[1]
	d3 := distances[2]
	i1 := kenobi[0]
	i2 := skywalker[0]
	i3 := sato[0]
	j1 := kenobi[1]
	j2 := skywalker[1]
	j3 := sato[1]

	// Se calcula el valor de las coordenadas x,y según algoritmo de Trilateración
	x := (((math.Pow(d1, 2)-math.Pow(d2, 2))+(math.Pow(i2, 2)-math.Pow(i1, 2))+(math.Pow(j2, 2)-math.Pow(j1, 2)))*(2*j3-2*j2) - ((math.Pow(d2, 2)-math.Pow(d3, 2))+(math.Pow(i3, 2)-math.Pow(i2, 2))+(math.Pow(j3, 2)-math.Pow(j2, 2)))*(2*j2-2*j1)) / ((2*i2-2*i3)*(2*j2-2*j1) - (2*i1-2*i2)*(2*j3-2*j2))

	y := ((math.Pow(d1, 2) - math.Pow(d2, 2)) + (math.Pow(i2, 2) - math.Pow(i1, 2)) + (math.Pow(j2, 2) - math.Pow(j1, 2)) + x*(2*i1-2*i2)) / (2*j2 - 2*j1)

	println(x)
	println(y)
	s := fmt.Sprintf("La posicion del emisor es: (%f;%f)", x, y)
	return s
}

// Función para obtener el mensaje transmitido
func GetMessage(messages [3]string) string {
	// Variable de control para saber si no se pudo determinar el mensaje de origen
	var indeterminado bool = false

	// Mensajes de cada uno de los satelites convertidos en Arreglos de String, el caracter delimitador es " " (espacio en blanco)
	messageKenobi := strings.Split(messages[0], " ")
	messageSkywalker := strings.Split(messages[1], " ")
	messageSato := strings.Split(messages[2], " ")

	// Mensaje transmitido que retornara la función
	var mensajeTransmitido string

	// Verificamos que la longitud de los mensajes en cada uno de los satelites sea igual
	if len(messageKenobi) != len(messageSkywalker) || len(messageKenobi) != len(messageSato) {
		return "\nNo es posible determinar el mensaje de origen, ya que el mensaje recibido por cada satelite es de diferente longitud!"
	} else {
		// Ciclo principal para comparar cada una de las palabras recibidas en los satelites
		for i := 0; i < len(messageKenobi); i++ {
			// Variable para guardar cada una de las palabras
			var palabras [3]string
			palabras[0] = messageKenobi[i]
			palabras[1] = messageSkywalker[i]
			palabras[2] = messageSato[i]

			// Variable para controlar las veces que se repite una palabra
			var repeticiones int = 0
			// Variable para almacenar la primera palabra encontrada diferente de ""
			var palabraActual string = ""
			// Ciclo para buscar en los tres satelites las palabras recibidas en el mensaje
			for k := 0; k < 3; k++ {
				// Verificar que no sea indeterminada
				if palabras[k] != "''" {
					// Si aun no hemos encontrado la primera palabra
					if palabraActual == "" {
						// Se aumenta la cantidad de repeticiones
						palabraActual = palabras[k]
						repeticiones++
					} else {
						/*Si la palabra coincide en otro satelite entonces aumento las repeticiones,
						de lo contrario no se puede determinar en esa posicion porque hay dos palabras*/
						if palabraActual == palabras[k] {
							repeticiones++
						} else {
							repeticiones--
						}
					}
				}
			}
			/* Si hay mas de una repeticion entonces la palabra se pudo determinar,
			de lo contrario se reemplaza con XXXX
			*/
			if repeticiones >= 1 {
				mensajeTransmitido = mensajeTransmitido + palabraActual + " "
			} else {
				indeterminado = true
				mensajeTransmitido = mensajeTransmitido + "xxxxx "
			}
		}

		if indeterminado == true {
			mensajeTransmitido = "No es posible determinar el mensaje de origen: " + mensajeTransmitido
		} else {
			mensajeTransmitido = "El mensaje transmitido fue: " + mensajeTransmitido
		}
		return mensajeTransmitido
	}
}
