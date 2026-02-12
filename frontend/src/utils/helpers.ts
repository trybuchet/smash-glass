/**
 * Converts a date to a relative human-readable format
 * 
 * @param date - The date to convert
 */
function formatRelativeDate(date: string|number|Date) {
    const today = new Date();
            
    const dateObj = new Date(date);
    const day = dateObj.getDate();
    const month = dateObj.toLocaleString('default', { month: 'short' });
    const year = dateObj.getFullYear();

    let hour = dateObj.getHours();
    const minute = String(dateObj.getMinutes()).padStart(2, '0'); // add leading zero if needed
    const ampm = hour < 12 ? 'AM' : 'PM';

    hour = hour % 12;
    hour = hour ? hour : 12; // the hour '0' should be '12'

    // If today, return as "23 minutes ago" or "2 hours ago"
    if (dateObj.getDate() === today.getDate()) {
        const diff = today.getTime() - dateObj.getTime();
        const minutes = Math.floor(diff / 60000);
        const hours = Math.floor(minutes / 60);
        if (minutes < 1) {
            return `Just now`;
        } else
        if (minutes < 60) {
            return `${minutes} minute${ (minutes > 1 ? 's' : '')} ago`;
        } else {
            return `${hours} hour${ (hours > 1 ? 's' : '')} ago`;
        }
    }

    // Else if yesterday, return as "Yesterday at 12:00 PM"
    if (dateObj.getDate() === today.getDate() - 1) {
        return `Yesterday at ${hour}:${minute} ${ampm}`;
    }

    // Else return as "Jan 1, 2021 at 12:00 PM"
    return `${month} ${day}, ${year} at ${hour}:${minute} ${ampm}`;
};

/**
 * Debounces a function
 * 
 * @param func - The function to debounce
 * @param waitFor - The time to wait before calling the function
 */
function debounce<F extends (...args: any[]) => void>(func: F, waitFor: number) {
    let timeout: ReturnType<typeof setTimeout> | null = null;

    const debounced = (...args: Parameters<F>) => {
        if (timeout !== null) {
            clearTimeout(timeout);
            timeout = null;
        }
        timeout = setTimeout(() => func(...args), waitFor);
    };

    return debounced;
}

/**
 * Just a simple formatted log function. Maybe I might add
 * logging to file/database in the future.
 * 
 * @param message - The message to log
 */
function log(message: string) {
    console.log(`[Smash Soda] [${new Date().toLocaleTimeString()}] ${message}`);
}

function getKeyName(key: number) {
    const keyMap: { [key: number]: string } = {
        2: 'Alt',
        4: 'Control',
        8: 'Backspace',
        9: 'Tab',
        13: 'Enter',
        16: 'Shift',
        17: 'Control',
        18: 'Alt',
        19: 'Pause',
        20: 'CapsLock',
        27: 'Escape',
        32: 'Space',
        33: 'PageUp',
        34: 'PageDown',
        35: 'End',
        36: 'Home',
        37: 'ArrowLeft',
        38: 'ArrowUp',
        39: 'ArrowRight',
        40: 'ArrowDown',
        45: 'Insert',
        46: 'Delete',
        48: '0',
        49: '1',
        50: '2',
        51: '3',
        52: '4',
        53: '5',
        54: '6',
        55: '7',
        56: '8',
        57: '9',
        65: 'A',
        66: 'B',
        67: 'C',
        68: 'D',
        69: 'E',
        70: 'F',
        71: 'G',
        72: 'H',
        73: 'I',
        74: 'J',
        75: 'K',
        76: 'L',
        77: 'M',
        78: 'N',
        79: 'O',
        80: 'P',
        81: 'Q',
        82: 'R',
        83: 'S',
        84: 'T',
        85: 'U',
        86: 'V',
        87: 'W',
        88: 'X',
        89: 'Y',
        90: 'Z',
        91: 'Meta',
        93: 'ContextMenu',
        96: 'Numpad0',
        97: 'Numpad1',
        98: 'Numpad2',
        99: 'Numpad3',
        100: 'Numpad4',
        101: 'Numpad5',
        102: 'Numpad6',
        103: 'Numpad7',
        104: 'Numpad8',
        105: 'Numpad9',
        106: 'NumpadMultiply',
        107: 'NumpadAdd',
        109: 'NumpadSubtract',
        110: 'NumpadDecimal',
        111: 'NumpadDivide',
        112: 'F1',
        113: 'F2',
        114: 'F3',
        115: 'F4',
        116: 'F5',
        117: 'F6',
        118: 'F7',
        119: 'F8',
        120: 'F9',
        121: 'F10',
        122: 'F11',
        123: 'F12',
        144: 'NumLock',
        145: 'ScrollLock',
        186: 'Semicolon',
        187: 'Equal',
        188: 'Comma',
        189: 'Minus',
        190: 'Period',
        191: 'Slash',
        192: 'Backquote',
        219: 'BracketLeft',
        220: 'Backslash',
        221: 'BracketRight',
        222: 'Quote'
    };

    return keyMap[key] || `Unknown Key (${key})`
}

export {
    formatRelativeDate,
    debounce,
    log,
    getKeyName
};