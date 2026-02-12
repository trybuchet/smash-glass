
<script lang="ts" setup>
import { ref, onMounted, onBeforeUnmount, nextTick, defineProps } from 'vue';


const props = defineProps<{ noCloseOnTrigger?: boolean; noCloseOnMenuClick?: boolean }>();

function onMenuClick(event: MouseEvent) {
  if (props.noCloseOnMenuClick) {
    // Prevent closing if prop is set
    return;
  }
  isOpen.value = false;
}
const isOpen = ref(false);
const dropdownMenuStyle = ref<Record<string, string>>({});

const toggleDropdown = async () => {
  isOpen.value = !isOpen.value;
  if (isOpen.value) {
    await nextTick();
    adjustDropdownPosition();
  }
};

const dropdownRef = ref<HTMLElement | null>(null);
const menuRef = ref<HTMLElement | null>(null);
const triggerLabelRef = ref<HTMLElement | null>(null);

function adjustDropdownPosition() {
  nextTick(() => {
    requestAnimationFrame(() => {
      const triggerLabel = triggerLabelRef.value;
      const menu = menuRef.value;
      if (!triggerLabel || !menu || !isOpen.value) return;

      // Reset styles
      menu.style.left = '';
      menu.style.right = '';
      menu.style.top = '';
      menu.style.bottom = '';
      menu.style.zIndex = '5000';

      const triggerLabelRect = triggerLabel.getBoundingClientRect();
      const menuRect = menu.getBoundingClientRect();
      const vw = window.innerWidth;
      const vh = window.innerHeight;

      // Default: menu below trigger, flush
      let top = triggerLabelRect.bottom;
      let left = triggerLabelRect.left;

      // If dropdown goes off right edge, align right
      if (left + menuRect.width > vw) {
        left = Math.max(vw - menuRect.width - 8, 8); // 8px margin
      }
      // If dropdown goes off left edge, align left
      if (left < 0) {
        left = 8;
      }

      // If not enough space below, show above trigger (menu's bottom = trigger's top)
      const spaceBelow = vh - triggerLabelRect.bottom;
      if (spaceBelow < menuRect.height) {
        top = triggerLabelRect.top - menuRect.height;
        // Clamp only if menu would go off the top
        if (top < 8) top = 8;
      } else {
        
      }

      menu.style.position = 'fixed';
      menu.style.left = left + 'px';
      menu.style.top = top + 'px';
    });
  });
}


function onTriggerClick(event: MouseEvent) {
  if (props.noCloseOnTrigger && isOpen.value) {
    // If already open and prop is set, do not close
    return;
  }
  toggleDropdown();
}

function onDocumentClick(event: MouseEvent) {
  const target = event.target as HTMLElement;
  if (props.noCloseOnMenuClick && target.closest('.vf-multiselect-option')) {
    return;
  }
  if (!target.closest('.dropdown')) {
    isOpen.value = false;
  }
}

onMounted(() => {
  document.addEventListener('click', onDocumentClick);
});

onBeforeUnmount(() => {
  document.removeEventListener('click', onDocumentClick);
});

declare module 'vue' {
  interface GlobalComponents {
    /**
     * DropDown
     * A dropdown component that toggles visibility of a menu when the trigger 
     * is clicked. The menu is positioned intelligently to stay within the viewport.
     * 
     * @example
     * <DropDown>
     *  <template #trigger>
     *    <AppButton>Toggle</AppButton>
     *  </template>
     *  <template #menu>
     *    <div>Menu Content</div>
     *  </template>
     * </DropDown>
     */
    DropDown: Object;
  }
}
</script>

<template>
  <div class="dropdown" ref="dropdownRef">
    <div class="dropdown-label" ref="triggerLabelRef" @click.stop="onTriggerClick">
      <slot name="trigger"></slot>
    </div>
    <div v-if="isOpen" class="dropdown-menu" ref="menuRef"
      @click.stop="onMenuClick"
    >
      <slot name="menu"></slot>
    </div>
  </div>
</template>
<style scoped lang="scss">
.dropdown {
    position: relative;
    display: inline-block;
    user-select: none;
    cursor: pointer;

  .dropdown-menu {
    background: rgba(0, 0, 0, 0.9);
    border: solid 1px #903721;
    border-radius: 4px;
    box-shadow: 0 0 8px rgba(0, 0, 0, 0.1);
    z-index: 1000;
    overflow-y: auto;
    overflow-x: hidden;
    width: max-content;
    min-width: 100px;
  }
}
</style>