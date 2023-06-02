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
}
