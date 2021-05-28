package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

// Estructura de un satelite
type Satellite struct {
	Name     string   `json:name,omitempty`
	Distance float64  `json:distance,omitempty`
	Message  []string `json:message,omitempty`
}

// Estructura de arreglo de los satelites
type Satellites struct {
	Satellites []Satellite `json:satellites`
}

// Estructura de respuesta del API
type Response struct {
	Position Position `json:position,omitempty`
	Message  string   `json:message,omitempty`
}

// Estrucutra de position para la respuesta
type Position struct {
	X float64 `json:x,omitempty`
	Y float64 `json:y,omitempty`
}

// Estructura para la petición de TopSecretSplit
type RequestSplit struct {
	Distance float64  `json:distance,omitempty`
	Message  []string `json:message,omitempty`
}

var InformacionSatelites Satellites

// Funcion principal
func main() {

	// Creación de endpoints
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/topsecret", TopSecretEndpoint)
	http.HandleFunc("/topsecret_split", TopSecretSplit)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default port %s", port)
	}
	// log.Fatal(http.ListenAndServe(":3000", router))
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// Función que devuelve la localizacion del emisor en sus coordenadas
func GetLocation(distances [3]float64) [2]float64 {
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

	var positions [2]float64
	positions[0] = x
	positions[1] = y

	return positions
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
		return "Error"
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

		if indeterminado {
			mensajeTransmitido = "Error"
		}
		return mensajeTransmitido
	}
}

// Metodo del Endpoint
func Inicio(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Operación Fuego de Quasar\nAutor:Daniel Calderon")
}

// Metodo del Endpoint
func TopSecretEndpoint(w http.ResponseWriter, req *http.Request) {
	// Variable para capturar los datos recibidos
	var satelites Satellites
	_ = json.NewDecoder(req.Body).Decode(&satelites)
	response, tipo := ConsumirMetodos(satelites)
	if tipo == -1 {
		http.Error(w, "No fue posible determinar los datos", http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(response)
	}
}

// Metodo del Endpoint
func TopSecretSplit(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		if len(InformacionSatelites.Satellites) != 3 {
			http.Error(w, "No fue posible determinar los datos", http.StatusNotFound)
		} else {
			response, tipo := ConsumirMetodos(InformacionSatelites)
			if tipo == -1 {
				http.Error(w, "No hay suficiente información para deterimnar los datos", http.StatusNotFound)
			} else {
				json.NewEncoder(w).Encode(response)
			}
		}
	} else {
		keys := req.URL.Query()["satellite"]
		nombre := keys[0]
		if nombre == "kenobi" || nombre == "skywalker" || nombre == "sato" {
			var nuevoSatelite RequestSplit
			_ = json.NewDecoder(req.Body).Decode(&nuevoSatelite)
			var sateliteToAdd Satellite
			sateliteToAdd.Name = nombre
			sateliteToAdd.Distance = nuevoSatelite.Distance
			sateliteToAdd.Message = nuevoSatelite.Message
			sateliteexiste := -1
			if InformacionSatelites.Satellites != nil {
				for i := 0; i < len(InformacionSatelites.Satellites); i++ {
					if InformacionSatelites.Satellites[i].Name == sateliteToAdd.Name {
						sateliteexiste = i
					}
				}
			}
			if sateliteexiste == -1 {
				InformacionSatelites.Satellites = append(InformacionSatelites.Satellites, sateliteToAdd)
			} else {
				InformacionSatelites.Satellites[sateliteexiste] = sateliteToAdd
			}

			json.NewEncoder(w).Encode("Se actualizo la información del satelite: " + nombre + " correctamente")
		} else {
			http.Error(w, "El nombre de satelite ingresado no existe", http.StatusNotFound)
		}
	}

}

// Metodo de apoyo para consumir funciones
func ConsumirMetodos(datos Satellites) (Response, int) {
	// Variable para control de error
	var error int = 0

	// Variables para consumo de las funciones principales de la aplicación
	var distances [3]float64
	var messages [3]string
	for i := 0; i < 3; i++ {
		distances[i] = datos.Satellites[i].Distance
		var mensajeActual string
		for j := 0; j < len(datos.Satellites[i].Message); j++ {
			if datos.Satellites[i].Message[j] == " " || datos.Satellites[i].Message[j] == "" {
				mensajeActual += "''" + " "
			} else {
				mensajeActual += datos.Satellites[i].Message[j] + " "
			}
		}
		messages[i] = strings.TrimSpace(mensajeActual)
	}

	// Invoca las funciones que realizan las operaciones
	posicionEmisor := GetLocation(distances)
	mensajeEnviado := GetMessage(messages)

	// Crea la respuesta que retornara el metodo del API
	var response Response

	response.Position.X = posicionEmisor[0]
	response.Position.Y = posicionEmisor[1]
	response.Message = mensajeEnviado
	if mensajeEnviado == "Error" || error == -1 {
		return response, -1
	} else {
		return response, 0
	}
}
