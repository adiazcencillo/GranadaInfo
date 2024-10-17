package monumento

import (
    "fmt"
    "strings"
)

type Horario struct {
    DiasCerrado     []bool
    HorariosApertura map[string][]string
}

func (h Horario) MostrarHorario(nombreHorario string) {
    fmt.Printf("Horario de %s:\n", nombreHorario)
    diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo", "Festivo"}

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
