package horario

import (
    "fmt"
    "strings"
)

// Horario representa los horarios de apertura y cierre de un monumento.
// Es un objeto valor ya que su identidad depende de sus atributos, no tiene identidad propia.
type Horario struct {
    HorariosApertura map[string][]string // Horarios de apertura por día de la semana. Permite varios horarios por día para el caso de cierre al mediodía.
}

// MostrarHorarios imprime los horarios de apertura por día de la semana.
func MostrarHorarios(h Horario) {
    for dia, horarios := range h.HorariosApertura {
        fmt.Printf("%s: %s\n", dia, strings.Join(horarios, ", "))
    }
}