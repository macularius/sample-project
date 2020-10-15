import { LibraryCardContextMenu, LibraryCardView } from "./LibraryCardView.js"
import bookModel from "../../../models/bookModel.js"

// класс окна "Читательский билет"
export class CLibraryCard {
    constructor() {
        this.view       // объект для быстрого доступа к представлениям
    }

    // метод инициализации компонента
    init(toBook, toEvent) {
        this.toBook = toBook // функция перехода к книге, возвращает ссылку на CBookTab
        this.toEvent = toEvent // функция перехода к событию, возвращает ссылку на CJournalTab 
    }

    // метод получения webix конфигурации компонента
    config() {
        // т.к. window и popup расположены не в дереве приложения, а поверх слоев, его нужно отрисовывать отдельно
        webix.ui(LibraryCardContextMenu())
        
        // вызов функции представления
        return LibraryCardView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            window: $$('libraryCard'),
            windowLabel: $$('libraryCardLabel'),
            datatable: $$('libraryCardDatatable'),
            windowCancelBtn: $$('libraryCardwCancelBtn'),
            datatableContextMenu: $$('libraryCardDatatableContextMenu')
        }

        // прикрепление контекстного меню к таблице
        this.view.datatableContextMenu.attachTo(this.view.datatable)

        // загрузка первичных данных в таблицу
        this.refreshTable()

        // обработка закрытия окна
        this.view.windowCancelBtn.attachEvent('onItemClick', () => {
            this.view.window.hide()
        })

        // обработка события нажатия на пункт контекстного меню
        this.view.datatableContextMenu.attachEvent('onMenuItemClick', (itemID) => {
            // проверка вложенности выбранного пункта меню
            if (this.view.datatableContextMenu.getItem(itemID)) {
                // получение значения пункта, на которое произошло нажатие    
                let item = this.view.datatableContextMenu.getItem(itemID).value
                this.handleContextMenu(item)
            }

        })
    }

    // обработка выбора в контекстном меню
    handleContextMenu(item) {
        // получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem()

        switch (item) {
            case LIBRARY_CARD_CONTEXT_MENU.toBook: // переход к книге
                if (!selected) {
                    webix.message('Выделите строку')
                    return
                }
                if (!selected.ID) {
                    console.error('Incorrect ID of item:', selected.ID)
                    return
                }
                let cBookTab = this.toBook()
                cBookTab.showByBookID(selected.ID)
                this.hide()
                break;
            case LIBRARY_CARD_CONTEXT_MENU.toEvent: // переход к событию
                if (!selected) {
                    webix.message('Выделите строку')
                    return
                }
                if (!selected.ID) {
                    console.error('Incorrect ID of item:', selected.ID)
                    return
                }
                let cJournalTab = this.toEvent()
                cJournalTab.showByBookID(selected.ID)
                this.hide()
                break;
            default:
                console.error(`Неизвестное значение пункта меню: ${item}.`)
                break
        }
    }

    // функция обновления таблицы книг
    refreshTable(books) {
        if (books) {
            this.view.datatable.clearAll()
            this.view.datatable.parse(books)
            return
        } else {
            bookModel.getBooks().then((books) => {
                // проверка наличия данных
                if (books) {
                    // преобразование даты издания
                    books.map((book) => {
                        let time = new Date(book.year)
                        book.year = time.getFullYear()
                    })
                }

                // заполнение таблицы окна данными книги
                this.view.datatable.clearAll()
                this.view.datatable.parse(books)
            })
        }
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
    show() {
        this.view.window.show()
    }

    // метод сокрытия окна
    hide() {
        this.view.window.hide()
    }
}

// допустимые значения пунктов контекстного меню читательского билета
export const LIBRARY_CARD_CONTEXT_MENU = {
    toBook: 'К книге',
    toEvent: 'К событию',
    take: 'Сдать книгу',
}