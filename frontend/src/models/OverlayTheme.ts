/**
 * Represents an overlay theme
 */
export default class OverlayTheme {
    /**
     * Unique ID of the theme
     */
    id: string;

    /**
     * Information about the theme
     */
    meta: {
        name: string;
        author: string;
        description: string;
    };

    constructor(data?: Partial<OverlayTheme>) {
        if (!data) return;
        this.id = data.id;
        this.meta = data.meta;
    }
}
