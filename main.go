package main

import (
    "fmt"
    "strings"
)

type Ubicacion struct {
    Calle  string
    Numero string
}

type Horario struct {
    DiasCerrado     []bool
    HorariosApertura map[string][]string
}

type Monumento struct {
    Nombre         string
    Ubicacion      Ubicacion
    HorarioVerano  Horario
    HorarioInvierno Horario
    SiempreAbierto bool
}

func NuevoMonumento(nombre string, ubicacion Ubicacion, horarioVerano, horarioInvierno Horario, siempreAbierto bool) Monumento {
    return Monumento{
        Nombre:         nombre,
        Ubicacion:      ubicacion,
        HorarioVerano:  horarioVerano,
        HorarioInvierno: horarioInvierno,
        SiempreAbierto: siempreAbierto,
    }
}

// Función para mostrar la información del monumento
func (m Monumento) MostrarInfo() {
    fmt.Printf("Nombre: %s\nUbicación: %s, %s\n",
        m.Nombre, m.Ubicacion.Calle, m.Ubicacion.Numero)
    
    if m.SiempreAbierto {
        fmt.Println("Este monumento está siempre abierto.")
        return
    }

    diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo", "Festivo"}

    fmt.Printf("Horario de Verano:\n")
    for i, dia := range diasSemana {
        estado := "Abierto"
        if m.HorarioVerano.DiasCerrado[i] {
            estado = "Cerrado"
            fmt.Printf("  %s: %s\n", dia, estado)
        } else {
            horarios := m.HorarioVerano.HorariosApertura[dia]
            fmt.Printf("  %s: %s - %s\n", dia, estado, strings.Join(horarios, ", "))
        }
    }

    fmt.Printf("Horario de Invierno:\n")
    for i, dia := range diasSemana {
        estado := "Abierto"
        if m.HorarioInvierno.DiasCerrado[i] {
            estado = "Cerrado"
            fmt.Printf("  %s: %s\n", dia, estado)
        } else {
            horarios := m.HorarioInvierno.HorariosApertura[dia]
            fmt.Printf("  %s: %s - %s\n", dia, estado, strings.Join(horarios, ", "))
        }
    }
}

func main() {
    // Crear una instancia de Ubicacion
    ubicacionAlhambra := Ubicacion{
        Calle:  "Calle Real de la Alhambra",
        Numero: "s/n",
    }

    // Crear horarios de ejemplo
    horarioVerano := Horario{
        DiasCerrado: []bool{false, false, false, false, false, false, false, false},
        HorariosApertura: map[string][]string{
            "Lunes":    {"09:00-13:00", "15:00-18:00"},
            "Martes":   {"09:00-13:00", "15:00-18:00"},
            "Miércoles": {"09:00-13:00", "15:00-18:00"},
            "Jueves":   {"09:00-13:00", "15:00-18:00"},
            "Viernes":  {"09:00-13:00", "15:00-18:00"},
            "Sábado":   {"09:00-13:00", "15:00-18:00"},
            "Domingo":  {"09:00-13:00", "15:00-18:00"},
            "Festivo":  {"09:00-13:00", "15:00-18:00"},
        },
    }

    horarioInvierno := Horario{
        DiasCerrado: []bool{false, false, false, false, false, false, false, false},
        HorariosApertura: map[string][]string{
            "Lunes":    {"10:00-14:00", "16:00-19:00"},
            "Martes":   {"10:00-14:00", "16:00-19:00"},
            "Miércoles": {"10:00-14:00", "16:00-19:00"},
            "Jueves":   {"10:00-14:00", "16:00-19:00"},
            "Viernes":  {"10:00-14:00", "16:00-19:00"},
            "Sábado":   {"10:00-14:00", "16:00-19:00"},
            "Domingo":  {"10:00-14:00", "16:00-19:00"},
            "Festivo":  {"10:00-14:00", "16:00-19:00"},
        },
    }

    // Crear una instancia de Monumento
    alhambra := NuevoMonumento("La Alhambra", ubicacionAlhambra, horarioVerano, horarioInvierno, false)

    // Mostrar la información del monumento
    alhambra.MostrarInfo()
}