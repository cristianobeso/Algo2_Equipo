from tdas.pilaYcola import Pila, Cola
from tdas.matrizAGrafo import Grafo

if __name__ == "__main__":
    pila = Pila()
    cola = Cola()
    print("Agregar:")
    for i in range(5):
        pila.apilar(i)
        cola.encolar(i)
        print(f"Tope pila: {pila.ver_tope()}, primero cola: {cola.ver_primero()}")
    print("Quitar: ")
    for i in range(5):
        print(f"Salida pila: {pila.desapilar()}, salida cola:{cola.desencolar()}")
    cola.encolar(30)
    print(f"Salida: {cola.ver_primero()}")

    grafo = Grafo()
    
