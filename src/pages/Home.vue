<template>
  <Layout>
    <template #link class="pre-barre">
      <div class="header-content">
        <Searchbar class="barre" @update:search="updateSearchQuery" />
        <div class="auth-section">
          <a v-if="!isAuthenticated" 
             href="http://localhost:5000/auth/signin" 
             class="login-link">
             Se connecter
          </a>
          <div v-else class="user-section">
            <span class="user-status">Connecté</span>
            <a href="http://localhost:5000/auth/signout" class="logout-link">
              Se déconnecter
            </a>
          </div>
        </div>
      </div>
    </template>
    <template #content-title>
      <h1>{{ title }}</h1>
    </template>
    <template #content-subtitle>
      <p>Vous pouvez trouver ici toutes les applications TRIGO</p>
    </template>
    <template #content>
      <div class="main">
        <CardList :cards="filteredCards" :isLoading="isLoading" />
      </div>
    </template>
  </Layout>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import Layout from "../components/Layout.vue";
import CardList from "../components/CardList.vue";
import Searchbar from "../components/Searchbar.vue";
import { useAuth } from '../composables/useAuth';


const { isAuthenticated, checkAuth } = useAuth()

onMounted(async () => {
  const auth = await checkAuth()
  if (!auth) {
    window.location.href = 'http://localhost:5000/auth/signin' // A changer avec le lien de connexion de l'API
  }
})

import app1 from '../assets/app1.webp';
import app2 from '../assets/app2.webp';
import app3 from '../assets/app3.webp';

const cardData = [
  {
    image: app1,
    title: 'A 1',
    description: 'Description pour la carte 1.',
    link: '/test'
  },
  {
    image: app2,
    title: 'B 2',
    description: 'Description pour la carte 2.',
    link: '#'
  },
  {
    image: app3,
    title: 'C 3',
    description: 'Description pour la carte 3.',
    link: '#'
  },
  {
    image: app3,
    title: 'C 3',
    description: 'Description pour la carte 3.',
    link: '#'
  },
  {
    image: app3,
    title: 'C 3',
    description: 'Description pour la carte 3.',
    link: '#'
  },
  {
    image: app3,
    title: 'C 3',
    description: 'Description pour la carte 3.',
    link: '#'
  },
  {
    image: app3,
    title: 'C 3',
    description: 'Description pour la carte 3.',
    link: '#'
  },
  {
    image: app3,
    title: 'C 3',
    description: 'Description pour la carte 3.',
    link: '#'
  },
  {
    image: app3,
    title: 'C 3',
    description: 'Description pour la carte 3.',
    link: '#'
  }
];

const searchQuery = ref('');
const isLoading = ref(false);
const filteredCards = ref([]);
const title = ref('Applications');

const fetchCards = async (query) => {
  if (!query.trim()) {
    filteredCards.value = cardData;
  } else {
    filteredCards.value = cardData.filter(card =>
      card.title.toLowerCase().includes(query.toLowerCase())
    );
  }
};

const updateSearchQuery = (query) => {
  searchQuery.value = query;
  fetchCards(query);
};

fetchCards('');

</script>

<style scoped>
.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.auth-section {
  margin-left: 1rem;
}

.login-link, .logout-link {
  color: #007bff;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border: 1px solid #007bff;
  border-radius: 4px;
}

.logout-link {
  margin-left: 1rem;
  color: #dc3545;
  border-color: #dc3545;
}

.user-status {
  color: #28a745;
}

.user-section {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.content-main {
  transition: all 0.3s ease;
  margin-top: 3%;
}

/* Contenu adapté au comportement de la sidebar */
body.sidebar-closed .content-main {
  margin-left: 60px;
}
.searchbar{
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(183%, -50%);
}

</style>