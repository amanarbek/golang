# Day 00 - Go Boot camp

## Statistics being handy

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00: Anscombe](#exercise-00-anscombe)


<h2 id="chapter-i" >Chapter I</h2>
<h2 id="general-rules" >General rules</h2>

- Your programs should not <ins>quit unexpectedly</ins>(завершиться неожиданно) (giving an error on a valid input). If this happens, your project will be considered non functional and will receive a 0 during the <ins>evaluation</ins>(оценивание).
- We encourage you to create test programs for your project even though this work won't have to be submitted and won't be graded. It will give you a chance to easily test your work and your peers' work. You will find those tests especially useful during your defence. Indeed, during defence, you are free to use your tests and/or the tests of the peer you are evaluating.
- Submit your work to your assigned git repository. Only the work in the git repository will be graded.
- If your code is using external dependencies, it should use [Go Modules](https://go.dev/blog/using-go-modules) for managing them

<h2 id="chapter-ii" >Chapter II</h2>
<h2 id="rules-of-the-day" >Rules of the day</h2>

- You should only <ins>turn in</ins>(сдавать) `*.go` files and (in case of external dependencies) `go.mod` + `go.sum`
- Your program should accept a sequence of numbers separated by newlines to its standard input. One number is also a sequence.
- You may assume it should only work with integers (the output can be floats though, <ins>rounded to 2 decimal places</ins>(округлённый до двух знаков после запятой)).
- Nevertheless it should print a <ins>meaningful error message</ins>(осмысленное сообщение об ошибке) without runtime panic if <ins>fed</ins>(быть снабжённым) some unexpected input, say, <ins>number out of bounds</ins>(число вне допустимого диапазона), letter characters or empty string.
- Your code for this task should be buildable with just `go build`

<h2 id="chapter-iii" >Chapter III</h2>
<h2 id="intro" >Intro</h2>

Go isn't generally considered a language of Data Science. But it doesn't mean that it can't crunch
numbers. In fact, it's comparable to C on basic tasks. Besides, it can be a lot easier to write, 
partially because GC handles <ins>memory management</ins>(управление памятью) and also Go's standard library is pretty good.
We're constantly taught that it can be a bad idea to just <ins>trust the gut</ins>(доверять интуиции) when <ins>dealing with</ins>(иметь дело с чем-то) 
important data. <ins>To make heads or tails of</ins>(разобраться в чём-либо) a sample of numbers it's usually better to use
statistical <ins>approach</ins>(подход). Data sometimes be can be deceptive, too, like, for example,
[Anscombe's quartet](https://en.wikipedia.org/wiki/Anscombe%27s_quartet), but the more metrics
we get - the more weighted decision we will able to make at the end, isn't it?


<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Exercise 00: Anscombe</h3>


So, let's say we have <ins>a bunch of integer numbers</ins>(куча целых чисел), strictly between -100000 and 100000. It may 
probably be a large set, so <ins>let's assume</ins>(давай предположим) our application will read it from a standard input, 
separated by newlines. Right now let's think of four major statistical metrics that we can <ins>derive
from this data</ins>(выводить на основе данных), so by default we can print all of them as a result, for example, like this:

```
Mean: 8.2
Median: 9.0
Mode: 3
SD: 4.35   (Standard deviation (SD) – стандартное отклонение)
```

The order and actual format doesn't really matter as long as we can understand <ins>which is which</ins>.(что есть что) 
A couple of notes, though:

1) Input data may or may not be sorted. You don't need to write your own sorting algorithm,
luckily, Go already has one in standard library and it works for integers.
2) Median is a middle number of a sorted sequence if its size is odd, and an average between
two middle ones if their count is even.
3) Mode is a number which is occurring most frequently, and if there are several, <ins>the smallest one
among those</ins>(наименьшее из них) is returned. You may think of using some structure for storing number counts, and some
Go standard container may help.
4) You can use both population and regular standard deviation, whichever you prefer.
5) <ins>Calling someone "average" can be mean.</ins>(Называть кого-то средним может быть подло)

It will also make sense for user to be able to choose specifically, which of these four parameters
to print, so need to implement this as well. By default it's all of them, but there should be 
a way to specify whether print just one, two or three specific metrics out of four when running
the program (<ins>without recompilation</ins>(без перекомпиляции)). There is a <ins>built-in module</ins>(встроенный модуль) in standard library that allows you 
<ins>to parse additional parameters</ins>(разбирать дополнительные параметры).

### Чеклист 

- [x] Функция считывания из стандартного потока.
- [x] Обработка ошибок считанного.
- [x] Сортировка полученного массива чисел.
- [x] Функция вычисления среднего арифметического значения (mean).
- [x] Вычисления медианы (median). Если нечетная длина массива то число по середине, иначе среднее из 2 чисел.
- [x] Функция определения частотного числа в массиве (mode).
- [x] Реализовать 2 функции для вычисления стандартного отклонения:
  - [x] Выборка (sample standard deviation).
  - [x] Генеральная совокупность (population standard deviation).
- [x] Написать тесты для функций.
- [x] Реализовать работу с флагами для выбора какие поля выводить.
- [ ] Использовать многопоточность для вычисления значений.