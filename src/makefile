# Имя бинарника
BINARY = readDB

# Исходники
SRC = .

# Цель по умолчанию
.PHONY: all
all: build1 build2 build3

# Сборка первого задания
.PHONY: build1
build1:
	go build -o ./exercise00/$(BINARY) $(SRC)/exercise00

# Сборка второго задания
.PHONY: build2
build2:
	go build -o ./exercise01/$(BINARY) $(SRC)/exercise01

# Сборка третьего задания
.PHONY: build3
build3:
	go build -o ./exercise002/$(BINARY) $(SRC)/exercise02

# Тесты
.PHONY: test
test:
	go test -v ./...

# Очистка полная
.PHONY: cleanall
cleanall:
	rm -f ./exercise00/$(BINARY) ./exercise01/$(BINARY) ./exercise02/$(BINARY) 
