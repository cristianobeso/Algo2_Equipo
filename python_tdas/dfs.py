from tdas.matrizAGrafo import Grafo

def DFS(grafo, v, padre, visitados):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados[w] = True
            padre[w] = v
            DFS(grafo,w,padre,visitados)

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
    padre[3] = None
    visitados[3] = True
    DFS(grafo,3,padre,visitados)
    print(padre)
