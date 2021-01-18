import { EVENT_CONTEXT_MENU } from './CJournalTab.js';

// возвращает webix конфигурации таба для работы с событиями
export function JournalTabView() {
    return {
        id: 'journalTab',
        rows: [
            {
                view: 'datatable',
                id: 'eventTabDatatable',
                select: true,
                columns: [
                    { id: 'ID', header: ['', { content: 'textFilter' }], hidden: true },
                    { id: 'bookString', header: ['Книга', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'employeeString', header: ['Сотрудник', { content: 'textFilter' }], sort: 'string', fillspace: true, },
                    { id: 'type', header: ['Тип события', { content:'selectFilter' }], sort: 'string', width: 120, },
                    { id: 'date', header: ['Дата события', { content: 'textFilter' }], sort: 'date', width: 160, format: webix.Date.dateToStr("%Y-%m-%d %H:%i") },
                ],
                data: [],
                onContext: {},
            },
        ]
    }
}

// возвращает webix конфигурацию контекстного меню таба
export function EventTabContextMenu() {
    return {
        view: 'contextmenu',
        id: 'eventTabDatatableContextMenu',
        data: [EVENT_CONTEXT_MENU.toEmployee, EVENT_CONTEXT_MENU.toBook],
    }
}

// элементы управления для таба
export function TabControllsView() {
    return {
        id: 'journaltab-controlls',
        hidden: true,
        cols: [
            {
                id: 'journaltab-to-employee-btn',
                view: 'icon',
                tooltip: 'Перейти к пользователю',
                icon: 'user',
                width: 30,
            },
            {
                id: 'journaltab-to-book-btn',
                view: 'icon',
                tooltip: 'Перейти к книге',
                icon: 'book',
                width: 30,
            },
            { width: 30 },
        ]
    }
}