package cola_prioridad_test

import (
	cola_prioridad "tdas/cola_prioridad"
	"testing"
)

// Función de comparación
func cmpInt(a, b int) int {
	return a - b
}

func TestCrearHeapVacio(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	if !heap.EstaVacia() {
		t.Errorf("Esperaba que el heap estuviera vacío al crearse, pero no lo está.")
	}
}

func TestEncolarYVerMax(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)

	max := heap.VerMax()
	if max != 20 {
		t.Errorf("Esperaba que el máximo fuera 20, pero obtuve %d", max)
	}
}

func TestDesencolar(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)

	desencolado := heap.Desencolar()
	if desencolado != 20 {
		t.Errorf("Esperaba desencolar 20, pero obtuve %d", desencolado)
	}

	max := heap.VerMax()
	if max != 10 {
		t.Errorf("Esperaba que el máximo fuera 10, pero obtuve %d", max)
	}
}

func TestCantidad(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	heap.Encolar(10)
	heap.Encolar(20)

	if heap.Cantidad() != 2 {
		t.Errorf("Esperaba 2 elementos en el heap, pero hay %d", heap.Cantidad())
	}

	heap.Desencolar()

	if heap.Cantidad() != 1 {
		t.Errorf("Esperaba 1 elemento en el heap después de desencolar, pero hay %d", heap.Cantidad())
	}
}

func TestHeapSort(t *testing.T) {
	arreglo := []int{5, 12, 11, 13, 4, 6, 7}
	esperado := []int{4, 5, 6, 7, 11, 12, 13}

	cola_prioridad.HeapSort(arreglo, cmpInt)

	for i, val := range arreglo {
		if val != esperado[i] {
			t.Errorf("HeapSort: esperada posición %d con valor %d, pero obtuve %d", i, esperado[i], val)
		}
	}
}

func TestCrearHeapArr(t *testing.T) {
	arreglo := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	heap := cola_prioridad.CrearHeapArr(arreglo, cmpInt)

	max := heap.VerMax()
	if max != 9 {
		t.Errorf("Esperaba que el máximo fuera 9, pero obtuve %d", max)
	}
}

func TestDesencolarHeapVacio(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Esperaba un pánico al desencolar de un heap vacío, pero no ocurrió.")
		}
	}()
	heap := cola_prioridad.CrearHeap(cmpInt)
	heap.Desencolar()
}

func TestVerMaxHeapVacio(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Esperaba un pánico al ver el máximo de un heap vacío, pero no ocurrió.")
		}
	}()
	heap := cola_prioridad.CrearHeap(cmpInt)
	heap.VerMax()
}

func TestDesencolarReduceTamanio(t *testing.T) {
	heap := cola_prioridad.CrearHeap(cmpInt)
	elementos := []int{20, 15, 10, 8, 7, 9}
	for _, elem := range elementos {
		heap.Encolar(elem)
	}

	if heap.Cantidad() != len(elementos) {
		t.Fatalf("Error en tamaño inicial: esperado %d, pero obtenido %d", len(elementos), heap.Cantidad())
	}

	for i := len(elementos); i > 0; i-- {
		heap.Desencolar()
		if heap.Cantidad() != i-1 {
			t.Errorf("Error en cantidad después de desencolar: esperado %d, obtenido %d", i-1, heap.Cantidad())
		}
	}
}
