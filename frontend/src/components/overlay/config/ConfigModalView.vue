<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useConfigStore } from '@/stores/configStore';
import { useOverlayStore } from '@/stores/overlayStore';
import ConfigGeneralView from './sections/ConfigGeneralView.vue';
import ConfigWidgetsView from './sections/ConfigWidgetsView.vue';
import ConfigHotkeysView from './sections/ConfigHotkeysView.vue';
import ConfigPluginsView from './sections/ConfigPluginsView.vue';

const overlayStore = useOverlayStore();
const configStore = useConfigStore();
const defaultConfig = ref(configStore.app.clone());

const group = ref('general');
const emit = defineEmits();

function setGroup(g: string) {
    group.value = g;
}

function close() {
    emit('cancel');

    // Rollback changes
    configStore.app = defaultConfig.value.clone();
    
    // For each widget, update the config
    overlayStore.widgets.forEach(widget => {
        widget.position = configStore.app.overlay.widgets.default[widget.name].position;
        widget.enabled = configStore.app.overlay.widgets.default[widget.name].enabled;
        widget.cfg = configStore.app.overlay.widgets.default[widget.name].cfg;
    });

    overlayStore.customWidgets.forEach(widget => {
        widget.position = configStore.app.overlay.widgets.custom[widget.name].position;
        widget.enabled = configStore.app.overlay.widgets.custom[widget.name].enabled;
        widget.cfg = configStore.app.overlay.widgets.custom[widget.name].cfg;

        if (widget.id) {
            window.$eventBus.emit(`plugin:updated:${widget.id}`, widget.cfg);
        }
    });

    // Stop webcam if it was enabled but not anymore
    if (!configStore.app.overlay.widgets.default['webcam'].enabled && overlayStore.webcamStream) {
        overlayStore.webcamStream.getTracks().forEach(track => track.stop());
        overlayStore.webcamStream = null;
    }

    configStore.saveConfig();
}

function save() {
    
    emit('confirm');

    configStore.saveConfig();
}

onMounted(() => {
    
});
</script>
<template>
    <div class="modal-overlay">
        <div class="modal-container">
            <div class="modal window raised has-sidebar">
                <div class="window-sidebar">
                    <div class="window-sidebar-item" @click="setGroup('general')" :class="{ active: group === 'general' }">
                        <i class="fas fa-cog"></i>
                        <span>General</span>
                    </div>
                    <div class="window-sidebar-item" @click="setGroup('widgets')" :class="{ active: group === 'widgets' }">
                        <i class="fas fa-window-restore"></i>
                        <span>Widgets</span>
                    </div>
                    <div class="window-sidebar-item" @click="setGroup('hotkeys')" :class="{ active: group === 'hotkeys' }">
                        <i class="fas fa-keyboard"></i>
                        <span>Hotkeys</span>
                    </div>
                    <div class="window-sidebar-item" @click="setGroup('plugins')" :class="{ active: group === 'plugins' }">
                        <i class="fas fa-plug"></i>
                        <span>Plugins</span>
                    </div>
                </div>
                <div class="window-content">
                    <div class="window-header font-primary">
                        Settings
                    </div>
                    <div class="window-body font-primary">
                        
                        <ConfigGeneralView v-if="group === 'general'" />
                        <ConfigWidgetsView v-if="group === 'widgets'" />
                        <ConfigHotkeysView v-if="group === 'hotkeys'" />
                        <ConfigPluginsView v-if="group === 'plugins'" />

                    </div>
                    <div class="window-footer font-primary">
                        <div class="btn btn-secondary" @click="close">Cancel</div>
                        <div class="btn btn-primary" @click="save">Submit</div>
                    </div>
                </div>

                <div class="modal-close">
                    <div class="btn btn-round" @click="close">
                        <i class="fas fa-times"></i>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<style lang="scss" scoped>
h2 {
    margin: 0;
}

.modal-overlay {
    background: rgba(0, 0, 0, 0.2);
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.modal-container {
    width: 100%;
    height: 100%;
    max-height: 480px;
    max-width: 640px;
    scrollbar-color: var($color-tertiary) var($color-primary) !important;
}

.modal {
    width: 100%;
    height: 100%;
    position: relative;

    &.has-sidebar {
        display: flex;

        .modal-content {
            flex: 1;
        }
    }

    .modal-close {
        position: absolute;
        top: -12px;
        right: -12px;
        cursor: pointer;
    }
}

.widgets {
    margin-top: 1rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

.form-help {
    color: $color-tertiary;
    margin-top: 0.25rem;
}

.hotkey-group {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5rem;

    &:nth-child(even) {
        background: rgba(255, 255, 255, 0.05);
    }
}

.plugin-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
}

.plugin-group {
    padding: .5rem;
    border: solid 1px rgba(255, 255, 255, 0.1);
    border-radius: 0.5rem;
    margin-bottom: 1rem;
    position: relative;

    .plugin-label {
        font-weight: bold;
    }

    .plugin-delete {
        position: absolute;
        top: 0.5rem;
        right: 0.5rem;
        cursor: pointer;
        transition: all 0.2s ease-in-out;

        &:hover {
            transform: scale(1.1);
        }
    }
}
</style>