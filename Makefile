dockerBuild:
	docker-compose up --build

dockerBuildBac: 
	docker-compose up -d --build

dockerUp:
	docker-compose up

dockerDown:
	docker-compose down

goBuildR:
	go build main.go && ./main

goInstallMod:
	go mod download