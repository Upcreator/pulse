<template>
  <div class="app">
    <div class="filter-controls">
      <select v-model="filter" @change="filterComments">
        <option value="all">Все пожелания</option>
        <option value="true">Отображаемые пожелания</option>
        <option value="false">Скрытые пожелания</option>
      </select>
    </div>

    <div v-if="filteredComments.length === 0" class="no-comments">
      Нет таких пожеланий.
    </div>

    <CommentCard
      v-for="comment in filteredComments"
      :key="comment.id"
      :comment="comment"
      @updateComment="fetchComments"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios';
import CommentCard from './components/CommentCard.vue';

const comments = ref([]);
const filter = ref('all');

const fetchComments = async () => {
  try {
    const response = await axios.get('https://puls-dobra.ru/api/comments');
    comments.value = response.data;
  } catch (error) {
    console.error('Error fetching comments:', error);
  }
};

const filteredComments = computed(() => {
  if (filter.value === 'all') {
    return comments.value;
  }
  return comments.value.filter(comment => comment.visibility.toString() === filter.value);
});

onMounted(fetchComments);
</script>

<style>
.app {
  padding: 20px;
}

.filter-controls {
  margin-bottom: 20px;
}

.no-comments {
  text-align: center;
  color: #888;
}
</style>
