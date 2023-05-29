package main

import (
	"context"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

func exampleScanySelect(ctx context.Context, pool *pgxpool.Pool) {
	// sqlx не совместим с pgx (можно сделать их совместимыми используя github.com/jackc/pgx/v4/stdlib, но вы теряете все преимущества pgx)
	//
	// github.com/georgysavva/scany - Библиотека для сканирования данных из базы данных в структуры  (аля sqlx)

	const query = `SELECT id, first_name, last_name, age FROM students`

	// делаем SELECT в пару строк без циклов и сканирования)
	var students []Student
	if err := pgxscan.Select(ctx, pool, &students, query); err != nil {
		log.Fatal(err)
	}

	log.Println(students)
}

func exampleScanyQueryV1(ctx context.Context, pool *pgxpool.Pool) {
	const query = `SELECT id, first_name, last_name, age FROM students`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// ScanAll - Если у нас необычынй(особый) объект, то можем для него реализовать свой метод Scan()
	// в данном случае используется StructScan по умолчанию
	var students []Student
	if err := pgxscan.ScanAll(&students, rows); err != nil {
		log.Fatal(err)
	}

	log.Println(students)
}

func exampleScanyQueryV2(ctx context.Context, pool *pgxpool.Pool) {
	const query = `SELECT id, first_name, last_name, age FROM students`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	students := make(map[int64]Student)

	// построчная обработка
	rs := pgxscan.NewRowScanner(rows)
	for rows.Next() {
		var st Student
		if err := rs.Scan(&st); err != nil {
			log.Fatal(err)
		}
		// do something here
		students[st.ID] = st
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(students)
}
