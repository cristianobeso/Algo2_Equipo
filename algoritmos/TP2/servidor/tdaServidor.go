package servidor

type DatosServidor interface {
	AgregarArchivo(nombre_archivo string) error
	VerVisitantes(desde, hasta string)
	VerMasVisitados(n int)
}
