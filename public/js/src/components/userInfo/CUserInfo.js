import UserInfoView from "./UserInfoView.js"

export class CUserInfo {
    constructor() {
        this.view
    }

    // метод инициализации компонента
    init() {}

    // метод получения webix конфигурации компонента
    config() {
        return UserInfoView()
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {}
}


