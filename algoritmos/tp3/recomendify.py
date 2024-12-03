#!/usr/bin/python3
import sys
from grafo import Grafo
from biblioteca import bfs_corto, pagerank, dfs_ciclo

def cargar_grafo(tsv_path):
    grafo_bipartito = Grafo()
    grafo_canciones = Grafo()
    with open(tsv_path, 'r') as archivo:
        for linea in archivo:
            datos = linea.strip().split('\t')
            user, track, artist, playlist = datos[1], datos[2], datos[3], datos[4]
            grafo_bipartito.agregar_arista(user, f"{track} - {artist}")
            grafo_canciones.agregar_arista(f"{track} - {artist}", playlist)
    return grafo_bipartito, grafo_canciones

def main():
    if len(sys.argv) < 2:
        print("Uso: ./recomendify <archivo_tsv>")
        sys.exit(1)
    archivo = sys.argv[1]
    grafo_bipartito, grafo_canciones = cargar_grafo(archivo)
    while True:
        try:
            entrada = input("> ").strip()
            if not entrada:
                continue
            comando, *args = entrada.split(" ", 1)
            if comando == "camino":
                origen, destino = args[0].split(" >>>> ")
                resultado = bfs_corto(grafo_bipartito, origen, destino)
                print(" --> ".join(resultado) if resultado else "No se encontro recorrido")
            elif comando == "mas_importantes":
                n = int(args[0])
                pr = pagerank(grafo_canciones)
                print("; ".join(c[0] for c in pr[:n]))
            elif comando == "ciclo":
                n, cancion = args[0].split(" ", 1)
                resultado = dfs_ciclo(grafo_canciones, cancion, int(n))
                print(" --> ".join(resultado) if resultado else "No se encontro recorrido")
            else:
                print("Comando no reconocido")
        except EOFError:
            break

if __name__ == "__main__":
    main()
