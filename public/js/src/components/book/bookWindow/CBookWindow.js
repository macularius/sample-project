import BookWindowView from "./BookWindowView.js"

// компонент окна для работы с сущностью книги
export class CBookWindow {
    constructor() {
        this.view
        this.isShow = false
    }

    // метод инициализации компонента
    init() { }

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
            form: $$('bookWindowForm')
        }

        // обрабтка закрытия окна
        this.view.windowCancelBtn.attachEvent('onItemClick', () => {
            this.view.window.hide()
        })
    }

    // метод вызова модального окна
    switch(type) {
        switch (this.isShow) {
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
                break;
            case BOOK_WINDOW_TYPE.update:
                this.view.windowLabel.setHTML('Редактирование книги')
                break;
            case BOOK_WINDOW_TYPE.delete:
                this.view.windowLabel.setHTML('Удаление книги')
                break;
            default:
                console.error('Неизвестный тип отображения окна для рабоыт с сущностью книги');
                return;
        }

        this.view.window.show()
    }

    // метод сокрытия окна
    hide() { }
}

// типы отображения модального окна для работы с сущностью книги
export const BOOK_WINDOW_TYPE = {
    create: 'CREATE',
    update: 'UPDATE',
    delete: 'DELETe',
}

