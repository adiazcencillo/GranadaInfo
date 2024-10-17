package main

import (
    "GranadaInfo/monumento"
)

func main() {
    // Crear una instancia de Ubicacion
    ubicacionAlhambra := monumento.Ubicacion{
        Calle:  "Calle Real de la Alhambra",
        Numero: "s/n",
    }

    // Crear horarios de ejemplo para verano e invierno
    horarioVerano := monumento.Horario{
        DiasCerrado: []bool{false, false, false, false, false, false, false, false},
        HorariosApertura: map[string][]string{
            "Lunes":    {"09:00-13:00", "15:00-18:00"},
            "Martes":   {"09:00-13:00", "15:00-18:00"},
            "Miércoles": {"09:00-13:00", "15:00-18:00"},
            "Jueves":   {"09:00-13:00", "15:00-18:00"},
            "Viernes":  {"09:00-13:00", "15:00-18:00"},
            "Sábado":   {"09:00-13:00", "15:00-18:00"},
            "Domingo":  {"09:00-13:00", "15:00-18:00"},
            "Festivo":  {"09:00-13:00"},
        },
    }

    horarioInvierno := monumento.Horario{
        DiasCerrado: []bool{false, false, false, false, false, false, false, true},
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
    alhambra := monumento.NuevoMonumento("La Alhambra", ubicacionAlhambra, horarioVerano, horarioInvierno, false)

    // Mostrar la información del monumento
    alhambra.MostrarInfo()
}
