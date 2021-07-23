package service

import (
	"github.com/Screen17/catalog/appcontext"
	"github.com/Screen17/catalog/model"
	"github.com/Screen17/catalog/model/dto"
	"github.com/Screen17/catalog/repository"
	"github.com/Screen17/catalog/util"
)

// BookService is a service for managing books.
type BookService struct {
	context appcontext.Context
}

// NewBookService is constructor.
func NewBookService(context appcontext.Context) *BookService {
	return &BookService{context: context}
}

// FindByID returns one record matched book's id.
func (b *BookService) FindByID(id string) *model.Book {
	if !util.IsNumeric(id) {
		return nil
	}

	rep := b.context.GetRepository()
	book := model.Book{}
	result, err := book.FindByID(rep, util.ConvertToUint(id))
	if err != nil {
		b.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// FindAllBooks returns the list of all books.
func (b *BookService) FindAllBooks() *[]model.Book {
	rep := b.context.GetRepository()
	book := model.Book{}
	result, err := book.FindAll(rep)
	if err != nil {
		b.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// FindAllBooksByPage returns the page object of all books.
func (b *BookService) FindAllBooksByPage(page string, size string) *model.Page {
	rep := b.context.GetRepository()
	book := model.Book{}
	result, err := book.FindAllByPage(rep, page, size)
	if err != nil {
		b.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// FindBooksByTitle returns the page object of books matched given book title.
func (b *BookService) FindBooksByTitle(title string, page string, size string) *model.Page {
	rep := b.context.GetRepository()
	book := model.Book{}
	result, err := book.FindByTitle(rep, title, page, size)
	if err != nil {
		b.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}

// CreateBook register the given book data.
func (b *BookService) CreateBook(dto *dto.BookDto) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := b.context.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep repository.Repository) error {
			var err error
			book := dto.Create()

			category := model.Category{}
			if book.Category, err = category.FindByID(txrep, dto.CategoryID); err != nil {
				return err
			}

			format := model.Format{}
			if book.Format, err = format.FindByID(txrep, dto.FormatID); err != nil {
				return err
			}

			if result, err = book.Create(txrep); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			b.context.GetLogger().GetZapLogger().Errorf(err.Error())
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}

// UpdateBook updates the given book data.
func (b *BookService) UpdateBook(dto *dto.BookDto, id string) (*model.Book, map[string]string) {
	errors := dto.Validate()

	if errors == nil {
		rep := b.context.GetRepository()
		var result *model.Book

		err := rep.Transaction(func(txrep repository.Repository) error {
			var err error
			var book *model.Book

			b := model.Book{}
			if book, err = b.FindByID(txrep, util.ConvertToUint(id)); err != nil {
				return err
			}

			book.Title = dto.Title
			book.Isbn = dto.Isbn
			book.CategoryID = dto.CategoryID
			book.FormatID = dto.FormatID

			category := model.Category{}
			if book.Category, err = category.FindByID(txrep, dto.CategoryID); err != nil {
				return err
			}

			format := model.Format{}
			if book.Format, err = format.FindByID(txrep, dto.FormatID); err != nil {
				return err
			}

			if result, err = book.Update(txrep); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			b.context.GetLogger().GetZapLogger().Errorf(err.Error())
			return nil, map[string]string{"error": "transaction error"}
		}

		return result, nil
	}

	return nil, errors
}

// DeleteBook deletes the given book data.
func (b *BookService) DeleteBook(id string) (*model.Book, map[string]string) {
	rep := b.context.GetRepository()
	var result *model.Book

	err := rep.Transaction(func(txrep repository.Repository) error {
		var err error
		var book *model.Book

		b := model.Book{}
		if book, err = b.FindByID(txrep, util.ConvertToUint(id)); err != nil {
			return err
		}

		if result, err = book.Delete(txrep); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		b.context.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, map[string]string{"error": "transaction error"}
	}

	return result, nil
}
