#Ejemplo de como usar clases en un diccionario
class MiDiccionario:
    def __init__(self):
        self.datos = {}

    def __getitem__(self, clave):
        """Permite acceder como diccionario[clave]."""
        if clave in self.datos:
            return self.datos[clave]
        raise KeyError(f"La clave '{clave}' no existe.")

    def __setitem__(self, clave, valor):
        """Permite asignar valores como diccionario[clave] = valor."""
        self.datos[clave] = valor
        print(f"✔ Se ha agregado/modificado: {clave} -> {valor}")

    def __delitem__(self, clave):
        """Permite eliminar elementos con del diccionario[clave]."""
        if clave in self.datos:
            print(f"✖ Se ha eliminado la clave: {clave}")
            del self.datos[clave]
        else:
            raise KeyError(f"La clave '{clave}' no existe.")

    def __contains__(self, clave):
        """Permite usar 'clave in diccionario'."""
        return clave in self.datos

    def __len__(self):
        """Permite usar len(diccionario)."""
        return len(self.datos)

    def __iter__(self):
        """Permite iterar sobre las claves del diccionario."""
        return iter(self.datos)

    def __repr__(self):
        """Representación visual del diccionario."""
        return str(self.datos)
