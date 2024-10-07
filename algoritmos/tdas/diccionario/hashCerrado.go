// // Se que no vamos a hacer el cerrado pero queria tener una base sobre la cual comenzar
// // hay que centrarnos en la funcion de hash y en la funcion para desplazarmos en caso de choque
package diccionario

// import "fmt"

// const (
// 	OCUPADO = 1
// 	VACIO   = 0
// 	BORRADO = -1
// )

// type celdaHash[K comparable, V any] struct {
// 	clave  K
// 	dato   V
// 	estado int
// }

// type hashCerrado[K comparable, V any] struct {
// 	tabla    []celdaHash[K, V]
// 	cantidad int
// 	tam      int
// 	borrados int
// }

// func CrearHashCerrado[K comparable, V any]() Diccionario[K, V] {
// 	return &hashCerrado[K, V]{tabla: make([]celdaHash[K, V], 7), cantidad: 0, tam: 7, borrados: 0}
// }

// func indiceHash[K comparable](clave string, tam int) int {
// 	return len(clave) % tam // xd
// }

// func convertirABytes[K comparable](clave K) []byte {
// 	return []byte(fmt.Sprintf("%v", clave)) // no se como se usa
// }

// func redimensionar[K comparable, V any](tabla []celdaHash[K, V]) []celdaHash[K, V] {
// 	nuevoTabla := make([]celdaHash[K, V], len(tabla)*2) // por ahora se duplica
// 	for _, valor := range tabla {
// 		if valor.estado == VACIO {
// 			indice := indiceHash[string](valor.clave, len(nuevoTabla))
// 			if nuevoTabla[indice].estado == VACIO {
// 				nuevoTabla[indice].clave = valor.clave
// 				nuevoTabla[indice].dato = valor.dato
// 			}
// 		}
// 	}
// }

// func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
// 	indice := indiceHash[string](clave, hash.tam)
// 	if (hash.cantidad+hash.borrados)/hash.tam > 0.5 { //aprox
// 		hash.tabla = redimensionar(hash.tabla)
// 		hash.tam = len(hash.tabla)
// 		indice = indiceHash[string](clave, hash.tam) // calcula un nuevo indice en caso de redimension
// 		hash.borrados = 0                            // despues de la redimencion ya no se guardan los borrados
// 	}
// 	if hash.tabla[indice].estado == VACIO {
// 		hash.tabla[indice].clave = clave
// 		hash.tabla[indice].dato = dato
// 	}
// 	hash.cantidad++
// }

// func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
// 	indice := indiceHash[string](clave, hash.tam)
// 	if hash.tabla[indice].clave == clave {
// 		return true
// 	}
// 	return false
// }

// func (hash *hashCerrado[K, V]) Obtener(clave K) V {
// 	indice := indiceHash[string](clave, hash.tam)
// 	if hash.tabla[indice].clave == clave {
// 		return hash.tabla[indice].dato
// 	}
// 	panic("no se encontro")
// }

// func (hash *hashCerrado[K, V]) Borrar(clave K) V {
// 	indice := indiceHash[string](clave, hash.tam)
// 	if hash.tabla[indice].clave == clave {
// 		valor := hash.tabla[indice].dato
// 		hash.tabla[indice].estado = BORRADO
// 		hash.borrados++
// 		hash.cantidad--
// 		return valor
// 	}
// 	panic("no se encontro")
// }

// func (hash *hashCerrado[K, V]) Cantidad() int {
// 	return hash.cantidad
// }
