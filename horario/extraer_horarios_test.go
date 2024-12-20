package horario

import(
	"testing"
	
	"golang.org/x/net/html"

	"github.com/onsi/gomega"
)

func TestExtraerNodosH3(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "../data/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosH3, err := extraerNodosH3(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(nodosH3).To(gomega.HaveLen(21))
	
}

func TestExtraerNodosHorarios(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "../data/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosHorario, err := extraerNodosHorarios(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(nodosHorario).To(gomega.HaveLen(21))
}	

func TestExtraerStringNodo(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "../data/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosH3, err := extraerNodosH3(doc)
	nodosHorario, err := extraerNodosHorarios(doc)

	runTest := func(t *testing.T, nodo *html.Node, expected string, shouldError bool, description string) {
		t.Run(description, func(t *testing.T) {
			g := gomega.NewWithT(t)
			texto, err := extraerStringNodo(nodo)
			if shouldError {
				g.Expect(err).To(gomega.HaveOccurred())
				g.Expect(texto).To(gomega.BeEmpty())
			} else {
				g.Expect(err).NotTo(gomega.HaveOccurred())
				g.Expect(texto).To(gomega.Equal(expected))
			}
		})
	}

	runTest(t, nodosH3[0], "Alhambra", false, "Extraer texto del primer nodo <h3>")
	runTest(t, nodosHorario[0], "Invierno: 08:30 - 18:00 y 19:00 - 20:45 ; Verano: 08:30 - 20:00 y 22:00 - 23:30", false, "Extraer texto del primer nodo de horarios")
	runTest(t, &html.Node{}, "", true, "Extraer texto de un nodo vacío")

}

func TestExtraerHorario(t *testing.T) {
	g := gomega.NewWithT(t)

	var stringHorario string = "Invierno: 10:00 - 14:00 y 16:00 - 18:00 (lunes-viernes), 10:00 - 18:00 (sábado, domingo y festivos); Verano: 10:00 - 14:00 y 18:00 - 20:00 (lunes-viernes); 10:00 - 20:00 (sábado, domingo y festivos)"

	horario, err := extraerHorario(stringHorario)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(horario).NotTo(gomega.BeNil())

	verifyHorario := func(dia, valorEsperado string) {
		g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario(dia)))
		g.Expect(horario.HorariosApertura[ClaveHorario(dia)]).To(gomega.Equal(valorEsperado))
	}

	verifyHorario("Lunes", "10:00 - 14:00 y 16:00 - 18:00")
	verifyHorario("Martes", "10:00 - 14:00 y 16:00 - 18:00")
	verifyHorario("Miércoles", "10:00 - 14:00 y 16:00 - 18:00")
	verifyHorario("Jueves", "10:00 - 14:00 y 16:00 - 18:00")
	verifyHorario("Viernes", "10:00 - 14:00 y 16:00 - 18:00")
	verifyHorario("Sábado", "10:00 - 18:00")
	verifyHorario("Domingo", "10:00 - 18:00")

	g.Expect(horario.DiasCerrado).To(gomega.Equal([]bool{false, false, false, false, false, false, false}))
}

func TestExtraerHorarioInvierno(t *testing.T) {
	g := gomega.NewWithT(t)

	var stringHorario string = "Invierno: 10:00 - 14:00 y 16:00 - 18:00 (lunes-viernes), 10:00 - 18:00 (sábado, domingo y festivos); Verano: 10:00 - 14:00 y 18:00 - 20:00 (lunes-viernes); 10:00 - 20:00 (sábado, domingo y festivos)"

	horarioInvierno, err := extraerHorarioInvierno(stringHorario)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(horarioInvierno).NotTo(gomega.BeNil())
	g.Expect(horarioInvierno).To(gomega.Equal(" 10:00 - 14:00 y 16:00 - 18:00 (lunes-viernes), 10:00 - 18:00 (sábado, domingo y festivos)"))
}

func TestExtraerHorarioDias(t *testing.T) {
	g := gomega.NewWithT(t)

	dia := "martes"
	stringHorario := "10:00 - 14:00 y 16:00 - 18:00 (lunes-viernes), 10:00 - 18:00 (sábado, domingo y festivos)"
	horarioSabado, err := extraerHorarioDias(dia, stringHorario)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(horarioSabado).NotTo(gomega.BeNil())
	g.Expect(horarioSabado).To(gomega.Equal("10:00 - 14:00 y 16:00 - 18:00"))

}

func TestExtraerMonumentos(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "../data/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	monumentos, err := extraerMonumentos(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(monumentos).To(gomega.HaveLen(21))

}


