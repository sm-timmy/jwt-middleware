module local/model

go 1.21.1

require (
	golang.org/x/crypto v0.13.0
	gopkg.in/reform.v1 v1.5.1
	local/database v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace local/database => ../database
