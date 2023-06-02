package tickets

import (
	"log"
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
		precio, err := strconv.ParseFloat(e[5], 64)
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