package horario

import(
	"fmt"
	"os"
	"regexp"

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

	return nodosHorario, nil
}


func extraerStringNodo(n *html.Node) (string, error) {

	if n == nil {
		return "", fmt.Errorf("El nodo proporcionado es nil")
	}
	var texto string

	var extraerTexto func(*html.Node)
	extraerTexto = func(n *html.Node) {
		if n.Type == html.TextNode {
			texto += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extraerTexto(c)
		}
	}
	extraerTexto(n)

	if texto == "" {
		return "", fmt.Errorf("El nodo no contiene texto")
	}

	return texto, nil
}


func extraerHorarioInvierno(cadena string) (string, error) {
	re := regexp.MustCompile(`(Invierno)?:?([^;]+);?`)
	matches := re.FindStringSubmatch(cadena)

	if len(matches) == 0 {
		return "", fmt.Errorf("No se encontraron horarios de invierno en la cadena proporcionada")
	}

	return matches[2], nil
}

func extraerHorarioDias(dia string, cadena string) (string, error) {

	re := regexp.MustCompile(fmt.Sprintf(`(\d{2}:\d{2}) - (\d{2}:\d{2})(?: y (\d{2}:\d{2}) - (\d{2}:\d{2}))? \((.*?%s.*?)\)`, dia))
	matches := re.FindStringSubmatch(cadena)

	if len(matches) < 4 {
		diaInicio, diaFin := extraerIntervaloDias(cadena)

		if(diaInicio == "mismo horario") {
			return cadena, nil
		} else {
			if(contenidoIntervaloDias(dia, diaInicio, diaFin)) {
				horario := regexp.MustCompile(`((\d{2}:\d{2}) - (\d{2}:\d{2})(?: y (\d{2}:\d{2}) - (\d{2}:\d{2}))?) \((.*?)-(.*?)\)`)
				matches_horario := horario.FindStringSubmatch(cadena)
				
				return matches_horario[1], nil
			} else {
				return "vacio", nil
			}
		}	

	} else {

		horario_regex := regexp.MustCompile(fmt.Sprintf(`((\d{2}:\d{2}) - (\d{2}:\d{2})(?: y (\d{2}:\d{2}) - (\d{2}:\d{2}))?) \((.*?%s.*?)\),`, dia))
		matches_horario := horario_regex.FindStringSubmatch(cadena)
		horario := ""

		if len(matches_horario) == 0 {
			horario_regex = regexp.MustCompile(fmt.Sprintf(`, ((\d{2}:\d{2}) - (\d{2}:\d{2})(?: y (\d{2}:\d{2}) - (\d{2}:\d{2}))?) \((.*?%s.*?)\)`, dia))
			matches_horario = horario_regex.FindStringSubmatch(cadena)

			horario = matches_horario[1]
		} else {
			horario = matches_horario[1]
		}

		return horario, nil
	}
}


func extraerIntervaloDias(cadena string) (string, string) {

	re_dias := regexp.MustCompile(`\((.*?)-(.*?)\)`)
	match := re_dias.FindStringSubmatch(cadena)

	if len(match) == 0 {

		return "mismo horario", ""

	} else {

		diaInicio := match[1]
		diaFin := match[2]

		return diaInicio, diaFin

	}

}

func contenidoIntervaloDias(dia string, inicio string, fin string) (bool) {

	diasSemana := map[string]int{
		"lunes":    0,
		"martes":   1,
		"miércoles": 2,
		"jueves":   3,
		"viernes":  4,
		"sábado":   5,
		"domingo":  6,
	}

	numeroDia := diasSemana[dia]
	numeroDiaInicio := diasSemana[inicio]
	numerioDiaFin := diasSemana[fin]

	if(numeroDia > numeroDiaInicio && numeroDia < numerioDiaFin) {
		return true
	} else {
		return false
	}
}

func extraerHorario(cadena string) (*Horario, error) {

	diasSemana := []string{"lunes", "martes", "miércoles", "jueves", "viernes", "sábado", "domingo"}
	horarios := make(map[ClaveHorario]string)
	cerrado := make([]bool, len(diasSemana))

	var err error
	for i, dia := range diasSemana {
		horarioDia, errDia := extraerHorarioDias(dia, cadena)
		if errDia != nil {
			err = errDia
		}

		if horarioDia == "vacio" {
			cerrado[i] = true
		} else {
			horarios[ClaveHorario(dia)] = horarioDia
		}

	}

	horario := NuevoHorario(cerrado, horarios)
	return horario, err
}

func extraerMonumentos(n *html.Node) ([]*Monumento, error) {

	nodosHorarios, err := extraerNodosHorarios(n)
	nodosH3, err := extraerNodosH3(n)
	if err != nil {
		return nil, fmt.Errorf("Error al extraer los nodos")
	}

	if len(nodosHorarios) != len(nodosH3) {
		return nil, fmt.Errorf("El número de nodos de horarios y títulos no coincide")
	}

	var monumentos []*Monumento

	

	for i:=0; i<len(nodosH3); i++ {
		nombreMonumento, err:= extraerStringNodo(nodosH3[i])

		if err != nil {
			return nil, fmt.Errorf("Error al extraer el nombre del monumento en índice %d: %v", i, err)
		}

		horarioTexto, err := extraerStringNodo(nodosHorarios[i])
		if err != nil {
			return nil, fmt.Errorf("error al extraer el horario en índice %d: %v", i, err)
		}

		horario, err := extraerHorario(horarioTexto)
		if err != nil {
			return nil, fmt.Errorf("error al analizar el horario en índice %d: %v", i, err)
		}

		monumento := NuevoMonumento(nombreMonumento, *horario)
		monumentos = append(monumentos, monumento)
	}

	return monumentos, nil
}