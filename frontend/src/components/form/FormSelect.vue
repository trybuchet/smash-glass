<template>
    <div class="toggle-group">
        <label class="form-label">{{ label }}</label>
        <div class="form-select-wrapper">
            <div ref="inputWrapper" class="form-select-input" @click="toggleDropdown" :class="{ active }">
                <input 
                    ref="textInput"
                    type="text" 
                    :name="name" 
                    v-model="searchQuery"
                    :disabled="busy"
                    :readonly="!active"
                    @click.stop="onInputClick"
                    @input="onInputChange"
                    placeholder="Type to search..."
                />
                <i v-if="!active" class="fas fa-caret-down"></i>
                <i v-else class="fas fa-caret-up"></i>
            </div>

            <div v-if="active" class="form-select-dropdown" :style="dropdownStyle">
                <div
                    v-for="option in filteredOptions"
                    :key="option.value"
                    @click="selectOption(option)"
                    :class="{ active: option.value === value }"
                    class="form-select-option"
                >
                    <img v-if="option.image" :src="option.image" />
                    <span>{{ option.label }}</span>
                </div>

                <div v-if="filteredOptions.length === 0" class="form-select-option disabled">
                    No results
                </div>
            </div>

            <div class="form-help">
                <slot></slot>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: 'FormSelect',
    props: {
        label: { type: [String, Number], default: 'Toggle' },
        value: { type: [String, Number], default: '' },
        name: { type: String, required: true },
        options: {
            type: Array as () => { label: string; value: string; image?: string }[],
            default: () => [],
        },
        debounce: { type: Number, default: 0 },
    },
    data() {
        return {
            active: false,
            searchQuery: '',
            isUserTyping: false,
            busy: false,
            inputTimeout: null,
            debounceTimer: null,
            dropdownStyle: {
                top: '0px',
                left: '0px',
                width: '0px',
            },
        };
    },
    computed: {
        filteredOptions() {
            // If user hasn't typed, always show all options
            if (!this.isUserTyping || !this.searchQuery) {
                return this.options;
            }

            return this.options.filter(option =>
                option.label.toLowerCase().includes(this.searchQuery.toLowerCase())
            );
        },
    },
    methods: {
        getLabel(value) {
            return this.options.find(opt => opt.value === value)?.label || '';
        },

        toggleDropdown() {
            if (this.active) {
                this.active = false;
                this.isUserTyping = false;
                return;
            }

            this.openDropdown();
        },

        hideDropdown(e) {
            if (!this.$el.contains(e.target)) {
                this.active = false;
                this.isUserTyping = false;
            }
        },

        onInputChange(event) {
            this.isUserTyping = true;
            this.onSearch(event);
        },

        onInputClick() {
            if (!this.active) {
                this.openDropdown();
            }
        },

        onSearch(event) {
            if (this.busy) return;

            if (this.debounce === 0) {
                this.$emit('onsearch', event.target.value);
                return;
            }

            clearTimeout(this.inputTimeout);
            this.inputTimeout = setTimeout(() => {
                clearTimeout(this.debounceTimer);
                this.busy = true;
                this.$emit('onsearch', event.target.value);
                this.debounceTimer = setTimeout(() => {
                    this.busy = false;
                }, this.debounce);
            }, 2000);
        },

        selectOption(option) {
            this.$emit('oninput', option.value);
            this.searchQuery = option.label;
            this.isUserTyping = false;
            this.active = false;
        },

        openDropdown() {
            this.closeOtherDropdowns();
            this.active = true;
            this.isUserTyping = false;
            this.$nextTick(() => {
                this.updateDropdownPosition();
                const input = this.$refs.textInput as HTMLInputElement | undefined;
                input?.focus();
            });
        },

        closeOtherDropdowns() {
            const current = this.$refs.inputWrapper as HTMLElement | undefined;
            document.querySelectorAll('.form-select-input.active').forEach((el) => {
                if (current && el === current) return;
                (el as HTMLElement).click();
            });
        },

        updateDropdownPosition() {
            const wrapper = this.$refs.inputWrapper as HTMLElement | undefined;
            if (!wrapper) return;

            const rect = wrapper.getBoundingClientRect();
            this.dropdownStyle = {
                top: `${rect.bottom}px`,
                left: `${rect.left}px`,
                width: `${rect.width}px`,
            };
        },
    },

    mounted() {
        document.addEventListener('click', this.hideDropdown);
        window.addEventListener('resize', this.updateDropdownPosition);
        window.addEventListener('scroll', this.updateDropdownPosition, true);

        this.searchQuery = this.getLabel(this.value);
    },

    beforeUnmount() {
        document.removeEventListener('click', this.hideDropdown);
        window.removeEventListener('resize', this.updateDropdownPosition);
        window.removeEventListener('scroll', this.updateDropdownPosition, true);
    },

    watch: {
        value(newVal) {
            this.searchQuery = this.getLabel(newVal);
            this.isUserTyping = false;
        },
        options() {
            this.searchQuery = this.getLabel(this.value);
            this.isUserTyping = false;
        },
    },
};
</script>

<style scoped lang="scss">
.toggle-group {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.form-select-wrapper {
  position: relative;
  width: 100%;
}

.form-select-input {
  display: flex;
  align-items: center;
  background: #222;
  border: 1px solid rgba(255, 255, 255, 0.2);
  padding: 0.5rem;
  border-radius: 8px;
  cursor: pointer;

  &.active {
    border-radius: 8px 8px 0 0;
  }

  input {
    flex: 1;
    background: none;
    border: none;
    outline: none;
    color: #fff;
  }
}

.form-select-dropdown {
  position: fixed;
  z-index: 1000;
  background: #111;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-top: none;
  border-radius: 0 0 8px 8px;
  max-height: 200px;
  overflow-y: auto;

  .form-select-option {
    padding: 0.25rem 0.5rem;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 0.5rem;

    &.disabled {
      opacity: 0.5;
      pointer-events: none;
    }

    &.highlighted,
    &:hover {
      background: #333;
    }

    &.selected {
      background: linear-gradient(90deg, #555, #777);
    }
  }
}
</style>
