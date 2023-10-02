module local/api

go 1.21.1

replace local/controller => ./../controller

replace local/middleware => ../middleware

replace local/database => ../database

replace local/auth => ../auth

replace local/model => ../model

require (
	github.com/labstack/echo/v4 v4.11.1
	local/controller v0.0.0-00010101000000-000000000000
	local/middleware v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gopkg.in/reform.v1 v1.5.1 // indirect
	local/auth v0.0.0-00010101000000-000000000000 // indirect
	local/database v0.0.0-00010101000000-000000000000 // indirect
	local/model v0.0.0-00010101000000-000000000000 // indirect
)
