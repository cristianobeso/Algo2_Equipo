package main

import (
	"TP2/servidor"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "agregar_archivo":
			if len(parts) != 2 {
				fmt.Fprintln(os.Stderr, "Error en comando agregar_archivo")
				return
			}
			err := servidor.AgregarArchivo(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		case "ver_visitantes":
			if len(parts) != 3 {
				fmt.Fprintln(os.Stderr, "Error en comando ver_visitantes")
				return
			}
			servidor.VerVisitantes(parts[1], parts[2])
		case "ver_mas_visitados":
			if len(parts) != 2 {
				fmt.Fprintln(os.Stderr, "Error en comando ver_mas_visitados")
				return
			}
			n := servidor.Atoi(parts[1])
			servidor.VerMasVisitados(n)
		default:
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", parts[0])
			return
		}
	}
}
