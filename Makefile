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

run-all-terms:
	gnome-terminal --title="Auth Service" -- bash -c "cd services/authentication-service && go run main.go; exec bash" & \
	gnome-terminal --title="Order Service" -- bash -c "cd services/order-service && go run main.go; exec bash" & \
	gnome-terminal --title="Product Service" -- bash -c "cd services/product-service && go run main.go; exec bash"

dev-docker-compose:
	sudo docker-compose -f docker-compose.dev.yml up --build

production-docker-compose:
	docker-compose up --build

