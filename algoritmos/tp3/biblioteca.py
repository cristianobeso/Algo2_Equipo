from collections import deque
import heapq

def bfs_corto(grafo, origen, destino):
    if origen not in grafo.vertices or destino not in grafo.vertices:
        return None
    visitados = set()
    cola = deque([(origen, [])])
    while cola:
        actual, camino = cola.popleft()
        if actual in visitados:
            continue
        visitados.add(actual)
        nuevo_camino = camino + [actual]
        if actual == destino:
            return nuevo_camino
        for vecino, _ in grafo.obtener_adyacentes(actual):
            if vecino not in visitados:
                cola.append((vecino, nuevo_camino))
    return None

def pagerank(grafo, iteraciones=100, d=0.85):
    pr = {v: 1 / len(grafo.vertices) for v in grafo.vertices}
    for _ in range(iteraciones):
        nuevo_pr = {}
        for vertice in grafo.vertices:
            suma = sum(pr[ady] / len(grafo.obtener_adyacentes(ady)) for ady, _ in grafo.obtener_adyacentes(vertice))
            nuevo_pr[vertice] = (1 - d) / len(grafo.vertices) + d * suma
        pr = nuevo_pr
    return sorted(pr.items(), key=lambda x: x[1], reverse=True)

def dfs_ciclo(grafo, inicio, longitud):
    def dfs(vertice, visitados, camino):
        if len(camino) == longitud and camino[0] == vertice:
            return camino
        if len(camino) > longitud:
            return None
        for vecino, _ in grafo.obtener_adyacentes(vertice):
            if vecino not in visitados or (len(camino) == longitud - 1 and vecino == camino[0]):
                resultado = dfs(vecino, visitados | {vecino}, camino + [vecino])
                if resultado:
                    return resultado
        return None
    return dfs(inicio, {inicio}, [inicio])
