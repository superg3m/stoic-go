<script setup>
import Popover from 'primevue/popover';
import {ref} from "vue";
import {updatePrimaryPalette} from "@primeuix/themes";

const activeIndex = ref(0);
const colors = {
  "emerald": "#10b981",
  "green" : "#22c55e",
  "lime" : "#84cc16",
  "red" : "#Ff4444",
  "orange" : "#f97316",
  "amber" : "#f59e0b",
  "yellow" : "#facc15",
  "teal" : "#14b8a6",
  "cyan" : "#06b6d4",
  "sky" : "#0ea5e9",
  "blue" : "#3b82f6",
  "indigo" : "#6366f1",
  "violet" : "#8b5cf6",
  "purple" : "#a855f7",
  "fuchsia" : "#d946ef",
  "pink" : "#ec4899",
  "rose" : "#f43f5e",
  "slate" : "#64748b",
  "gray" : "#6b7280",
  "zinc" : "#8e8e93",
  "neutral" : "#d1d5db",
  "stone" : "#525259"
};

function changePrimaryColor(colorName, index) {
  updatePrimaryPalette({
    50: `{${colorName}.50}`,
    100: `{${colorName}.100}`,
    200: `{${colorName}.200}`,
    300: `{${colorName}.300}`,
    400: `{${colorName}.400}`,
    500: `{${colorName}.500}`,
    600: `{${colorName}.600}`,
    700: `{${colorName}.700}`,
    800: `{${colorName}.800}`,
    900: `{${colorName}.900}`,
    950: `{${colorName}.950}`
  });

  activeIndex.value = index
}

const colorPanel = ref()
function toggleColorPanel(event) {
  colorPanel.value.toggle(event);
}

</script>

<template>
  <Button style="width: 32px; height: 32px"  icon="pi pi-palette" label="" @click="toggleColorPanel($event)" />
  <Popover ref="colorPanel">
    <div class="color-container">

      <button v-for="(colorName, index) in Object.keys(colors)" v-tooltip.top="{ value: colorName, showDelay: 100, hideDelay: 100 }" @click="changePrimaryColor(colorName, index)" :key="index" :class="{ 'activeColor': index === activeIndex, 'passiveColor': index !== activeIndex }" :style="{ '--color': colors[colorName] }" />
    </div>
  </Popover>
</template>

<style scoped>

@property --color {
  syntax: "<color>";
  inherits: false;
  initial-value: #10b981;
}

.activeColor {
  background-color: var(--color);
  border-radius: 100%;
  width: 24px;
  height: 24px;
  cursor: pointer;
  border: 3px solid black;
}

.passiveColor {
  background-color: var(--color);
  border-radius: 100%;
  width: 24px;
  height: 24px;
  cursor: pointer;
  transition: background-color 0.4s ease;
}

.passiveColor:hover {
  filter: brightness(0.75);
}

.color-container {
  display: flex;
  max-width: 200px;
  flex-wrap: wrap;
  gap: 8px;
}
</style>
