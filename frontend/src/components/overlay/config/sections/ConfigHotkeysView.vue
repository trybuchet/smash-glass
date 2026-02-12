<script lang="ts" setup>
import { createConfirmDialog } from 'vuejs-confirm-dialog';
import { ChangeHotkey } from '@bindings/hotkeyservice';
import { ChangeHotkeyArgs } from '@bindings/models';
import { useConfigStore } from '@/stores/configStore';
import { getKeyName } from '@/utils/helpers';
import HotkeyDialog from '../HotkeyDialog.vue';

const configStore = useConfigStore();

function getHotkeyName(hotkey: string) {
    return hotkey.split(':').slice(1).join(' ').replace(/\b\w/g, (c: string) => c.toUpperCase());
}

function getHotkeyString(hotkey: { modifiers: number[], key: number, event: string }) {
    return hotkey.modifiers.map((m: number) => getKeyName(m)).join(' + ') + ' + ' + getKeyName(hotkey.key);
}

function setHotkey(event: string) {

    const { reveal, onConfirm } = createConfirmDialog(HotkeyDialog, {
    })

    onConfirm((data: { modifiers: number[], key: number }) => {
        configStore.app.hotkeys[event].modifiers = data.modifiers;
        configStore.app.hotkeys[event].key = data.key;

        configStore.saveConfig();

        const hotkey = new ChangeHotkeyArgs();
        hotkey.Name = event;
        hotkey.Modifiers = data.modifiers;
        hotkey.Key = data.key;

        ChangeHotkey(hotkey).then(() => {
            console.log('Hotkey changed');
        }).catch((err: any) => {
            console.error(err);
        });
    })
    reveal()

}
</script>
<template>

    <div style="margin-bottom: 1rem;">
        Here you can configure the keyboard shortcut keys for the overlay.
    </div>
    <div v-for="(hotkey, index) in configStore.app.hotkeys" class="hotkey-group">
        <div class="hotkey">
            <div class="hotkey-label">{{ getHotkeyName(index) }}</div>
            <div class="form-help">
                {{ getHotkeyString(hotkey) }}
            </div>
        </div>
        <div class="btn btn-secondary" @click="setHotkey(hotkey.event)">Change</div>
    </div>

</template>
<style lang="scss" scoped>
.hotkey-group {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0.5rem;

    &:nth-child(even) {
        background: rgba(255, 255, 255, 0.05);
    }
}
</style>