<script setup>
import {onBeforeMount, ref} from 'vue';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import router from '@/router';
import {isAuthorized} from "@/UserStore.ts";
</script>


<template>
  <div class="Card">
    <Card style="width: 25rem; overflow: hidden">
      <template #header>
        <!--<img alt="user header" src="https://primefaces.org/cdn/primevue/images/usercard.png" />
        -->
      </template>
      <template #title>
        <h2 class="text-center">Update User Password</h2>
      </template>
      <template #content>
        <div class="entryboxes">

        </div>
      </template>
    </Card>
  </div>

</template>


<style scoped>
.Card {
  display: flex;
  justify-content: center; /* Centers horizontally */
  align-items: center;    /* Centers vertically */
  height: 100vh;          /* Full viewport height for proper centering */
  padding: 1rem;          /* Optional, adds space around the card */
  background-color: #222222;
}
.text-center{
  text-align: center;
}
.button-group div {
  text-align: center;
  margin-bottom: 1rem;
}
.button-group div:last-child {
  margin-bottom: 0;
}
</style>
