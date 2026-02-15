package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// =========================================================
// CUMPLIMIENTO PASO 2.3: IMPLEMENTACIÓN DE INTERFACES
// =========================================================
// Tema Unidad 3: Interfaces y Polimorfismo.
// Esta interfaz define el comportamiento obligatorio para cualquier
// elemento que quiera ser parte de nuestra biblioteca digital.
type Digital interface {git remote add ....
	ObtenerFicha() string // Método para listar información
	Descargar()           // Método para simular la descarga/acceso
}

// =========================================================
// CUMPLIMIENTO PASO 1.3: IDENTIFICACIÓN DE CLASES (STRUCTS)
// =========================================================
// Tema Unidad 1 y 2: Estructuras y Composición.
// Se identifica 'publicacion' como la clase base que contiene
// los atributos comunes para evitar duplicidad de código.

// Tema Unidad 3: Encapsulamiento.
// Los atributos (titulo, autor, anio) inician con minúscula para ser
// PRIVADOS. Esto impide que se modifiquen directamente desde fuera,
// obligando al uso de métodos controlados (Setters).
type publicacion struct {
	titulo string
	autor  string
	anio   int
}

// =========================================================
// CUMPLIMIENTO PASO 2.3: MANEJO DE ERRORES
// =========================================================
// Se implementan métodos Setters que retornan 'error' para validar datos.

// Setter para Título
func (p *publicacion) SetTitulo(t string) error {
	t = strings.TrimSpace(t)
	if t == "" {
		return errors.New("el título no puede estar vacío")
	}
	p.titulo = t
	return nil
}

// Setter para Autor
func (p *publicacion) SetAutor(a string) error {
	a = strings.TrimSpace(a)
	if a == "" {
		return errors.New("el autor es obligatorio")
	}
	p.autor = a
	return nil
}

// CUMPLIMIENTO PASO 2.4: COMENTARIOS EN LÓGICA COMPLEJA
// Setter para Año: Contiene lógica de negocio específica.
// Validamos que el año esté dentro de un rango histórico lógico
// (desde la invención de la imprenta moderna aprox. 1400 hasta la actualidad).
func (p *publicacion) SetAnio(a int) error {
	anioActual := time.Now().Year()
	if a < 1400 || a > anioActual {
		return errors.New("el año ingresado no es válido (Rango: 1400 - Actualidad)")
	}
	p.anio = a
	return nil
}

// Getters (Accesores Públicos)
func (p *publicacion) GetTitulo() string { return p.titulo }
func (p *publicacion) GetAutor() string  { return p.autor }

// =========================================================
// CUMPLIMIENTO PASO 1.3: CLASES DERIVADAS
// =========================================================

// --- CLASE 1: EBOOK (Libro de Texto) ---
// Hereda de 'publicacion' mediante composición.
type Ebook struct {
	publicacion        // Herencia
	formato     string // PDF, EPUB, MOBI
	paginas     int
}

// Constructor NuevoEbook: Centraliza la creación y validación.
func NuevoEbook(titulo, autor string, anio int, formato string, paginas int) (*Ebook, error) {
	e := &Ebook{}

	// Reutilizamos la lógica de validación de la clase padre
	if err := e.SetTitulo(titulo); err != nil {
		return nil, err
	}
	if err := e.SetAutor(autor); err != nil {
		return nil, err
	}
	if err := e.SetAnio(anio); err != nil {
		return nil, err
	}

	// Validaciones específicas de Ebook
	formato = strings.ToUpper(strings.TrimSpace(formato))
	validos := map[string]bool{"PDF": true, "EPUB": true, "MOBI": true}
	if !validos[formato] {
		return nil, errors.New("formato inválido (Use: PDF, EPUB o MOBI)")
	}
	e.formato = formato

	if paginas <= 0 {
		return nil, errors.New("el número de páginas debe ser positivo")
	}
	e.paginas = paginas

	return e, nil
}

// Implementación de la Interfaz 'Digital'
func (e Ebook) ObtenerFicha() string {
	return fmt.Sprintf("[TEXTO] %s | Autor: %s (%d) | Formato: %s (%d págs)",
		e.titulo, e.autor, e.anio, e.formato, e.paginas)
}

func (e Ebook) Descargar() {
	fmt.Printf(">> Descargando archivo: %s.%s ... [COMPLETADO]\n",
		strings.ReplaceAll(e.titulo, " ", "_"), strings.ToLower(e.formato))
}

// --- CLASE 2: AUDIOLIBRO ---
type AudioLibro struct {
	publicacion
	narrador string
	duracion int // Minutos
}

// Constructor NuevoAudioLibro
func NuevoAudioLibro(titulo, autor string, anio int, narrador string, duracion int) (*AudioLibro, error) {
	a := &AudioLibro{}
	if err := a.SetTitulo(titulo); err != nil {
		return nil, err
	}
	if err := a.SetAutor(autor); err != nil {
		return nil, err
	}
	if err := a.SetAnio(anio); err != nil {
		return nil, err
	}

	if strings.TrimSpace(narrador) == "" {
		return nil, errors.New("el narrador es obligatorio")
	}
	a.narrador = narrador

	if duracion <= 0 {
		return nil, errors.New("la duración debe ser mayor a 0")
	}
	a.duracion = duracion

	return a, nil
}

// Implementación de la Interfaz 'Digital'
func (a AudioLibro) ObtenerFicha() string {
	return fmt.Sprintf("[AUDIO] %s | Voz: %s | Duración: %d min",
		a.titulo, a.narrador, a.duracion)
}

func (a AudioLibro) Descargar() {
	fmt.Printf(">> Reproduciendo stream de audio: '%s' ... [EN PROGRESO]\n", a.titulo)
}

// =========================================================
// LÓGICA PRINCIPAL DEL SISTEMA (MAIN)
// =========================================================

// CUMPLIMIENTO PASO 1.2: DESARROLLO DEL SISTEMA SELECCIONADO
var biblioteca []Digital // Slice polimórfico

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Datos de prueba (Avance significativo del software)
	l1, _ := NuevoEbook("Don Quijote", "Cervantes", 1605, "PDF", 1200)
	a1, _ := NuevoAudioLibro("1984", "George Orwell", 1949, "Francisco M.", 480)
	biblioteca = append(biblioteca, l1, a1)

	for {
		limpiarPantalla()
		fmt.Println("=== SISTEMA DE GESTIÓN DE LIBROS ELECTRÓNICOS ===")
		fmt.Println("1. Registrar nuevo Ebook")
		fmt.Println("2. Registrar nuevo Audiolibro")
		fmt.Println("3. Consultar Catálogo")
		fmt.Println("4. Acceder a Contenido (Polimorfismo)")
		fmt.Println("5. Salir")
		fmt.Print(">> Seleccione opción: ")

		var op int
		fmt.Scanln(&op)
		reader.ReadString('\n')

		switch op {
		case 1:
			agregarMaterial(reader, "Ebook")
		case 2:
			agregarMaterial(reader, "Audio")
		case 3:
			listarMaterial()
		case 4:
			accederMaterial(reader)
		case 5:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción inválida.")
		}

		fmt.Println("\n[Presione Enter para continuar...]")
		reader.ReadString('\n')
	}
}

func agregarMaterial(r *bufio.Reader, tipo string) {
	fmt.Printf("\n--- REGISTRAR %s ---\n", strings.ToUpper(tipo))

	fmt.Print("Título: ")
	tit, _ := r.ReadString('\n')
	fmt.Print("Autor: ")
	aut, _ := r.ReadString('\n')
	fmt.Print("Año de publicación: ")
	var anio int
	fmt.Scanln(&anio)
	r.ReadString('\n')

	var err error
	var nuevoItem Digital // Variable de tipo Interfaz

	if tipo == "Ebook" {
		fmt.Print("Formato (PDF/EPUB/MOBI): ")
		form, _ := r.ReadString('\n')
		fmt.Print("Total de páginas: ")
		var pags int
		fmt.Scanln(&pags)

		// Llamada al constructor con validación de errores
		nuevoItem, err = NuevoEbook(tit, aut, anio, strings.TrimSpace(form), pags)
	} else {
		fmt.Print("Nombre del Narrador: ")
		narr, _ := r.ReadString('\n')
		fmt.Print("Duración (minutos): ")
		var dur int
		fmt.Scanln(&dur)

		nuevoItem, err = NuevoAudioLibro(tit, aut, anio, strings.TrimSpace(narr), dur)
	}

	// CUMPLIMIENTO PASO 2.3: MANEJO DE ERRORES EN CAPA USUARIO
	if err != nil {
		fmt.Printf("\n[ERROR CRÍTICO]: No se pudo registrar.\nDetalle: %v\n", err)
	} else {
		biblioteca = append(biblioteca, nuevoItem)
		fmt.Println("\n[ÉXITO] Material registrado correctamente.")
	}
}

func listarMaterial() {
	fmt.Println("\n--- CATÁLOGO ---")
	if len(biblioteca) == 0 {
		fmt.Println("Biblioteca vacía.")
		return
	}
	for i, item := range biblioteca {
		// Polimorfismo: Go decide qué 'ObtenerFicha' ejecutar
		fmt.Printf("%d. %s\n", i+1, item.ObtenerFicha())
	}
}

func accederMaterial(r *bufio.Reader) {
	listarMaterial()
	if len(biblioteca) == 0 {
		return
	}

	fmt.Print("\nID a descargar: ")
	var idx int
	fmt.Scanln(&idx)

	if idx > 0 && idx <= len(biblioteca) {
		item := biblioteca[idx-1]
		item.Descargar() // Ejecuta Descargar() según sea Ebook o AudioLibro
	} else {
		fmt.Println("[!] ID no encontrado.")
	}
}

func limpiarPantalla() {
	fmt.Print("\033[H\033[2J")
}
