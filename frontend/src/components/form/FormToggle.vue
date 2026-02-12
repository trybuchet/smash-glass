<template>
    <div class="toggle-container">
        <div class="toggle-group" :class="{ inline }">
            <label>{{ label }}</label>
            <div class="form-toggle-wrapper">
                <div class="form-toggle">
                    <input type="checkbox" :id="`toggle-${name}`" :checked="value"
                    @input="$emit('oninput', !value)"
                    /><label :for="`toggle-${name}`"></label>
                </div>
            </div>
        </div>
        <div class="form-help">
            <slot></slot>
        </div>
    </div>
</template>
<script lang="ts">

export default {
    name: 'FormToggle',
    props: {
        label: {
            type: [String, Number],
            default: 'Toggle',
        },
        value: {
            type: Boolean,
            default: false,
        },
        inline: {
            type: Boolean,
            default: true,
        },
        name: {
            type: String,
            required: true,
        },
    },
    data() {
        return {
        };
    },
    computed: {
        
    },
    methods: {
        
    },
    mounted() {
    },
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

.form-toggle-wrapper {
    height: 24px;
    position: relative;

    .form-toggle {
        height: 24px;

        input[type=checkbox]{
            height: 0;
            width: 0;
            visibility: hidden;
        }

        label {
            cursor: pointer;
            text-indent: -9999px;
            width: 48px;
            height: 24px;
            background: grey;
            display: block;
            border-radius: 100px;
            position: absolute;
            top: 0;
            left: 0;
        }

        label:after {
            content: '';
            position: absolute;
            top: 4px;
            left: 4px;
            width: 16px;
            height: 16px;
            background: #fff;
            border-radius: 90px;
            transition: 0.3s;
        }

        input:checked + label {
            background: lighten($color-quaternary, 10%);
        }

        input:checked + label:after {
            left: calc(100% - 5px);
            transform: translateX(-100%);
        }

        label:active:after {
            width: 24px;
        }
    }

}
</style>