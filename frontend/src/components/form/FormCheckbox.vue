<template>
    <div class="toggle-group">
        <label>{{ label }}</label>
        <div class="form-checkbox-wrapper">
            <div class="form-checkbox-input" v-for="option in options">
                <input 
                    type="checkbox"
                    :name="name"
                    :id="option.value"
                    :value="option.value"
                    :checked="value.includes(option.value)"
                    @change="$emit('oninput', ($event.target as any).checked ? [...value, option.value] : value.filter(v => v !== option.value))"
                />
                <label>{{ option.label }}</label>
            </div>
        </div>
    </div>
</template>
<script lang="ts">

export default {
    name: 'FormCheckbox',
    props: {
        label: {
            type: [String, Number],
            default: 'Text',
        },
        value: {
            type:  Array as () => string[],
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

.form-checkbox-wrapper {
    margin-top: .25rem;
    position: relative;
    display: flex;
    flex-direction: column;
    gap: .25rem;

    .form-checkbox-input {
        display: flex;
        align-items: center;
        cursor: pointer;
        position: relative;

        input[type="checkbox"] {
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
        input[type="checkbox"] {
            position: absolute;
            top: 0;
            left: 0;
            width: 20px;
            height: 20px;
            background: $color-primary;
            border: solid 1px rgba(255, 255, 255, 0.2);
            cursor: pointer;
        }
    }
}
</style>