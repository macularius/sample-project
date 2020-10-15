import UserInfoView from "./UserInfoView.js"
import { getCookie } from "../../../helpers/cookies.js"
import authModel from "../../models/authModel.js"

// класс компонента информации о пользователе
export class CUserInfo {
    constructor() {
        this.view                           // быстрый доступ к представлениям компонента
        this.currentEmployee = undefined    // сотрудник, соответствующий текущему пользователю
    }

    // метод инициализации компонента
    init() { }

    // метод получения webix конфигурации компонента
    config() {
        // проверяем наличие сотрудника текущего пользователя в куках
        if (!(this.currentEmployee = getCookie('current_employee'))) {
            // отложенное обновление информации о пользователе
            authModel.getCurrentEmployee().then((emp) => {
                // проверка наличия данных
                if (!emp) {
                    return
                }

                this.currentEmployee = emp
                this.refreshEmployeeLabel(emp)
            })
        }

        // если удалось получить текущего сотрудника, то отображаем его данные, иначе "ждем загрузку данных"
        return this.currentEmployee ? UserInfoView(`${employee.lastname} ${employee.firstname}`) : UserInfoView('загрузка...')
    }

    // функция обновления информации о текущем пользователе
    refreshEmployeeLabel(employee) {
        this.view.userLabel.setValue(`${employee.lastname} ${employee.firstname}`)
    }

    // метод инициализации обработчиков событий компонента
    attachEvents() {
        // инициализация используемых представлений
        this.view = {
            userLabel: $$('userInfoLabel'),
            logoutBtn: $$('logoutBtn'),
        }

        // выход
        this.view.logoutBtn.attachEvent('onItemClick', () => {
            authModel.logout().then(() => { document.location.replace('/') })
        })
    }
}


