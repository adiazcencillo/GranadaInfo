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
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "nombre-monumento" {
					nodosH3 = append(nodosH3, n)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			recorrerNodos(c)
		}
	}

	recorrerNodos(doc)

	if len(nodosH3) == 0 {
		return nil, fmt.Errorf("No se encontraron nodos <h3> con id 'nombre-monumento'")
	}

	return nodosH3, nil
}

func extraerNodosHorarios(doc *html.Node) ([]*html.Node, error) {

	var nodosHorario []*html.Node

	var recorrerNodos func(*html.Node)
	recorrerNodos = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "p" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "horario" {
					nodosHorario = append(nodosHorario, n)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			recorrerNodos(c)
		}
	}

	recorrerNodos(doc)

	if len(nodosHorario) == 0 {
		return nil, fmt.Errorf("No se encontraron nodos <p> con id 'horario'")
	}

	return nodosH3, nil
}


func extraerStringNodo(n *html.Node) (string, error) {
	if n == nil {
		return "", fmt.Errorf("El nodo proporcionado es nil")
	}

	var texto string
	if n.Type == html.TextNode {
		texto = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		subTexto, err := extraerStringNodo(c)
		if err != nil {
			return "", err
		}
		texto += subTexto
	}

	if texto == "" {
		return "", fmt.Errorf("El nodo no contiene texto")
	}
	return texto, nil
}


func extraerHorarioInvierno(s string) {

}

