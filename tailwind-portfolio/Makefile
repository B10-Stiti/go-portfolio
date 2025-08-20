# Local development
tailwind:
	npx @tailwindcss/cli -i ./web/static/input.css -o ./web/static/output.css --watch

run:
	air

# Production / Render
build:
	npx @tailwindcss/cli -i ./web/static/input.css -o ./web/static/output.css
	go build -o main main.go

start:
	./main

clean:
	rm -f main
