<template>
    <div class="app-drawer">
        <div class="drawer-content" :class="{ open }">
            <div class="drawer-header" @click="open = !open">
                <span class="drawer-title">{{ title }}</span>
                <i class="fas fa-chevron-up" v-if="open"></i>
                <i class="fas fa-chevron-down" v-else></i>
            </div>
            <div v-if="open" class="drawer-body">
                <slot></slot>
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue';

defineProps({
    title: {
        type: String,
        default: 'Drawer Title',
    },
    startOpen: {
        type: Boolean,
        default: false,
    },
})

const open = ref(false);
</script>
<style lang="scss" scoped>
.drawer-content {
    display: flex;
    flex-direction: column;

    .drawer-header {
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: .5rem .75rem;
        background: lighten($color-primary, 4%);
        border: solid 1px rgba(255, 255, 255, 0.1);
        border-radius: 8px;
        cursor: pointer;
    }

    &.open .drawer-header {
        border-radius: 8px 8px 0 0;
    }

    .drawer-body {
        padding: 1rem;
        background: lighten($color-primary, 8%);
        border: solid 1px rgba(255, 255, 255, 0.1);
        border-top: none;
        border-radius: 0 0 8px 8px;
        display: flex;
        flex-direction: column;
        gap: .75rem;
    }
}
</style>