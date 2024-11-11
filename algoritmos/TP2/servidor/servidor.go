package servidor

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	visitasPorIP          = make(map[string]int)
	solicitudesPorRecurso = make(map[string]int)
	sospechosasDoS        = make(map[string]bool)
	logEntries            []LogEntry
	timeLayout            = "2006-01-02T15:04:05-07:00"
)

type entrada struct {
	tiempo             time.Time
	metodo             string
	recurso            string
	cantidadPetisiones int
	denegado           bool
}

// Almacenaremos todo en un abb para tener las busquedas por rangos
type Servidor struct {
	visitantes TDADiccionario.DiccionarioOrdenado[string, entrada]
	visitados  TDADiccionario.Diccionario[string, int]
	denegados  []string
}

func CrearEstructuraDatos() DatosServidor {
	return &Servidor{visitantes: TDADiccionario.CrearABB[string, entrada](CompareIPs), visitados: TDADiccionario.CrearHash[string, int](), denegados: make([]string, 0)}
}

func (serv *Servidor) AgregarArchivo(archivo string) error {
	file, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("Error al abrir el archivo: %v", err)
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
		cant := 0
		deng := false

		if serv.visitantes.Pertenece(ip) {
			datos := serv.visitantes.Obtener(ip)

			// tengo que hacer una funcion para que me de en segundos la diferencia
			if tiempo-datos.tiempo >= 2 { // si la diferencia es mayor a 2 segundos
				cant = datos.cantidadPetisiones + 1
			}
			if cant >= 5 && !datos.denegado {
				serv.denegados = append(serv.denegados, ip)
				deng = true
			}
		}
		cantVist := 0
		if serv.visitados.Pertenece(recurso) {
			datos := serv.visitados.Obtener(ip)
			cantVist = datos + 1

		}
		serv.visitados.Guardar(recurso, cantVist)
		serv.visitantes.Guardar(ip, entrada{tiempo: tiempo, metodo: metodo, recurso: recurso, cantidadPetisiones: cant, denegado: deng})

	}
	TDAHeap.HeapSort(serv.denegados, CompareIPs)
	for i := len(serv.denegados) - 1; i >= 0; i-- {
		fmt.Printf(serv.denegados[i])
	}

	fmt.Println("OK")
	return nil
}

func (serv *Servidor) VerVisitantes(desde, hasta string) {
	serv.visitantes.IterarRango(&desde, &hasta, func(clave string, dato entrada) bool { fmt.Printf("\t%s\n", clave); return true })
	fmt.Printf("OK")
}

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
	}
	TDAHeap.HeapSort(arr, func() int {}) // funcion de comparacion para las cantidades
	for i := range n {
		fmt.Printf(arr[i].recurso)
	}
}

// Procesa un archivo de log y detecta posibles ataques de DoS.
func AgregarArchivo(nombre_archivo string) error {
	file, err := os.Open(nombre_archivo)
	if err != nil {
		return fmt.Errorf("Error al abrir el archivo: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ipRequests := make(map[string][]time.Time) // Para detectar DoS
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) != 4 {
			return fmt.Errorf("Formato incorrecto en el log: %s", line)
		}

		ip := parts[0]
		dateTime, err := time.Parse(timeLayout, parts[1])
		if err != nil {
			return fmt.Errorf("Error al parsear fecha: %v", err)
		}
		resource := parts[3]

		logEntries = append(logEntries, LogEntry{IP: ip, DateTime: dateTime, Resource: resource})
		visitasPorIP[ip]++
		solicitudesPorRecurso[resource]++

		// Detectar posibles ataques de DoS
		ipRequests[ip] = append(ipRequests[ip], dateTime)
		if len(ipRequests[ip]) >= 5 {
			if ipRequests[ip][len(ipRequests[ip])-1].Sub(ipRequests[ip][len(ipRequests[ip])-5]) < 2*time.Second {
				if !sospechosasDoS[ip] {
					fmt.Printf("DoS: %s\n", ip)
					sospechosasDoS[ip] = true
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error al leer el archivo: %v", err)
	}

	fmt.Println("OK")
	return nil
}

// Muestra todas las IPs que realizaron alguna petición dentro de un rango.
func VerVisitantes(desde, hasta string) {
	ipList := make(map[string]bool)

	for _, entry := range logEntries {
		if CompareIPs(entry.IP, desde) >= 0 && CompareIPs(entry.IP, hasta) <= 0 {
			ipList[entry.IP] = true
		}
	}

	var sortedIPs []string
	for ip := range ipList {
		sortedIPs = append(sortedIPs, ip)
	}
	sort.Slice(sortedIPs, func(i, j int) bool {
		return CompareIPs(sortedIPs[i], sortedIPs[j]) < 0
	})

	fmt.Println("Visitantes:")
	for _, ip := range sortedIPs {
		fmt.Printf("\t%s\n", ip)
	}
	fmt.Println("OK")
}

// Muestra los n recursos más visitados.
func VerMasVisitados(n int) {
	type recurso struct {
		nombre  string
		visitas int
	}
	var recursos []recurso
	for nombre, visitas := range solicitudesPorRecurso {
		recursos = append(recursos, recurso{nombre, visitas})
	}
	sort.Slice(recursos, func(i, j int) bool {
		return recursos[i].visitas > recursos[j].visitas
	})

	fmt.Println("Sitios más visitados:")
	for i := 0; i < n && i < len(recursos); i++ {
		fmt.Printf("\t%s - %d\n", recursos[i].nombre, recursos[i].visitas)
	}
	fmt.Println("OK")
}

// Compara dos IPs numéricamente.
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

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
