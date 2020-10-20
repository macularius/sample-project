import UserInfoView from './ToolbarView.js'
import authModel from '../../models/authModel.js'

// класс компонента информации о пользователе
export class CToolbar {
    constructor() {
        this.view                           // быстрый доступ к представлениям компонента
        this.currentEmployee = undefined    // сотрудник, соответствующий текущему пользователю
        this.onLogout                       // callback функция при логауте пользователя
    }

    // метод инициализации компонента
    init(onLogout) { 
        this.onLogout = onLogout
    }

    // метод получения webix конфигурации компонента
    config() {
        // если удалось получить текущего сотрудника, то отображаем его данные, иначе 'ждем загрузку данных'
        return UserInfoView()
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

        // отложенное обновление информации о пользователе
        authModel.getCurrentEmployee().then((emp) => {
            // проверка наличия данных
            if (!emp) {
                return
            }

            this.currentEmployee = emp
            this.refreshEmployeeLabel(emp)
        })

        // выход
        this.view.logoutBtn.attachEvent('onItemClick', () => {
            this.onLogout()
        })
    }
}


