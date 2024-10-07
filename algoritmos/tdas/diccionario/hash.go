package diccionario

import (
	"fmt"
	"hash/fnv"
	"tdas/lista"
)

const capacidadInicial = 10
const factorCargaMaximo = 0.75 // Factor de carga, donde era 1 creo el max pero teniamos que poner menos por las dudas no?

type entradaDiccionario[K comparable, V any] struct {
	clave K
	dato  V
}

type diccionarioHashAbierto[K comparable, V any] struct {
	tablas   []lista.Lista[entradaDiccionario[K, V]] // Listas enlazadas para colisiones
	cantidad int
}

type iteradorDiccionarioHashAbierto[K comparable, V any] struct {
	diccionario *diccionarioHashAbierto[K, V]
	indice      int
	iterLista   lista.IteradorLista[entradaDiccionario[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &diccionarioHashAbierto[K, V]{
		tablas: make([]lista.Lista[entradaDiccionario[K, V]], capacidadInicial),
	}
	for i := 0; i < capacidadInicial; i++ {
		hash.tablas[i] = lista.CrearListaEnlazada[entradaDiccionario[K, V]]()
	}
	return hash
}

// calcula el hash con fnv que teoricamente anda bien, pero no se que fuente pondriamos
// en la pagina oficial de GO estaria igualmente:  https://pkg.go.dev/hash/fnv#New128a
func calcularHash[K comparable](clave K) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", clave)))
	return h.Sum32()
}

func (diccionario *diccionarioHashAbierto[K, V]) redimensionar(nuevaCapacidad int) {

	nuevasTablas := make([]lista.Lista[entradaDiccionario[K, V]], nuevaCapacidad)

	for i := 0; i < nuevaCapacidad; i++ {
		nuevasTablas[i] = lista.CrearListaEnlazada[entradaDiccionario[K, V]]()
	}

	for i := range diccionario.tablas {

		diccionario.tablas[i].Iterar(func(nodo entradaDiccionario[K, V]) bool {

			nuevoHash := calcularHash(nodo.clave) % uint32(nuevaCapacidad)
			nuevasTablas[nuevoHash].InsertarUltimo(nodo)
			return true
		})
	}
	diccionario.tablas = nuevasTablas
}

func (diccionario *diccionarioHashAbierto[K, V]) Guardar(clave K, dato V) {

	if float64(diccionario.cantidad)/float64(len(diccionario.tablas)) > factorCargaMaximo {
		diccionario.redimensionar(len(diccionario.tablas) * 2) // Duplicamo ?
	}

	hash := calcularHash(clave) % uint32(len(diccionario.tablas))
	lista := diccionario.tablas[hash]

	iter := lista.Iterador()
	actual := false

	for iter.HaySiguiente() {
		entrada := iter.VerActual()

		if entrada.clave == clave {

			iter.Borrar()
			actual = true
			break // no se si usar los breaks
		}
		iter.Siguiente()
	}

	if actual {
		lista.InsertarUltimo(entradaDiccionario[K, V]{clave, dato})
	} else {
		// Si no se encontró la clave, se agrega como nueva entrada
		lista.InsertarUltimo(entradaDiccionario[K, V]{clave, dato})
		diccionario.cantidad++
	}
}

func (diccionario *diccionarioHashAbierto[K, V]) Obtener(clave K) V {

	hash := calcularHash(clave) % uint32(len(diccionario.tablas))
	lista := diccionario.tablas[hash]

	var valor V
	encontrado := false

	lista.Iterar(func(nodo entradaDiccionario[K, V]) bool {

		if nodo.clave == clave {
			valor = nodo.dato
			encontrado = true
			return false
		}
		return true
	})

	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	return valor
}

func (diccionario *diccionarioHashAbierto[K, V]) Pertenece(clave K) bool {

	hash := calcularHash(clave) % uint32(len(diccionario.tablas))
	lista := diccionario.tablas[hash]

	encontrado := false
	lista.Iterar(func(nodo entradaDiccionario[K, V]) bool {

		if nodo.clave == clave {
			encontrado = true
			return false
		}
		return true
	})

	return encontrado
}

func (diccionario *diccionarioHashAbierto[K, V]) Borrar(clave K) V {

	hash := calcularHash(clave) % uint32(len(diccionario.tablas))
	lista := diccionario.tablas[hash]

	var valor V
	encontrado := false
	iter := lista.Iterador()

	for iter.HaySiguiente() {

		if iter.VerActual().clave == clave {
			valor = iter.VerActual().dato
			iter.Borrar()
			encontrado = true
			diccionario.cantidad--
			break
		}

		iter.Siguiente()
	}

	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	return valor
}

func (diccionario *diccionarioHashAbierto[K, V]) Cantidad() int {
	return diccionario.cantidad
}

/**************  TODO ESTO MAL, LO HIZO CHAT GPT Y NO PASA LOS TESTS, fallan unos 4 tests maso **************/

// Iterar aplica una función a cada par clave-dato en el diccionario.
func (d *diccionarioHashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	for _, lista := range d.tablas {
		lista.Iterar(func(nodo entradaDiccionario[K, V]) bool {
			return visitar(nodo.clave, nodo.dato)
		})
	}
}

// Iterador devuelve un iterador para el diccionario.
func (d *diccionarioHashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorDiccionarioHashAbierto[K, V]{
		diccionario: d,
		indice:      0,
		iterLista:   d.tablas[0].Iterador(), // Podría necesitar ajustes si la tabla está vacía
	}
	iter.avanzar() // Mueve al primer elemento válido
	return iter
}

// avanzar mueve el iterador a la siguiente lista con entradas.
func (iter *iteradorDiccionarioHashAbierto[K, V]) avanzar() {
	for iter.indice < len(iter.diccionario.tablas) {
		if iter.iterLista.HaySiguiente() {
			return // Si hay un siguiente elemento, no hacemos nada más.
		}
		iter.indice++
		if iter.indice < len(iter.diccionario.tablas) {
			iter.iterLista = iter.diccionario.tablas[iter.indice].Iterador()
		}
	}
}

// HaySiguiente verifica si hay un siguiente elemento.
func (iter *iteradorDiccionarioHashAbierto[K, V]) HaySiguiente() bool {
	return iter.indice < len(iter.diccionario.tablas) && iter.iterLista.HaySiguiente()
}

// VerActual devuelve el elemento actual del iterador.
func (iter *iteradorDiccionarioHashAbierto[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador terminó de iterar")
	}
	nodo := iter.iterLista.VerActual()
	return nodo.clave, nodo.dato
}

// Siguiente avanza al siguiente elemento en el iterador.
func (iter *iteradorDiccionarioHashAbierto[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador terminó de iterar")
	}
	iter.iterLista.Siguiente()
	iter.avanzar() // Solo avanzar si hay un siguiente elemento.
}
