class Heap:
    def __init__(self, func):
        self.datos = []
        self.cant = 0
        self.func = func

    def esta_vacia(self):
        return self.cantidad == 0
    
    def ver_max(self):
        if self.esta_vacia():
            raise IndexError("La cola está vacía")
        return self.datos[0]

    def encolar(self, elemento):
        self.datos.append(elemento)
        self.cant+=1
        self.heapify_up(self.cant-1)

    def desencolar(self):
        if self.esta_vacia():
            raise IndexError("La cola está vacía")
        
        self.swap(0, self.cant-1)
        self.cant-=1
        elemento = self.datos.pop()
        self.heapify_down(0)
        return elemento
    
    def cantidad(self):
        return self.cant
    
    def heapify_up(self, pos):
        while pos > 0:
            padre = int((pos - 1)/2)
            if self.func(self.datos[pos],self.datos[padre]) <= 0:
                break
            self.swap(pos,padre)
            pos = padre

    def heapify_down(self, pos):
        ultimo = self.cant - 1
        while True:
            hijoIzq = pos*2 + 1
            hijoDer = pos*2 + 2
            mayor = pos

            if hijoIzq <= ultimo and self.func(self.datos[hijoIzq], self.datos[mayor]) > 0:
                mayor = hijoIzq
            if hijoDer <= ultimo and self.func(self.datos[hijoDer], self.datos[mayor]) > 0:
                mayor = hijoDer
            if mayor == pos:
                break
            self.swap(pos,mayor)
            pos = mayor

    def swap(self,a,b):
        self.datos[a],self.datos[b]=self.datos[b],self.datos[a]
    
    def __repr__(self):
        return str(self.datos)

# def cmp(a,b): return a - b

# if __name__ == "__main__":
#     h = Heap(cmp)
#     h.encolar(10)
#     h.encolar(5)
#     h.encolar(13)
#     print(h)
#     print(h.desencolar())
#     print(h)
#     print(h.desencolar())
#     print(h)