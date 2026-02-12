<template>
    <div class="toggle-group">
        <label>{{ label }}</label>
        <div class="form-radio-wrapper">
            <div class="form-radio-input" v-for="option in options">
                <input type="radio" :name="name" :value="option.value" :checked="option.value === value" @change="$emit('oninput', option.value)" />
                <label>{{ option.label }}</label>
            </div>
        </div>
    </div>
</template>
<script lang="ts">

export default {
    name: 'FormRadio',
    props: {
        label: {
            type: [String, Number],
            default: 'Text',
        },
        value: {
            type:  [String, Number],
            default: false,
        },
        name: {
            type: String,
            required: true,
        },
        options: {
            type: Array as () => {label: string, value: string}[],
            default: () => [],
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

.form-radio-wrapper {
    margin-top: .25rem;
    position: relative;
    display: flex;
    flex-direction: column;
    gap: .25rem;

    .form-radio-input {
        display: flex;
        align-items: center;
        cursor: pointer;
        position: relative;

        input[type="radio"] {
            cursor: pointer;
        }

        label {
            cursor: pointer;
            margin-left: 2rem;
            position: relative;
            top: 0.1rem;
            color: $color-tertiary;
        }

        // Style the radio button
        input[type="radio"] {
            position: absolute;
            top: 0;
            left: 0;
            width: 20px;
            height: 20px;
            border-radius: 50%;
            background: $color-primary;
            border: solid 1px rgba(255, 255, 255, 0.2);
            cursor: pointer;
        }

        // Create custom radio button
        input[type="radio"]:checked {
            background-color: $color-primary;
            border: solid 1px rgba(255, 255, 255, 0.2);
        }

        input[type="radio"]:checked::after {
            content: '';
            position: absolute;
            top: 50%;
            left: 50%;
            width: 10px;
            height: 10px;
            background-color: $color-primary;
            border-radius: 50%;
            transform: translate(-50%, -50%);
        }

        input[type="radio"]:hover {
            border: solid 1px lighten($color-quaternary, 10%);
        }

        input[type="radio"]:checked:hover {
            border: solid 1px lighten($color-quaternary, 10%);
        }
    }
}
</style>