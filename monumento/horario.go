package monumento

import (
    "fmt"
    "strings"
)

// Horario representa los horarios de apertura y cierre de un monumento.
type Horario struct {
    DiasCerrado     []bool              // Indica qué días está cerrado el monumento. Incluye los festivos así como un día extra para festivos.
    HorariosApertura map[string][]string // Horarios de apertura por día de la semana. Permite varios horarios por día para el caso de cierre al mediodía.
}

// MostrarHorario imprime el horario del monumento para un nombre de horario específico para diferenciar entre verano e invierno.
func (h Horario) MostrarHorario(nombreHorario string) {
    fmt.Printf("Horario de %s:\n", nombreHorario)
    diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo", "Festivo"}

    // Itera sobre los días de la semana y muestra si está abierto o cerrado, y los horarios de apertura (sólo si esta abierto).
    for i, dia := range diasSemana {
        estado := "Abierto"
        if h.DiasCerrado[i] {
            estado = "Cerrado"
            fmt.Printf("  %s: %s\n", dia, estado)
        } else {
            horarios := h.HorariosApertura[dia]
            fmt.Printf("  %s: %s - %s\n", dia, estado, strings.Join(horarios, ", "))
        }
    }
}
