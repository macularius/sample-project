import BookTabView from "./BookTabView"
import { CBookWindow } from "./bookWindow/CBookWindow";

export class CBookTab {
    constructor() {
        this.#view
        this.#window
    }

    // метод инициализации компонента
    init() {
        this.#window = new CBookWindow();
        this.#window.init()
    }

    // метод получения webix конфигурации компонента
    config() {
        BookTabView(this.#window.config())
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {}
    
    // метод вызова модального окна
    #showByID(id) {}
}


