// main.go
package main

import "fmt"

// Definir la estructura Ubicacion
type Ubicacion struct {
    Calle     string
    Numero    string
    Ciudad    string
    Provincia string
    Pais      string
}

// Función para inicializar una instancia de Ubicacion
func NuevaUbicacion(calle, numero string) Ubicacion {
    return Ubicacion{
        Calle:     calle,
        Numero:    numero,
        Ciudad:    "Granada",
        Provincia: "Granada",
        Pais:      "España",
    }
}

// Definir la estructura Horario
type Horario struct {
    HorariosApertura map[string]string // Mapa de días de la semana con sus horarios
    DiasCerrado      [7]bool           // Vector de 7 elementos para los días de la semana
}

// Función para inicializar una instancia de Horario
func NuevoHorario(horariosApertura map[string]string, diasCerrado [7]bool) Horario {
    return Horario{
        HorariosApertura: horariosApertura,
        DiasCerrado:      diasCerrado,
    }
}

// Definir la estructura Monumento
type Monumento struct {
    Nombre         string
    Ubicacion      Ubicacion
    HorarioVerano  Horario
    HorarioInvierno Horario
}

// Función para inicializar una instancia de Monumento
func NuevoMonumento(nombre string, ubicacion Ubicacion, horarioVerano, horarioInvierno Horario) Monumento {
    return Monumento{
        Nombre:         nombre,
        Ubicacion:      ubicacion,
        HorarioVerano:  horarioVerano,
        HorarioInvierno: horarioInvierno,
    }
}

// Función para mostrar la información del monumento
func (m Monumento) MostrarInfo() {
    fmt.Printf("Nombre: %s\nUbicación: %s, %s, %s, %s, %s\n",
        m.Nombre, m.Ubicacion.Calle, m.Ubicacion.Numero, m.Ubicacion.Ciudad,
        m.Ubicacion.Provincia, m.Ubicacion.Pais)
    fmt.Printf("Horario de Verano:\n")
    for dia, horario := range m.HorarioVerano.HorariosApertura {
        fmt.Printf("  %s: %s\n", dia, horario)
    }
    diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
    fmt.Printf("Días Cerrado en Verano:\n")
    for i, cerrado := range m.HorarioVerano.DiasCerrado {
        estado := "Abierto"
        if cerrado {
            estado = "Cerrado"
        }
        fmt.Printf("  %s: %s\n", diasSemana[i], estado)
    }
    fmt.Printf("Horario de Invierno:\n")
    for dia, horario := range m.HorarioInvierno.HorariosApertura {
        fmt.Printf("  %s: %s\n", dia, horario)
    }
    fmt.Printf("Días Cerrado en Invierno:\n")
    for i, cerrado := range m.HorarioInvierno.DiasCerrado {
        estado := "Abierto"
        if cerrado {
            estado = "Cerrado"
        }
        fmt.Printf("  %s: %s\n", diasSemana[i], estado)
    }
}

// Ejemplo de funcionamiento
func main() {
    // Crear una instancia de Ubicacion
    ubicacionAlhambra := NuevaUbicacion("Calle Real de la Alhambra", "s/n")

    // Crear una instancia de Horario de Verano
    horariosAperturaVerano := map[string]string{
        "Lunes":    "08:30 - 14:00",
        "Martes":   "08:30 - 14:00",
        "Miércoles": "08:30 - 14:00",
        "Jueves":   "08:30 - 14:00",
        "Viernes":  "08:30 - 14:00",
        "Sábado":   "08:30 - 14:00",
        "Domingo":  "08:30 - 14:00",
    }
    diasCerradoVerano := [7]bool{false, false, false, false, false, false, true} // Cerrado los domingos en verano
    horarioVerano := NuevoHorario(horariosAperturaVerano, diasCerradoVerano)

    // Crear una instancia de Horario de Invierno
    horariosAperturaInvierno := map[string]string{
        "Lunes":    "09:00 - 13:00",
        "Martes":   "09:00 - 13:00",
        "Miércoles": "09:00 - 13:00",
        "Jueves":   "09:00 - 13:00",
        "Viernes":  "09:00 - 13:00",
        "Sábado":   "09:00 - 13:00",
        "Domingo":  "09:00 - 13:00",
    }
    diasCerradoInvierno := [7]bool{false, false, false, false, false, true, true} // Cerrado los sábados y domingos en invierno
    horarioInvierno := NuevoHorario(horariosAperturaInvierno, diasCerradoInvierno)

    // Crear una instancia de Monumento
    alhambra := NuevoMonumento("La Alhambra", ubicacionAlhambra, horarioVerano, horarioInvierno)

    // Mostrar la información del monumento
    alhambra.MostrarInfo()
}