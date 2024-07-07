ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ifeq ($(OS),Windows_NT)
	MAIN_PATH = /tmp/bin/main.exe
	SYNC_ASSETS_COMMAND =	@go run github.com/makiuchi-d/arelo@v1.13.1 \
	--target "./public" \
	--pattern "**/*.js" \
	--pattern "**/*.css" \
	--delay "100ms" \
	--templ generate --notify-proxy
else
	MAIN_PATH = tmp/bin/main
	SYNC_ASSETS_COMMAND =	@go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "public" \
	--build.include_ext "js,css" \
	--screen.clear_on_rebuild true \
	--log.main_only true
endif

templ:
	@templ generate --watch --proxy="http://localhost:3000" --open-browser=false

server:
	@go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build --tags dev -o ${MAIN_PATH} ./" --build.bin "${MAIN_PATH}" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true \
	--screen.clear_on_rebuild true \
	--log.main_only true

dev:
	@make -j5 templ server
