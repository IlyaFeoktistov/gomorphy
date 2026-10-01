# gomorphy

> Fork of [jus1d/gomorphy](https://github.com/jus1d/gomorphy) used by the
> [Rech](https://github.com/IlyaFeoktistov/Rech) compiler. Adds
> [`Parses`](#all-parses-with-tags): every parse of a word form with a tag on
> each form of its lexeme. Otherwise follows upstream.

Russian morphological analyzer for Go, backed by [pymorphy3](https://github.com/no-plagiarism/pymorphy3) / [OpenCorpora](http://opencorpora.org/) dictionaries embedded at compile time.

## Installation

```
go get github.com/IlyaFeoktistov/gomorphy
```

## Usage

```go
import morph "github.com/IlyaFeoktistov/gomorphy"

a, err := morph.Default()
if err != nil {
	log.Fatal(err)
}

// All grammatical forms of a word
forms := a.WordForms("кошка")
// [кошка кошки кошке кошку кошкой кошке кошки кошек кошкам кошек кошками кошках]

// OpenCorpora tag for a word
tag := a.Tag("кошка")
// "NOUN,inan,femn sing,nomn"

// All forms of a phrase (or a word) with adjective–noun agreement
forms = a.PhraseFormsConcordant("красивая кошка")
// [красивая кошка красивой кошки красивой кошке красивую кошку ...]
```

### All parses with tags

`Tag` and `WordForms` return the best tag and untagged forms. `Parses` returns
every parse of an exact word form (no "е"/"ё" substitution) together with the
whole lexeme, each form carrying its own tag:

```go
for _, p := range a.Parses("стали") {
	fmt.Println(p.Tag, p.Lexeme[0].Word)
	// VERB,perf,intr plur,past,indc стать
	// NOUN,inan,femn sing,gent сталь
	// NOUN,inan,femn sing,datv сталь
	// ...
}
```

## Dictionary

The embedded dictionary is built from the OpenCorpora v0.92 dataset (revision 417127) compiled by pymorphy2 v0.9.1. It contains:

- 5 140 055 word entries
- 3 456 paradigms
- 5 532 grammatical tags

> Please note that embedding the dictionary into the executable increases its size by ~8.8 MB.

## License

The **Go source code** is licensed under the [MIT License](LICENSE).

The **embedded dictionary data** (`data/`) is derived from [OpenCorpora](http://opencorpora.org/)
and is licensed under [CC BY-SA 4.0](data/LICENSE). If you distribute a binary that embeds
this data, you must comply with the CC BY-SA 4.0 terms (attribution + ShareAlike).
