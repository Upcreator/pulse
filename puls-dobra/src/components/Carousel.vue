<template>
  <section id="carousel">
    <h2 class="text-4xl font-bold text-blue-700 text-center pt-8 pb-10">Отчёт о проделанной работе</h2>
    <div id="default-carousel" class="relative w-full z-0" data-carousel="slide">
      <!-- Carousel wrapper -->
      <div class="relative h-56 overflow-hidden rounded-lg md:h-96">
        <!-- Slides -->
        <div
          v-for="(item, index) in items"
          :key="index"
          class="absolute top-0 left-0 w-full h-full transition-opacity duration-700 ease-in-out"
          :style="{ opacity: currentSlide === index ? 1 : 0 }"
          v-show="currentSlide === index"
        >
          <div class="w-full h-full bg-blue-400 flex items-center justify-center text-white text-xl">
            {{ index + 1 }}
          </div>
        </div>
      </div>
      <!-- Slider indicators -->
      <div class="absolute z-20 flex -translate-x-1/2 bottom-5 left-1/2 space-x-3 rtl:space-x-reverse">
        <button
          v-for="(item, index) in items"
          :key="'indicator-' + index"
          type="button"
          class="w-3 h-3 rounded-full"
          :class="{ 'bg-blue-500': currentSlide === index, 'bg-gray-300': currentSlide !== index }"
          @click="goToSlide(index)"
          :aria-current="currentSlide === index ? 'true' : 'false'"
          :aria-label="'Slide ' + (index + 1)"
        ></button>
      </div>
      <!-- Slider controls -->
      <button
        type="button"
        class="absolute top-0 start-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
        @click="prevSlide"
      >
        <span class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none">
          <svg class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4"/>
          </svg>
          <span class="sr-only">Previous</span>
        </span>
      </button>
      <button
        type="button"
        class="absolute top-0 end-0 z-30 flex items-center justify-center h-full px-4 cursor-pointer group focus:outline-none"
        @click="nextSlide"
      >
        <span class="inline-flex items-center justify-center w-10 h-10 rounded-full bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none">
          <svg class="w-4 h-4 text-white dark:text-gray-800 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
          </svg>
          <span class="sr-only">Next</span>
        </span>
      </button>
    </div>
  </section>
</template>

<script setup>
import { ref } from 'vue';

const items = new Array(5).fill(null).map((_, index) => ({ index }));

const currentSlide = ref(0);

const goToSlide = (index) => {
  currentSlide.value = index;
};

const nextSlide = () => {
  currentSlide.value = (currentSlide.value + 1) % items.length;
};

const prevSlide = () => {
  currentSlide.value = (currentSlide.value - 1 + items.length) % items.length;
};

setInterval(nextSlide, 3000); // Слайд меняется каждые 3 сек
</script>

<style scoped>
#default-carousel {
  max-width: 100%;
}

.transition-opacity {
  transition: opacity 0.7s ease-in-out;
}
</style>
