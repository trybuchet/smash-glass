import OverlayWidget from "../OverlayWidget";
import Position from "@/interfaces/Position";

export default class WebcamWidget extends OverlayWidget {

    name: string = 'webcam';
    
    enabled: boolean = false;

    position: Position = { position: 'bottom-right', x: 0, y: 0, w: 280, h: 180 };

}