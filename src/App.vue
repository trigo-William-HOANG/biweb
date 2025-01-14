<script setup>
import { ref, computed } from 'vue';
import Home from './pages/Home.vue';
import { useAuth } from './composables/useAuth';

const { checkAuth } = useAuth();

// Define routes with auth requirements
const routes = {
  '/': {
    component: Home,
    requiresAuth: true
  }
};

// Reactive variable for the current path
const currentPath = ref(window.location.pathname);

// Function to update currentPath with auth check
const updatePath = async () => {
  const path = window.location.pathname;
  const route = routes[path];
  
  if (route && route.requiresAuth) {
    const isAuthenticated = await checkAuth();
    if (!isAuthenticated) {
      window.location.href = 'http://localhost:5000/auth/signin';
      return;
    }
  }
  
  currentPath.value = path;
};

// Listen for changes to the URL
window.addEventListener("popstate", updatePath);

// Initial auth check
updatePath();

// Compute the component to render based on the current path
const currentView = computed(() => {
  return (routes[currentPath.value] || routes['/404']).component;
});
</script>

<template>
  <component :is="currentView" />
</template>
<style>
*{
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Helvetica', sans-serif;
    border-radius: 1%;
}

.Cards{
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 75px;
    margin: 10px;
    padding: 50px;
  }

  
body{
    background-color: var(--background);    
}
main{
  min-height: 86dvh;
    max-height: fit-content;
    margin: 1%;
    padding: 2%;
    background-color: white;
    box-shadow: 0 .5rem 1rem rgba(0, 0, 0, .15);
}
.page {
    margin-top: 20px;
  }
  
  button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
  </style>