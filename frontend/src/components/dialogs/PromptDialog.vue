<template>
    <div class="modal-overlay">
        <div class="modal-container">
            <div class="modal window raised">
                <div class="window-content">
                    <div class="window-header font-primary">
                        {{ title }}
                    </div>
                    <div class="window-body font-primary">
                        <div class="dialog-message" v-html="message"></div>
                        <FormText
                        @oninput="response = $event"
                        name="dialog-response"
                        />
                    </div>
                    <div class="window-footer font-primary">
                        <div class="btn btn-secondary" @click="close">Cancel</div>
                        <div class="btn btn-primary" @click="confirm">Set</div>
                    </div>
                </div>

                <div class="modal-close">
                    <div class="btn btn-round" @click="close">
                        <i class="fas fa-times"></i>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue';
import FormText from '@/components/form/FormText.vue';

export default defineComponent({
    name: 'PromptDialog',
    components: {
        FormText
    },
    props: {
        title: {
            type: String,
            required: false
        },
        message: {
            type: String,
            required: false
        },
    },
    data() {
        return {
            response: '',
        }
    },
    methods: {
      close() {
        this.$emit('cancel');
      },
      confirm() {
        if (!this.response || this.response.length === 0) return;
        this.$emit('confirm', this.response);
      },
    },
    mounted() {

    },
    beforeUnmount() {
    }
});
</script>
<style lang="scss" scoped>
.modal-overlay {
    background: rgba(0, 0, 0, 0.2);
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal-container {
    width: 100%;
    height: 100%;
    max-height: 240px;
    max-width: 400px;
}

.modal {
    width: 100%;
    height: 100%;
    position: relative;

    &.has-sidebar {
        display: flex;

        .modal-content {
            flex: 1;
        }
    }

    .modal-close {
        position: absolute;
        top: -12px;
        right: -12px;
        cursor: pointer;
    }
}

.form-help {
color: $color-tertiary;
margin-top: 0.25rem;
}
</style>