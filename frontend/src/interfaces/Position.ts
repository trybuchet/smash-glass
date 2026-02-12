export default interface Position {
    position: 'top-left' | 'top-right' | 'bottom-left' | 'bottom-right' | 'top-center' | 'bottom-center' | 'center' | 'custom';
    autoWidth?: boolean;
    autoHeight?: boolean;
    minWidth?: number;
    minHeight?: number;
    fullWidth?: boolean;
    fullHeight?: boolean;
    x: number;
    y: number;
    w: number;
    h: number;
}