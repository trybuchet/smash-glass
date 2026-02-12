<script setup lang="ts">
import OverlayView from './overlay/OverlayView.vue';
</script>

<template>
    <div class="app-container">
        <OverlayView />
    </div>
    <DialogsWrapper />
</template>

<style scoped>
.app-container {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    box-sizing: border-box;
    overflow: hidden;
}
</style>
