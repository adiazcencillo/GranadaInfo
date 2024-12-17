package horario

import(
	"testing"

	"golang.org/x/net/html"

	"github.com/onsi/gomega"
)

func TestCargarDocumento(t *testing.T, filePath string) {
	g := gomega.NewWithT(t)

	t.Run("Cargar documento HTML correctamente", func(t *testing.T) {

		doc, err := cargarDocumento(filePath)

		g.Expect(err).NotTo(gomega.HaveOccurred()) 
		g.Expect(doc).NotTo(gomega.BeNil())         
	})
}

