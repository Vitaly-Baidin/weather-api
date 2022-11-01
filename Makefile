check_install:
	which swagger || echo 'swagger is not installed'
swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models
compose-up:
	docker-compose up --build -d && docker-compose logs -f
