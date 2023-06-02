package tickets

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Estructura para manejar cada ticket
type Ticket struct {
	ID          int
	Nombre      string
	Email       string
	PaisDestino string
	HoraVuelo   string
	Precio      float64
}

// Funcion para manejo de datos
func CargarTickets(archivo string) ([]Ticket, error) {

	//Leo archivo
	res, err := os.ReadFile(archivo)
	if err != nil {
		log.Fatal(err)
	}

	//Separo mis objetos por cada salto de linea
	datos := strings.Split(string(res), "\n")

	//Declaro lista de tickets
	var tickets []Ticket

	//Recorro datos
	for _, v := range datos {

		//Separo mis elementos por cada coma
		e := strings.Split(v, ",")

		//Convierto string a int
		id, err := strconv.Atoi(e[0])
		if err != nil {
			log.Fatal(err)
		}
		nombre := e[1]
		email := e[2]
		paisDestino := e[3]
		horaVuelo := e[4]
		//Convierto string a float
		parseando := strings.TrimSpace(e[5])
		precio, err := strconv.ParseFloat(parseando, 64)
		if err != nil {
			log.Fatal(err)
		}

		//Guardo valores
		ticket := Ticket{
			ID:          id,
			Nombre:      nombre,
			Email:       email,
			PaisDestino: paisDestino,
			HoraVuelo:   horaVuelo,
			Precio:      precio,
		}

		tickets = append(tickets, ticket)
	}

	//Retorno lista de tickets
	return tickets, nil
}

// Requerimiento 1
func ContarPersonasPorDestino(tickets []Ticket, pais string) (int, string, error) {

	//Declaro contador
	contador := 0
	//Declaro pais elegido para mostrar en mensaje
	paisElegido := pais

	//Recorro tickets
	for _, ticket := range tickets {
		if ticket.PaisDestino == pais {
			contador++
		}
	}

	//Retorno contador
	return contador, paisElegido, nil
}

// Requerimiento 2
func CantidadTicketsPorFranjaHoraria(tickets []Ticket) (int, int, int, int, error) {

	//Declaro contadores
	madrugada, manana, tarde, noche := 0, 0, 0, 0

	//Recorro tickets
	for _, ticket := range tickets {
		//Obtengo hora y convierto a int
		fraccionarHora := strings.Split(ticket.HoraVuelo, ":")
		hora, err := strconv.Atoi(fraccionarHora[0])
		if err != nil {
			log.Fatal(err)
		}

		//Comparo hora: madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19), y noche (20 → 23)
		switch {
		case hora >= 0 && hora <= 6:
			madrugada++
		case hora >= 7 && hora <= 12:
			manana++
		case hora >= 13 && hora <= 19:
			tarde++
		case hora >= 20 && hora <= 23:
			noche++
		}
	}

	//Retorno contadores
	return madrugada, manana, tarde, noche, nil
}

// Requerimiento 3
func PorcentajeDePasajerosPaisPorDia(tickets []Ticket, pais string) (float64, string, error) {

	//Declaro contador
	contador := 0
	//Declaro pais elegido para mostrar en mensaje
	paisElegido := pais

	//Recorro tickets
	for _, ticket := range tickets {
		if ticket.PaisDestino == pais {
			contador++
		}
	}

	//Calculo porcentaje
	porcentaje := float64(contador) * 100 / float64(len(tickets))

	//Retorno porcentaje
	return porcentaje, paisElegido, nil
}

// Requerimiento 4
func ProcesarFunciones(listaTickets []Ticket, canal ...chan string) {
	//Genero una lista con los paises de todos los tickets
	var listaDePaises []string
	for _, ticket := range listaTickets {
		listaDePaises = append(listaDePaises, ticket.PaisDestino)
	}
	//Genero un mapa para quitar los paises repetidos de la lista
	elementosUnicos := make(map[string]bool)
	listaDePaisesSinRepetir := []string{}

	//Agrego paises y si esta repetido lo ignoro
	for _, consultarPaisRepetido := range listaDePaises {
		if !elementosUnicos[consultarPaisRepetido] {
			elementosUnicos[consultarPaisRepetido] = true
			listaDePaisesSinRepetir = append(listaDePaisesSinRepetir, consultarPaisRepetido)
		}
	}

	//Genero un bucle que puede ser infinito pero lo seteo en 3
	for i := 0; i < 3; i++ {
		//Genero un numero aleatorio que va a representar un pais al azar
		numeroAleatorio := rand.Intn(len(listaDePaisesSinRepetir))

		//Requerimiento 1
		totalPersonasPorDestino, destino, err := ContarPersonasPorDestino(listaTickets, listaDePaisesSinRepetir[numeroAleatorio])
		if err != nil {
			log.Fatal(err)
		}
		//Muestro total
		canal[0] <- fmt.Sprintf("\n############### Requerimiento 1 (como GOROUTINE) ###############\nEl total de personas que viajan a %s es: %d\n", destino, totalPersonasPorDestino)

		//Requerimiento 2
		madrugada, manana, tarde, noche, err := CantidadTicketsPorFranjaHoraria(listaTickets)
		if err != nil {
			log.Fatal(err)
		}
		//Muestro cantidad de pasajeros por franja horaria
		canal[1] <- fmt.Sprintf("\n############### Requerimiento 2 (como GOROUTINE) ###############\nCantidad de pasajeros:\nPor la madrugada: %d\nPor la mañana: %d\nPor la tarde: %d\nPor la noche: %d\n", madrugada, manana, tarde, noche)

		//Requerimiento 3
		porcentaje, pais, err := PorcentajeDePasajerosPaisPorDia(listaTickets, listaDePaisesSinRepetir[numeroAleatorio])
		if err != nil {
			log.Fatal(err)
		}
		//Muestro porcentaje
		canal[2] <- fmt.Sprintf("\n############### Requerimiento 3 (como GOROUTINE) ###############\nEl porcentaje de personas hacia %s en el dia de hoy es: %%%.1f\n", pais, porcentaje)
	}
}
