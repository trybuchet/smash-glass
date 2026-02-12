import { createConfirmDialog } from 'vuejs-confirm-dialog';
import ConfirmDialog from '@/components/dialogs/ConfirmDialog.vue';
import PromptDialog from '@/components/dialogs/PromptDialog.vue';

/**
 * Asks the user to confirm an action, such as deleting an item.
 * 
 * @param action    The action to perform when the user confirms.
 * @param message   The message to display in the dialog.
 * @param title     The title of the dialog.
 * 
 * @example
 * confirm(() => console.log('Confirm confirmed'))
 */
function confirm(action: Function, message: string = 'Are you sure?', title: string = 'Confirm'): void {
    const { reveal, onConfirm } = createConfirmDialog(ConfirmDialog, {
        question: message,
        title: title,
    })
 
    onConfirm(action as any)
    reveal()
}

/**
 * Asks the user to confirm an action, such as deleting an item.
 * 
 * @param action    The action to perform when the user confirms.
 * @param message   The message to display in the dialog.
 * @param title     The title of the dialog.
 * 
 * @example
 * confirm(() => console.log('Confirm confirmed'))
 */
function prompt(action: Function, message: string = 'Give a value', title: string = 'New'): void {
    const { reveal, onConfirm } = createConfirmDialog(PromptDialog, {
        message: message,
        title: title,
    })
 
    onConfirm(action as any)
    reveal()
}

export {
    confirm,
    prompt
}

declare global {
    interface Window {
        $dialog: {
            confirm: typeof confirm,
            prompt: typeof prompt
        }
    }
}