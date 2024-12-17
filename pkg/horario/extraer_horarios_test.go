package horario

import(
	"testing"

	"golang.org/x/net/html"

	"github.com/onsi/gomega"
)

func TestCargarDocumento(t *testing.T) {
	g := gomega.NewWithT(t)

	t.Run("Cargar documento HTML correctamente", func(t *testing.T) {

		filePath := "~/horarios.html"

		doc, err := cargarDocumento(filePath)

		g.Expect(err).NotTo(gomega.HaveOccurred()) 
		g.Expect(doc).NotTo(gomega.BeNil())         
	})
}

func TestExtraerNodosH3(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "~/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosH3, err := extraerNodosH3(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(nodosH3).To(gomega.HaveLen(21))
	
}

func TestExtraerNodosHorarios(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "~/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosHorario, err := extraerNodosHorario(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(nodosH3).To(gomega.HaveLen(21))
}	

func TestExtraerStringNodo(t *testing.T) {
	g := gomega.NewWithT(t)

	filePath := "~/horarios.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosH3 := extraerNodosH3(doc)
	nodosHorario := extraerNodosHorario(doc)

	texto, err := extraerStringNodo(nodosH3[0])
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(texto).To(gomega.Equal("Alhambra"))

	texto, err := extraerStringNodo(nodosHorario[0])
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(texto).To(gomega.Equal("08:30 - 18:00, 19:00 - 20:45 (invierno); 08:30 - 20:00, 22:00 - 23:30 (verano)"))

}

