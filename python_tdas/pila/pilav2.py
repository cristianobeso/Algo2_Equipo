class Nodo:
    def __init__(self, valor):
        self.valor = valor
        self.proximo = None

class Pila:

    def __init__(self):
        self.tope = None

    def esta_vacia(self):
        return self.tope == None
    
    def ver_tope(self):
        if self.esta_vacia():
            raise Exception("La pila está vacía")
        return self.tope.valor

    def apilar(self, valor):
        nodo = Nodo(valor)
        nodo.proximo = self.tope
        self.tope = nodo
        

    def desapilar(self):
        if self.esta_vacia():
            raise Exception("La pila está vacía")
        remove = self.tope
        self.tope = remove.proximo
        
        return remove.valor
    
if __name__ == "__main__":
    pila = Pila()
    print("Apilar: ")
    for i in range(5):
        pila.apilar(i)
        print(f"Ultimo elemento: {pila.ver_tope()}")
    print("Desapilar: ")
    for i in range(5):
        print(f"Salida de elementos: {pila.desapilar()}")
    print("Apilar: ")
    pila.apilar(20)
    print(f"Ultimo elemento: {pila.ver_tope()}")