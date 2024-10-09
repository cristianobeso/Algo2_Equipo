package diccionario

import (
	"fmt"
	"hash/fnv"
	"tdas/lista"
)

const capacidadInicial = 5
const factorCargaMaximo = 2.3

type entradaDiccionario[K comparable, V any] struct {
	clave K
	dato  V
}

type diccionarioHashAbierto[K comparable, V any] struct {
	tablas   []lista.Lista[entradaDiccionario[K, V]]
	cantidad int
	tam      int
}

type iteradorDiccionarioHashAbierto[K comparable, V any] struct {
	diccionario *diccionarioHashAbierto[K, V]
	indice      int
	iterLista   lista.IteradorLista[entradaDiccionario[K, V]]
}

// Esta funcion devuelve el primo que sea mayor al doble del primo recivido
// Funciona mejor si se comienza con 5
// Hasta 409.597 siendo este primo se acierta con 90%
// En adelante se acierta con 30%
// Ultimo primo verificado 1.677.721.597
// Mejor forma que encontramos para que sea de tiempo constante
func posiblePrimo(num int) int {
	nuevoNumero := num*2 + 1
	if nuevoNumero%10 == 5 {
		nuevoNumero += 2
	}
	return nuevoNumero
}

/*
Precondiciones: Ninguna.

Postcondiciones: Devuelve un diccionario hash vacío con capacidad inicial definida. Las tablas están
inicializadas como listas enlazadas vacías.
*/
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &diccionarioHashAbierto[K, V]{
		tablas: make([]lista.Lista[entradaDiccionario[K, V]], capacidadInicial), tam: capacidadInicial,
	}
	for i := 0; i < capacidadInicial; i++ {
		hash.tablas[i] = lista.CrearListaEnlazada[entradaDiccionario[K, V]]()
	}
	return hash
}

/*
Precondiciones: La clave debe ser un valor comparable.

Postcondiciones: Devuelve la representación en bytes de la clave.
*/
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

/*
Precondiciones: La clave debe ser un valor comparable.

Postcondiciones: Devuelve el hash de la clave como un valor uint32, utilizando el algoritmo FNV.
En la pagina oficial de GO:  https://pkg.go.dev/hash/fnv#New128a
*/
func calcularHash[K comparable](clave K) uint32 {
	h := fnv.New32a()
	h.Write(convertirABytes(clave))
	return h.Sum32()
}

/*
Precondiciones: nuevaCapacidad debe ser un valor positivo mayor que 0.

Postcondiciones: Redimensiona el diccionario hash duplicando su capacidad, moviendo las entradas
existentes a la nueva tabla según su nuevo índice de hash.
*/
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
	diccionario.tam = nuevaCapacidad
}

/*
Precondiciones: La clave debe ser un valor comparable.
El diccionario no debe estar lleno (el factor de carga no debe superar el máximo).

Postcondiciones: Si la clave ya existe, el dato asociado es reemplazado.
Si la clave no existe, se inserta un nuevo par clave-dato y se incrementa la cantidad de elementos.
Si se supera el factor de carga máximo, se redimensiona la tabla antes de insertar.
*/
func (diccionario *diccionarioHashAbierto[K, V]) Guardar(clave K, dato V) {

	if float64(diccionario.cantidad)/float64(diccionario.tam) > factorCargaMaximo {
		diccionario.redimensionar(posiblePrimo(diccionario.tam))
	}

	hash := calcularHash(clave) % uint32(diccionario.tam)
	lista := diccionario.tablas[hash]

	iter := lista.Iterador()
	encontrado := false

	for iter.HaySiguiente() {
		entrada := iter.VerActual()

		if entrada.clave == clave {

			iter.Borrar()
			encontrado = true
			break
		}
		iter.Siguiente()
	}
	lista.InsertarUltimo(entradaDiccionario[K, V]{clave, dato})
	if !encontrado {
		diccionario.cantidad++
	}
}

/*
Precondiciones: La clave debe ser un valor comparable.
La clave debe existir en el diccionario.

Postcondiciones: Devuelve el valor asociado a la clave. Si la clave no existe, se genera un pánico.
*/
func (diccionario *diccionarioHashAbierto[K, V]) Obtener(clave K) V {

	hash := calcularHash(clave) % uint32(diccionario.tam)
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

/*
Precondiciones: La clave debe ser un valor comparable.

Postcondiciones: Devuelve true si la clave existe en el diccionario, false si no.
*/
func (diccionario *diccionarioHashAbierto[K, V]) Pertenece(clave K) bool {

	hash := calcularHash(clave) % uint32(diccionario.tam)
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

/*
Precondiciones: La clave debe ser un valor comparable. La clave debe existir en el diccionario.

Postcondiciones: Elimina la entrada asociada a la clave y devuelve el valor almacenado.
Si la clave no existe, se genera un pánico.
*/
func (diccionario *diccionarioHashAbierto[K, V]) Borrar(clave K) V {

	hash := calcularHash(clave) % uint32(diccionario.tam)
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

/*
Precondiciones: Ninguna.

Postcondiciones: Devuelve la cantidad actual de elementos en el diccionario.
*/
func (diccionario *diccionarioHashAbierto[K, V]) Cantidad() int {
	return diccionario.cantidad
}

/*
Precondiciones: El diccionario debe estar correctamente inicializado. Visitar es una función que toma una clave
y un valor, devolviendo true para continuar la iteración y false para detenerla.

Postcondiciones: Aplica la función visitar a cada par clave-valor en el diccionario. Si visitar devuelve false,
se detiene la iteración antes de finalizar.
*/
func (d *diccionarioHashAbierto[K, V]) Iterar(visitar func(K, V) bool) {
	for _, lista := range d.tablas {
		salirAntes := false
		lista.Iterar(func(nodo entradaDiccionario[K, V]) bool {
			if visitar(nodo.clave, nodo.dato) {
				return true
			} else {
				salirAntes = true
				return false
			}
		})
		if salirAntes {
			break
		}
	}
}

/*
Precondiciones: Ninguna.

Postcondiciones: Devuelve un iterador externo para recorrer el diccionario.
*/
func (d *diccionarioHashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorDiccionarioHashAbierto[K, V]{
		diccionario: d,
		indice:      0,
		iterLista:   d.tablas[0].Iterador(),
	}
	iter.avanzar()
	return iter
}

/*
Precondiciones: El iterador debe estar asociado a un diccionario válido.

Postcondiciones: Avanza el iterador al siguiente elemento válido en el diccionario, o lo deja en un estado
inválido si no hay más elementos.
*/
func (iter *iteradorDiccionarioHashAbierto[K, V]) avanzar() {
	for iter.indice < iter.diccionario.tam {
		if iter.iterLista.HaySiguiente() {
			return
		}
		iter.indice++
		if iter.indice < iter.diccionario.tam {
			iter.iterLista = iter.diccionario.tablas[iter.indice].Iterador()
		}
	}
}

/*
Precondiciones: El iterador debe estar asociado a un diccionario válido.

Postcondiciones: Devuelve true si hay un siguiente elemento disponible en el iterador, false si no.
*/
func (iter *iteradorDiccionarioHashAbierto[K, V]) HaySiguiente() bool {
	return iter.indice < iter.diccionario.tam && iter.iterLista.HaySiguiente()
}

/*
Precondiciones: El iterador debe estar en un estado válido. Debe haber un elemento disponible.

Postcondiciones: Devuelve la clave y el valor del elemento actual. Si no hay un elemento válido, genera un pánico.
*/
func (iter *iteradorDiccionarioHashAbierto[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.iterLista.VerActual()
	return nodo.clave, nodo.dato
}

/*
Precondiciones: El iterador debe estar en un estado válido. Debe haber un siguiente elemento disponible.

Postcondiciones: Avanza el iterador al siguiente elemento en el diccionario. Si no hay más elementos, genera un pánico.
*/
func (iter *iteradorDiccionarioHashAbierto[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.iterLista.Siguiente()
	iter.avanzar()
}
