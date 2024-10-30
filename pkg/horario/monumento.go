package horario

// Monumento representa un monumento con su información básica.
// Un Monumento tiene atributos que describen su identidad inmutable.
// Por ejemplo, el atributo Nombre garantiza la unicidad del monumento, considerándose así una entidad.
type Monumento struct {
	Nombre         string  // Nombre del monumento.
	Horario        Horario // Horario de apertura y cierre del monumento.
}

// NuevoMonumento crea y devuelve una nueva instancia de Monumento.
// Recibe el nombre del monumento y su horario
func NuevoMonumento(nombre string, horario Horario) Monumento {
	return Monumento{
		Nombre:         nombre,
		Horario:       horario,
	}
}
