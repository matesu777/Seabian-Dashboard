.PHONY: dev air tailwind build run clean

dev:
	@$(MAKE) -j2 air tailwind

air:
	air

tailwind:
	npm run dev

build:
	npm run build
	go build -o bin/seabian-dashboard .

run: build
	./bin/seabian-dashboard

clean:
	rm -rf bin/
