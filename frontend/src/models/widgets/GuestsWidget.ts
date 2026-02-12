import OverlayWidget from "../OverlayWidget";
import Position from "@/interfaces/Position";

export default class GuestsWidget extends OverlayWidget {

    name: string = 'guests';
    
    enabled: boolean = true;

    position: Position = { position: 'top-right', x: 0, y: 0, w: 250, h: 32, autoHeight: true, minHeight: 32 };

    cfg = {
        showLatency: true
    };

}