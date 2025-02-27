class Grafo:
    def __init__(self):
        self.grafo = {}
    
    def matrizInc(self, matriz): # matriz de incidencia
        dimx = len(matriz)
        dimy = len(matriz[0])
        for i in range(dimy):
            self.grafo[i] = {}
        for i in range(dimx):
            # se supone que solo se conectan dos vectores por cada arista
            lista = []
            for j in range(dimy):
                if matriz[i][j] != 0:
                    lista.append(j) # se añaden esos vectores en la lista
            self.grafo[lista[0]].update({lista[1] : matriz[i][lista[0]]})
            self.grafo[lista[1]].update({lista[0] : matriz[i][lista[1]]})

    
    def matrizAdy(self, matriz): #matriz de adyacencia
        dim = len(matriz)
        for i in range(dim):
            aris = {}
            for j in range(dim):
                if matriz[i][j] != 0:
                    aris[j] = None
            self.grafo[i] = aris

    def __iter__(self):
        return iter(self.grafo)

    def vertices(self):
        return list(self.grafo.keys())
    
    def adyacentes(self,v):
        return list(self.grafo[v].keys())
        #return iter(self.grafo[v]) podría ser, no estoy seguro, aun no probado
    
    def agregar_vertice(self, vertice):
        if not vertice in self.grafo:
            self.grafo[vertice] = {}

    def agregar_arista(self, vertice1, vertice2, valor):
        if vertice1 in self.grafo and vertice2 in self.grafo:
            self.grafo[vertice1].update({vertice2:valor})
            self.grafo[vertice2].update({vertice1:valor})

    def peso(self,v,w):
        return self.grafo[v][w]

# mat_ady = [[0,1,1,0],
#            [1,0,1,0],
#            [1,1,0,1],
#            [0,0,1,0]]

# mat_inc = [[ 1, 1, 0, 0, 0],
#            [ 1, 0, 1, 0, 0],
#            [ 0, 1, 1, 0, 0],
#            [ 0, 0, 1, 1, 0]]

# if __name__ == "__main__":
#     grafo = Grafo()
#     grafo.matrizInc(mat_inc)
#     for v in grafo:
#         print(f"v: {v}")
#         for w in grafo.adyacentes(v):
#             print(f"ady: {w}")