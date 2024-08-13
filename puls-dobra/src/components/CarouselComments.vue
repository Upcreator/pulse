<template>
  <section id="carousel-section">
    <div class="carousel-container">
      <h2>Ваши пожелания</h2>

      <!-- Desktop Carousel -->
      <div v-if="isDesktop" class="desktop-carousel">
        <Carousel
          :items-to-show="3"
          :items-to-scroll="1"
          :autoplay="3"
          :loop="true"
          :pagination="true"
          :navigation="true"
          class="carousel"
        >
          <Slide v-for="(comment, index) in comments" :key="comment.id">
            <div class="carousel-item">
              <p class="comment-text">{{ comment.text }}</p>
              <p class="comment-author">- {{ comment.author }}</p>
            </div>
          </Slide>
          <template #addons>
            <Navigation class="carousel-navigation" />
            <Pagination class="carousel-pagination" />
          </template>
        </Carousel>
      </div>

      <!-- Mobile Carousel -->
      <div v-else class="mobile-carousel">
        <Carousel
          :items-to-show="1"
          :items-to-scroll="1"
          :autoplay="false"
          :loop="true"
          :pagination="true"
          :navigation="true"
          class="carousel"
        >
          <Slide v-for="(comment, index) in comments" :key="comment.id">
            <div class="carousel-item">
              <p class="comment-text">{{ comment.text }}</p>
              <p class="comment-author">- {{ comment.author }}</p>
            </div>
          </Slide>
          <template #addons>
            <Navigation class="carousel-navigation" />
            <Pagination class="carousel-pagination" />
          </template>
        </Carousel>
      </div>
    </div>
  </section>
</template>

<script>
import { Carousel, Slide, Navigation, Pagination } from 'vue3-carousel';
import 'vue3-carousel/dist/carousel.css';
import { ref, onMounted, onUnmounted } from 'vue';
import axios from 'axios';

export default {
  components: {
    Carousel,
    Slide,
    Navigation,
    Pagination
  },
  setup() {
    const comments = ref([]);
    const isDesktop = ref(window.innerWidth >= 800);

    const fetchComments = async () => {
      try {
        const response = await axios.get('http://87.228.19.168:8080/comments');
        comments.value = response.data.filter(comment => comment.visibility); // Пожелания сразу фильтруются на основе visibility
      } catch (error) {
        console.error('Error fetching comments:', error);
      }
    };

    const handleResize = () => {
      isDesktop.value = window.innerWidth >= 800;
    };

    onMounted(() => {
      fetchComments(); 
      window.addEventListener('resize', handleResize); 
    });

    onUnmounted(() => {
      window.removeEventListener('resize', handleResize); 
    });

    return { comments, isDesktop };
  }
};
</script>

<style scoped>
#carousel-section {
  padding: 2rem;
  background-color: #f9f9f9;
  text-align: center;
}

.carousel-container {
  max-width: 1000px;
  margin: 0 auto;
}

.carousel-container h2 {
  font-size: 2rem;
  margin-bottom: 1rem;
  color: #333;
}

.carousel-item {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.2);
  padding: 1rem;
  text-align: center;
}

.comment-text {
  font-size: 1.2rem;
  color: #333;
}

.comment-author {
  font-size: 1rem;
  color: #777;
  margin-top: 0.5rem;
}

.carousel-navigation {
}

.carousel-pagination {
}

@media (min-width: 800px) {
  .mobile-carousel {
    display: none;
  }
}

@media (max-width: 799px) {
  .desktop-carousel {
    display: none;
  }
}
</style>
