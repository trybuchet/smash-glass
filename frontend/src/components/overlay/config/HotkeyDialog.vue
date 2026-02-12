<template>
    <div class="modal-overlay">
        <div class="modal-container">
            <div class="modal window raised">
                <div class="window-content">
                    <div class="window-header font-primary">
                        Set Hotkey
                    </div>
                    <div class="window-body font-primary">

                        Press some modifier keys and a key to set a custom hotkey.
                        <div v-if="modifiers.length > 0" class="form-help">
                        {{ hotkeyString }}
                        </div>

                    </div>
                    <div class="window-footer font-primary">
                        <div class="btn btn-secondary" @click="close">Cancel</div>
                        <div class="btn btn-primary" @click="confirm">Confirm</div>
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
<script lang="ts" setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import { getKeyName } from '@/utils/helpers';

const modifiers = ref([]);
const key = ref(-1);

const emit = defineEmits(['cancel', 'confirm']);

// Computed hotkey string
const hotkeyString = computed(() => {
    if (key.value === -1) {
        return modifiers.value.map((m: number) => getKeyName(m)).join(' + ');
    }
    return modifiers.value.map((m: number) => getKeyName(m)).join(' + ') + ' + ' + getKeyName(key.value);
});

function onKeyDown(event: KeyboardEvent) {
    event.preventDefault();

    if (event.key === 'Escape') {
        close();
        return;
    }
    if (event.key === 'Control' || event.key === 'Shift' || event.key === 'Alt') {
        if (!modifiers.value.includes(event.keyCode)) {
        modifiers.value.push(event.keyCode);
        }
    } else {
        key.value = event.keyCode;
    }
}

function close() {
    emit('cancel');
}

function confirm() {
    emit('confirm', { modifiers: modifiers.value, key: key.value });
}

onMounted(() => {
    document.addEventListener('keydown', onKeyDown);
});

onBeforeUnmount(() => {
    document.removeEventListener('keydown', onKeyDown);
});
</script>
<style lang="scss" scoped>
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
    z-index: 1000;
}

.modal-container {
    width: 100%;
    height: 100%;
    max-height: 220px;
    max-width: 400px;
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

.window-body {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>