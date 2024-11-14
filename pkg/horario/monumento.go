package horario

type Monumento struct {
	Nombre         string
	Horario        Horario
}

func NuevoMonumento(nombre string, horario Horario) *Monumento {
	return &Monumento{
		Nombre:  nombre,
		Horario: horario,
	}
}
