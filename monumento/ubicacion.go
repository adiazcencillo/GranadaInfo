package monumento

// Ubicacion representa la dirección de un monumento. No se añade cuidad ni país ya que todos están en Granada.
// Una ubicación no tiene identidad propia, basa su identidad en sus atributos, por tanto es un objeto valor
type Ubicacion struct {
    Calle  string  // Nombre de la calle
    Numero string  // Número de la calle
}
