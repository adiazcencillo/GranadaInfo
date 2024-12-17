package horario

import(
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func cargarDocumento(filePath string) (*html.Node, error) {

	archivo, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("No se pudo abrir el archivo: %w", err)
	}
	defer archivo.Close()

	contenido, err := html.Parse(archivo)
	if err != nil {
		return nil, fmt.Errorf("Error parseando el HTML: %w", err)
	}

	return contenido, nil

}

func extraerNodosH3(doc *html.Node) ([]*html.Node, error) {
	var nodosH3 []*html.Node

	var recorrerNodos func(*html.Node)
	recorrerNodos = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h3" {
			texto := extraerStringNodo(n)

			if texto != "MONUMENTOS DE GRANADA" {
				nodosH3 = append(nodosH3, n)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			recorrerNodos(c)
		}
	}

	recorrerNodos(doc)

	return nodosH3, nil
}
