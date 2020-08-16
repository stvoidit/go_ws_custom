build:
	mkdir build
	go build -o ./build .
	cd ws && yarn && yarn build