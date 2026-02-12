<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps<{
    icon?: string;
    title?: string;
    active?: boolean;
}>();
</script>

<template>
    <div class="dropdown-item">
        <i v-if="props.icon" :class="props.icon"></i>
        <span>{{ props.title }}</span>
    </div>
</template>

<style lang="scss" scoped>
.dropdown-item {
    display: flex;
    align-items: center;
    padding: 6px 12px;
    cursor: pointer;
    transition: all 0.2s ease;
    border-top: solid 1px rgba(255, 255, 255, 0.1);
    border-bottom: solid 1px rgba(0, 0, 0, 0.1);
    min-width: 175px;
    width: 100%;

    &:first-child {
        border-radius: 5px 5px 0 0;
        border-top: none;
    }

    &:last-child {
        border-radius: 0 0 5px 5px;
        border-bottom: none;
    }

    &:hover {
        background-color: rgba(255, 255, 255, 0.1);
    }

    &.active {
        //background-color: $c_pink !important;
        color: white;
    }

    i {
        margin-right: 8px;
        font-size: .875rem;
        width: 20px;
        text-align: center;
    }

    span {
        font-size: 14px;
    }
}
</style>