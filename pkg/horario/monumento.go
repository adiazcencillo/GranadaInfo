package horario

import "fmt"

// Monumento representa un monumento con su información básica.
// Monumento tiene un conjunto de atributos que describen el monumento, es decir, tiene una identidad inmutable, siendo monumento una entidad.
// Por ejemplo, tiene Nombre como propiedad clave que garantiza su unicidad, considerandose entidad.
type Monumento struct {
    Nombre          string  // Nombre del monumento
	HorarioInvierno Horario  // Horario de invierno del monumento
    HorarioVerano   Horario  // Horario de verano del monumento. Será el mismo que el de invierno si tiene el mismo todo el año.
	SiempreAbierto  bool  // Indica si el monumento está siempre abierto
}

// NuevoMonumento crea y devuelve una nueva instancia de Monumento.
func NuevoMonumento(nombre string, horarioVerano, horarioInvierno Horario, siempreAbierto bool) Monumento {
    return Monumento{
        Nombre:          nombre,
        HorarioVerano:   horarioVerano,
        HorarioInvierno: horarioInvierno,
		SiempreAbierto:  siempreAbierto,
    }
}

// MostrarInfo imprime la información del monumento.
func (m Monumento) MostrarInfo() {
    fmt.Printf("Nombre: %s\n",
        m.Nombre)

	if m.SiempreAbierto {
		fmt.Println("Este monumento está siempre abierto.")
		return
	} 
	// Muestra los horarios de verano e invierno si el monumento no está siempre abierto.
	m.HorarioVerano.MostrarHorario("Verano")
	m.HorarioInvierno.MostrarHorario("Invierno")
}