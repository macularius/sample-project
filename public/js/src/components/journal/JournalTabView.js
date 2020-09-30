// возвращает webix конфигурации таба для работы с событиями
export default function JournalTabView() {
    return {
        id: 'journalTab',
        rows: [
            {
                view: 'datatable',
                columns: [
                    { id: 'book', header: 'Книга', fillspace: true, },
                    { id: 'employee', header: 'Сотрудник', fillspace: true, },
                    { id: 'type', header: 'Тип события', fillspace: true, },
                    { id: 'data', header: 'Дата события', fillspace: true, },
                ],
                data: []
            },
        ]
    }
}