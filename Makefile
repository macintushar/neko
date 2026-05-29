.PHONY: hello build ui clean

build: ui
	cd server && go build -o ../neko .

ui:
	cd ui && bun ci && bun run build
	rm -rf server/ui-dist
	cp -r ui/dist server/ui-dist

clean:
	rm -rf server/static ui/dist neko
