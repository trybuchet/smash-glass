import ChatMessage from "./ChatMessage";
import User from "../User";

export default class ChatBubble extends ChatMessage {

    /**
     * The ID of the timeout for the message
     */
    exitTimeout: any = null;

    /**
     * The ID of the timeout for the message
     */
    deleteTimeout: any = null;

    /**
     * The bubble is exiting
     */
    exit: boolean = false;

    constructor(data?: Partial<ChatBubble>) {
        super(data);
    }

}