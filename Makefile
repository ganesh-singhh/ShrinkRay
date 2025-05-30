all: clean backend frontend

.DEFAULT_GOAL := all

FRONTEND_PORT ?= 3000

clean:
	@echo "Cleaning up..."
	rm -rf builds && \
	mkdir -p builds/frontend-builds/
	@echo "Cleaning done..!\n"

backend: 
	@echo "Starting backend build..."
	cd backend/cmd && \
	go mod tidy && \
	go build -o ../../builds/shrinkray-backend && \
	cp .env ../../builds/.env
	@echo "\nBackend build done and uploaded to builds directory..!\n"

frontend:
	@echo "Starting frontend build..."
	cd frontend/ && \
	npm install && \
	npm run build && \
	rm -rf ../builds/frontend-builds/ && \
	mkdir -p ../builds/frontend-builds/ && \
	mv build/* ../builds/frontend-builds/ && \
	rm -rf build
	@echo "Frontend build done..!\n"


serve-frontend:
	@echo "Starting frontend hosting..."
	serve -s builds/frontend-builds/ -l $(FRONTEND_PORT)


serve-backend:
	@echo "Starting backend server..."
	cd builds && ./shrinkray-backend
	

.PHONY:clean backend frontend all
