package monumento

import "fmt"

type Monumento struct {
    Nombre          string
    Ubicacion       Ubicacion
    HorarioVerano   Horario
    HorarioInvierno Horario
    SiempreAbierto  bool
}

func NuevoMonumento(nombre string, ubicacion Ubicacion, horarioVerano, horarioInvierno Horario, siempreAbierto bool) Monumento {
    return Monumento{
        Nombre:          nombre,
        Ubicacion:       ubicacion,
        HorarioVerano:   horarioVerano,
        HorarioInvierno: horarioInvierno,
        SiempreAbierto:  siempreAbierto,
    }
}

func (m Monumento) MostrarInfo() {
    fmt.Printf("Nombre: %s\nUbicación: %s, %s\n",
        m.Nombre, m.Ubicacion.Calle, m.Ubicacion.Numero)

    if m.SiempreAbierto {
        fmt.Println("Este monumento está siempre abierto.")
        return
    }

    m.HorarioVerano.MostrarHorario("Verano")
    m.HorarioInvierno.MostrarHorario("Invierno")
}