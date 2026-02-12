<template>
    <div class="toggle-group">
        <label v-if="label">{{ label }}</label>
        <div class="form-text-wrapper">
            <input :disabled="busy" :type="type" :name="name" :id="name" :value="value" @input="onInput" />
        </div>
        <div class="form-help">
            <slot></slot>
        </div>
    </div>
</template>
<script lang="ts">

export default {
    name: 'FormText',
    props: {
        label: {
            type: [String, Number],
            required: false
        },
        value: {
            type: [String, Number],
            default: "",
        },
        name: {
            type: String,
            required: true,
        },
        type: {
            type: String,
            default: 'text',
        },
        debounce: {
            type: Number,
            default: 0
        }
    },
    data() {
        return {
            inputTimeout: null,
            debounceTimer: null,
            busy: false
        };
    },
    computed: {
        
    },
    methods: {
        onInput(event: any) {
            if (this.busy) {
                return;
            }

            if (this.debounce === 0) {
                this.$emit('oninput', event.target.value);
                return;
            }

            clearTimeout(this.inputTimeout);
            this.inputTimeout = setTimeout(() => {
                clearTimeout(this.debounceTimer);
                this.busy = true;
                this.$emit('oninput', event.target.value);
                this.debounceTimer = setTimeout(() => {
                    this.busy = false;
                }, this.debounce);
            }, 2000);
        }
    },
    mounted() {
    },
    beforeUnmount() {
        clearTimeout(this.inputTimeout);
        clearTimeout(this.debounceTimer);
    }
}
</script>
<style lang="scss" scoped>
.toggle-group {
    &.inline {
        display: flex;
        flex-direction: row;
        gap: 1rem;
    }
}

.form-text-wrapper {
    display: flex;
    align-items: center;
    background: #222222;
    border: solid 1px rgba(255, 255, 255, 0.2);
    padding: 0.5rem;
    border-radius: 8px;
    cursor: pointer;
    user-select: none;
    margin-top: 0.5rem;

    &.active {
        border-radius: 8px 8px 0 0;
    }

    &:hover {
        border: solid 1px lighten($color-quaternary, 10%);
    }

    input {
        cursor: pointer;
        flex: 1;
        color: $color-secondary;
        font-family: 'Roboto', sans-serif;
        user-select: none;
        outline: none;
        border: none;
        background: none;
    }
}

.form-help {
    color: $color-tertiary;
    margin-top: 0.25rem;
}
</style>