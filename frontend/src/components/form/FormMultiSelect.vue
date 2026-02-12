<template>
  <div class="toggle-group">
    <label class="form-label">{{ label }}</label>
    <div class="form-select-wrapper">
      <div class="form-select-input" :class="{ active }" @click="toggleDropdown">
        <div class="tags-container">
          <span v-for="(tag, index) in selectedLabels" :key="tag + index" class="tag">
            {{ tag }}
            <i class="fas fa-times" @click.stop="removeTag(index)"></i>
          </span>

          <input
            autocomplete="off"
            ref="inputEl"
            type="text"
            v-model="searchQuery"
            :disabled="busy || reachedMax"
            @click.stop
            @input="onSearch"
            @keydown="onKeyDown"
            @blur="onBlur"
            placeholder="Type or select..."
          />
        </div>

        <i v-if="!active" class="fas fa-caret-down"></i>
        <i v-else class="fas fa-caret-up"></i>
      </div>

      <div v-if="active" class="form-select-dropdown">
        <div
          v-for="option in filteredOptions"
          :key="option.value"
          @click="selectOption(option)"
          :class="{
            active: selectedValues.includes(option.value),
            disabled: reachedMax
          }"
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
  name: "FormMultiSelect",
  props: {
    label: { type: [String, Number], default: "Select" },
    value: { type: Array as () => string[], default: () => [] },
    name: { type: String, required: true },
    options: {
      type: Array as () => { label: string; value: string; image?: string }[],
      default: () => [],
    },
    debounce: { type: Number, default: 0 },
    maxTags: { type: Number, default: 5 },
  },
  data() {
    return {
      active: false,
      searchQuery: "",
      busy: false,
      inputTimeout: null as any,
      debounceTimer: null as any,
      selectedValues: [...this.value],
    };
  },
  computed: {
    selectedLabels(): string[] {
      return this.selectedValues.map(
        (v) => this.options.find((o) => o.value === v)?.label || v
      );
    },
    filteredOptions(): { label: string; value: string; image?: string }[] {
      if (!this.searchQuery) return this.options;
      return this.options.filter((option) =>
        option.label.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
    reachedMax(): boolean {
      return this.selectedValues.length >= this.maxTags;
    },
  },
  methods: {
    toggleDropdown() {
      if (this.reachedMax) return;
      this.active = !this.active;
      this.$nextTick(() => (this.$refs.inputEl as HTMLInputElement)?.focus());
    },
    hideDropdown(e: Event) {
      if (!this.$el.contains(e.target as Node)) this.active = false;
    },

    onSearch(event: any) {
      const val: string = event.target.value;

      // If user types or pastes commas, process
      if (val.includes(",")) {
        this.addTypedTags();
        return;
      }

      if (this.busy) return;

      if (this.debounce === 0) {
        this.$emit("onsearch", val);
        return;
      }

      clearTimeout(this.inputTimeout);
      this.inputTimeout = setTimeout(() => {
        clearTimeout(this.debounceTimer);
        this.busy = true;
        this.$emit("onsearch", val);
        this.debounceTimer = setTimeout(() => {
          this.busy = false;
        }, this.debounce);
      }, 500);
    },

    onKeyDown(event: KeyboardEvent) {
      const key = event.key;

      // Commit tags on Enter or comma
      if (key === "Enter" || key === ",") {
        event.preventDefault();
        this.addTypedTags();
        return;
      }

      // Backspace removes last tag if input empty
      if (key === "Backspace" && this.searchQuery.length === 0) {
        this.removeTag(this.selectedValues.length - 1);
      }
    },

    onBlur() {
      // short delay so clicks on dropdown still register first
      setTimeout(() => {
        this.addTypedTags();
        this.active = false;
      }, 100);
    },

    addTypedTags() {
      const raw = this.searchQuery.trim();
      if (!raw) return;

      const parts = raw
        .split(",")
        .map((p) => p.trim())
        .filter((p) => p.length > 0);

      for (const part of parts) {
        if (this.reachedMax) break;

        // find valid option (case-insensitive by label)
        const optionMatch = this.options.find(
          (o) =>
            o.label.toLowerCase() === part.toLowerCase() ||
            o.value.toLowerCase() === part.toLowerCase()
        );

        // if no match → skip entirely (discard)
        if (!optionMatch) continue;

        const val = optionMatch.value;
        if (!this.selectedValues.includes(val)) {
          this.selectedValues.push(val);
        }
      }

      this.searchQuery = "";
      this.emitChange();
    },

    selectOption(option: { label: string; value: string }) {
      if (this.reachedMax) return;
      if (!this.selectedValues.includes(option.value)) {
        this.selectedValues.push(option.value);
        this.emitChange();
      }
      this.searchQuery = "";
      this.active = false;
    },

    removeTag(index: number) {
      if (index < 0) return;
      this.selectedValues.splice(index, 1);
      this.emitChange();
      this.$nextTick(() => (this.$refs.inputEl as HTMLInputElement)?.focus());
    },

    emitChange() {
      this.$emit("oninput", [...this.selectedValues]);
      this.$emit("update:modelValue", [...this.selectedValues]);
    },
  },
  mounted() {
    document.addEventListener("click", this.hideDropdown);
  },
  beforeUnmount() {
    document.removeEventListener("click", this.hideDropdown);
  },
  watch: {
    value(newVal: string[]) {
      this.selectedValues = [...newVal];
    },
  },
};
</script>

<style lang="scss" scoped>
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
  background: $color-primary;
  border: solid 1px rgba(255, 255, 255, 0.2);
  padding: 0.5rem;
  border-radius: 8px;
  cursor: text;
  user-select: none;
  flex-wrap: wrap;
  gap: 0.25rem;

  input {
    flex: 1;
    min-width: 120px;
    color: $color-secondary;
    background: none;
    border: none;
    outline: none;
    font-family: "Roboto", sans-serif;
  }

  &.active {
    border-radius: 8px 8px 0 0;
  }

  &:hover {
    border: solid 1px lighten($color-quaternary, 10%);
  }

  .tags-container {
    display: flex;
    flex-wrap: wrap;
    gap: 0.25rem;
    flex: 1;
    align-items: center;
  }

  .tag {
    background: lighten($color-primary, 5%);
    border-radius: 4px;
    padding: 0.15rem 0.35rem;
    display: flex;
    align-items: center;
    gap: 0.25rem;
    font-size: 0.85rem;
    color: $color-secondary;

    i {
      cursor: pointer;
      font-size: 0.7rem;
    }
  }
}

.form-select-dropdown {
  background: darken($color-primary, 2%);
  border: solid 1px rgba(255, 255, 255, 0.2);
  border-top: none;
  border-radius: 0 0 8px 8px;
  max-height: 200px;
  overflow-y: auto;

  .form-select-option {
    padding: 0.25rem 0.5rem;
    cursor: pointer;
    user-select: none;
    display: flex;
    align-items: center;
    gap: 0.5rem;

    &:hover {
      background: lighten($color-primary, 2%);
    }

    &.active {
      background: linear-gradient(
        90deg,
        darken($color-quaternary, 10%),
        darken($color-quaternary, 5%)
      );
      text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
    }

    &.disabled {
      opacity: 0.5;
      cursor: not-allowed;
    }

    &:last-child {
      border-radius: 0 0 8px 8px;
    }
  }
}

.form-help {
  color: $color-tertiary;
  margin-top: 0.25rem;
}
</style>