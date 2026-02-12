<template>
  <div class="toggle-group">
    <label v-if="label">{{ label }}</label>

    <div class="range-slider" ref="sliderRef">
      <input
        ref="inputRef"
        type="range"
        :min="min"
        :max="max"
        :step="step"
        v-model="valueInternal"
        @input="onInput"
      />
      <span data-value>{{ valueInternal }}</span>
    </div>

    <div class="form-help">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  ref,
  watch,
  onMounted,
  onBeforeUnmount,
  nextTick,
} from "vue";

const props = defineProps<{
  label?: string;
  modelValue: number;
  min?: number;
  max?: number;
  step?: number;
}>();

const emit = defineEmits();

const min = props.min ?? 0;
const max = props.max ?? 100;
const step = props.step ?? 1;

const valueInternal = ref(props.modelValue);

const sliderRef = ref<HTMLElement | null>(null);
const inputRef = ref<HTMLInputElement | null>(null);

watch(
  () => props.modelValue,
  (v) => (valueInternal.value = v)
);

function onInput() {
  emit("oninput", valueInternal.value);
}

/* Keep CSS variable --range-width updated to match input width (px).
   This preserves the original box-shadow math used for the thumb/track design. */
let ro: ResizeObserver | null = null;

function setRangeWidth() {
  const input = inputRef.value;
  const slider = sliderRef.value;
  if (!input || !slider) return;
  // use input.offsetWidth as the effective range width
  const w = input.offsetWidth || input.getBoundingClientRect().width;
  // set as pixels so calc in CSS works correctly
  slider.style.setProperty("--range-width", `${Math.round(w)}px`);
}

onMounted(async () => {
  await nextTick();
  setRangeWidth();

  // ResizeObserver on the input element (works when parent/layout changes)
  if (inputRef.value) {
    ro = new ResizeObserver(() => {
      setRangeWidth();
    });
    ro.observe(inputRef.value);
  }

  // also observe the slider container (covers some edge-cases)
  if (sliderRef.value && ro && inputRef.value !== sliderRef.value) {
    ro.observe(sliderRef.value);
  }

  // ensure correct width if fonts/images load later
  window.addEventListener("load", setRangeWidth);
  // fallback for when window resizes
  window.addEventListener("resize", setRangeWidth);
});

onBeforeUnmount(() => {
  if (ro) {
    if (inputRef.value) ro.unobserve(inputRef.value);
    if (sliderRef.value) ro.unobserve(sliderRef.value);
    ro.disconnect();
    ro = null;
  }
  window.removeEventListener("load", setRangeWidth);
  window.removeEventListener("resize", setRangeWidth);
});
</script>

<style scoped lang="scss">
.range-slider {
  --range-width: 220px;

  --track-clr: #222222;

  --thumb-w: 10px;
  --thumb-h: 30px;
  --thumb-radius: 0;
  --thumb-space: 20px;
  --thumb-clr: #df5533;
  --thumb-clr-active: #df5533;
  --thumb-scale-y-active: 0.75;

  display: flex;
  align-items: center;
  gap: 0.5rem;
  width: 100%;
  box-sizing: border-box;
  overflow: hidden;
}

.range-slider > [data-value] {
  font-size: 0.8rem;
  width: 40px;
  opacity: 0.7;
}

.range-slider > input[type="range"] {
  width: 100%;
  height: var(--thumb-h);
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  background: #222222;
  cursor: pointer;
  overflow: hidden;
  outline: none;
  border: none;
  border: solid 1px rgba(255, 255, 255, 0.2);
}

.range-slider > input[type="range"]::-webkit-slider-thumb {
  cursor: grab;
  width: var(--thumb-w);
  height: var(--thumb-h);
  background: var(--thumb-clr);
  background-color: var(--thumb-clr);
  outline: none;
  border: none;
  border-radius: var(--thumb-radius);
  box-shadow:
    calc((var(--range-width) / 2 + var(--thumb-space)) * -1) 0 0
      calc(var(--range-width) / 2) var(--thumb-clr),
    calc((var(--range-width) / 2 + var(--thumb-space))) 0 0
      calc(var(--range-width) / 2) var(--track-clr);
  -webkit-appearance: none;
  appearance: none;
  transition-property: scale, box-shadow, background-color;
  transition-duration: 300ms;
  transition-timing-function: ease-in-out;
  scale: 1 var(--thumb-scale-y, 1);
}

.range-slider > input[type="range"]::-moz-range-thumb {
  cursor: grab;
  width: var(--thumb-w);
  height: var(--thumb-h);
  background-color: var(--thumb-clr);
  outline: none;
  border: none;
  border-radius: var(--thumb-radius);
  /* track  - either side of the thumb using box-shadow */
  box-shadow:
    calc((var(--range-width) / 2 + var(--thumb-space)) * -1) 0 0
      calc(var(--range-width) / 2) var(--thumb-clr),
    calc((var(--range-width) / 2 + var(--thumb-space))) 0 0
      calc(var(--range-width) / 2) var(--track-clr);
  -moz-appearance: none;
  appearance: none;
  transition-property: scale, box-shadow, background-color;
  transition-duration: 300ms;
  transition-timing-function: ease-in-out;
  scale: 1 var(--thumb-scale-y, 1);
}

.range-slider > input[type="range"]:hover,
.range-slider > input[type="range"]:focus-within {
  --thumb-clr: var(--thumb-clr-active);
  --thumb-scale-y: var(--thumb-scale-y-active);
}

.range-slider:has(input[type="range"]:focus-visible) {
  outline: 1px dashed white;
  outline-offset: 4px;
}
</style>