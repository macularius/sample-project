// возвращает webix конфигурацию рабочего пространства приложения
export default function WorkedPlaceView(bookTab, employeeTab, journalTab, userInfo) {
    return {
        id: 'workedPlace',
        rows: [
            // header
            userInfo.config(),
            {
                cols: [
                    {
                        view: 'tabbar',
                        id: 'main-tabbar',
                        value: 'listView',
                        multiview: true,
                        options: [
                            { id: 'bookTab', value: 'Книги' },
                            { id: 'employeeTab', value: 'Сотрудники' },
                            { id: 'journalTab', value: 'Журнал событий' },
                        ]

                    },
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