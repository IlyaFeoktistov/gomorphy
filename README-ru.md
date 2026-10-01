# gomorphy

[EN](README.md) | **RU**

> Форк [jus1d/gomorphy](https://github.com/jus1d/gomorphy), который использует
> компилятор [Rech](https://github.com/IlyaFeoktistov/Rech). Добавляет
> [`Parses`](#все-разборы-с-тегами): все разборы словоформы с тегом у каждой
> формы её лексемы. В остальном совпадает с оригиналом.

Морфологический анализатор русского языка для Go на основе словарей [pymorphy3](https://github.com/no-plagiarism/pymorphy3) / [OpenCorpora](http://opencorpora.org/), встраиваемых в бинарник при компиляции.

## Установка

```
go get github.com/IlyaFeoktistov/gomorphy
```

## Использование

```go
import morph "github.com/IlyaFeoktistov/gomorphy"

a, err := morph.Default()
if err != nil {
	log.Fatal(err)
}

// Все грамматические формы слова
forms := a.WordForms("кошка")
// [кошка кошки кошке кошку кошкой кошке кошки кошек кошкам кошек кошками кошках]

// Тег OpenCorpora для слова
tag := a.Tag("кошка")
// "NOUN,inan,femn sing,nomn"

// Все формы словосочетания (или слова) с согласованием прилагательного и существительного
forms = a.PhraseFormsConcordant("красивая кошка")
// [красивая кошка красивой кошки красивой кошке красивую кошку ...]
```

### Все разборы с тегами

`Tag` и `WordForms` возвращают лучший тег и формы без тегов. `Parses`
возвращает все разборы точной словоформы (без подстановки «е»/«ё») вместе со
всей лексемой, у каждой формы — свой тег:

```go
for _, p := range a.Parses("стали") {
	fmt.Println(p.Tag, p.Lexeme[0].Word)
	// VERB,perf,intr plur,past,indc стать
	// NOUN,inan,femn sing,gent сталь
	// NOUN,inan,femn sing,datv сталь
	// ...
}
```

## Словарь

Встроенный словарь собран из набора данных OpenCorpora v0.92 (ревизия 417127), скомпилированного pymorphy2 v0.9.1. Он содержит:

- 5 140 055 словоформ
- 3 456 парадигм
- 5 532 грамматических тега

> Встраивание словаря увеличивает размер исполняемого файла примерно на 8,8 МБ.

## Лицензия

**Исходный код на Go** распространяется по [лицензии MIT](LICENSE).

**Встроенные словарные данные** (`data/`) производны от [OpenCorpora](http://opencorpora.org/)
и распространяются по лицензии [CC BY-SA 4.0](data/LICENSE). Распространяя бинарник со
встроенными данными, нужно соблюдать условия CC BY-SA 4.0 (указание авторства + ShareAlike).
