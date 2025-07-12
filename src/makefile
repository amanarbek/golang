# Имя бинарника
FIRST = readDB
SECOND = compareDB
THIRD = compareFS

# Исходники
SRC = .

# Цель по умолчанию
.PHONY: all
all: build1 build2 build3

# Сборка первого задания
.PHONY: build1
build1:
	go build -o ./exercise00/$(FIRST) $(SRC)/exercise00

# Тест первого задания
.PHONY: test1
test1:
	./exercise00/$(FIRST) -f recipes/recipes.json
	./exercise00/$(FIRST) -f recipes/recipes.xml

# Сборка второго задания
.PHONY: build2
build2:
	go build -o ./exercise01/$(SECOND) $(SRC)/exercise01

# Тест второго задания
.PHONY: test2
test2:
	./exercise01/$(SECOND) --old recipes/recipes.xml --new recipes/recipes.json

# Сборка третьего задания
.PHONY: build3
build3:
	go build -o ./exercise002/compareDB $(SRC)/exercise02

# Тесты reader
.PHONY: test
test:
	go test -v ./...

# Очистка полная
.PHONY: cleanall
cleanall:
	rm -f ./exercise00/$(FIRST) ./exercise01/$(SECOND) ./exercise02/$(THIRD) 
