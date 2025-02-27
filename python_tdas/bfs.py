from tdas.pilaYcola import Cola
from tdas.matrizAGrafo import Grafo

def BFS(grafo, v, padre, visitados):
    cola = Cola()
    cola.encolar(v)
    padre[v] = None
    visitados[v] = True

    while not cola.esta_vacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados[w] = True
                padre[w] = v
                cola.encolar(w)



if __name__ == "__main__":
    matriz = [[0,1,0,0,0,0,0],
              [1,0,0,1,0,0,1],
              [0,0,0,0,1,1,0],
              [0,1,0,0,1,1,1],
              [0,0,1,1,0,1,0],
              [0,0,1,1,1,0,1],
              [0,1,0,1,0,1,0]]
    grafo = Grafo()
    grafo.matrizAdy(matriz)
    padre = {}
    visitados = {}
    BFS(grafo,3,padre,visitados)
    print(padre)
