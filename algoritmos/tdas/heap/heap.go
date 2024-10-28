package cola_prioridad

type Heap[T any] struct {
	elementos []T
	cmp       func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &Heap[T]{elementos: []T{}, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := &Heap[T]{elementos: arreglo, cmp: funcion_cmp}
	for i := len(h.elementos)/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
	return h
}

func (h *Heap[T]) EstaVacia() bool {
	return len(h.elementos) == 0
}

func (h *Heap[T]) Encolar(valor T) {
	h.elementos = append(h.elementos, valor)
	h.heapifyUp(len(h.elementos) - 1)
}

func (h *Heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola está vacía")
	}
	return h.elementos[0]
}

func (h *Heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola está vacía")
	}
	max := h.elementos[0]
	h.intercambiar(0, len(h.elementos)-1)
	h.elementos = h.elementos[:len(h.elementos)-1]
	h.heapifyDown(0)
	return max
}

func (h *Heap[T]) Cantidad() int {
	return len(h.elementos)
}

func (h *Heap[T]) intercambiar(i, j int) {
	h.elementos[i], h.elementos[j] = h.elementos[j], h.elementos[i]
}

func (h *Heap[T]) heapifyUp(pos int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if h.cmp(h.elementos[pos], h.elementos[padre]) <= 0 {
			break
		}
		h.intercambiar(pos, padre)
		pos = padre
	}
}

func (h *Heap[T]) heapifyDown(pos int) {
	ultimo := len(h.elementos) - 1
	for {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
		mayor := pos

		if hijoIzq <= ultimo && h.cmp(h.elementos[hijoIzq], h.elementos[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer <= ultimo && h.cmp(h.elementos[hijoDer], h.elementos[mayor]) > 0 {
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
	heap := CrearHeapArr(elementos, funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		elementos[i] = heap.Desencolar()
	}
}
