<script lang="ts" setup>
import { computed } from 'vue';
import Pad from '@/models/Pad';

const props = withDefaults(defineProps<{
    pad?: Pad | null
}>(), {
    pad: null
});

// Use a blank pad as fallback if none provided
const displayPad = computed(() => props.pad ?? new Pad());

// Axis movement calculation
const setAxis = (axis: number) => {
    // Determine which deadzone index to use: 0 for left stick, 1 for right stick
    const dzIndex = axis < 2 ? 0 : 1;
    const x = displayPad.value.axes[axis];
    const y = displayPad.value.axes[axis + 1];
    //const dz = window.$pads.preset().dz[dzIndex];
    const dz = 0.2; // Hardcoded deadzone for now
    return `transform: translate(${x * 16}px, ${-y * 16}px);`;
};

const getTriggerStyle = (axisIndex: number) => {
    const axisValue = Math.max(0, displayPad.value.axes[axisIndex] || 0);
    const t = Math.min(1, axisValue);

    if (t === 0) return {}; // <-- no inline fill, CSS controls it

    const start = [255, 255, 255]; // starting colour for blend
    const end = [255, 105, 180];   // pressed colour

    const r = Math.round(start[0] + (end[0] - start[0]) * t);
    const g = Math.round(start[1] + (end[1] - start[1]) * t);
    const b = Math.round(start[2] + (end[2] - start[2]) * t);

    return { fill: `rgb(${r}, ${g}, ${b})` };
};

const DEADZONE = 0.2;

const stickActive = (axis: number) => {
    const x = displayPad.value.axes[axis] ?? 0;
    const y = displayPad.value.axes[axis + 1] ?? 0;

    return Math.abs(x) > DEADZONE || Math.abs(y) > DEADZONE;
};
</script>
    
<template>
    <div :class="{locked: displayPad.locked}">
        <svg 
        version="1.1" id="Layer_1" 
        xmlns="http://www.w3.org/2000/svg" 
        xmlns:xlink="http://www.w3.org/1999/xlink" 
        x="0px" 
        y="0px"
        viewBox="0 0 580 580" 
        style="enable-background:new 0 0 580 580;" 
        class="gamepad"
        xml:space="preserve">

            <!-- Background -->
            <path class="pad-bg" d="M564.6,378.4c0,0-42.5-143.5-58.8-153.9c-6.3-4-8.4-7.4-9.6-10.8c-4,7.2-11.4,10.8-18.5,8L392.9,189
                c-5.6-2.2-9.6-7.7-11.1-14.4c-10.1,6.4-19.6,7.8-19.6,7.8h-72.1H290h-72.1c0,0-11-1.6-22-9.4c-1.2,7.3-5.4,13.6-11.4,15.9
                l-84.8,32.7c-6.2,2.4-12.6,0-16.8-5.6c-1.4,2.6-3.7,5.3-8.5,8.3C58,234.9,15.5,378.4,15.5,378.4S-40,538.2,59,557.5
                c0,0,24.2-15.3,45-40.1c20.8-24.7,61.4-59.9,83.1-60.4c21.3-0.5,99.4,0,102.8,0c0,0,0,0,0.1,0h0.1c3.4,0,81.5-0.5,102.8,0
                c21.8,0.5,62.3,35.6,83.1,60.4c20.8,24.7,45,40.1,45,40.1C620,538.2,564.6,378.4,564.6,378.4z M443.4,225.5
                c17.4,0,31.5,14.1,31.5,31.5c0,17.4-14.1,31.5-31.5,31.5s-31.5-14.1-31.5-31.5C411.8,239.6,426,225.5,443.4,225.5z M395.1,273.7
                c17.4,0,31.5,14.1,31.5,31.5s-14.1,31.6-31.5,31.6c-17.4,0-31.5-14.1-31.5-31.5S377.7,273.7,395.1,273.7z M349.8,288.2
                c0,16.5-13.4,29.9-29.9,29.9S290,304.7,290,288.2s13.4-29.9,29.9-29.9S349.8,271.7,349.8,288.2z M281.8,199.7
                c12.6,0,22.9,10.3,22.9,22.9c0,12.6-10.3,22.9-22.9,22.9c-12.6,0-22.9-10.3-22.9-22.9S269.2,199.7,281.8,199.7z M81.3,274.1
                c0-30.6,24.9-55.6,55.6-55.6c30.6,0,55.6,24.9,55.6,55.6c0,30.6-24.9,55.6-55.6,55.6C106.3,329.7,81.3,304.7,81.3,274.1z
                M265.6,401.5c0,2.4-1.9,4.4-4.4,4.4h-31.6v31.7c0,2.4-1.9,4.4-4.4,4.4h-34.1c-2.4,0-4.4-1.9-4.4-4.4v-31.7h-31.6
                c-2.4,0-4.4-1.9-4.4-4.4v-34.1c0-2.4,1.9-4.4,4.4-4.4h31.6v-31.7c0-2.4,1.9-4.4,4.4-4.4h34.1c2.4,0,4.4,1.9,4.4,4.4V363h31.6
                c2.4,0,4.4,1.9,4.4,4.4V401.5z M240.5,318.1c-16.5,0-29.9-13.4-29.9-29.9s13.4-29.9,29.9-29.9s29.9,13.4,29.9,29.9
                C270.5,304.7,257,318.1,240.5,318.1z M338.9,445.5c-30.3,0-54.9-24.6-54.9-54.9s24.6-54.9,54.9-54.9s54.9,24.6,54.9,54.9
                C393.8,420.8,369.2,445.5,338.9,445.5z M443.4,387.8c-17.4,0-31.5-14.1-31.5-31.5s14.1-31.5,31.5-31.5s31.5,14.1,31.5,31.5
                C474.9,373.5,460.8,387.8,443.4,387.8z M495.4,336.9c-17.4,0-31.5-14.1-31.5-31.5s14.1-31.5,31.5-31.5c17.4,0,31.5,14.1,31.5,31.5
                C527,322.7,512.8,336.9,495.4,336.9z"
            />

            <!-- Dpad -->
            <polygon class="Up Dpad" :class="{'gamepad-btn-pressed': displayPad.buttons[12]?.pressed}" points="218.6,362.9 218.6,338.3 198,338.3 198,363.3 208.2,371.6 	"/>
            <polygon class="Down Dpad" :class="{'gamepad-btn-pressed': displayPad.buttons[13]?.pressed}" points="198,405.3 198,430 218.6,430 218.6,405.1 208.2,396.6 	"/>
            <polygon class="Left Dpad" :class="{'gamepad-btn-pressed': displayPad.buttons[14]?.pressed}" points="186.9,373.8 162,373.8 162,394.8 187.2,394.8 195.6,384.3 	"/>
            <polygon class="Right Dpad" :class="{'gamepad-btn-pressed': displayPad.buttons[15]?.pressed}" points="230.1,395.3 254.9,395.3 254.9,374.4 229.7,374.4 221.2,384.8 	"/>

            <!-- Sticks -->
            <path class="gp-btn btn-10 lstick" :style="`${setAxis(0)}`" :class="{'gamepad-btn-pressed': displayPad.buttons[10]?.pressed || stickActive(0)}" d="M136.9,238.8c-19.4,0-35.3,15.9-35.3,35.3s15.9,35.3,35.3,35.3s35.3-15.9,35.3-35.3C172.2,254.7,156.3,238.8,136.9,238.8z"/>
            <path class="gp-btn btn-11 rstick Cam" :style="`${setAxis(2)}`" :class="{'gamepad-btn-pressed': displayPad.buttons[11]?.pressed || stickActive(2)}" d="M338.9,355.3c-19.4,0-35.3,15.9-35.3,35.3s15.9,35.3,35.3,35.3s35.3-15.9,35.3-35.3S358.3,355.3,338.9,355.3z"/>
            
            <!-- Bumpers -->
            <path class="gp-btn btn-4 BumperL" :class="{'gamepad-btn-pressed': displayPad.buttons[4]?.pressed}" d="M93.6,164.3l73.6-28.4c7.6-2.9,15.7,2.5,18,12.1l0.5,2c2.3,9.6-2,19.8-9.6,22.8l-73.6,28.4c-7.6,2.9-15.7-2.5-18-12.1l-0.5-2C81.7,177.4,86,167.2,93.6,164.3z"/>
            <path class="gp-btn btn-5 BumperR" :class="{'gamepad-btn-pressed': displayPad.buttons[5]?.pressed}" d="M394.3,149.9l0.5-2c2.3-9.6,10.4-15,18-12.1l73.6,28.4c7.6,2.9,11.9,13.1,9.6,22.8l-0.5,2c-2.3,9.6-10.4,15-18,12.1l-73.6-28.4C396.3,169.7,392,159.5,394.3,149.9z"/>

            <!-- Triggers (rendered after bumpers so they appear on top) -->
            <path class="gp-btn gp-ltrigger TriggerL" :class="{'gamepad-btn-pressed': displayPad.buttons[6]?.pressed}" d="M119,27h36.8c9,0,16.3,7.3,16.3,16.3V103c0,9-7.3,16.3-16.3,16.3H119c-9,0-16.3-7.3-16.3-16.3V43.3C102.7,34.3,110,27,119,27z"/>
            <path class="gp-btn gp-rtrigger TriggerR" :class="{'gamepad-btn-pressed': displayPad.buttons[7]?.pressed}" d="M474.5,43.3V103c0,9-7.3,16.3-16.3,16.3h-36.8c-9,0-16.3-7.3-16.3-16.3V43.3c0-9,7.3-16.3,16.3-16.3h36.8C467.1,27,474.5,34.3,474.5,43.3z"/>

            <!-- Start/Select -->
            <circle class="gp-btn btn-8 MenuL" cx="240.5" cy="288.2" r="18" :class="{'gamepad-btn-pressed': displayPad.buttons[8]?.pressed}"/>
            <circle class="gp-btn btn-9 MenuR" cx="319.9" cy="288.2" r="18" :class="{'gamepad-btn-pressed': displayPad.buttons[9]?.pressed}"/>

            <!-- Face Buttons -->
            <circle class="gp-btn btn-0 ActionA" cx="443.4" cy="356.2" r="18" :class="{'gamepad-btn-pressed': displayPad.buttons[0]?.pressed}"/>
            <circle class="gp-btn btn-1 ActionB" cx="495.4" cy="305.3" r="18" :class="{'gamepad-btn-pressed': displayPad.buttons[1]?.pressed}"/>
            <circle class="gp-btn btn-2 ActionH" cx="395.1" cy="305.3" r="18" :class="{'gamepad-btn-pressed': displayPad.buttons[2]?.pressed}"/>
            <circle class="gp-btn btn-3 ActionV" cx="443.4" cy="257" r="18" :class="{'gamepad-btn-pressed': displayPad.buttons[3]?.pressed}"/>

            <!-- Guide button -->
            <path class="gp-btn Exit" :class="{'gamepad-btn-pressed': displayPad.buttons[16]?.pressed}" d="M281.8,208.9c-7.6,0-13.7,6.2-13.7,13.7c0,7.6,6.2,13.7,13.7,13.7s13.7-6.2,13.7-13.7C295.5,215,289.4,208.9,281.8,208.9z"/>
        </svg>

        <div class="padlock" v-if="displayPad.locked">
            <i class="fas fa-lock"></i>
        </div>
    </div>
</template>

<style scoped lang="scss">
.gamepad {
    position: relative;
    width: 100%;

    &.locked {
        svg {
            opacity: 0.25;
        }

        .padlock {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 2rem;
            color: #fff;
        }
    }
}
</style>