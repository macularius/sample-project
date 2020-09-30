import { CUserInfo } from './src/components/userInfo/CUserInfo.js';
import { CBookTab } from './src/components/book/CBookTab.js';
import { CEmployeeTab } from './src/components/employee/CEmployeeTab.js';
import { CJournalTab } from './src/components/journal/CJournalTab.js';

export class Application {
    constructor() {
        this.userInfo = new CUserInfo();
        this.bookTab = new CBookTab();
        this.employeeTab = new CEmployeeTab();
        this.journalTab = new CJournalTab();
    }

    init() {
        this.userInfo.init()
        this.bookTab.init()
        this.employeeTab.init()
        this.journalTab.init()
    }

    run() {
        webix.ui(this.config())
    }

    config() {
        return {
            rows: [
                // header
                {
                    cols: [
                        {
                            view: 'tabbar',

                        },
                        {
                            cols: [
                                // фио
                                {
                                    view: 'label',
                                    template: 'Фамилия Имя Отчество',
                                },
                                // кнопка выхода
                                {
                                    view: 'button',
                                    text: 'Выход',
                                },
                            ]
                        },
                    ],
                },
                // содержимое табов
                {

                },
            ],
        }
    }
}

webix.ready(()=>{
    let app = new Application();
    app.init()
    app.run()
})