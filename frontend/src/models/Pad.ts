import User from "./User";

/**
 * Represents a gamepad (controller) input device.
 */
export default class Pad {

    /** Index of the gamepad */
    index = -1;

    /** Whether the gamepad is currently locked */
    locked = false;

    /** Axes (standardised to 8) */
    axes: number[] = Array(8).fill(0);

    /** Buttons (standardised to 17) */
    buttons: Array<{
        pressed: boolean;
        touched: boolean;
        value: number;
    }> = Array(17).fill(null).map(() => ({
        pressed: false,
        touched: false,
        value: 0,
    }));

    /** Owner of the gamepad */
    owner: {
        id: string;
        name: string;
        hotseatTime?: number;
    } = null;

    constructor(data?: Partial<Pad>) {
        if (!data) return;
        Object.assign(this, data);

        if (Array.isArray(data.axes)) {
            for (let i = 0; i < this.axes.length; i++) {
                this.axes[i] = data.axes[i] ?? 0;
            }
        }

        if (Array.isArray(data.buttons)) {
            for (let i = 0; i < this.buttons.length; i++) {
                const b = data.buttons[i];
                if (!b) continue;
                this.buttons[i].pressed = !!b.pressed;
                this.buttons[i].touched = !!b.touched;
                this.buttons[i].value = b.value ?? 0;
            }
        }
    }
}