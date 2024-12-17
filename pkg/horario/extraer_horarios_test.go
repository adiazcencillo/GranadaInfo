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

	filePath := "test_single_block.html"
	doc, err := cargarDocumento(filePath)
	g.Expect(err).NotTo(gomega.HaveOccurred())

	nodosH3, err := extraerNodosH3(doc)
	g.Expect(err).NotTo(gomega.HaveOccurred())
	g.Expect(nodosH3).To(gomega.HaveLen(21))
}



