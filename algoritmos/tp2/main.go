package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp2/servidor"
)

/*
Precondiciones: Entrada estándar habilitada con comandos válidos en los formatos:
 1. "agregar_archivo <archivo>" - El archivo debe existir y ser accesible.
 2. "ver_visitantes <IP_desde> <IP_hasta>" - Ambas deben ser IPs válidas.
 3. "ver_mas_visitados <n>" - `n` debe ser un número entero positivo.

Postcondiciones:
- Ejecuta el comando correspondiente o imprime un mensaje de error en `stderr` si el comando es inválido o falla.
- Finaliza al encontrar un error o al procesar todos los comandos correctamente.
*/
func main() {
	datServ := servidor.CrearEstructuraDatos()
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
			err := datServ.AgregarArchivo(parts[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		case "ver_visitantes":
			if len(parts) != 3 {
				fmt.Fprintln(os.Stderr, "Error en comando ver_visitantes")
				return
			}
			datServ.VerVisitantes(parts[1], parts[2])
		case "ver_mas_visitados":
			if len(parts) != 2 {
				fmt.Fprintln(os.Stderr, "Error en comando ver_mas_visitados")
				return
			}
			n := servidor.Atoi(parts[1])
			datServ.VerMasVisitados(n)
		default:
			fmt.Fprintf(os.Stderr, "Error en comando %s\n", parts[0])
			return
		}
	}
}
