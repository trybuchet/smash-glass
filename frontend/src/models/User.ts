export default class User {
    id: string|number = -1;
    name: string = 'JoeBloggs';
    role: string = 'user';
    latency: number = 0;

    constructor(data?: Partial<User>) {
        Object.assign(this, data);
    }

    /**
     * Helper method to truncate the user's name
     * @param length    The maximum length of the name
     * @returns 
     */
    truncateName(length: number): string {
        if (this.name.length <= length) {
            return this.name;
        }

        return this.name.substr(0, length) + '...';
    }
}