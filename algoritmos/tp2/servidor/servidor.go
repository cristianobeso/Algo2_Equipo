package servidor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	"time"
)

type LogEntry struct {
	IP       string
	DateTime time.Time
	Resource string
}

var (
	timeLayout = "2006-01-02T15:04:05-07:00"
)

type entrada struct {
	tiempo             time.Time
	metodo             string
	recurso            string
	cantidadPetisiones int
}

// Almacenaremos todo en un abb para tener las busquedas por rangos
type Servidor struct {
	visitantes TDADiccionario.DiccionarioOrdenado[string, entrada]
	visitados  TDADiccionario.Diccionario[string, int]
	denegados  TDADiccionario.Diccionario[string, string]
}

/*
Precondiciones: No requiere parámetros ni dependencias previas.

Postcondiciones: Devuelve una instancia de Servidor inicializada con, un diccionario ordenado para visitantes
y un diccionario hash para visitados.
*/
func CrearEstructuraDatos() DatosServidor {
	return &Servidor{visitantes: TDADiccionario.CrearABB[string, entrada](CompareIPs), visitados: TDADiccionario.CrearHash[string, int]()}
}

/*
Precondiciones: El archivo especificado por el parámetro archivo debe existir y ser accesible.
El formato del archivo debe consistir en líneas tabuladas con cuatro columnas, Dirección IP (string),
Fecha y hora en el formato "2006-01-02T15:04:05-07:00", Método de la solicitud (string) y Recurso solicitado (string).

Postcondiciones: Los visitantes y sus datos asociados se almacenan en visitantes.
Los recursos solicitados y sus conteos se actualizan en visitados.
Las IPs sospechosas de DoS se almacenan en denegados.
Si no hay errores, imprime DoS para las IPs sospechosas en orden y retorna nil.
Si hay errores en el archivo o el formato, retorna un error descriptivo.
*/
func (serv *Servidor) AgregarArchivo(archivo string) error {
	serv.denegados = TDADiccionario.CrearHash[string, string]()
	file, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) != 4 {
			return fmt.Errorf("Formato incorrecto en el log: %s", line)
		}

		ip := parts[0]
		tiempo, err := time.Parse(timeLayout, parts[1])
		if err != nil {
			return fmt.Errorf("Error al parsear fecha: %v", err)
		}
		metodo := parts[2]
		recurso := parts[3]
		cant := 1

		if serv.visitantes.Pertenece(ip) {
			datos := serv.visitantes.Obtener(ip)

			if tiempo.Sub(datos.tiempo).Seconds() < 2 && tiempo.Sub(datos.tiempo).Seconds() >= 0 {
				cant = datos.cantidadPetisiones + 1
				tiempo = datos.tiempo
				if cant >= 5 && !serv.denegados.Pertenece(ip) {
					serv.denegados.Guardar(ip, ip)

				}
			}

		}
		cantVist := 1
		if serv.visitados.Pertenece(recurso) {
			datos := serv.visitados.Obtener(recurso)
			cantVist = datos + 1

		}
		serv.visitados.Guardar(recurso, cantVist)
		serv.visitantes.Guardar(ip, entrada{tiempo: tiempo, metodo: metodo, recurso: recurso, cantidadPetisiones: cant})

	}

	iter := serv.denegados.Iterador()
	arr := make([]string, 0)
	for iter.HaySiguiente() {
		_, ip_g := iter.VerActual()
		arr = append(arr, ip_g)
		iter.Siguiente()
	}
	TDAHeap.HeapSort(arr, CompareIPs)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("DoS: %s\n", arr[i])
	}

	fmt.Println("OK")
	return nil
}

/*
Precondiciones: desde y hasta deben ser direcciones IP válidas.

Postcondiciones: Imprime en la consola todas las IPs entre el rango [desde, hasta] presentes en el diccionario visitantes.
*/
func (serv *Servidor) VerVisitantes(desde, hasta string) {
	fmt.Println("Visitantes:")
	serv.visitantes.IterarRango(&desde, &hasta, func(clave string, dato entrada) bool { fmt.Printf("\t%s\n", clave); return true })
	fmt.Println("OK")
}

/*
Precondiciones: n debe ser un número entero mayor o igual a 1.
visitados debe contener al menos n recursos distintos si se desean mostrar n elementos.

Postcondiciones: Imprime los n recursos más visitados en la consola junto con sus conteos.
Si hay menos de n recursos, muestra todos los disponibles.
*/
func (serv *Servidor) VerMasVisitados(n int) {
	type Dato struct {
		cantidad int
		recurso  string
	}
	iter := serv.visitados.Iterador()
	arr := make([]Dato, 0)
	for iter.HaySiguiente() {
		rec, cant := iter.VerActual()
		arr = append(arr, Dato{cantidad: cant, recurso: rec})
		iter.Siguiente()
	}
	TDAHeap.HeapSort(arr, func(elemento1, elemento2 Dato) int { return elemento2.cantidad - elemento1.cantidad })
	fmt.Println("Sitios más visitados:")
	for i := 0; i < len(arr); i++ {
		if i >= n {
			break
		}
		fmt.Printf("\t%s - %d\n", arr[i].recurso, arr[i].cantidad)
	}
	fmt.Println("OK")

}

/*
Precondiciones: ip1 y ip2 deben ser cadenas que representen direcciones IP.

Postcondiciones: Devuelve,un valor negativo si ip1 es menor que ip2, cero si son iguales,
un valor positivo si ip1 es mayor que ip2.
*/
func CompareIPs(ip1, ip2 string) int {
	parts1 := strings.Split(ip1, ".")
	parts2 := strings.Split(ip2, ".")
	for i := 0; i < 4; i++ {
		num1, _ := strconv.Atoi(parts1[i])
		num2, _ := strconv.Atoi(parts2[i])
		if num1 != num2 {
			return num1 - num2
		}
	}
	return 0
}

/*
Precondiciones: `s` debe ser una cadena de texto que represente un número entero válido en base 10.

Postcondiciones: Devuelve el número entero representado por la cadena `s`. Si la cadena no es un número
entero válido, devuelve 0.
*/
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
