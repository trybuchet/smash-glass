<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useConfigStore } from '@/stores/configStore';
import { useOverlayStore } from '@/stores/overlayStore';
import { GetMonitors, MoveMainWindowToMonitor, Focus } from '@bindings/windowservice';
import { GetOverlayStyles } from '@bindings/styleservice';

const configStore = useConfigStore();
const defaultConfig = ref(configStore.app.clone());
const overlayStore = useOverlayStore();

const themes = ref([]);
const displays = ref([]);

async function loadThemes() {
    const styles = await GetOverlayStyles();

    themes.value = [
        { label: 'Default', value: 'default' },
        ...styles.map((s) => {
            const label = s.Meta?.name || s.Meta?.Name || s.ID || s.ID || 'Unknown';
            const value = s.ID || s.ID || s.Meta?.id || s.Meta?.Id || label;
            return {
                label,
                value
            }
        })
    ];
}

async function getDisplays() {
    const monitors = await GetMonitors();
    displays.value = monitors.map((m, i) => {
        return {
            label: m.Name,
            value: i
        }
    })
}

async function setDisplay(display: string|number) {
    await MoveMainWindowToMonitor(display as number);
    await Focus();
}

async function setTheme(theme: string) {
    configStore.app.overlay.theme = theme;
    await overlayStore.applyTheme(theme);
}

onMounted(async () => {
    await loadThemes();
    await getDisplays();
})
</script>
<template>

    <form class="form">
        <FormSelect
        v-if="themes.length > 0"
        label="Theme"
        name="theme"
        :value="configStore.app.overlay.theme"
        :options="themes"
        @oninput="setTheme($event)"
        >
            Select the custom theme for the overlay.
        </FormSelect>

        <FormRange
        label="Opacity"
        name="opacity"
        :modelValue="(configStore.app.overlay.opacity * 100).toFixed(0)"
        @oninput="configStore.app.overlay.opacity = $event / 100"
        >
            Set the opacity of the overlay.
        </FormRange>

        <FormSelect
        v-if="displays.length > 0"
        label="Display"
        name="theme"
        :value="configStore.app.overlay.display"
        :options="displays"
        @oninput="setDisplay($event)"
        >
            Select the display to show the overlay on.
        </FormSelect>
    </form>

</template>
<style lang="scss" scoped>
.form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}
</style>
