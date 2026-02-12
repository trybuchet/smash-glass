<script lang="ts" setup>
import { ref } from 'vue';
import { useOverlayStore } from '@/stores/overlayStore';
import OverlayWidgetView from '../OverlayWidgetView.vue';
import GamepadsWidget from '@/models/widgets/GamepadsWidget';
import GamepadSVG from '../input/GamepadSVG.vue';

defineProps<{ widget: GamepadsWidget }>();

const overlayStore = useOverlayStore();
</script>
<template>
    <div class="widget">
        <OverlayWidgetView :widget="widget">
            <div v-if="overlayStore.pads.length > 0" class="widget-content panel">
                <div v-for="pad in overlayStore.pads" class="pad">
                    <GamepadSVG class="pad-svg" :pad="pad" />
                    <div v-if="pad.owner" class="owner guest small">
                        {{ pad.owner.name }}
                    </div>
                    <div v-if="pad.owner && widget.cfg.showHotseatTime" class="muted small">
                        {{ pad.owner.hotseatTime }}
                    </div>
                </div>
            </div>
        </OverlayWidgetView>
    </div>
</template>
<style scoped lang="scss">
.widget-content {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    padding: 8px 10px;
}

.pad {
    width: 100px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.pad-svg {
    width: 80px;
    height: 80px;
    position: relative;
}

.owner {
    width: 100px;
    overflow: hidden;
    text-align: center;
    text-overflow: ellipsis;
    white-space: nowrap;
}
</style>