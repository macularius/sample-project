import UserInfoView from "./UserInfoView"

export class CUserInfo {
    constructor() {
        this.#view
    }

    // метод инициализации компонента
    init() {}

    // метод получения webix конфигурации компонента
    config() {
        UserInfoView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {}
}


