package horario

type Horario struct {
    DiasCerrado     []bool
    HorariosApertura map[ClaveHorario]string
}

func NuevoHorario(diasCerrado []bool, horariosApertura map[ClaveHorario][]string) *Horario {
    return &Horario{
        DiasCerrado:      diasCerrado,
        HorariosApertura: horariosApertura,
    }
}