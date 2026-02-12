import type mitt from "mitt";

declare global {
    interface Window {
        $eventBus: mitt;
    }
}