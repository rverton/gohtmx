all: tailwind
	go build -o gohtmx ./cmd/gohtmx

run:
	go run ./cmd/gohtmx

watch:
	nodemon --watch '*' -e html,go  --exec go run ./cmd/gohtmx --signal SIGTERM

tailwind:
	cd public && npx tailwindcss -i ./styles.css -o ./output.css

tailwind-watch:
	cd public && npx tailwindcss -i ./styles.css -o ./output.css --watch

test:
	go test ./...
