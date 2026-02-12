/**
 * WebSocket message interface
 * @interface WSMessage
 * 
 * @property {string} event - The event type of the message
 * @property {any} data - The data associated with the event
 */
export default class WSData {

    /**
     * The event type of the message
     * @type {string}
     */
    event: string;

    /**
     * The data associated with the event
     * @type {any}
     */
    data: any;

    /**
     * Creates a new WSData instance
     * @param {string} json - The json string to parse
     */
    constructor(json: string) {
        const obj = JSON.parse(json);
        this.event = obj.event;
        this.data = obj.data;
    }
    
}