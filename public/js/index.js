import { Application } from "./src/components/Application.js";

webix.ready(() => {
    let app = new Application();
    app.init()
    webix.ui(app.config())
    app.attachEvents()
})