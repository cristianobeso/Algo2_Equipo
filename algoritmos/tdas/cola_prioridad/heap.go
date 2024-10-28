package cola_prioridad

type Heap[T any] struct { // cambie el struct porq asi lo decia en el video
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &Heap[T]{datos: []T{}, cant: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := &Heap[T]{datos: arreglo, cant: len(arreglo), cmp: funcion_cmp}
	for i := len(h.datos)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return h
}

func (h *Heap[T]) EstaVacia() bool {
	return h.cant == 0
}

func (h *Heap[T]) Encolar(valor T) {
	h.datos = append(h.datos, valor)
	h.cant++
	h.heapifyUp(h.cant - 1)
}

func (h *Heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola está vacía")
	}
	return h.datos[0]
}

func (h *Heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola está vacía")
	}
	max := h.datos[0]
	h.intercambiar(0, h.cant-1)
	h.datos = h.datos[:h.cant-1]
	h.cant--
	h.heapifyDown(0)
	return max
}

func (h *Heap[T]) Cantidad() int {
	return h.cant
}

func (h *Heap[T]) intercambiar(i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}

func (h *Heap[T]) heapifyUp(pos int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if h.cmp(h.datos[pos], h.datos[padre]) <= 0 {
			break
		}
		h.intercambiar(pos, padre)
		pos = padre
	}
}

func (h *Heap[T]) heapifyDown(pos int) {
	ultimo := h.cant - 1
	for {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
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

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heap := CrearHeapArr(elementos, funcion_cmp).(*Heap[T])

	for i := len(elementos) - 1; i > 0; i-- {
		heap.intercambiar(0, i)
		heap.cant--
		heap.heapifyDown(0)
	}
}
