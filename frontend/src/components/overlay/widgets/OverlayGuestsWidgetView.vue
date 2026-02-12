<script lang="ts" setup>
import { ref } from 'vue';
import { useOverlayStore } from '@/stores/overlayStore';
import OverlayWidgetView from '../OverlayWidgetView.vue';
import GuestsWidget from '@/models/widgets/GuestsWidget';

defineProps<{ widget: GuestsWidget }>();

const overlayStore = useOverlayStore();
</script>
<template>
    <div class="widget">
        <OverlayWidgetView :widget="widget">
            <div v-if="overlayStore.guests.length > 0" id="widget-guests" class="widget-content panel">
                <div v-for="guest in overlayStore.guests" class="guest-item">
                    <div class="guest">
                        {{ guest.name }}
                    </div>
                    <div v-if="widget.cfg.showLatency" class="latency">
                        {{ guest.latency }}ms
                    </div>
                </div>
            </div>
        </OverlayWidgetView>
    </div>
</template>
<style scoped lang="scss">
#widget-guests {
    width: 100%;
    height: 100%;
}

.guest-item {
    display: flex;
    justify-content: space-between;
    margin-bottom: .5rem;
    gap: .5rem;

    &:last-child {
        margin-bottom: 0;
    }

    .guest {
        flex: 1;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }
}
</style>