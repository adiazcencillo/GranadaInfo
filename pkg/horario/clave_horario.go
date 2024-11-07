package horario

type ClaveHorario string

const (
    Lunes     ClaveHorario = "Lunes"
    Martes    ClaveHorario = "Martes"
    Miércoles ClaveHorario = "Miércoles"
    Jueves    ClaveHorario = "Jueves"
    Viernes   ClaveHorario = "Viernes"
    Sábado    ClaveHorario = "Sábado"
    Domingo   ClaveHorario = "Domingo"
    Festivo   ClaveHorario = "Festivo"
)

var ClavesValidas = map[ClaveHorario]bool{
    Lunes:     true,
    Martes:    true,
    Miércoles: true,
    Jueves:    true,
    Viernes:   true,
    Sábado:    true,
    Domingo:   true,
    Festivo:   true,
}
