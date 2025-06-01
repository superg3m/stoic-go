<script>
import { defineComponent } from 'vue';
import Card from 'primevue/card';
import InputText from 'primevue/inputtext';

export default defineComponent({
  name: 'ResourcesPage',
  components: {
    Card,
    InputText
  },
  data() {
    return {
      searchTerm: '',
      resources: [
        {
          title: 'Vue.js Documentation',
          description: 'Learn the fundamentals and advanced topics of Vue.js.',
          url: 'https://vuejs.org/'
        },
        {
          title: 'PrimeVue',
          description: 'Explore the PrimeVue UI components for Vue.js applications.',
          url: 'https://www.primefaces.org/primevue/'
        },
        {
          title: 'JavaScript Info',
          description: 'A tutorial on JavaScript with examples.',
          url: 'https://javascript.info/'
        }
      ]
    };
  },
  computed: {
    filteredResources() {
      if (!this.searchTerm) {
        return this.resources;
      }
      const term = this.searchTerm.toLowerCase();
      return this.resources.filter(resource =>
        resource.title.toLowerCase().includes(term) ||
        resource.description.toLowerCase().includes(term)
      );
    }
  }
});
</script>

<template>
  <div class="page-container">
    <div class="header_text">
      <h1>Resources Page</h1>
    </div>
    <div class="search-container">
      <span class="p-input-icon-left">
        <i class="pi pi-search"></i>
        <InputText v-model="searchTerm" placeholder="Search resources..." />
      </span>
    </div>
    <transition-group name="card" tag="div" class="card-container">
      <Card
        v-for="(resource) in filteredResources"
        :key="resource.title"
        class="resource-card">
        <template #title>
          {{ resource.title }}
        </template>
        <template #content>
          <p>{{ resource.description }}</p>
          <a :href="resource.url" target="_blank">
            Learn More <i class="pi pi-external-link"></i>
          </a>
        </template>
      </Card>
    </transition-group>
  </div>
</template>

<style scoped>
.page-container {
  background: linear-gradient(135deg, #f5f7fa, #c3cfe2);
  min-height: 100vh;
  padding: 20px;
}

.header_text {
  font-size: xx-large;
  text-align: center;
  margin-bottom: 20px;
  color: black;
}

.search-container {
  text-align: center;
  margin-bottom: 30px;
}

.card-container {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
}

.resource-card {
  width: 300px;
  transition: transform 0.3s, box-shadow 0.3s;
}
.resource-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 8px 16px rgba(0,0,0,0.2);
}

.card-enter-active, .card-leave-active {
  transition: all 0.5s;
}
.card-enter-from, .card-leave-to {
  opacity: 0;
  transform: translateY(20px);
}
</style>
