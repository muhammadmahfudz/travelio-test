package domain

type Book struct {
	title     string
	thumbnail string
	author    []string
}

func BookBuilder() *Book {
	return &Book{}
}

func (b *Book) BookBuild() *Book {

	return &Book{
		title:     b.title,
		thumbnail: b.thumbnail,
		author:    b.author,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetThumbnail(thumbnail string) {
	b.thumbnail = thumbnail
}

func (b *Book) SetAuthor(author []string) {
	b.author = author
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetThumbnail() string {
	return b.thumbnail
}

func (b *Book) GetAuthor() []string {
	return b.author
}
