import FormCheckbox from "./FormCheckbox.vue";
import FormMultiSelect from "./FormMultiSelect.vue";
import FormRadio from "./FormRadio.vue";
import FormRange from "./FormRange.vue";
import FormSelect from "./FormSelect.vue";
import FormText from "./FormText.vue";
import FormToggle from "./FormToggle.vue";

export default {
    FormCheckbox,
    FormMultiSelect,
    FormRadio,
    FormRange,
    FormSelect,
    FormText,
    FormToggle
};

declare module 'vue' {
    export interface GlobalComponents {
        FormCheckbox: typeof FormCheckbox
        FormMultiSelect: typeof FormMultiSelect
        FormRadio: typeof FormRadio
        FormRange: typeof FormRange
        FormSelect: typeof FormSelect
        FormText: typeof FormText
        FormToggle: typeof FormToggle
    }
}