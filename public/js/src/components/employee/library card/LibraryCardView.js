import { LIBRARY_CARD_CONTEXT_MENU } from "./ClibraryCard.js"

// возвращает webix конфигурации таба для работы с событиями
export function LibraryCardView() {
    return {
        view: 'window',
        id: 'libraryCard',
        head: {
            view: 'template',
            id: 'libraryCardLabel',
            template: 'Читательский билет',
            css: 'webix_template'
        },
        fullscreen: true,
        body: {
            rows: [
                {
                    view: 'datatable',
                    id: 'libraryCardDatatable',
                    select: true,
                    columns: [
                        { id: 'ID', header: ['', { content: "textFilter" }], hidden: true },
                        { id: 'ISBN', header: ['ISBN', { content: "textFilter" }], width: 140, },
                        { id: 'name', header: ['Название', { content: "textFilter" }], fillspace: true, },
                    ],
                    data: [],
                    onContext: {
                        // обработка вызова контекстного меню при остутствии данных
                        webix_view: function (e, id) {
                            id = this.locate(e.target || e.srcElement);
                            if (!id) {
                                $$("libraryCardDatatableContextMenu").setContext({ obj: webix.$$(e) });
                                $$("libraryCardDatatableContextMenu").show(e);
                                webix.html.preventEvent(e);
                            }
                        }
                    },
                },
                {
                    view: 'button',
                    id: 'libraryCardwCancelBtn',
                    value: 'Закрыть',
                },
                {
                    height: 10,
                }
            ]
        }
    }
}

// возвращает webix конфигурацию контекстного меню
export function LibraryCardContextMenu() {
    return {
        view: 'contextmenu',
        id: 'libraryCardDatatableContextMenu',
        data: [
            LIBRARY_CARD_CONTEXT_MENU.take,
            LIBRARY_CARD_CONTEXT_MENU.toBook,
            LIBRARY_CARD_CONTEXT_MENU.toEvent
        ],
    }
}