import { reactive } from 'vue'

export const UserStore = reactive({
  Authorized: false,
  User: null,
  PFP: "https://i.imgur.com/670yOS4.png"
})

async function getUser(userId) {
  try {
    const response = await fetch(`http://localhost:8080/User?id=${userId}`, {
      method: "GET",
      credentials: "include",
    });

    const text = await response.text()
    UserStore.User = JSON.parse(text)
  } catch (e) {
    console.warn(e)
  }
}

export async function isAuthorized() {
  try {
    const response = await fetch("http://localhost:8080/User/Authorized", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      }
    });

    if (!response.ok) {
      UserStore.Authorized = false
      return false
    }

    const text = await response.text()
    await getUser(parseInt(text))
  } catch (e) {
    console.error(e)
    UserStore.Authorized = false
    return false
  }

  UserStore.Authorized = true
  return true
}
