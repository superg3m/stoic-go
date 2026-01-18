<script setup>
import {onBeforeMount, ref} from "vue";
import router from "@/router/index.js";
import {isAuthorized, UserStore} from "@/UserStore.ts";
import PrimeColorPicker from "@/components/PrimeColorPicker.vue";


const sideBarVisible = ref(false);
const authorized = ref(false);

function toggleDropdown() {
  sideBarVisible.value = !sideBarVisible.value;
}

function onSettings() {
  sideBarVisible.value = false;
  router.push("/settings");
}

async function onLogout() {
  sideBarVisible.value = false;

  await fetch("http://localhost:8080/User/Logout", {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    }
  });

  await router.push("/")
}
</script>

<template>
  <div class="container">
    <Menubar>
      <template #start>
        <svg class="clickable" @click="router.push('/home')" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="117" height="24" viewBox="0 0 117 24">
          <image xlink:href="data:image/png;base64,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" x="0" y="0" width="117" height="24"/>
        </svg>
        <nav class="nav-bar">
            <ul class="nav-list">
              <li class="nav-item" @click="router.push('/resources')">Resources</li>
            </ul>
          </nav>
      </template>
      <template #end>
        <div class="right_container">
          <PrimeColorPicker></PrimeColorPicker>

          <div v-if="UserStore.Authorized" class="avatar-container" @click="toggleDropdown">
            <Avatar
              :style="{width: '42px', height: '42px'}"
              :image="UserStore.PFP"
              shape="circle"
              class="clickable scale-on-hover avatar"
            />
          </div>
        </div>
      </template>
    </Menubar>

    <Drawer v-model:visible="sideBarVisible" position="right" :style="{height: 'fit-content', position: 'absolute', top: 0}">
      <div class="drawer-header">
        <Avatar image="https://i.imgur.com/670yOS4.png" :style="{width: '48px', height: '48px'}" shape="circle" class="avatar-large" />
        <h1 class="user-name">{{UserStore.User.Email}}</h1>
      </div>
      <template #footer>
        <div class="drawer-footer">
          <Button @click="onSettings" label="Settings" icon="pi pi-cog" outlined></Button>
          <Button @click="onLogout" label="Logout" icon="pi pi-sign-out" severity="danger" outlined></Button>
        </div>
      </template>
    </Drawer>
  </div>
</template>

<style scoped>
.container {
  position: relative;
}

.clickable {
  cursor: pointer;
  transition: transform 0.2s ease;
}

.scale-on-hover:hover {
  transition: transform 0.2s ease;
}

.right_container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
}

.scale-on-hover:hover {
  transform: scale(1.05);
}

.avatar-container {
  display: flex;
  align-items: center;
  justify-content: center;
}

.avatar {
  width: 48px;
  height: 48px;
}

.avatar-large {
  width: 56px;
  height: 56px;
}

.drawer-header {
  display: flex;
  align-items: center;
  width: 100%;
  justify-content: space-around;
}

.user-name {
  font-size: 1rem;
  color: white;
}

.drawer-footer {
  display: flex;
  padding: 0 16px;
  justify-content: space-between;
}

.nav-bar {
  margin-left: 1rem;
}

.nav-list {
  list-style: none;
  display: flex;
  gap: 1rem;
  margin: 0;
  padding: 0;
}

.nav-item {
  cursor: pointer;
  padding: 0.5rem 1rem;
  transition: background-color 0.2s;
}

.nav-item:hover {
  background-color: rgba(0, 0, 0, 0.1);
  border-radius: 4px;
}
</style>
