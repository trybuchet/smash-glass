<script lang="ts" setup>
import { computed, onMounted } from 'vue';
import { useOverlayStore } from '@/stores/overlayStore';
import { useConfigStore } from '@/stores/configStore';
import { useSocketStore } from '@/stores/socketStore';

import OverlayChatWidgetView from './widgets/OverlayChatWidgetView.vue';
import OverlayGuestsWidgetView from './widgets/OverlayGuestsWidgetView.vue';
import OverlayGamepadsWidgetView from './widgets/OverlayGamepadsWidgetView.vue';
import OverlayWebcamWidgetView from './widgets/OverlayWebcamWidgetView.vue';
import OverlayWidgetView from './OverlayWidgetView.vue';

const overlayStore = useOverlayStore();
const configStore = useConfigStore();
const socketStore = useSocketStore();

const chatWidget = computed(() => overlayStore.findWidgetByName('chat'));
const guestsWidget = computed(() => overlayStore.findWidgetByName('guests'));
const gamepadsWidget = computed(() => overlayStore.findWidgetByName('gamepads'));
const webcamWidget = computed(() => overlayStore.findWidgetByName('webcam'));

const overlayStyle = computed(() => ({
    opacity: configStore.app.overlay.opacity,
}));

onMounted(() => {
    socketStore.connect();
})
</script>

<template>
    <div id="app" class="overlay" :style="overlayStyle">
        <OverlayChatWidgetView v-if="chatWidget?.enabled" :widget="chatWidget" />
        <OverlayGuestsWidgetView v-if="guestsWidget?.enabled" :widget="guestsWidget" />
        <OverlayGamepadsWidgetView v-if="gamepadsWidget?.enabled" :widget="gamepadsWidget" />
        <OverlayWebcamWidgetView v-if="webcamWidget?.enabled" :widget="webcamWidget" />

        <!-- Custom Widgets -->
        <template v-for="widget in overlayStore.customWidgets">
            <OverlayWidgetView v-if="widget.enabled" :widget="widget">
                <div v-html="widget.html"></div>
            </OverlayWidgetView>
        </template>
    </div>
</template>

<style scoped lang="scss">
.overlay {
    position: fixed;
    inset: 0;
    overflow: hidden;
}
</style>