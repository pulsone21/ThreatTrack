module frontend

go 1.21.5

require (
	github.com/gorilla/mux v1.8.1
	github.com/pulsone21/threattrack v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.5.0 // indirect

replace github.com/pulsone21/threattrack => ../
