run-auth:
	go run services/authentication-service/main.go

run-order:
	go run services/order-service/main.go

run-product:
	go run services/product-service/main.go

run-all:
	$(MAKE) run-auth & \
	$(MAKE) run-order & \
	$(MAKE) run-product

dev-docker-compose:
	sudo docker-compose -f docker-compose.dev.yml up --build

production-docker-compose:
	docker-compose up --build

