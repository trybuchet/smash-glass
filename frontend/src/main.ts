import { createApp } from 'vue'
import { createPinia } from 'pinia';
import mitt from 'mitt';
import '@/styles/css/legacy.css';

import { MoveMainWindowToMonitor } from '@bindings/windowservice';

import App from './components/App.vue'
import * as ConfirmDialog from 'vuejs-confirm-dialog';
import form from './components/form';
import * as dialog from './utils/dialog';

import { useConfigStore } from './stores/configStore';
import { useOverlayStore } from './stores/overlayStore';
import { useSocketStore } from './stores/socketStore';

async function initApp() {

    // Hook eventbus
    window.$eventBus = mitt();

    const app = createApp(App)
    app.use(createPinia())
    app.use(ConfirmDialog)

    // Register form components
    for (const key in form) {
        app.component(key, form[key]);
    }

    const configStore = useConfigStore()
    const overlayStore = useOverlayStore()
    const socketStore = useSocketStore()

    await configStore.init();
    await socketStore.init();
    await overlayStore.init();

    // Dialog windows
    window.$dialog = dialog;

    // Move to the default monitor
    MoveMainWindowToMonitor(configStore.app.overlay.display);

    app.mount('#smashglass')
}

initApp()
