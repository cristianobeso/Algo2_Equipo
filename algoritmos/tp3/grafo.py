class Grafo:
    def __init__(self, es_dirigido=False):
        self.vertices = {}
        self.es_dirigido = es_dirigido

    def agregar_vertice(self, vertice):
        if vertice not in self.vertices:
            self.vertices[vertice] = set()

    def agregar_arista(self, origen, destino, peso=1):
        self.agregar_vertice(origen)
        self.agregar_vertice(destino)
        self.vertices[origen].add((destino, peso))
        if not self.es_dirigido:
            self.vertices[destino].add((origen, peso))

    def obtener_adyacentes(self, vertice):
        return self.vertices.get(vertice, set())

    def eliminar_vertice(self, vertice):
        if vertice in self.vertices:
            del self.vertices[vertice]
            for adyacentes in self.vertices.values():
                adyacentes = {(v, p) for v, p in adyacentes if v != vertice}

    def eliminar_arista(self, origen, destino):
        self.vertices[origen] = {(v, p) for v, p in self.vertices[origen] if v != destino}
        if not self.es_dirigido:
            self.vertices[destino] = {(v, p) for v, p in self.vertices[destino] if v != origen}
