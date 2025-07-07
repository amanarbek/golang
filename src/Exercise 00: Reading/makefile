# Имя бинарника
BINARY = readDB

# Исходники
SRC = .

# Цель по умолчанию
.PHONY: all
all: build

# Сборка бинарника
.PHONY: build
build:
	go build -o $(BINARY) $(SRC)

# Тесты
.PHONY: test
test:
	go test -v ./...

# Очистка
.PHONY: clean
clean:
	rm -f $(BINARY)
