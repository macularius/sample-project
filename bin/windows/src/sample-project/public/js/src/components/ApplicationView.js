// возвращает webix конфигурацию рабочего пространства приложения
export default function WorkedPlaceView(bookTab, employeeTab, journalTab, userInfo) {
    return {
        id: 'workedPlace',
        rows: [
            // header
            userInfo.config(),
            {
                id: 'main',
                cols: [
                    {
                        view: 'tabbar',
                        id: 'main-tabbar',
                        value: 'listView',
                        width: 600,
                        multiview: true,
                        options: [
                            { id: 'bookTab', value: 'Книги' },
                            { id: 'employeeTab', value: 'Сотрудники' },
                            { id: 'journalTab', value: 'Журнал событий' },
                        ]

                    },
                    {},
                    {
                        id: 'tab-controlls',
                        rows: [
                            bookTab.configTabControlls(),
                            employeeTab.configTabControlls(),
                            journalTab.configTabControlls(),
                        ]
                    }, // элементы управления табов
                ],
            },
            // содержимое табов
            {
                view: 'multiview',
                id: 'main-views',
                cells: [
                    bookTab.config(),
                    employeeTab.config(),
                    journalTab.config(),
                ]
            },
        ],
    }
}