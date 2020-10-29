module routes

go 1.15

replace person => ../person

require (
	github.com/gin-gonic/gin v1.6.3
	person v0.0.0-00010101000000-000000000000
)
