import OverlayWidget from "../OverlayWidget";
import Position from "@/interfaces/Position";

export default class GamepadsWidget extends OverlayWidget {

    name: string = 'gamepads';

    enabled: boolean = true;

    position: Position = { position: 'bottom-center', x: 0, y: 0, w: 100, h: 100, minWidth: 100, minHeight: 100, autoWidth: true, autoHeight: true };

    canResize: boolean = false;

    cfg = {
        showHotseatTime: true
    };

}