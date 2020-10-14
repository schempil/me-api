module main

go 1.15

replace person => ./person

require (
	github.com/gorilla/mux v1.8.0
	person v0.0.0-00010101000000-000000000000 // indirect
)
