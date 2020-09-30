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
        $$('main-tabbar').setValue('bookTab')
    }

    attachEvents() {
        this.userInfo.attachEvents()
        this.bookTab.attachEvents()
        this.employeeTab.attachEvents()
        this.journalTab.attachEvents()
    }

    config() {
        return {
            rows: [
                // header
                {
                    cols: [
                        {
                            view: 'tabbar',
                            id: 'main-tabbar',
                            value: 'listView',
                            width: 700,
                            multiview: true,
                            options: [
                                { id: 'bookTab', value: 'Книги' },
                                { id: 'employeeTab', value: 'Сотрудники' },
                                { id: 'journalTab', value: 'Журнал событий' },
                            ]

                        },
                        {},
                        this.userInfo.config(),
                    ],
                },
                // содержимое табов
                {
                    view: "multiview",
                    id: "main-views", 
                    cells: [
                        this.bookTab.config(),
                        this.employeeTab.config(),
                        this.journalTab.config(),
                    ]
                },
            ],
        }
    }
}

webix.ready(() => {
    let app = new Application();
    app.init()
    app.run()
    app.attachEvents()
})