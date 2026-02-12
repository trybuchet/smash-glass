// src/models/OverlayWidget.ts
import Position from "@/interfaces/Position";

let widgetId = 1;

export default class OverlayWidget {

    /**
     * Auto incrementing id
     */
    id: number;

    /**
     * Whether the widget is a custom widget or not
     */
    custom: boolean = false;

    /**
     * Whether the widget is enabled or not
     **/
    enabled: boolean = false;

    /**
     * Whether the widget is active or not
     */
    active: boolean = false;

    /**
     * Name of the widget
     */
    name: string = 'widget';

    /**
     * Position of the widget
     */
    position: Position = { position: 'custom', x: 0, y: 0, w: 200, h: 150 };

    /**
     * Whether the widget can be resized or not
     */
    canResize: boolean = true;

    /**
     * HTML for the widget
     */
    html: string = '';

    /**
     * Custom styles for the widget
     */
    css: string = '';

    /**
     * Various configuration options
     */
    cfg: any = {};

    /**
     * Creates a new OverlayWidget instance.
     */
    constructor(data?: Partial<OverlayWidget>) {
        if (!data) return;
        this.id = widgetId++;
        Object.assign(this, data);
    }

    /**
     * Resolve the position of the widget based on the viewport width and height.
     * 
     * @param viewportW - The width of the viewport.
     * @param viewportH - The height of the viewport.
     */
    resolvePosition(viewportW: number, viewportH: number) {
        let { x, y, w, h, position, fullWidth, fullHeight } = this.position;

        if (fullWidth) w = viewportW;
        if (fullHeight) h = viewportH;

        if (fullWidth || fullHeight) {
            const allowed = new Set(['top-center', 'center', 'bottom-center']);
            if (!allowed.has(position)) position = 'top-center';
        }

        switch (position) {
            case 'top-left':
                x = 0;
                y = 0;
                break;
            case 'top-right':
                x = viewportW - w;
                y = 0;
                break;
            case 'bottom-left':
                x = 0;
                y = viewportH - h;
                break;
            case 'bottom-right':
                x = viewportW - w;
                y = viewportH - h;
                break;
            case 'center':
                x = (viewportW - w) / 2;
                y = (viewportH - h) / 2;
                break;
            case 'top-center':
                x = (viewportW - w) / 2;
                y = 0;
                break;
            case 'bottom-center':
                x = (viewportW - w) / 2;
                y = viewportH - h;
                break;
            case 'custom':
                break;
        }

        return { x, y };
    }

    /**
     * Move the widget to a specific position.
     * 
     * @param x - The x-coordinate of the new position.
     * @param y - The y-coordinate of the new position.
     */
    moveTo(x: number, y: number) {
        this.position.position = 'custom';
        this.position.x = x;
        this.position.y = y;
    }

    /**
     * Resize the widget.
     * 
     * @param w - The new width of the widget.
     * @param h - The new height of the widget.
     */
    resize(w: number, h: number) {
        this.position.w = Math.max(50, w);
        this.position.h = Math.max(50, h);
    }
}
