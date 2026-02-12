import Position from "@/interfaces/Position";

export default class AppConfig {
    // Group of overlay configurations
    overlay = {
        theme: 'default',
        opacity: 0.9,
        zoom: .5,
        display: 0,
        widgets: {
            default: {
                "chat": {
                    enabled: true,
                    position: { position: 'top-left', x: 0, y: 0, w: 400, h: 200, autoHeight: true, minHeight: 200 },
                    cfg: {
                        showHistory: true,
                        showTimestamps: true,
                        scrollSpeed: 1
                    }
                },
                "guests": {
                    enabled: true,
                    position: { position: 'top-right', x: 0, y: 0, w: 250, h: 32, autoHeight: true, minHeight: 32 },
                    cfg: {
                        showLatency: true
                    }
                },
                "gamepads": {
                    enabled: true,
                    position: { position: 'bottom-center', x: 0, y: 0, w: 100, h: 100, autoWidth: true, autoHeight: true },
                    cfg: {
                        showHotseatTime: true
                    }
                },
                "webcam": {
                    enabled: false,
                    position: { position: 'bottom-right', x: 0, y: 0, w: 280, h: 180 },
                    cfg: {
                        deviceId: "-1",
                    }
                }
            },
            custom: {},
        },
    };

    headless: {
        name: string,
        widgets: {}[],
    }[] = [
        {
            name: 'Game',
            widgets: []
        },
        {
            name: 'Be Right Back',
            widgets: []
        },
        {
            name: 'Coming Soon',
            widgets: []
        }
    ]

    activeScene: number = 0;

    hotkeys = {
        'hotkey:chat': { modifiers: [17, 18], key: 67, event: 'hotkey:chat' }, // CTRL + ALT + C
        'hotkey:opacity:in': { modifiers: [17, 18], key: 38, event: 'hotkey:opacity:in' }, // CTRL + ALT + UP
        'hotkey:opacity:out': { modifiers: [17, 18], key: 40, event: 'hotkey:opacity:out' },
        'hotkey:zoom:out': { modifiers: [17, 18], key: 37, event: 'hotkey:zoom:out' },
        'hotkey:zoom:in': { modifiers: [17, 18], key: 39, event: 'hotkey:zoom:in' },
        'hotkey:menu': { modifiers: [17, 18], key: 112, event: 'hotkey:menu' },
        'hotkey:move': { modifiers: [17, 18], key: 77, event: 'hotkey:move' },
    };

    constructor(data?: Partial<AppConfig>) {
        if (data) {
            Object.assign(this, data);

            // Ensure display is a number
            if (typeof this.overlay.display === 'string') {
                this.overlay.display = parseInt(this.overlay.display);
            }
        }
    }

    static from(data: Partial<AppConfig>) {
        return new AppConfig(data);
    }

    clone() {
        return new AppConfig(JSON.parse(JSON.stringify(this)));
    }
}