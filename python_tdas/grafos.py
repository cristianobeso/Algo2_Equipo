mat_ady = [[0,1,1,0],
           [1,0,1,0],
           [1,1,0,1],
           [0,0,1,0]]


def mat_a_dic(matriz, dim):
    dic_ver = {}
    for i in range(dim):
        aris = {}
        for j in range(dim):
            if matriz[i][j] != 0:
                aris[j] = None # peso de cada arista
        dic_ver[i] = aris
    return dic_ver

def grafo_vertices(grafo):
    lista_vertices = []
    for vertice in grafo:
        lista_vertices.append(vertice)
    return lista_vertices

def grafo_adyacentes(grafo,v):
    lista_adyacentes = []
    for vertice in grafo[v]:
        lista_adyacentes.append(vertice)
    return lista_adyacentes

grafo = mat_a_dic(mat_ady,4) # crea el grafo
lv = grafo_vertices(grafo)
la = grafo_adyacentes(grafo,2)
print(grafo)
print(lv)
print(la)