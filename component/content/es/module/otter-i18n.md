# otter/i18n
Utilidades I18n

## Configuración
Primero asegúrese de tener algunos archivos de traducción.

```json
// en.json
{
"helloWorld": "Hello World",
}
```
```json
// es.json
{
"helloWorld": "Hola Mundo",
}
```

Luego lea el archivo en tiempo de ejecución o en tiempo de compilación y páselo a i18n.AddLocale()

### Leer archivos de traducción en tiempo de compilación
```go
// translations.go
package translations

importar (
_ "embed"

"github.com/martinmunillas/otter/i18n"
)
//go:embed en.json
var EnJson []byte

//go:embed es.json
var EsJson []byte


func inicio() {
i18n.AddLocale("es", bytes.NewReader(traducciones.EnJson))
i18n.AddLocale("es", bytes.NewReader(traducciones.EsJson))
}
```
### Leer archivos de traducción en tiempo de ejecución
```go
// main.go
package translations

import (
    "os"
    "github.com/martinmunillas/otter/i18n"
)

func main() {
    file, _ := os.Open("./translations/en.json")
    i18n.AddLocale("en", file)
    _ = file.Close()

    file, _ = os.Open("./translations/es.json")
    i18n.AddLocale("es", file)
    _ = file.Close()
}
```
Incluso podrías obtener las traducciones desde una fuente remota. Una vez que hayas agregado tus configuraciones regionales, asegúrate de usar i18n.Middleware()
```go
// main.go
func main() {
    //...
    _ = http.ListenAndServe(server.PortString(8080), i18n.Middleware(mux))
}
```
Este middleware se asegurará de configurar la configuración regional del usuario según el contexto, haciéndolo disponible con `i18n.FromContext()`

## Cambio de configuración regional
La configuración regional del usuario está determinada por la cookie de `otter-lang`, pero cuando no se ha configurado, se utilizará el encabezado "Aceptar idioma".

La cookie `otter-lang` se configurará solo cuando el usuario decida cambiar su configuración regional predeterminada.

El middleware habilitará un punto final /set-locale que leerá la clave de configuración regional del cuerpo de la solicitud y la configurará en la cookie de otter-lang.

Puede crear sus propios configuradores locales, pero ya existe un componente i18n.LanguageSelector que envuelve este punto final. Puede diseñarlo desde css ya que es un elemento `<select />` con una clase "selector de idioma".

## Uso de traducciones
Para el uso de traducciones en sus archivos templ solo necesita llamar a i18n.T() con su contexto y la clave de traducción.

```go
// hola_mundo.templ
templ HolaMundo() {
    <h1>
        @i18n.T(ctx, "holaMundo")
    </h1>
}
```
Si sus traducciones contienen HTML sin formato, por ejemplo `"hola <b>mundo</b>"`, puede usar `i18n.RawT()`.
```go
// hola_mundo.templ
templ HolaMundo() {
    <h1>
        @i18n.RawT(ctx, "holaMundo")
    </h1>
}
```
También hay reemplazos disponibles, que son muy útiles cuando el contenido necesita más estilo o contenido dinámico.
```json
// en.json
{
    "hello": "¡Hola {name}, somos {logo} y estamos encantados de tenerte con nosotros!"
}
```
```go
// hola.templ
css logoStyles() {
    color: tomate;
    font-weight: 900;
    font-size: 0.875rem;
}

templ logo() {
    <p class={ logoStyles() }>IncreíbleStartup</p>
}

templ Hola(nombre: string) {
    <h1>
        @i18n.T(ctx, "hola", mapa[cadena]cualquier{
        nombre: nombre,
        logotipo: logo(),
        })
    </h1>
}
```
