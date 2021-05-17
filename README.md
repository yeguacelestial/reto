# Reto: API Rest

Hace menos de 7 días, lo único que sabía de Go fue que era un lenguaje hecho por Google. Este proyecto de verdad
me sirvió no solo para postularme a la vacante de desarrollador, sino también para introducirme a este lenguaje. Sin duda alguna, hoy me encuentro bastante interesado en profundizar mi aprendizaje en Go.

La API fue desarrollada en la versión Go 1.16.4 (go version go1.16.4 linux/amd64).

Se desarrollaron unit tests para casi todo, pero le di enfoque a testear los buenos y peores casos al realizar llamadas a los endpoints.

# Pre-requisitos

## Clonar el repositorio

Para poder trabajar con el proyecto, es necesario clonar el repositorio:

`git clone https://github.com/yeguacelestial/reto.git`

## Crear archivo de variables de entorno

Antes de ejecutar el servidor localmente, se requiere crear un archivo `.env` en la carpeta `api/login/` con el siguiente contenido:

```bash
JWT_SECRET='nomada'
```

Esta es necesaria para poder generar y validar los JWT de forma correcta, ya que no está escrita directamente en el código.

## Unit tests

Las pruebas unitarias del proyecto se realizaron con GitHub Actions. Cada que se realiza un commit a la branch de `main`, estas pruebas se ejecutan en Windows, Mac OS y Linux de forma paralela, en las versiones de Go `1.15.x` y `1.16.x`.

Para ver los resultados de los tests, basta con hacer click en la palomita verde (o tachita roja si algo no salió bien) del último commit realizado.

# Probando la API localmente

## Instalar dependencias

Antes de correr el servidor, es necesario instalar las dependencias. Para esto, dentro de la carpeta raíz del proyecto `api/`, se ejecuta el siguiente comando:

```bash
go mod tidy
```

## Correr el servidor

Una vez que se cumplieron los pre-requisitos, basta con posicionarse en la carpeta raíz del proyecto `api/` y ejecutar el siguiente comando:

```bash
go run .
```

Si en la consola vemos el siguiente contenido:

```bash
[*] REST API - Mux Router
[*] Serving on port :10000

[*] Created default user on database with email: demo@usuario.com
```

...entonces la API está lista para ser consumida por HTTP.

## Correr unit tests localmente

Los tests principales de la API se encuentran en `api/main_test.go`.

Para correr los unit tests del proyecto de la API, basta con posicionarse en la carpeta `api/` y ejecutar el comando
```bash
go test .
```
