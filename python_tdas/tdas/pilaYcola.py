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
            raise IndexError("La pila está vacía")
        return self.tope.valor

    def apilar(self, valor):
        nodo = Nodo(valor)
        nodo.proximo = self.tope
        self.tope = nodo
        
    def desapilar(self):
        if self.esta_vacia():
            raise IndexError("La pila está vacía")
        remove = self.tope
        self.tope = remove.proximo
        return remove.valor
    
class Cola:

    def __init__(self):
        self.primero = None
        self.ultimo = None

    def esta_vacia(self):
        return self.primero == None
    
    def ver_primero(self):
        if self.esta_vacia():
            raise IndexError("La cola está vacía")
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
            raise IndexError("La cola está vacía")
        elemento = self.primero.valor
        self.primero = self.primero.proximo
        return elemento