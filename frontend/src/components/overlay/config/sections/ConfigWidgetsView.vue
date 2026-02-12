<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useConfigStore } from '@/stores/configStore';
import { useOverlayStore } from '@/stores/overlayStore';
import AppDrawer from '@/components/common/AppDrawer.vue';

const overlayStore = useOverlayStore();
const configStore = useConfigStore();
const positions = [
    { label: 'Top Left', value: 'top-left' },
    { label: 'Top Right', value: 'top-right' },
    { label: 'Bottom Left', value: 'bottom-left' },
    { label: 'Bottom Right', value: 'bottom-right' },
    { label: 'Top Center', value: 'top-center' },
    { label: 'Bottom Center', value: 'bottom-center' },
    { label: 'Center', value: 'center' },
    { label: 'Custom', value: 'custom' },
];
const mediaDevices = ref<{label: string, value: string}[]>([]);

function setWidgetPosition(widgetName: string, position: any) {
    const widget = overlayStore.findWidgetByName(widgetName);
    if (widget) {
        widget.position.position = position as any;
        configStore.app.overlay.widgets[widget.custom ? 'custom' : 'default'][widgetName].position.position = position;

        if (widget.custom) {
            window.$eventBus.emit(`plugin:updated:${widget.id}`, widget.cfg);
        }
    }
}

function setWidgetProp(widgetName: string, key: string, value: any) {
    const widget = overlayStore.findWidgetByName(widgetName);
    if (widget) {
        widget[key] = value;
        configStore.app.overlay.widgets[widget.custom ? 'custom' : 'default'][widgetName][key] = value;
    }
}

function setWidgetConfig(widgetName: string, key: string, value: any) {
    const widget = overlayStore.findWidgetByName(widgetName);
    if (widget) {
        if (widget.custom) {
            widget.cfg[key].value = value;
            configStore.app.overlay.widgets['custom'][widgetName].cfg[key].value = value;
            window.$eventBus.emit(`plugin:updated:${widget.id}`, widget.cfg);
        } else {
            widget.cfg[key] = value;
            configStore.app.overlay.widgets['default'][widgetName].cfg[key] = value;
        }
    }
}

async function getMediaDevices() {
    await navigator.mediaDevices.enumerateDevices().then(devices => {
        mediaDevices.value = [ 
            { label: "None", value: "-1" },
            ...devices.map(device => {
                return {
                    label: device.label,
                    value: device.deviceId
                }
            })
        ];
    })
}

function setWebcamDevice(deviceId: any) {
    configStore.app.overlay.widgets.default['webcam'].cfg.deviceId = deviceId;

    if (deviceId === "-1") {
        
        // Stop the webcam stream
        if (overlayStore.webcamStream) {
            overlayStore.webcamStream.getTracks().forEach(track => track.stop());
            overlayStore.webcamStream = null;
        }

        return;
    }

    const device = mediaDevices.value.find(device => device.value === deviceId);
    if (device) {
        navigator.mediaDevices.getUserMedia({ video: { deviceId: { exact: deviceId } } }).then(stream => {
            overlayStore.webcamStream = stream;
        });
    }
}

function toggleWebcamStream(value: boolean) {
    
    configStore.app.overlay.widgets.default['webcam'].enabled = value;
    const widget = overlayStore.findWidgetByName('webcam');
    if (widget) {
        widget.enabled = value;
    }

    if (!configStore.app.overlay.widgets.default['webcam'].enabled && overlayStore.webcamStream) {
        overlayStore.webcamStream.getTracks().forEach(track => track.stop());
        overlayStore.webcamStream = null;
    } else if (configStore.app.overlay.widgets.default['webcam'].cfg.deviceId !== "-1") {
        navigator.mediaDevices.getUserMedia({ video: { deviceId: { exact: configStore.app.overlay.widgets.default['webcam'].cfg.deviceId } } }).then(stream => {
            overlayStore.webcamStream = stream;
        });
    }
}

function pressWidgetButton(widgetId: string|number, action: string) {
    window.$eventBus.emit(`plugin:button:${widgetId}`, action);
}

onMounted(async () => {
    await getMediaDevices();
})
</script>
<template>
    <div class="config-widgets">
        <AppDrawer title="Chat" startOpen>
            <FormSelect
            label="Position"
            name="chatPosition"
            :value="(configStore.app.overlay.widgets.default['chat'].position.position)"
            :options="positions"
            @oninput="setWidgetPosition('chat', $event)"
            />
            <FormToggle 
            label="Enabled"
            :value="configStore.app.overlay.widgets.default['chat'].enabled"
            @oninput="setWidgetProp('chat', 'enabled', $event)"
            name="chatEnabled"
            />
            <FormToggle 
            label="Show Timestamps"
            :value="configStore.app.overlay.widgets.default['chat'].cfg.showTimestamps"
            @oninput="setWidgetConfig('chat', 'showTimestamps', $event)"
            name="showTimestamps"
            />
            <FormToggle 
            label="Show Message History"
            :value="configStore.app.overlay.widgets.default['chat'].cfg.showHistory"
            @oninput="setWidgetConfig('chat', 'showHistory', $event)"
            name="chatShowHistory"
            />
        </AppDrawer>

        <AppDrawer title="Guests" startOpen>
            <FormSelect
            label="Position"
            name="chatPosition"
            :value="(configStore.app.overlay.widgets.default['guests'].position.position)"
            :options="positions"
            @oninput="setWidgetPosition('guests', $event)"
            />
            <FormToggle 
            label="Enabled"
            :value="configStore.app.overlay.widgets.default['guests'].enabled"
            @oninput="setWidgetProp('guests', 'enabled', $event)"
            name="guestsEnabled"
            />
            <FormToggle 
            label="Show Latency"
            :value="configStore.app.overlay.widgets.default['guests'].cfg.showLatency"
            @oninput="setWidgetConfig('guests', 'showLatency', $event)"
            name="showLatency"
            />
        </AppDrawer>

        <AppDrawer title="Gamepads" startOpen>
            <FormSelect
            label="Position"
            name="chatPosition"
            :value="(configStore.app.overlay.widgets.default['gamepads'].position.position)"
            :options="positions"
            @oninput="setWidgetPosition('gamepads', $event)"
            />
            <FormToggle 
            label="Enabled"
            :value="configStore.app.overlay.widgets.default['gamepads'].enabled"
            @oninput="setWidgetProp('gamepads', 'enabled', $event)"
            name="gamepadsEnabled"
            />
            <FormToggle 
            label="Show Latency"
            :value="configStore.app.overlay.widgets.default['gamepads'].cfg.showHotseatTime"
            @oninput="setWidgetConfig('gamepads', 'showHotseatTime', $event)"
            name="hotseatTime"
            />
        </AppDrawer>

        <AppDrawer v-for="widget in overlayStore.customWidgets" :title="widget.name" startOpen>

            <FormToggle 
            label="Enabled"
            :value="configStore.app.overlay.widgets.custom[widget.name].enabled"
            @oninput="setWidgetProp(widget.name, 'enabled', $event); overlayStore.enableCustomWidget(widget.name, $event)"
            name="pluginEnabled"
            />

            <FormSelect
            label="Position"
            name="position"
            :value="(configStore.app.overlay.widgets.custom[widget.name].position.position)"
            :options="positions"
            @oninput="setWidgetPosition(widget.name, $event)"
            />

            <template v-for="(prop, index) in widget.cfg">

                <FormText
                v-if="prop.type === 'text'"
                :label="prop.label"
                :value="configStore.app.overlay.widgets.custom[widget.name].cfg[index].value"
                @oninput="setWidgetConfig(widget.name, (index as any), $event)"
                :name="prop.label"
                >
                    {{ prop.help }}
                </FormText>

                <FormToggle 
                v-if="prop.type === 'toggle'"
                :label="prop.label"
                :value="configStore.app.overlay.widgets.custom[widget.name].cfg[index].value"
                @oninput="setWidgetConfig(widget.name, (index as any), $event)"
                :name="prop.label"
                >
                    {{ prop.help }}
                </FormToggle>

                <FormRange
                v-if="prop.type === 'range'"
                :label="prop.label"
                :name="prop.name"
                :min="configStore.app.overlay.widgets.custom[widget.name].cfg[index].min"
                :max="configStore.app.overlay.widgets.custom[widget.name].cfg[index].max"
                :modelValue="Number(configStore.app.overlay.widgets.custom[widget.name].cfg[index].value)"
                @oninput="setWidgetConfig(widget.name, (index as any), $event)"
                >
                    {{ prop.help }}
                </FormRange>

                <FormSelect
                v-if="prop.type === 'select'"
                :label="prop.label"
                :name="prop.label"
                :value="(configStore.app.overlay.widgets.custom[widget.name].cfg[index].value)"
                :options="configStore.app.overlay.widgets.custom[widget.name].cfg[index].options"
                @oninput="setWidgetConfig(widget.name, (index as any), $event)"
                >
                    {{ prop.help }}
                </FormSelect>

                <div v-if="prop.type === 'button'" class="prop-button">
                    <div
                    class="btn btn-secondary" 
                    @click="pressWidgetButton(widget.id, index as any)"
                    >
                        {{ prop.label }}
                    </div>
                    <div class="form-help">
                        {{ prop.help }}
                    </div>
                </div>

            </template>

        </AppDrawer>
    </div>

</template>
<style lang="scss" scoped>
.config-widgets {
    display: flex;
    flex-direction: column;
    gap: 2rem;
}
</style>