import { BOOK_CONTEXT_MENU } from "./CBookTab.js";

// возвращает webix конфигурации таба для работы с книгами
export function BookTabView() {
    return {
        id: 'bookTab',
        rows: [
            {
                view: 'datatable',
                id: 'bookTabDatatable',
                select: true,
                columns: [
                    { id: 'ID', hidden: true },
                    { id: 'isbn', header: 'ISBN', fillspace: true, },
                    { id: 'name', header: 'Название', fillspace: true, },
                    { id: 'author', header: 'Автор', fillspace: true, },
                    { id: 'publisher', header: 'Издательство', fillspace: true, },
                    { id: 'year', header: 'Год издания', fillspace: true, },
                    { id: 'status', header: 'Статус', fillspace: true, },
                ],
                data: [],
                onContext: {},
            },
        ]
    }
}

// возвращает webix конфигурации контекстного меню таба
export function BookTabContextMenu() {
    return {
        view: 'contextmenu',
        id: 'bookTabDatatableContextMenu',
        data: [BOOK_CONTEXT_MENU.toEvent, BOOK_CONTEXT_MENU.toEmployee, BOOK_CONTEXT_MENU.add, BOOK_CONTEXT_MENU.edit, BOOK_CONTEXT_MENU.remove],
    }
}