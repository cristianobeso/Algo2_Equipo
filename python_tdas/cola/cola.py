class Nodo:
    def __init__(self, valor):
        self.valor = valor
        self.proximo = None

class Cola:

    def __init__(self):
        self.primero = None
        self.ultimo = None

    def esta_vacia(self):
        return self.primero == None
    
    def ver_primero(self):
        if self.esta_vacia():
            raise Exception("La cola está vacía")
        return self.primero.valor

    def encolar(self, valor):
        nodo = Nodo(valor)
        if self.esta_vacia():
            self.primero = nodo
        else:
            self.ultimo.proximo = nodo
        self.ultimo = nodo
        

    def desencolar(self):
        if self.esta_vacia():
            raise Exception("La cola está vacía")
        elemento = self.primero.valor
        self.primero = self.primero.proximo
        return elemento
    
if __name__ == "__main__":
    cola = Cola()
    print("Encolar:")
    for i in range(5):
        cola.encolar(i)
        print(f"Primer elemento: {cola.ver_primero()}")
    print("Desencolar: ")
    for i in range(5):
        print(f"Salida de elementos: {cola.desencolar()}")