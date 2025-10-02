package main

import (
	"context"      // para context.Background()
	"database/sql" // para sql.Open()
	"fmt"          // para fmt.Println()
	"log"
	sqlc "tp/bd/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://abril:1234@localhost:5432/bd?sslmode=disable"
	//connStr := "user=abril password=1234 dbname=bd"
	bd, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer bd.Close()
	queries := sqlc.New(bd)
	ctx := context.Background()

	//CREAR LIBRO 1
	createdLibro, err := queries.CreateLibro(ctx, // Create
		sqlc.CreateLibroParams{
			Titulo:          "La tia Cosima",
			Autor:           "Florencia Bonelli",
			Descripcion:     "Cósima es una mujer en la plenitud de la vida. Psicóloga de profesión y especializada en el tratamiento del autismo infantil, posee una fundación de enorme prestigio, donde se respira un ambiente cuidado y buen humor",
			Valoracion:      sql.NullInt32{Int32: 4, Valid: true},
			Anio:            int32(2020),
			GeneroPrincipal: "Romance",
		})
	if err != nil {
		log.Fatalf("Falla en crear libro: %v", err)
	}
	fmt.Printf("Crear libro: %+v\n", createdLibro)

	libro, err := queries.GetLibroByID(ctx, createdLibro.ID) // Read One
	if err != nil {
		log.Fatalf("Falla en obtener libro: %v", err)
	}
	fmt.Printf("Recuperar libro: %+v\n", libro)

	//UPDATE LIBRO
	err1 := queries.UpdateLibro(ctx, sqlc.UpdateLibroParams{
		ID:              1,
		Titulo:          "La Tia Cosima",
		Autor:           "Florencia Bonelli",
		Descripcion:     "Cósima es una mujer en la plenitud de la vida. Psicóloga de profesión y especializada en el tratamiento del autismo infantil, posee una fundación de enorme prestigio, donde se respira un ambiente cuidado y buen humor. Allí trabaja con perros especialmente adiestrados para ayudar a los niños con alguna condición del espectro autista",
		Valoracion:      sql.NullInt32{Int32: 4, Valid: true},
		Anio:            int32(2020),
		GeneroPrincipal: "Romance",
	})
	if err1 != nil {
		log.Fatalf("Falla en actualizar libro: %v", err1)
	}
	libro, err2 := queries.GetLibroByID(ctx, createdLibro.ID) // Read One
	if err2 != nil {
		log.Fatalf("Falla en obtener libro: %v", err2)
	}
	fmt.Printf("Recuperar libro: %+v\n", libro)

	//CREAR LIBRO 2
	createdLibro, err3 := queries.CreateLibro(ctx, // Create
		sqlc.CreateLibroParams{
			Titulo:          "El Principe Cruel",
			Autor:           "Holly Black",
			Descripcion:     "Jude tenía siete años cuando sus padres fueron asesinados y, junto con sus dos hermana, fue trasladada a la traicionera Corte Suprema del Reino Feérico. Diez años más tarde, lo único que Jude desea, a pesar de ser una mera mortal, es sentir que pertenece a ese lugar.",
			Valoracion:      sql.NullInt32{Int32: 5, Valid: true},
			Anio:            int32(2018),
			GeneroPrincipal: "Fantasia",
		})
	if err3 != nil {
		log.Fatalf("Falla en crear libro: %v", err3)
	}
	fmt.Printf("Crear libro: %+v\n", createdLibro)

	//LISTAR LIBROS
	libros, err4 := queries.ListLibros(ctx)

	if err4 != nil {
		log.Fatalf("Falla en listar libros: %v", err4)
	}
	fmt.Printf("Recuperar libros: %+v\n", libros)

	//ELIMINAR LIBROS
	err5 := queries.DeleteLibro(ctx, 1)
	if err5 != nil {
		log.Fatalf("Falla en eliminar libro: %v", err5)
	}
	err5 = queries.DeleteLibro(ctx, 2)
	if err5 != nil {
		log.Fatalf("Falla en eliminar libro: %v", err5)
	}

	//LISTAR LIBROS
	libros2, err6 := queries.ListLibros(ctx)

	if err6 != nil {
		log.Fatalf("Falla en listar libros: %v", err6)
	}
	fmt.Printf("Recuperar libros: %+v\n", libros2)
}