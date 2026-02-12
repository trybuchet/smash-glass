<script lang="ts" setup>
import { useOverlayStore } from '@/stores/overlayStore';

const overlayStore = useOverlayStore();

function downloadPlugins() {
    
}
</script>
<template>

    <div class="plugin-header">
        <h2>Installed Plugins</h2>
        <div class="btn btn-primary" @click="overlayStore.getPlugins()">Get More!</div>
    </div>
    <div v-for="(plugin, index) in overlayStore.plugins" class="plugin-group">
        <div class="plugin">
            <div>
                <span class="plugin-label">{{ plugin.Meta.name }}</span>
            </div>
            <div style="margin-bottom: 0.5rem;">
                <i>by <span class="form-help">{{ plugin.Meta.author }}</span></i>
            </div>
            <div>
                {{ plugin.Meta.description }}
            </div>
        </div>
        <div class="plugin-delete">
            <i class="fas fa-trash" @click="overlayStore.deletePlugin(plugin.Meta.id)"></i>
        </div>
    </div>

</template>
<style lang="scss" scoped>
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