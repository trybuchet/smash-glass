import Position from "@/interfaces/Position";
import OverlayWidget from "../OverlayWidget";

export default class ChatWidget extends OverlayWidget {

    name: string = 'chat';

    enabled: boolean = true;

    active: boolean = false;

    position: Position = { position: 'top-left', x: 0, y: 0, w: 400, h: 200, autoHeight: true, minHeight: 200 };

    cfg = {
        showHistory: true,
        showTimestamps: true,
        scrollSpeed: 1
    };

}