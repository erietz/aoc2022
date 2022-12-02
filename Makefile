run: day02

day01:
	go run main.go < ./inputs/day01/part1.txt

day02:
	go run main.go < ./inputs/day02/part1.txt

test:
	go test -v ./...
