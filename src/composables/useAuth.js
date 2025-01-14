import { ref } from 'vue'

export function useAuth() {
  const isAuthenticated = ref(false)

  const checkAuth = async () => {
    try {
      const response = await fetch('http://localhost:5000/auth/status', {
        credentials: 'include' // Required for sending cookies
      })
      const data = await response.json()
      isAuthenticated.value = data.isAuthenticated
      console.log('Auth status:', isAuthenticated.value ? 'Logged in' : 'Not logged in')
      return data.isAuthenticated
    } catch (error) {
      console.error('Auth check failed:', error)
      isAuthenticated.value = false
      return false
    }
  }

  return {
    isAuthenticated,
    checkAuth
  }
}