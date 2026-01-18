<script setup>
import { ref } from 'vue';
import InputGroup from 'primevue/inputgroup';
import InputGroupAddon from 'primevue/inputgroupaddon';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
import router from '@/router';

const toast = useToast();
const email = ref('');
const password = ref('');
const isRegister = ref(false);

const toggleForm = () => {
  isRegister.value = !isRegister.value;
};

async function tryLogin() {
  const user = {
    "Email": email.value,
    "Password": password.value
  };

  try {
    const response = await fetch("http://localhost:8080/User/Login", {
      method: "POST",
      credentials: "include",
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

    await router.push('/home');
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Network Error",
      detail: error.message,
      life: 3000
    });
  }
}

async function tryRegister() {
  const user = {
    "Email": email.value,
    "Password": password.value
  };

  try {
    const response = await fetch("http://localhost:8080/User", {
      method: "POST",
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

      return
    }

    const successText = await response.text();
    toast.add({
      severity: "success",
      summary: "Success",
      detail: successText,
      life: 3000
    });

    await tryLogin();
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Network Error",
      detail: error.message,
      life: 3000
    });
  }
}
</script>

<template>
  <Toast />
  <div class="auth-container">
    <div class="auth-card">
      <h2 class="auth-title">{{ isRegister ? 'Register' : 'Sign In' }}</h2>
      <div class="auth-form">
        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-envelope"></i>
          </InputGroupAddon>
          <InputText placeholder="Email" v-model="email" />
        </InputGroup>

        <InputGroup>
          <InputGroupAddon>
            <i class="pi pi-lock"></i>
          </InputGroupAddon>
          <InputText placeholder="Password" type="password" v-model="password" />
        </InputGroup>

        <Button v-if="!isRegister" label="Sign In" @click="tryLogin()"/>
        <Button v-if="isRegister" label="Register" @click="tryRegister()"/>
      </div>

      <Button
        variant="link"
        :label="`Switch to ${isRegister ? 'Sign In' : 'Register'}`"
        @click="toggleForm()"
      />
    </div>
  </div>
</template>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100vw;
  background-color: #afb1b3;
  margin: 0;
}

.auth-card {
  width: 100%;
  max-width: 450px;
  border-radius: 10px;
  background-color: #d1d0d0;
  padding: 2px;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
}

.auth-card:hover {
  box-shadow: 0 6px 18px rgba(0, 0, 0, 0.5);
}

.auth-title {
  font-size: 2rem;
  font-weight: 700;
  color: #333;
  text-align: center;
  margin-bottom: 1.5rem;
}

.auth-form {
  display: flex;
  flex-direction: column;
  gap: 1.2rem;
  width: 80%;
  padding: 20px;
  border-radius: 8px;
  background-color: #f9f9f9;
}

.submit-button {
  margin-top: 1rem;
  width: 100%;
  height: 3rem;
  font-weight: 600;
  background-color: #007bff;
  border: none;
  color: white;
  border-radius: 8px;
  transition: background-color 0.3s ease;
}

.submit-button:hover {
  background-color: #0056b3;
}
</style>
