package horario

import(
	"testing"
	
	"golang.org/x/net/html"

	"github.com/onsi/gomega"
)

func TestCargarDocumento(t *testing.T) {
	g := gomega.NewWithT(t)

	t.Run("Cargar documento HTML correctamente", func(t *testing.T) {

		filePath := "/home/adiazcencillo/horarios.html"

		contenido, err := cargarDocumento(filePath)

		g.Expect(err).NotTo(gomega.HaveOccurred()) 
		g.Expect(contenido).NotTo(gomega.BeNil())         
	})

	t.Run("Archivo no encontrado", func(t *testing.T) {

		filePath := "noexiste.html"

		doc, err := cargarDocumento(filePath)

		g.Expect(err).To(gomega.HaveOccurred())
		g.Expect(doc).To(gomega.BeNil())
	})
}

func TestExtraerNodosH3(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "/home/adiazcencillo/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosH3, err := extraerNodosH3(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(nodosH3).To(gomega.HaveLen(21))
	
}

func TestExtraerNodosHorarios(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "/home/adiazcencillo/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosHorario, err := extraerNodosHorarios(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(nodosHorario).To(gomega.HaveLen(21))
}	

func TestExtraerStringNodo(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "/home/adiazcencillo/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosH3, err := extraerNodosH3(doc)
	nodosHorario, err := extraerNodosHorarios(doc)

	t.Run("Extraer texto del primer nodo <h3>", func(t *testing.T) {
		texto, err := extraerStringNodo(nodosH3[0])
		g.Expect(err).NotTo(gomega.HaveOccurred())
		g.Expect(texto).To(gomega.Equal("Alhambra"))
	})

	t.Run("Extraer texto del primer nodo de horarios", func(t *testing.T) {
		texto, err := extraerStringNodo(nodosHorario[0])
		g.Expect(err).NotTo(gomega.HaveOccurred()) 
		g.Expect(texto).To(gomega.Equal("Invierno: 08:30 - 18:00 y 19:00 - 20:45 ; Verano: 08:30 - 20:00 y 22:00 - 23:30"))
	})

	t.Run("Extraer texto de un nodo vacío", func(t *testing.T) {
		nodoVacio := &html.Node{} // Crear un nodo vacío
		texto, err := extraerStringNodo(nodoVacio)
		g.Expect(err).To(gomega.HaveOccurred()) 
		g.Expect(texto).To(gomega.BeEmpty())  
	})

}

func TestExtraerHorario(t *testing.T) {
	g := gomega.NewWithT(t)

	var stringHorario string = "Invierno: 10:00 - 14:00 y 16:00 - 18:00 (lunes-viernes), 10:00 - 18:00 (sábado, domingo y festivos); Verano: 10:00 - 14:00 y 18:00 - 20:00 (lunes-viernes); 10:00 - 20:00 (sábado, domingo y festivos)"

	horario, err := extraerHorario(stringHorario)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(horario).NotTo(gomega.BeNil())

	g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario("Lunes")))
	g.Expect(horario.HorariosApertura[ClaveHorario("Lunes")]).To(gomega.Equal("10:00 - 14:00 y 16:00 - 18:00"))

	g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario("Martes")))
	g.Expect(horario.HorariosApertura[ClaveHorario("Martes")]).To(gomega.Equal("10:00 - 14:00 y 16:00 - 18:00"))

	g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario("Miércoles")))
	g.Expect(horario.HorariosApertura[ClaveHorario("Miércoles")]).To(gomega.Equal("10:00 - 14:00 y 16:00 - 18:00"))

	g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario("Jueves")))
	g.Expect(horario.HorariosApertura[ClaveHorario("Jueves")]).To(gomega.Equal("10:00 - 14:00 y 16:00 - 18:00"))

	g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario("Viernes")))
	g.Expect(horario.HorariosApertura[ClaveHorario("Viernes")]).To(gomega.Equal("10:00 - 14:00 y 16:00 - 18:00"))

	g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario("Sábado")))
	g.Expect(horario.HorariosApertura[ClaveHorario("Sábado")]).To(gomega.Equal("10:00 - 18:00"))

	g.Expect(horario.HorariosApertura).To(gomega.HaveKey(ClaveHorario("Domingo")))
	g.Expect(horario.HorariosApertura[ClaveHorario("Domingo")]).To(gomega.Equal("10:00 - 18:00"))

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

	dia := "lunes"
	stringHorario := "10:00 - 14:00 y 16:00 - 18:00 (lunes-viernes), 10:00 - 18:00 (sábado, domingo y festivos)"
	horarioSabado, err := extraerHorarioDias(dia, stringHorario)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(horarioSabado).NotTo(gomega.BeNil())
	g.Expect(horarioSabado).To(gomega.Equal("10:00 - 14:00 y 16:00 - 18:00")

}		


