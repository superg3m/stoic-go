<script setup>
import {onBeforeMount, ref} from 'vue';
import Button from 'primevue/button';
import Card from 'primevue/card';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import router from '@/router';
import { UserStore } from "@/UserStore.ts"
import Dialog from 'primevue/dialog';

const toast = useToast();
const newPassword = ref('');
const oldPassword = ref('');
const newEmail = ref('');

const visibleEmail = ref(false);
const visiblePass = ref(false);

async function tryDelete() {
  const user = {
    "ID": UserStore.User.ID
  };

  try {
    const response = await fetch("http://localhost:8080/User", {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      const errorsJson = await response.json();
      const errors = errorsJson.Errors;
      for (let i = 0; i < errors.length; i++) {
        toast.add({
          severity: "error",
          summary: "Error",
          detail: errors[i],
          life: 3000
        });
      }

      return;
    }

    await onLogout();
    await router.push('/');

    const successText = await response.text();
    toast.add({
      severity: "success",
      summary: "Success",
      detail: successText,
      life: 3000
    });
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Network Error",
      detail: error.message,
      life: 3000
    });
  }
}

async function onLogout() {
  await fetch("http://localhost:8080/User/Logout", {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    }
  });

  await router.push("/")
}

async function updatePFP() {
  let number = Math.floor(Math.random()*10)
  const profilePictures = [
    "https://i.imgur.com/670yOS4.png", // case 0
    "https://i.imgur.com/xCHrjdR.png", // case 1
    "https://i.imgur.com/P3oXQFm.png", // case 2
    "https://i.imgur.com/U0vao8u.png", // case 3
    "https://i.imgur.com/U0vao8u.png", // case 4
    "https://i.imgur.com/U0vao8u.png", // case 5
    "https://i.imgur.com/oQGgbvv.png", // case 6
    "https://i.imgur.com/YF2yS0b.png", // case 7
    "https://i.imgur.com/YF2yS0b.png", // case 8
    "https://i.imgur.com/tA4BaUh.png"  // case 9
  ];
  UserStore.PFP = profilePictures[number];
}

async function updatePassword() {
  const user = {
    "ID": UserStore.User.ID,
    "Email": "",
    "NewPassword": newPassword.value,
    "OldPassword": oldPassword.value
  }

  visiblePass.value = false;

  try{
    const response = await fetch("http://localhost:8080/User", {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      const errorsJson = await response.json();
      const errors = errorsJson.Errors;
      for (let i = 0; i < errors.length; i++) {
        toast.add({
          severity: "error",
          summary: "Error",
          detail: errors[i],
          life: 3000
        });
      }

      return;
    }

    visiblePass.value = false;

    const successText = await response.text();
    toast.add({
      severity: "success",
      summary: "Success",
      detail: successText,
      life: 3000
    });
  } catch (error){

  }
}


async function updateEmail() {
  const user = {
    "ID": UserStore.User.ID,
    "Email": newEmail.value,
    "OldPassword": '', // if empty don't update it
    "NewPassword": '', // if empty don't update it
  }

  visibleEmail.value = false;

  try {
    const response = await fetch("http://localhost:8080/User", {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      const errorsJson = await response.json();
      const errors = errorsJson.Errors;
      for (let i = 0; i < errors.length; i++) {
        toast.add({
          severity: "error",
          summary: "Error",
          detail: errors[i],
          life: 3000
        });
      }

      return;
    }

    const successText = await response.text();
    toast.add({
      severity: "success",
      summary: "Success",
      detail: successText,
      life: 3000
    });
  } catch (error){

  }
}
</script>

<template>
  <Toast />
  <div class="Card">
    <Card style="width: 25rem; overflow: hidden">
      <template #header>
      </template>
      <template #title>
        <h2 class="text-center">GoGarden Settings</h2>
      </template>
      <template #content>
        <p class="text-center">
          Here's where you can edit settings
        </p>
      </template>
      <template #footer>
        <div class="button-group">
          <div class="card flex justify-center">
            <Button label="Update Email" @click="visibleEmail = true" />
            <Dialog v-model:visible="visibleEmail" modal header="Edit User Email" :style="{ width: '25rem' }">
              <div class="flex items-center gap-4 mb-8">
                <label for="newEmail" class="font-semibold w-24">New Email</label>
                <div> </div>
                <InputText id="newEmail" class="flex-auto" v-model="newEmail" autocomplete="off" />
              </div>
              <div class="buttons">
                <div class="flex justify-end gap-2">
                  <Button type="button" label="Cancel" severity="secondary" @click="visibleEmail = false"></Button>
                  <Button type="button" label="Save" @click=updateEmail></Button>
                </div>
              </div>

            </Dialog>
          </div>

          <div class="card flex justify-center">
            <Button label="Update Password" @click="visiblePass = true" />
            <Dialog v-model:visible="visiblePass" modal header="Edit User Password" :style="{ width: '25rem' }">
              <div class="flex items-center gap-4 mb-4">
                <label for="oldPassword" class="font-semibold w-24">Old Password</label>
                <div> </div>
                <InputText id="oldPassword" class="flex-auto" v-model="oldPassword" autocomplete="off" />
              </div>
              <div class="flex items-center gap-4 mb-8">
                <label for="newPassword" class="font-semibold w-24">New Password</label>
                <div> </div>
                <InputText id="newPassword" class="flex-auto" v-model="newPassword" autocomplete="off" />
              </div>
              <div class="buttons">
                <div class="flex justify-end gap-2">
                  <Button type="button" label="Cancel" severity="secondary" @click="visiblePass = false"></Button>
                  <Button type="button" label="Save" @click="updatePassword"></Button>
                </div>
              </div>

            </Dialog>
          </div>

          <div>
            <Button label="Random Profile Picture" @click="updatePFP"></Button>
          </div>

          <div>
            <Button label="Delete Account" outlined class="deletebutton" @click="tryDelete"></Button>
          </div>
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
.deletebutton {
  border-color: red !important;
  color: red !important;
}
.deletebutton:hover {
  background-color: red !important;
  color: white !important;
}
.buttons {
  margin-top: 1rem;
}
</style>
