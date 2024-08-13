<template>
  <div id="app">
    <LoadingScreen v-if="isLoading" />
    <div v-else>
      <Header />
      <HeroSection />
      <DateRange />
      <CharitySection />
      <CarouselComments />
      <ModalForm/>
      <Carousel />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import LoadingScreen from './components/LoadingScreen.vue';
import HeroSection from './components/HeroSection.vue';
import DateRange from './components/DateRange.vue';
import CharitySection from './components/CharitySection.vue';
import ProgressBar from './components/ProgressBar.vue';
import ModalForm from './components/ModalForm.vue';
import CarouselComments from './components/CarouselComments.vue';
import Carousel from './components/Carousel.vue';
import Header from './components/Header.vue';
const isLoading = ref(true);

onMounted(() => {
  setTimeout(() => {
    isLoading.value = false;
  }, 3000); // 3 сек для загрузочного экрана
});
</script>

<style>
body, html {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow-x: hidden;
}
</style>
