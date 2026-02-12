<script lang="ts" setup>
import { ref } from 'vue';
import { useOverlayStore } from '@/stores/overlayStore';
import OverlayWidgetView from '../OverlayWidgetView.vue';
import WebcamWidget from '@/models/widgets/WebcamWidget';

const overlayStore = useOverlayStore();

defineProps<{ widget: WebcamWidget }>();
</script>
<template>
    <div class="widget">
        <OverlayWidgetView :widget="widget">
            <div class="widget-content">
                <div class="webcam-placeholder" v-if="!overlayStore.webcamStream">
                    <i class="fa fa-video-slash"></i>
                </div>
                <video v-if="overlayStore.webcamStream" :srcObject="overlayStore.webcamStream" autoplay></video>
            </div>
        </OverlayWidgetView>
    </div>
</template>
<style scoped lang="scss">
.widget-content {
    width: 100%;
    height: 100%;
}

.webcam-placeholder {
    width: 100%;
    height: 100%;
    background-color: #302f2f;
    color: #FFF;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 32px;
}

video {
    width: 100%;
    height: 100%;
}
</style>