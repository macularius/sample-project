import BookWindowView from "./BookWindowView"

// компонент окна для работы с сущностью книги
export class CBookWindow {
    constructor() {
        this.view
        this.isShow
    }

    // метод инициализации компонента
    init() { }

    // метод получения webix конфигурации компонента
    config() {
        return BookWindowView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() { }

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

                break;
            case BOOK_WINDOW_TYPE.update:

                break;
            case BOOK_WINDOW_TYPE.delete:
                
                break;
            default:
                console.error('Неизвестный тип отображения окна для рабоыт с сущностью книги');
                break;
        }
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

