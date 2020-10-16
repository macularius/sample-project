import BookWindowView from "./BookWindowView.js"
import bookModel from "./../../../models/bookModel.js"

// компонент окна для работы с сущностью книги
export class CBookWindow {
    constructor() {
        this.view       // объект для быстрого доступа к представлениям
        this.type       // тип текущего отображения окна
        this.onChange   // callback функция при CUD операциях над книгой
    }

    // метод инициализации компонента
    init(onChange) {
        this.onChange = onChange // callback функция при CUD операциях над книгой
    }

    // метод получения webix конфигурации компонента
    config() {
        return BookWindowView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            window: $$('bookWindow'),
            windowLabel: $$('bookWindowLabel'),
            windowCancelBtn: $$('bookWindowCancelBtn'),
            windowConfirmBtn: $$('bookWindowConfirmBtn'),
            form: $$('bookWindowForm'),
            formfields: {
                ISBN: $$('bookWindowFormISBN'),
                name: $$('bookWindowFormName'),
                author: $$('bookWindowFormAuthor'),
                publisher: $$('bookWindowFormPublisher'),
                status: $$('bookWindowFormStatus'),
                year: $$('bookWindowFormYear'),
            }
        }

        // обработка закрытия окна
        this.view.windowCancelBtn.attachEvent('onItemClick', () => {
            this.view.window.hide()
        })

        // обработка события "принять"
        this.view.windowConfirmBtn.attachEvent('onItemClick', () => {
            // валидация введенных данных по обязательным полям
            if (!this.validate()) {
                webix.message('Заполните поля отмеченные *', 'error')
                return;
            }

            switch (this.type) {
                case BOOK_WINDOW_TYPE.create:
                    bookModel.createBook(this.fetch()).then(() => {
                        this.onChange()
                        this.hide()
                    })
                    break;
                case BOOK_WINDOW_TYPE.update:
                    bookModel.updateBook(this.fetch()).then(() => {
                        this.onChange()
                        this.hide()
                    })
                    break;
                case BOOK_WINDOW_TYPE.delete:
                    bookModel.deleteBook(this.fetch()).then(() => {
                        this.onChange()
                        this.hide()
                    })
                    break;
            }
        })
    }

    // метод вызова модального окна
    switch(type) {
        switch (this.view.window.isVisible()) {
            case true:
                this.hide()
                break;
            case false:
                this.show(type)
                break;
        }
    }

    // метод отображения окна
    show(type) {
        switch (type) {
            case BOOK_WINDOW_TYPE.create:
                this.view.windowLabel.setHTML('Добавление книги')
                this.view.formfields.status.hide()
                this.view.window.resize()
                break;
            case BOOK_WINDOW_TYPE.update:
                this.view.windowLabel.setHTML('Редактирование книги')
                break;
            case BOOK_WINDOW_TYPE.delete:
                this.view.windowLabel.setHTML('Удаление книги')
                this.view.formfields.ISBN.disable()
                this.view.formfields.name.disable()
                this.view.formfields.author.disable()
                this.view.formfields.publisher.disable()
                this.view.formfields.status.disable()
                this.view.formfields.year.disable()
                this.view.window.resize()
                break;
            default:
                console.error('Неизвестный тип отображения окна для работы с сущностью книги');
                return;
        }

        this.type = type
        this.view.window.show()
    }

    // метод сокрытия окна
    hide() {
        this.view.window.hide()
    }

    // метод получения сущности из формы окна
    fetch() {
        return this.view.form.getValues()
    }

    // метод размещения сущности в форме окна
    parse(values) {
        this.view.form.setValues(values)
    }

    // функция валидации формы
    validate() {
        let isValid = false

        // удаление пробелов в полях формы
        this.view.formfields.ISBN.setValue(this.view.formfields.ISBN.getValue().trim())
        this.view.formfields.name.setValue(this.view.formfields.name.getValue().trim())
        this.view.formfields.author.setValue(this.view.formfields.author.getValue().trim())
        this.view.formfields.publisher.setValue(this.view.formfields.publisher.getValue().trim())
        this.view.formfields.year.setValue(this.view.formfields.year.getValue().trim())

        // валидация webix
        isValid = this.view.form.validate()

        return isValid
    }
}

// типы отображения модального окна для работы с сущностью книги
export const BOOK_WINDOW_TYPE = {
    create: 'CREATE',
    update: 'UPDATE',
    delete: 'DELETe',
}

