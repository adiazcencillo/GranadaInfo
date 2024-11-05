package horario

type Horario struct {
    DiasCerrado     []bool              // Indica si el monumento está cerrado en cada día de la semana y festivos.
    HorariosApertura map[string][]string // Horarios de apertura por día de la semana. Permite múltiples horarios por día.
    SiempreAbierto   bool
}

// NuevoHorario crea una nueva instancia de Horario.
// Si siempreAbierto es verdadero, se ignoran diasCerrado y horariosApertura,
// y se establecen valores por defecto para una apertura continua.
func NuevoHorario(siempreAbierto bool, diasCerrado []bool, horariosApertura map[string][]string) *Horario {
    if siempreAbierto {
        horariosApertura = map[string][]string{
            "Lunes":    {"00:00-23:59"},
            "Martes":   {"00:00-23:59"},
            "Miércoles": {"00:00-23:59"},
            "Jueves":   {"00:00-23:59"},
            "Viernes":  {"00:00-23:59"},
            "Sábado":   {"00:00-23:59"},
            "Domingo":  {"00:00-23:59"},
            "Festivo":  {"00:00-23:59"},
        }
        diasCerrado = []bool{false, false, false, false, false, false, false, false}
    }
    return &Horario{
        DiasCerrado:     diasCerrado,
        HorariosApertura: horariosApertura,
        SiempreAbierto:  siempreAbierto,
    }
}