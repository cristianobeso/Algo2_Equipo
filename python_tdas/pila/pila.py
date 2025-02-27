class Pila:
    def __init__(self):
        self.datos = []
        self.cantidad = 0

    def esta_vacia(self):
        return self.cantidad == 0
    
    def ver_tope(self):
        if self.cantidad == 0:
            raise IndexError("La pila está vacía")
        return self.datos[self.cantidad-1]

    def apilar(self, elemento):
        self.datos.append(elemento)
        self.cantidad+=1

    def desapilar(self):
        if self.cantidad == 0:
            raise IndexError("La pila está vacía")
        self.cantidad-=1
        return self.datos.pop()

if __name__ == "__main__":
    pila = Pila()
    for i in range(10):
        pila.apilar(i)
        print(f"Ultimo elemento: {pila.ver_tope()}")
    print("Desapilar: ")
    for i in range(5):
        print(f"Salida de elementos: {pila.desapilar()}")