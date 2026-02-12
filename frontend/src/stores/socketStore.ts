import { defineStore } from 'pinia';
import { ref } from 'vue';
import { Application, Events } from '@wailsio/runtime';

import { GetProp } from '../../bindings/SmashGlass/services/configservice';

import WSData from '@/models/WSData';

export const useSocketStore = defineStore('socketStore', () => {
    /* ---------- state ---------- */

    const conn = ref<WebSocket | null>(null);
    const isConnected = ref(false);

    async function init() {
        Events.On('socket:message', (data: any) => {
            const msg = JSON.parse(data.data.data);
            window.$eventBus.emit(msg.event, msg.data);
        });
    }

    /* ---------- actions ---------- */

    async function connect() {
        //window.$helpers.log('Connecting to main app...');

        let port = 9002;
        const prop = await GetProp('Socket', 'port');
        if (prop[0]) {
            console.log(`Socket port: ${prop}`);
            port = parseInt(prop[0]);
        }

        const ws = new WebSocket(`ws://localhost:${port}`);
        conn.value = ws;

        ws.onopen = () => {
            console.log('Websocket connected');
            isConnected.value = true;
        };

        ws.onclose = () => {
            if (isConnected.value) {
                console.log('Websocket disconnected');
                isConnected.value = false;
                Application.Quit();
            }
        };

        ws.onerror = () => {
            console.error('Websocket error');
        };

        ws.onmessage = (event: MessageEvent) => {
            onMessage(event);
        };
    }

    async function disconnect() {
        conn.value?.close();
    }

    async function send(event: string, data: any) {
        conn.value?.send(JSON.stringify({ event: event, data: data }));
    }

    async function onMessage(event: MessageEvent) {
        const data = new WSData(event.data);
        window.$eventBus.emit(data.event, data.data);
    }

    /* ---------- expose ---------- */

    return {
        init,
        conn,
        isConnected,
        connect,
        disconnect,
        send,
        onMessage
    };
});