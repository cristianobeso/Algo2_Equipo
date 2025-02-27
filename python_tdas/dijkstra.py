from math import dist, inf
from tdas.heap import Heap
from tdas.matrizAGrafo import Grafo

def cmp(a,b): return a[1] - b[1]

def camino_minimo(grafo, origen):
    distancia, padre = {},{}
    for v in grafo:
        distancia[v] = inf
    distancia[origen] = 0
    padre[origen] = None
    h = Heap(cmp)
    h.encolar((origen,0))
    while not h.esta_vacia():
        v = h.desencolar()
        for w in grafo.adyacentes(v):
            if (distancia[v] + grafo.peso(v,w)<distancia[w]):
                distancia[w] = distancia[v] + grafo.peso(v,w)
                padre[w] = v
                h.encolar((w,distancia[w]))
    return padre,distancia