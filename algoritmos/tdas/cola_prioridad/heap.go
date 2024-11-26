package cola_prioridad

type heap[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

const FACTOR_REDIMENSION = 2

// CrearHeap crea un heap vacío con una función de comparación proporcionada.
// Precondición: funcion_cmp no es nil.
// Postcondición: Devuelve un heap vacío.
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{datos: make([]T, 1), cant: 0, cmp: funcion_cmp}
}

// CrearHeapArr construye un heap a partir de un arreglo y una función de comparación.
// Precondición: arreglo no es nil, funcion_cmp no es nil.
// Postcondición: Devuelve un heap que mantiene la propiedad de heap con los elementos de 'arreglo'.
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	copia := make([]T, len(arreglo))
	copy(copia, arreglo)

	h := &heap[T]{datos: copia, cant: len(copia), cmp: funcion_cmp}
	for i := len(h.datos)/FACTOR_REDIMENSION - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return h
}

// EstaVacia verifica si el heap está vacío.
// Precondición: Ninguna.
// Postcondición: Devuelve true si el heap está vacío, de lo contrario false.
func (h *heap[T]) EstaVacia() bool {
	return h.cant == 0
}

// Encolar agrega un elemento al heap.
// Precondición: Ninguna.
// Postcondición: Agrega 'valor' al heap, manteniendo la propiedad de heap.
func (h *heap[T]) Encolar(valor T) {
	if cap(h.datos) == h.cant {
		h.redimensionar(cap(h.datos) * FACTOR_REDIMENSION)
	}
	h.datos[h.cant] = valor
	h.cant++
	h.heapifyUp(h.cant - 1)
}

// VerMax devuelve el elemento máximo del heap sin eliminarlo.
// Precondición: El heap no está vacío.
// Postcondición: Devuelve el elemento máximo del heap.
func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

// Desencolar elimina y devuelve el elemento máximo del heap.
// Precondición: El heap no está vacío.
// Postcondición: Devuelve el elemento máximo y reduce el tamaño del heap, manteniendo la propiedad de heap.
func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := h.datos[0]
	h.intercambiar(0, h.cant-1)
	h.cant--
	h.datos = h.datos[:h.cant] //creo que esta linea no es necesaria
	h.heapifyDown(0)

	if h.cant*4 <= cap(h.datos) {
		h.redimensionar(cap(h.datos) / FACTOR_REDIMENSION)
	}

	//h.redimensionarSiEsNecesario()

	return max
}

// Cantidad devuelve el número de elementos en el heap.
// Precondición: Ninguna.
// Postcondición: Devuelve la cantidad de elementos en el heap
func (h *heap[T]) Cantidad() int {
	return h.cant
}

// intercambiar intercambia dos elementos en el slice 'datos' del heap.
// Precondición: i y j son índices válidos dentro del slice 'datos'.
// Postcondición: Intercambia los elementos en las posiciones i y j.
func (h *heap[T]) intercambiar(i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}

// heapifyUp restablece la propiedad de heap subiendo un elemento en el heap.
// Precondición: 'pos' es un índice válido en el heap.
// Postcondición: El elemento en 'pos' se ajusta en el heap, manteniendo la propiedad de heap.
func (h *heap[T]) heapifyUp(pos int) {
	for pos > 0 {
		padre := (pos - 1) / FACTOR_REDIMENSION
		if h.cmp(h.datos[pos], h.datos[padre]) <= 0 {
			break
		}
		h.intercambiar(pos, padre)
		pos = padre
	}
}

// heapifyDown restablece la propiedad de heap bajando un elemento en el heap.
// Precondición: 'pos' es un índice válido en el heap.
// Postcondición: El elemento en 'pos' se ajusta en el heap, manteniendo la propiedad de heap.
func (h *heap[T]) heapifyDown(pos int) {
	ultimo := h.cant - 1
	for {
		hijoIzq := FACTOR_REDIMENSION*pos + 1
		hijoDer := FACTOR_REDIMENSION*pos + FACTOR_REDIMENSION
		mayor := pos

		if hijoIzq <= ultimo && h.cmp(h.datos[hijoIzq], h.datos[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer <= ultimo && h.cmp(h.datos[hijoDer], h.datos[mayor]) > 0 {
			mayor = hijoDer
		}
		if mayor == pos {
			break
		}
		h.intercambiar(pos, mayor)
		pos = mayor
	}
}

// HeapSort ordena un arreglo usando el heap sort in-place.
// Precondición: elementos y funcion_cmp no son nil.
// Postcondición: Ordena el slice 'elementos' de acuerdo con la función de comparación 'funcion_cmp'.
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	h := &heap[T]{datos: elementos, cant: len(elementos), cmp: funcion_cmp}

	//  heap max-heap
	for i := len(elementos)/FACTOR_REDIMENSION - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}

	// Orden in-place
	for i := len(elementos) - 1; i > 0; i-- {
		h.intercambiar(0, i)
		h.cant--
		h.datos = h.datos[:h.cant]
		h.heapifyDown(0)
	}
}

func (h *heap[T]) redimensionar(tam int) {
	nuevaArreglo := make([]T, tam)
	copy(nuevaArreglo, h.datos)
	h.datos = nuevaArreglo
}
