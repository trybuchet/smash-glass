import User from "../User";

export default class ChatMessage {
    /**
     * The unique identifier for the message.
     */
    id: number;

    type: 'chat'|'log' = 'chat';

    /**
     * The content of the message.
     */
    content: string;

    /**
     * The timestamp when the message was created.
     */
    createdAt: Date = new Date();

    /**
     * The formatted createdAt date.
     */
    createdAtFormatted: string;

    /**
     * The user who sent the message.
     */
    user: User = new User();

    /**
     * Creates a new instance of ChatMessage.
     * @param id - The unique identifier for the message.
     * @param content - The content of the message.
     * @param createdAt - The timestamp when the message was created.
     * @param user - The user who sent the message.
     * @returns A new instance of ChatMessage.
     * @example
     * const message = new ChatMessage('1', 'Hello, world!', new Date(), new User('1', 'John Doe'));
     */
    constructor(data?: Partial<ChatMessage>) {
        if (!data) return;
        Object.assign(this, data);

        this.createdAt = new Date();
        this.createdAtFormatted = this.formatCreatedAt();
    }

    /**
     * Formats the createdAt date to a string.
     * @returns A string representing the formatted createdAt date.
     */
    formatCreatedAt(): string {
        if (!this.createdAt) return "";
        const options: Intl.DateTimeFormatOptions = {
            hour: '2-digit',
            minute: '2-digit',
            hour12: false,
        };
        return this.createdAt.toLocaleTimeString([], options);
    }
}