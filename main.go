package main

import (
	"fmt"
	"github.com/matiasRaspa/desafio-go-bases-raspa-matias/internal/tickets"
	"log"
)

func main() {
	
	//Cargo datos
	listaTickets, err := tickets.CargarTickets("tickets.csv")
	if err != nil {
		log.Fatal(err)
	}
	//Compruebo datos
	fmt.Println(listaTickets[0].PaisDestino)

	//Requerimiento 1
	totalPersonasPorDestino, destino, err := tickets.ContarPersonasPorDestino(listaTickets, "Australia")
	if err != nil {
		log.Fatal(err)
	}
	//Muestro total
	fmt.Printf("\n############### Requerimiento 1 ###############\nEl total de personas que viajan a %s es: %d\n", destino, totalPersonasPorDestino)

	//Requerimiento 2
	madrugada, manana, tarde, noche, err := tickets.CantidadTicketsPorFranjaHoraria(listaTickets)
	if err != nil {
		log.Fatal(err)
	}
	//Muestro cantidad de pasajeros por franja horaria
	fmt.Printf("\n############### Requerimiento 2 ###############\nCantidad de pasajeros:\nPor la madrugada: %d\nPor la mañana: %d\nPor la tarde: %d\nPor la noche: %d\n", madrugada, manana, tarde, noche)

	//Requerimiento 3
	porcentaje, pais, err := tickets.PorcentajeDePasajerosPaisPorDia(listaTickets, "Australia")
	if err != nil {
		log.Fatal(err)
	}
	//Muestro porcentaje
	fmt.Printf("\n############### Requerimiento 3 ###############\nEl porcentaje de personas hacia %s en el dia de hoy es: %%%.1f\n", pais, porcentaje)
}
