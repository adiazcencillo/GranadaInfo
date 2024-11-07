package horario

type Horario struct {
    DiasCerrado     []bool
    HorariosApertura map[ClaveHorario][]string
    SiempreAbierto   bool
}

func NuevoHorario(siempreAbierto bool, diasCerrado []bool, horariosApertura map[ClaveHorario][]string) (*Horario, error) {
    if siempreAbierto {
        horariosApertura = map[ClaveHorario][]string{
            Lunes:     {"00:00-23:59"},
            Martes:    {"00:00-23:59"},
            Miércoles: {"00:00-23:59"},
            Jueves:    {"00:00-23:59"},
            Viernes:   {"00:00-23:59"},
            Sábado:    {"00:00-23:59"},
            Domingo:   {"00:00-23:59"},
            Festivo:   {"00:00-23:59"},
        }
        diasCerrado = []bool{false, false, false, false, false, false, false, false}
    } 
    return &Horario{
        DiasCerrado:      diasCerrado,
        HorariosApertura: horariosApertura,
        SiempreAbierto:   siempreAbierto,
    }, nil
}