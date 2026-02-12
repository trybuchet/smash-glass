<template>
    <div class="modal-overlay">
        <div class="modal-container">
            <div class="modal window raised">
                <div class="window-content">
                    <div class="window-header font-primary">
                        {{ title }}
                    </div>
                    <div class="window-body font-primary">
                        {{ question }}
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
const emit = defineEmits(['cancel', 'confirm']);

defineProps({
    title: {
        type: String,
        required: false
    },
    question: {
        type: String,
        required: false
    },
});

function close() {
    emit('cancel');
}

function confirm() {
    emit('confirm');
}
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