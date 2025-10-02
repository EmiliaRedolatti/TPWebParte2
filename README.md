# TPWebParte2

TPWebParte2 es una aplicación en **Go** que permite crear y consultar libros almacenados en una base de datos **PostgreSQL**.  
La idea es que cada usuario pueda registrar libros, consultar sus detalles (autor, género, año, descripción, valoración) y obtener un listado completo de los libros cargados.

## Ejecución del proyecto

### Requisitos
- **Go** 
- **Docker** y **Docker Compose**  

###  Clonar el repositorio
```bash
git clone https://github.com/EmiliaRedolatti/TPWebParte2.git
cd TPWebParte2
```
### Levantar la base de datos con Docker
Si ya existe el contenedor tp_postgres, eliminarlo:
```bash
docker rm -f tp_postgres
```
Levantar el contenedor:
```bash
docker compose up -d
```
### Ejecutar
```bash
go run .
```
